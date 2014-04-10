package triblab

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"trib"
)

// Creates an RPC client that connects to addr.
func NewClient(addr string) trib.Storage {
	c := Client{addr: addr}
	/*
		if c.Connect() != nil {
			fmt.Println("fail to connect")
			//handle for connection failure
		}
	*/
	return &c
}

// Serve as a backend based on the given configuration
func ServeBack(b *trib.BackConfig) error {
	e := rpc.Register(b.Store)
	if e != nil {
		fmt.Println(e)
		return e
	}
	rpc.HandleHTTP()
	s, e := net.Listen("tcp", b.Addr)
	if e != nil {
		fmt.Println("server tcp error:", e)
	} else {
		go http.Serve(s, nil)
		b.Ready <- true
	}
	return nil
}
