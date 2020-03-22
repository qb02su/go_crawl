package main

import (
	"crawl/engine"
	"crawl/parser"
	"crawl/persist"
	"crawl/scheduler"
)

func main() {
	url:="https://movie.douban.com/top250"
	e:=engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkerCount:10,
		ItemChan:persist.ItemSave(),
	}
	e.Run(engine.Request{
		Url:url,
		ParseFunc:parser.ParseCityList,
	})

}


