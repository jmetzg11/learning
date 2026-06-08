// The httptest package lets you test HTTP handlers without starting a real server

// Server
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}

func TestHandler(t *testing.T) {
	// builds a fake *http.Request — no network involved
	req := httptest.NewRequest(http.MethodGet, "http://localhost",
		strings.NewReader("foo"))
	// ResponseRecorder captures whatever the handler writes (headers, body, status)
	w := httptest.NewRecorder()
	Handler(w, req)

	if got := w.Result().Header.Get("X-API-VERSION"); got != "1.0" {
		t.Errorf("api version: expected 1.0, got %s", got)
	}

	// w.Body is the raw bytes buffer; w.Result().Body is an io.ReadCloser — both work
	body, _ := ioutil.ReadAll(w.Body)
	if got := string(body); got != "hello foo" {
		t.Errorf("body: expected hello foo, got %s", got)
	}

	// w.Result() converts the recorder into a *http.Response for inspection
	if http.StatusOK != w.Result().StatusCode {
		t.FailNow()
	}
}


// Client — testing code that makes HTTP calls (opposite of the recorder above)
// NewRecorder: you call the handler directly; no server runs
// NewServer: a real local server starts; your client code hits it over the network
func (c DurationClient) GetDuration(url string, lat1, lng1, lat2, lng2 float64) (time.Duration, error) {
	resp, err := c.client.Post(
		url, "application/json",
		buildRequestBody(lat1, lng1, lat2, lng2),
	)
	if err != nil {
		return 0, err
	}

	return parseResponseBody(resp.Body)
}

func TestDurationClientGet(t *testing.T) {
	// spins up a real HTTP server on a random local port — no external service needed
	srv := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// fake response the client will parse
				_, _ = w.Write([]byte(`{"duration": 314}`))
			},
		),
	)
	defer srv.Close()

	client := NewDurationClient()
	duration, err :=
		// srv.URL is e.g. "http://127.0.0.1:54321" — points at the local test server
		client.GetDuration(srv.URL, 51.551261, -0.1221146, 51.57, -0.13)
	if err != nil {
		t.Fatal(err)
	}

	if duration != 314*time.Second {
		t.Errorf("expected 314 seconds, got %v", duration)
	}
}


// The iotest package
// two main uses: validate a custom Reader impl, and simulate flaky readers

// iotest.TestReader — verifies a custom io.Reader follows the contract:
//   reads the right bytes, signals EOF correctly, handles edge cases
func TestLowerCaseReader(t *testing.T) {
	err := iotest.TestReader(
		&LowerCaseReader{reader: strings.NewReader("aBcDeFgHiJ")},
		[]byte("acegi"), // what the reader should produce
	)
	if err != nil {
		t.Fatal(err)
	}
}

// iotest.TimeoutReader — wraps a reader so it returns a timeout error after the first read
// use it to test whether your code handles transient read errors

// foo1 uses io.ReadAll — gives up on any error, not resilient
func foo1(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err // TimeoutReader will trigger this — TestFoo1 fails
	}
	_ = b
	return nil
}

// TestFoo1 fails: foo1 can't handle the mid-stream timeout error
func TestFoo1(t *testing.T) {
	err := foo1(iotest.TimeoutReader(
		strings.NewReader(randomString(1024)),
	))
	if err != nil {
		t.Fatal(err)
	}
}

// foo2 uses a custom readAll with retries — handles the transient error and keeps going
func foo2(r io.Reader) error {
	b, err := readAll(r, 3)
	if err != nil {
		return err
	}
	_ = b
	return nil
}

// TestFoo2 passes: retries absorb the timeout error
func TestFoo2(t *testing.T) {
	err := foo2(iotest.TimeoutReader(
		strings.NewReader(randomString(1024)),
	))
	if err != nil {
		t.Fatal(err)
	}
}

// retries on non-EOF errors up to `retries` times before giving up
func readAll(r io.Reader, retries int) ([]byte, error) {
	b := make([]byte, 0, 512)
	for {
		if len(b) == cap(b) {
			b = append(b, 0)[:len(b)]
		}
		n, err := r.Read(b[len(b):cap(b)])
		b = b[:len(b)+n]
		if err != nil {
			if err == io.EOF {
				return b, nil
			}
			retries--
			if retries < 0 {
				return b, err
			}
		}
	}
}

