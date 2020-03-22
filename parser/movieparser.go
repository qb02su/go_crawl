package parser

import (
	"crawl/engine"
	"crawl/model"
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`<script\stype="application/ld\+json">([^<]+)`)

func ParseMovie(contents []byte) engine.ParseResult {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		json := match[1]
		film:=ParseJson(json)
		fmt.Printf("json : %s\n",film)

	}
	return engine.ParseResult{}
}

func ParseJson(json []byte)model.Film{
	re:=regexp.MustCompile(`"name":[^,]+,`)
	name:=re.FindSubmatch(json)
	re2:=regexp.MustCompile(`"datePublished":[^,]+,`)
	datepulished:=re2.FindSubmatch(json)
	re3:=regexp.MustCompile(`"ratingValue":[^}]+}`)
	rating:=re3.FindSubmatch(json)
	var film model.Film
	film.Name=string(name[0])
	film.Datepublished=string(datepulished[0])
	film.Ratingvalue=string(rating[0])
	return film
}