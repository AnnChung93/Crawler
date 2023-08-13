package main

import (
	"encoding/json"
	"strconv"

	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	var books []Book

	category := ""
	c.OnHTML(".mod_b.type02_l001-1 ul li.here a", func(e *colly.HTMLElement) {
		category = e.Text
	})

	c.OnHTML(".mod_a .item", func(e *colly.HTMLElement) {
		book := Book{}

		book.Category = category
		// Extracting 'no' field and converting to int
		noStr := e.ChildText("p.no_list strong.no")
		book.No, _ = strconv.Atoi(noStr)

		// Extracting other fields
		book.Title = e.ChildText("div.type02_bd-a h4 a")
		book.Author = e.ChildText("div.type02_bd-a ul li a")

		// Extracting discount percent and converting to int
		discountPercentStr := e.ChildText("li.price_a strong:nth-child(1)")
		book.Discount_percent, _ = strconv.Atoi(discountPercentStr)

		// Extracting discount price and converting to int
		discountPriceStr := e.ChildText("li.price_a strong:nth-child(2)")
		book.Discount_price, _ = strconv.Atoi(discountPriceStr)

		// Append book to the list
		books = append(books, book)
	})

	c.OnRequest(func(r *colly.Request) { // Set User-Agent
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
	})

	c.Visit("https://www.books.com.tw/web/sys_saletopb/books/") // Visit 要放最後

	json_str, _ := json.MarshalIndent(books, "", "\t")
	export_txt(string(json_str))
}

func export_txt(output_string string) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}

	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}

	i := strings.LastIndex(path, "/")

	file_name := time.Now().Format("20060102_150403")
	f, err := os.Create(string(path[0:i+1]) + "/" + file_name + ".txt") // create a file
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString(output_string) // write content to file
	if err != nil {
		panic(err)
	}
}

type Book struct {
	Category         string
	No               int
	Title            string
	Author           string
	Discount_percent int
	Discount_price   int
	Publication_date string
}
