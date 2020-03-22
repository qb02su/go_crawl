package parser

import (
	"crawl/engine"
	"regexp"
)
var(
 cityListRe=regexp.MustCompile(`<a href="(https://movie.douban.com/subject/[0-9]+/)" class="">\s+<span class="title">([^<]+)</span>`)
 nextpageRe=regexp.MustCompile(`<a href="([^"]+)"\s>[^&]+&gt;`)
)
func ParseCityList(content []byte)engine.ParseResult{
	//re:=regexp.MustCompile(cityListRe)
	all:=cityListRe.FindAllSubmatch(content,-1)
	result:=engine.ParseResult{}
	for _,c:=range all{
		result.Item=append(result.Item,string(c[2]))
		result.Request=append(result.Request,engine.Request{
			Url:       string(c[1]),
			ParseFunc: ParseMovie,
		})

	}
	all2:=nextpageRe.FindAllSubmatch(content,-1)
	for _,c:=range all2{
		result.Request=append(result.Request,engine.Request{
			Url:"https://movie.douban.com/top250"+string(c[1]),
			ParseFunc:ParseCityList,

		})
	}
	return result
}


