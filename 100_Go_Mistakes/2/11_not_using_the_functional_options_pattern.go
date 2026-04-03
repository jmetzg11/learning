// we have an api to build a server
func NewServer(addr string, port int) (*http.Server, error) {
	//...
}

// now we want optional parameter
// do we pass a struct, doesn't work well if Config.Port is optional
type Config struct {
	Port int
}

func NewServer(addr string, cfg Config) {
	//...
}

// Builder pattern
package main

import (
	"errors"
	"net/http"
)

const defaultHTTPPort = 8080

type Config struct {
	Port int
}

type ConfigBuilder struct {
	port *int
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b // return builder, possible to do chain calling then
}

func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port == nil {
		cfg.Port = defaultHTTPPort
	} else {
		if *b.port == 0 {
			cfg.Port = randomPort()
		} else if *b.port < 0 {
			return Config{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}

	return cfg, nil
}

func NewServer(addr string, config Config) (*http.Server, error) {
	return nil, nil
}

func client() error {
	builder := ConfigBuilder{}
	builder.Port(8080) // Client has the option to make a port
	cfg, err := builder.Build()
	if err != nil {
		return err
	}

	server, err := NewServer("localhost", cfg) // passing the config that didn't need a port declared
	if err != nil {
		return err
	}

	_ = server
	return nil
}

func randomPort() int {
	return 4 // Chosen by fair dice roll, guaranteed to be random.
}

// Functional options pattern
type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option{
	return func(options *options) err {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

func NewServer(add string, opts ..Option) (**http.Server, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options) // calls each option, which results in modifying the common options struct
		if err != nil {
			return nil, err
		}
	}

	var port int
	if options.port == nil {
		port = defaultHTTPPort
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}
}

fun main() {
	// srv, err := NewServer("localhost")
	srv, err := NewServer("localhost", WithPort(8080))
	// ...
}
