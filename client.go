package triblab

import (
	"net/rpc"
	"trib"
)

type Client struct {
	Addr string
}

type Connection struct {
	conn       *rpc.Client
	is_connect bool
}

func (self *Connection) connect(addr string) error {
	var e error
	if !self.is_connect {
		self.conn, e = rpc.DialHTTP("tcp", addr)
	}
	if e == nil {
		self.is_connect = true
	}
	return e
}

var globalConnection *Connection = new(Connection)

func stub(addr string, method string, args interface{}, reply interface{}) error {
	e := globalConnection.connect(addr)
	if e != nil {
		return e
	}

	e = globalConnection.conn.Call(method, args, reply)
	if e != nil {
		globalConnection.is_connect = false
		return e
	}
	return nil
}

func (self *Client) Get(key string, value *string) error {
	return stub(self.Addr, "Storage.Get", key, value)
}

func (self *Client) Set(kv *trib.KeyValue, succ *bool) error {
	return stub(self.Addr, "Storage.Set", kv, succ)
}

func (self *Client) Keys(p *trib.Pattern, list *trib.List) error {
	list.L = []string{}
	return stub(self.Addr, "Storage.Keys", p, list)
}

func (self *Client) ListGet(key string, list *trib.List) error {
	list.L = []string{}
	return stub(self.Addr, "Storage.ListGet", key, list)
}

func (self *Client) ListAppend(kv *trib.KeyValue, succ *bool) error {
	return stub(self.Addr, "Storage.ListAppend", kv, succ)
}

func (self *Client) ListRemove(kv *trib.KeyValue, n *int) error {
	return stub(self.Addr, "Storage.ListRemove", kv, n)
}

func (self *Client) ListKeys(p *trib.Pattern, list *trib.List) error {
	list.L = []string{}
	return stub(self.Addr, "Storage.ListKeys", p, list)
}

func (self *Client) Clock(atLeast uint64, ret *uint64) error {
	return stub(self.Addr, "Storage.Clock", atLeast, ret)
}
