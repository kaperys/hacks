// Copyright (c) 2020, Mike Kaperys <mike@kaperys.io>
// See LICENSE for licensing information

package main

import (
	"io"
	"net"
	"os"
)

func main() {
	port := "9090"
	if p := os.Getenv("VOID_PORT"); p != "" {
		port = p
	}

	l, err := net.Listen("tcp4", net.JoinHostPort("", port))
	if err != nil {
		panic(err)
	}
	defer l.Close()

	println("listening on :" + port)
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			// TODO(kaperys) configurable packet size?
			buf := make([]byte, 1024)
			for {
				len, err := c.Read(buf)
				if err != nil {
					if err == io.EOF {
						break
					}

					panic(err)
				}

				println(string(buf[:len]))
			}
		}(c)
	}
}
