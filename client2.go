package triblab

import (
	// "net/rpc"
	"trib"
	"trib/colon"
	"strings"
)

type Client2 struct {
	Addr string
	BinName string
	Client trib.Storage
	Escaped_Prefix string
}

// type Connection struct {
// 	conn       *rpc.Client
// 	is_connect bool
// }

// func (self *Connection) connect(addr string) error {
// 	var e error
// 	if !self.is_connect {
// 		self.conn, e = rpc.DialHTTP("tcp", addr)
// 	}
// 	if e == nil {
// 		self.is_connect = true
// 	}
// 	return e
// }

// var globalConnection *Connection = new(Connection)

// func stub(addr string, method string, args interface{}, reply interface{}) error {
// 	e := globalConnection.connect(addr)
// 	if e != nil {
// 		return e
// 	}

// 	e = globalConnection.conn.Call(method, args, reply)
// 	if e != nil {
// 		globalConnection.is_connect = false
// 		return e
// 	}
// 	return nil
// }



// func (self *Client2) Get(key string, value *string) error {

// 	key2 := clone.Escape(BinName)+"::"+clone.Escape(key)

// 	return stub(self.Addr, "Storage.Get", key2, value)
// }


// func getPrefix() string{
// 	return colon.Escape(self.BinName)+"::"
// }
func (self *Client2) Get(key string, value *string) error {
	//process key


	key2 := colon.Escape(self.BinName)+"::"+colon.Escape(key)

	return self.Client.Get(key2, value);
}

func (self *Client2) Set(kv *trib.KeyValue, succ *bool) error {
	kv.Key = colon.Escape(self.BinName)+"::"+colon.Escape(kv.Key)	
	return self.Client.Set(kv, succ);
}

func (self *Client2) Keys(p *trib.Pattern, list *trib.List) error {

	// BinName::abcd 

	// In Pat {a:b, c:d}
	// Pattern {BinName::a|:b, c|:d}
	p.Prefix = colon.Escape(self.BinName) + "::"+ colon.Escape(p.Prefix)
	p.Suffix = colon.Escape(p.Suffix)
	// return stub(self.Addr, "Storage.Keys", p, list)
	

	e := self.Client.Keys(p, list);
	for i,_ := range list.L{
		list.L[i] = strings.TrimLeft(list.L[i], colon.Escape(self.BinName)+"::")
	}
	return e
}


func (self *Client2) ListGet(key string, list *trib.List) error {
	list.L = []string{}
	key2 := colon.Escape(self.BinName)+"::"+colon.Escape(key)
	// return stub(self.Addr, "Storage.ListGet", key, list)
	return self.Client.ListGet(key2,list)
}

func (self *Client2) ListAppend(kv *trib.KeyValue, succ *bool) error {
	kv.Key = colon.Escape(self.BinName)+"::"+kv.Key
	return self.Client.ListAppend(kv, succ)
	// return stub(self.Addr, "Storage.ListAppend", kv, succ)
}

func (self *Client2) ListRemove(kv *trib.KeyValue, n *int) error {
	kv.Key = colon.Escape(self.BinName)+"::"+colon.Escape(kv.Key)
	return self.Client.ListRemove(kv,n)
	// return stub(self.Addr, "Storage.ListRemove", kv, n)
}

func (self *Client2) ListKeys(p *trib.Pattern, list *trib.List) error {
	list.L = []string{}
	p.Prefix = colon.Escape(self.BinName) + "::"+ colon.Escape(p.Prefix)
	p.Suffix = colon.Escape(p.Suffix)

	e := self.Client.ListKeys(p, list);
	for i,_ := range list.L{

		list.L[i] = list.L[i][len(colon.Escape(self.BinName)+"::"):]

		list.L[i] = colon.Unescape(list.L[i])
	}

	return e
	// return stub(self.Addr, "Storage.ListKeys", p, list)
}

func (self *Client2) Clock(atLeast uint64, ret *uint64) error {
	return self.Client.Clock(atLeast,ret)
	// return stub(self.Addr, "Storage.Clock", atLeast, ret)
}
