package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/unicode"
)

func Fetch(url string)([]byte,error){
	/*resp,err:=http.Get(url)
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()*/
	request, _:=http.NewRequest(http.MethodGet,url,nil)
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	resp,_:=http.DefaultClient.Do(request)
	defer resp.Body.Close()



	if resp.StatusCode!=http.StatusOK{
		return nil,fmt.Errorf("error:statue code:%d",resp.StatusCode)
	}
	bodyRead:=bufio.NewReader(resp.Body)
	r1:=determineEncoding(bodyRead)
	utf8reader:=transform.NewReader(bodyRead,r1.NewDecoder())
	return ioutil.ReadAll(utf8reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Ftcher error:%v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}