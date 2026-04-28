package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("h1", func(e *colly.HTMLElement){
		fmt.Println("Заголовок: ", e.Text)
	})

	c.OnHTML("article.product_pod", func(e *colly.HTMLElement){
		title := e.ChildText("h3 a");
		price := e.ChildText(".price_color")

		fmt.Println("Название: ", title)
		fmt.Println("Цена: ", price)
		fmt.Print("\n")
	})

	c.Visit("http://books.toscrape.com/")
}