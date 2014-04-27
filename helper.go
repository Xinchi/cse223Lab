package triblab

import (
	"trib"
	"hash/fnv"
	// "fmt"
)

type helper struct {
	Backs []string
}

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}

func (self *helper) Bin(name string) trib.Storage{
	//get the hash value
	a := hash(name)
	b := uint32(len(self.Backs))
	a = a%b
	//chord
	addr := self.Backs[a]
	client := NewClient(addr)
	return &Client2{BinName:name,Addr:addr, Client:client}
}

// storage.get("key")  -> value

// func hash(s string) uint32 {
//         h := fnv.New32a()
//         h.Write([]byte(s))
//         return h.Sum32()
// }

// func main() {
//         fmt.Println(hash("HelloWorld"))
//         fmt.Println(hash("HelloWorld."))
// }

