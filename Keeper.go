package triblab

import (
	// "net/rpc"
	"trib"
	"time"
	"math"
)

type Keeper struct{
	kc trib.KeeperConfig
	cls []trib.Storage
}

func (self *Keeper)Run() error {
	//init clients
	self.cls = make([]trib.Storage,len(self.kc.Backs))
	//check connectivity
	// for _,cl := range self.cls{
	// 	e := cl.Clock()
	// 	if e!=nil{
	// 		return e;
	// 	}
	// }
	
	//send Ready chan
	//heartbeat
	tick := time.Tick(500 * time.Millisecond)
	maxclock := uint64(64)
	for{
		select{
			case <-tick:
				tchan := make(chan uint64, len(self.cls))
					
				//fire clock query to each back, best concurrently
				for _, cl := range self.cls{
					go func(cl trib.Storage){
						var t uint64
						cl.Clock(maxclock, &t)
						tchan <- t
						}(cl)
				}
				go func(){
					for i := 0; i < len(self.kc.Backs); i++ {
						maxclock = uint64(math.Max(float64(maxclock), float64(<-tchan)))
					}
				}()
			default:
				time.Sleep(100*time.Millisecond)
		}
	}
}