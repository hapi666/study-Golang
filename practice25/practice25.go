package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

type Qus struct {
	Question string
	Answer   []string
}

var (
	err       error
	questions = make([]Qus, 0)
	question  Qus
)

func Decode1(src string) (dst string) {
	data, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(src)), simplifiedchinese.GBK.NewDecoder()))
	if err == nil {
		dst = string(data)
	}
	return
}

func main() {
	ch := make(chan string)
	el := make(chan string)
	qs := make([]string, 0)
	as := make([]string, 0)

	go func() {
		ch <- os.Args[1]
	}()
	go func(ch chan string) {
		var Url string
		Url = <-ch
		doc, err := goquery.NewDocument(Url)
		if err != nil {
			log.Fatal(err.Error())
		}

		doc.Find("td .style1 strong").Each(func(i int, s *goquery.Selection) {
			qs = append(qs, Decode1(strings.TrimSpace(s.Text())))
		})
		doc.Find("td .style1 .green").Each(func(i int, s *goquery.Selection) {
			as = append(as, Decode1(strings.TrimSpace(s.Text())))
		})
		for i, q := range qs {
			question.Question = q
			//与question:=Qus{Question:q}等效
			question.Answer = make([]string, 0)
			for j := i * 3; j < i*3+3; j++ {
				question.Answer = append(question.Answer, as[j])
			}

			questions = append(questions, question)
		}
		el <- "starting..."
	}(ch)
	fmt.Println(<-el)
	fmt.Printf("%v\n", questions)
}
