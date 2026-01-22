// This code demonstrates the functional options pattern in Go
package main

import "fmt"

// Server defines a struct with configuration options
type Server struct {
	host string
	port int
	tls  bool
}

// Option is a function type that takes a *Server pointer
type Option func(*Server)

// WithHost is an option function to set the host
func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

// WithTLS is an option function to enable TLS
func WithTLS(tls bool) Option {
	return func(s *Server) {
		s.tls = tls
	}
}

// NewServer creates a new Server with default and custom options
// Returning a pointer from NewServer is a common Go pattern because it is efficient,
// allows mutability, supports shared state, and aligns with idiomatic Go practices for constructors.
func NewServer(port int, options ...Option) *Server {
	// Set default values
	s := &Server{
		host: "localhost",
		port: port,
		tls:  false,
	}

	// Apply options
	for _, opt := range options {
		opt(s)
	}

	return s
}

func main() {
	// Initialize with default host and no TLS
	server1 := NewServer(8080)
	fmt.Printf("Server 1: Host: %s, Port: %d, TLS: %v\n", server1.host, server1.port, server1.tls)

	// Initialize with a custom host and TLS enabled
	server2 := NewServer(8443, WithHost("theapi.com"), WithTLS(true))
	fmt.Printf("Server 2: Host: %s, Port: %d, TLS: %v\n", server2.host, server2.port, server2.tls)
}
