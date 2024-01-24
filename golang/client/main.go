package main

import (
  "context"
  "log"
  "net"
  "time"
)

// borrowed from: https://go.dev/src/net/example_test.go
func ExampleDialer() {

	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", "localhost:2000")

	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	defer conn.Close()

	if _, err := conn.Write([]byte("Hello, World!")); err != nil {
		log.Fatal(err)
	}

}

func main() {
  ExampleDialer()
}
