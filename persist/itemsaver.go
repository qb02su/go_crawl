package persist

import "log"

func ItemSave()chan interface{}{
	out:=make(chan interface{})
	go func() {
		itemCount:=0
		for{
			item:=<-out
			log.Printf("Item Save Got %d Item %v",itemCount,item)
			itemCount++
		}
	}()
	return out
}
