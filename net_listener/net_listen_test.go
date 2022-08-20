package net_listener

import (
	"bytes"
	"log"
	"net"
	"testing"
)

// Listen on TCP port 12000 on all available unicast and
// anycast IP addresses of the local system.

func Test_Net_Listen(t *testing.T) {
	l, err := net.Listen("tcp", ":12000")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection
		conn, err := l.Accept()
		if err != nil {
			t.Fatal(err)
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			log.Printf("Got data from : %s", c.RemoteAddr())

			// var buf bytes.Buffer
			// buf.Grow(256)

			var databuf bytes.Buffer

			for {
				var data []byte = make([]byte, 512)

				// conn.Read() func would block but io.Copy(dst, src) wouldn't
				n, err := c.Read(data)
				databuf.Write(data)
				if n == 0 {
					c.Close()
					break
				}
				if err != nil {
					c.Close()
					log.Fatal("Can not read data from the connection")
				}
				log.Printf("Received %d bytes data : %s\n", n, databuf.String())
				databuf.Reset()
			}
		}(conn)
	}

}
