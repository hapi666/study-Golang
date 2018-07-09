package main

import (
	"strings"
	"log"
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func main()  {
	html := `<body id="1">
				<div class="name">
					<div class="g">
						<a>
							<b title="wangyiyun">laa</b>
						</a>
					</div>
				</div>
				<div>DIV2</div>
				<span id="ll" class="mm">SPAN</span>
			</body>
			`
	dom,err:=goquery.NewDocumentFromReader(strings.NewReader(html))
	if err!=nil{
		log.Fatalln(err)
	}
	dom.Find(".g a b").Each(func(i int, selection *goquery.Selection) {
		if value,ok:=selection.Attr("title");ok{
			fmt.Println(value)
		}

	})
}
