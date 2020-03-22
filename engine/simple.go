package engine

import (
	"crawl/fetcher"
	"log"
)
type SimpleEngine struct {

}


func (e *SimpleEngine)Run(seeds...Request){
	var requests []Request
	for _,req:=range seeds{
		requests=append(requests,req)
	}
	for len(requests)>0{
		r:=requests[0]
		requests=requests[1:]
		/*log.Printf("fetching:%s",r.Url)
		body,err:=fetcher.Fetch(r.Url)
		if err!=nil{
			log.Printf("fetch error when fetching url:%s%v",r.Url,err)
			continue
		}W
		ParseResult:= r.ParseFunc(body)*/
		ParseResult,err:=Worker(r)
		if err!=nil{
			continue
		}
		requests=append(requests,ParseResult.Request...)
		for _,item:=range ParseResult.Item{
			log.Printf("got item %v",item)
		}
	}
}
func Worker(r Request)(ParseResult,error){
	log.Printf("Fetching %s",r.Url)
	body,err:=fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("fetch error when fetching url:%s%v",r.Url,err)
		return ParseResult{},err
	}
	return r.ParseFunc(body),nil
}