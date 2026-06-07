// HTTP client 
// .DefaultClient is the default Client and is used by Get, Head, and Post
// var DefaultClient = &Client{}
client := &http.Cleint{}
resp, err := client.Get("https://golang.org/")
// or 
resp, err := http.Get("https://golang.org/")

// Five steps involved in an HTTP request:
// 1. Dial to establish a TCP connection
// 2. TLS handshake (if enabled)
// 3. Send the request 
// 4. Read the response headers
// 5. Read the response body []

// Default has no timeout. Can exhause resources 
// net.Dialer.Timeout
// 	maximum amount of time a dial will wait for a connection to complete 
// http.Transport.TLSHandeshakeTimeout
// 	maximum amount of time to wait for the TLS handshake 
// http.Transport.ResponseHeaderTimeout
// 	amount o ftime to wait for a server's response headers 
// http.Client.Timeout 
// 	time limit for a request. Includes all 5 steps 

client := &http.Client {
	Timout: 5 * time.Second, 
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: time.Second,
		}).DialContext,
		TLSHandshakeTimeout: time.Second,
		ResponseHeaderTimeout: time.Second,
	},
}

// http.Transport.IdelConnTimeout - default 90s
// 	how long an idle connection stays in the pool before being closed
// http.Transport.MaxIdleConns - default 100
// 	total idle connections across ALL hosts
// http.Transport.MaxIdleConnsPerHost - default 2
// 	idle connections per individual host — the binding constraint when hitting one host
// 	e.g. 100 idle allowed globally but only 2 kept for api.example.com
// 	increase this if your app hammers a single host with concurrent requests


// HTTP server 
server := &http.Server{}
server.Serve(Listener)

// 1. Wait for client to send the request 
// 2. TLS handshake (if enabled) (doesn't have to be repeated with an already established connection)
// 3. Read the request headers 
// 4. Read the request body 
// 5. Write the response  

// http.Server.ReadHeaderTimeout
// 	amount of time to read the header 
// http.Server.ReadTimeout 
// 	amount of time to read the entire request 
// http.TimeoutHandler 
// 	wrapper function that specifies the maximum amount of time for a handler to complete 

// should always set up in case a bad actor exhausts our resources

s := &http.Server {
	Addr: ":8080",
	ReadHeaderTimeout: 500 * time.Millisecond,
	ReadTimeout: 500 * time.Millisecond,
	Handler: http.TimeoutHandler(handler, time.Second, "foo"), // Wraps the HTTP handler 
}

// http.Server.IdleTimout
// 	how long the server keeps a keep-alive connection open waiting for the client's next request
// 	without this, a slow/idle client holds a connection open indefinitely
s := &http.Server {
	// ...
	IdleTimout: time.Second
}
