// http.Error doesn't stoep a hanlder execution and must be added manually

func handler(w http.ResponseWriter, req *http.Request) {
	err := foo(req)
	if err != nil {
		htpp.Error(w, "foo", http.StatusInternalServerError)
		// no return on err != nil, execution continues

		// solution:
		return
	}

	_, _ = w.Write([]byte("all good"))
	w.WriteHeader(http.StatusCreated)
}

// response contains error and response message
// foo
// all good
