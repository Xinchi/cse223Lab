package triblab

import (
	//	"net"
	"fmt"
	"net/rpc"
	"time"
	"trib"
)

type Client struct {
	addr string
	conn *rpc.Client
	err  error
}

var reg_name string = "Storage."
var max_attempt int = 5

func (self *Client) Connect() error {
	for i := 0; i < max_attempt; i++ {
		self.conn, self.err = rpc.DialHTTP("tcp", self.addr)
		if self.err != nil {
			fmt.Printf("connection attempt: %d : error = %s\n", i, self.err)
			time.Sleep(1 * time.Second)
		} else {
			return nil
		}
	}
	return self.err
}
func (self *Client) Clock(atLeast uint64, ret *uint64) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"Clock", atLeast, ret); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
func (self *Client) Get(key string, value *string) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"Get", key, value); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
func (self *Client) Set(kv *trib.KeyValue, succ *bool) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"Set", kv, succ); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
func (self *Client) Keys(p *trib.Pattern, r *trib.List) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"Keys", p, r); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
func (self *Client) ListKeys(p *trib.Pattern, r *trib.List) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"ListKeys", p, r); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
func (self *Client) ListGet(key string, ret *trib.List) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"ListGet", key, ret); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
func (self *Client) ListAppend(kv *trib.KeyValue, succ *bool) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"ListAppend", kv, succ); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
func (self *Client) ListRemove(kv *trib.KeyValue, n *int) error {
	if self.conn == nil {
		if self.err = self.Connect(); self.err != nil {
			return self.err
		}
	}
	for self.err = self.conn.Call(reg_name+"ListRemove", kv, n); self.err != nil; {
		self.err = self.Connect()
		if self.err != nil {
			break
		}
	}
	return self.err
}
