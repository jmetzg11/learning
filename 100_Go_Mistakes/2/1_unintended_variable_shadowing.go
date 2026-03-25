var client *http.Client
if tracing {
	client, err := createClientWithTracing()
	...
	log.Println(client)
} else {
	client, err := createDefaultClient()
	...
	log.Println(client)
}
// use client

// compiles because of the log, or else 'cleint declared and not used

// Solution # 1 - assign to temp variable and then assign to client

var client **http.Client
if tracing {
	c, err := createClientWithTracing()
	...
	client = c
} else {
	// same logic
}

// Solution # 2 - assign in scope, have to create an error variable too

var client *http.Client
var err error
if tracing {
	client, err = createClientWithTracing()
	...
} else {
	// same logic
}

var client *http.Client
var err error
if tracing {
	client, err = createClientWithTracing()
} else {
	cleint, err = createDefaultClient()
}
if err != nil {
	// common error handling
}
