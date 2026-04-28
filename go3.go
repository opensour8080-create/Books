package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {

	var temp string

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36"),
	)

	c.OnHTML("h1", func(e *colly.HTMLElement){
		location := e.Text
		fmt.Println(location)
	})

	c.OnHTML("#weather-now-number", func(e *colly.HTMLElement) {
		temp = e.Text

		
	})

	c.Visit("https://world-weather.ru/pogoda/russia/saratov/")

	fmt.Println("Температура:", temp)
}