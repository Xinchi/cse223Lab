package triblab

import (
	// "fmt"
	"net"
	"net/http"
	"net/rpc"
	"trib"
)

// Creates an RPC client that connects to addr.
func NewClient(addr string) trib.Storage {
	return &Client{Addr: addr}
}

// Serve as a backend based on the given configuration
// func ServeBack(b *trib.BackConfig) error {
// 	server := rpc.NewServer()
// 	server.Register(b.Store)
// 	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
// 	listener, e := net.Listen("tcp", b.Addr)
// 	if e != nil {
// 		if b.Ready != nil {
// 			b.Ready <- false
// 		}
// 		return e
// 	}
// 	if b.Ready != nil {
// 		b.Ready <- true
// 	}
// 	return http.Serve(listener, nil)
// 	// return nil
// }
// func ServeBack(b *trib.BackConfig) error {
// 	e := rpc.Register(b.Store)
// 	if e != nil {
// 		fmt.Println(e)
// 		return e
// 	}
// 	rpc.HandleHTTP()
// 	s, e := net.Listen("tcp", b.Addr)
// 	if e != nil {
// 		fmt.Println("server tcp error:", e)
// 	} else {
// 		go http.Serve(s, nil)
// 		b.Ready <- true
// 	}
// 	return nil
// }


// Serve as a backend based on the given configuration
func ServeBack(b *trib.BackConfig) error {
	server := rpc.NewServer()
	server.Register(b.Store)
	//	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", b.Addr)
	if e != nil {
		if b.Ready != nil {
			b.Ready <- false
		}
		return e
	}
	if b.Ready != nil {
		b.Ready <- true
	}
	return http.Serve(listener, server)
}