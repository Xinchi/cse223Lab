package triblab

import (
	"trib"
)

func NewBinClient(backs []string) trib.BinStorage {
	return &helper{Backs:backs}
}

func ServeKeeper(kc *trib.KeeperConfig) error {
	
	cls := make([]trib.Storage, len(kc.Backs))
	for index:= range kc.Backs{
		cls[index] = NewClient(kc.Backs[index])
	}
	//set true
	kc.Ready <- true

	// keeper := Keeper()
	// return keeper.Run()
	return nil
}

func NewFront(s trib.BinStorage) trib.Server {
	panic("todo")
}

