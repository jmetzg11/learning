// If the struct implements the io.Closer interface, we must eventually clal the Close method


// HTTP body
type handler struct {
	client http.Client
	url string
}

func (h handler) getBody() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	// Always close the body if http.Get doesn't return a error
	// otherwise it's a dataleak
	defer func() {
		err := resp.Body.Close()
		if err ! nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// keep the connection alive if we if we are interested in status code but not the body content
func (h handler) getStatusCode(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}

	// Close response body

	// reads the body but discards it. More efficient than io.ReadAll
	_, _ = io.Copy(io.Discard, resp.Body)

	return resp.StatusCode, nil
}


// sql.Rows
db, err := sql.Open("postgres", dataSourceName)
if err != nil {
	return err
}

rows, err := dbQuery("SELECT * FROM CUSTOMERS")
if err != nil {
	return err
}

// forgetting to close the rows means a connection leak,
// which prevent the database connection from being put back into the connection pool.
defer func() {
	if err := rows.Close(); err != nil {
		log.Printf("failed to close rows: %v\n", err)
	}
}

// Use rows

return nil


// os.File
f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
if err != nil {
	return err
}

// Best to close because we don't know when the GC will be triggered
// also errors will be surfaced sooner
defer func() {
	if err := f.Close(); err != nil {
		log.Printf("failed to close file: %v\n", err)
	}
}

func writeToFile(filename string, content []byte) (err error) {
	// Open file

	defer func() {
		closeErr := f.Close()
		if err == nil { // write succeeded, so surface the close error if any
			err = closeErr
		}
	}()

	_, err = f.Write(content) // sets named return err
	return                    // naked return; defer runs here and can still modify err
}

// if durability is a critical factor, we can use the Sync() method to commit a change
// slight impact on performance
func writeToFile(filename string, content []byte) error {
	// Open file

	defer func() {
		_ = f.Close() // Ignores possible errors
	}()

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return f.Sync()
}
