package main

import (
  "net"
  "log"
  "io"
  "fmt"
)

// borrowed from: https://go.dev/src/net/example_test.go
func ExampleListener() {

  // Listen on TCP port 2000 on all available unicast and
  // anycast IP addresses of the local system.
  l, err := net.Listen("tcp", ":2000")

  if err != nil {
    log.Fatal(err)
  }

  defer l.Close()

  for {

    // Wait for a connection.
    conn, err := l.Accept()

    if err != nil {
      log.Fatal(err)
    }

    // Handle the connection in a new goroutine.
    // The loop then returns to accepting, so that
    // multiple connections may be served concurrently.
    go func(c net.Conn) {

      buf, err := io.ReadAll(c)

      if err != nil {
        log.Fatal(err)
      }

      fmt.Printf("Received: %s\n", buf)

      // Echo all incoming data.
      io.Copy(c, c)

      // Shut down the connection.
      c.Close()
    }(conn)
  }
}

func main() {

  ExampleListener()

}
