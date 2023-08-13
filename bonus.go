package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"strconv"

// 	"os"
// 	"os/exec"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	browser "github.com/EDDYCJY/fake-useragent"
// 	"github.com/gocolly/colly"
// )

// func main() {
// 	c := colly.NewCollector()

// 	var books []Book

// 	category := ""
// 	c.OnHTML(".mod_b.type02_l001-1 ul li.here a", func(e *colly.HTMLElement) {
// 		category = e.Text
// 		// category = e.Attr("href")
// 	})
// 	c.OnHTML(".mod_a .item", func(e *colly.HTMLElement) {
// 		book := Book{}

// 		book.Category = category
// 		// Extracting 'no' field and converting to int
// 		noStr := e.ChildText("p.no_list strong.no")
// 		book.No, _ = strconv.Atoi(noStr)

// 		// Extracting other fields
// 		book.Title = e.ChildText("div.type02_bd-a h4 a")
// 		book.Author = e.ChildText("div.type02_bd-a ul li a")

// 		// Extracting discount percent and converting to int
// 		discountPercentStr := e.ChildText("li.price_a strong:nth-child(1)")
// 		book.Discount_percent, _ = strconv.Atoi(discountPercentStr)

// 		// Extracting discount price and converting to int
// 		discountPriceStr := e.ChildText("li.price_a strong:nth-child(2)")
// 		book.Discount_price, _ = strconv.Atoi(discountPriceStr)

// 		// 嘗試做出bonus
// 		book.Sub_url = e.ChildAttr("div.type02_bd-a h4 a", "href")

// 		// Append book to the list
// 		books = append(books, book)
// 	})

// 	c.OnRequest(func(r *colly.Request) { // Set User-Agent
// 		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
// 	})

// 	c.Visit("https://www.books.com.tw/web/sys_saletopb/books/") // Visit 要放最後

// 	for i, book := range books {
// 		// Visit the product URL to fetch its HTML content
// 		bonus_c := colly.NewCollector()

// 		bonus_c.OnHTML(".type02_p003 ul li", func(bonus_e *colly.HTMLElement) {
// 			// Check if the li element contains "出版日期"
// 			if strings.Contains(bonus_e.Text, "出版日期：") {
// 				publishDate := strings.TrimSpace(strings.TrimPrefix(bonus_e.Text, "出版日期："))
// 				fmt.Println("publishDate", publishDate)
// 				books[i].Publication_date = publishDate
// 			}
// 		})

// 		bonus_c.OnError(func(r *colly.Response, err error) {
// 			fmt.Println("error:", r.StatusCode, err)
// 		})

// 		// random User-Agent
// 		client := browser.Client{
// 			MaxPage: 3,
// 			Delay:   200 * time.Millisecond,
// 			Timeout: 10 * time.Second,
// 		}
// 		cache := browser.Cache{}
// 		b := browser.NewBrowser(client, cache)

// 		random := b.Random()
// 		fmt.Println(">>> random browser", random)

// 		bonus_c.OnRequest(func(r *colly.Request) { // Set User-Agent
// 			r.Headers.Set("User-Agent", random)
// 		})

// 		// Visit the bonus URL
// 		err := bonus_c.Visit(book.Sub_url)
// 		if err != nil {
// 			fmt.Println("Error visiting bonus URL:", err)
// 		}

// 		time.Sleep(1000 * time.Millisecond)
// 	}

// 	json_str, _ := json.MarshalIndent(books, "", "\t")
// 	export_txt(string(json_str))
// }

// func export_txt(output_string string) {
// 	file, err := exec.LookPath(os.Args[0])
// 	if err != nil {
// 		panic(err)
// 	}

// 	path, err := filepath.Abs(file)
// 	if err != nil {
// 		panic(err)
// 	}

// 	i := strings.LastIndex(path, "/")

// 	file_name := time.Now().Format("20060102_150403")
// 	f, err := os.Create(string(path[0:i+1]) + "/" + file_name + ".txt") // create a file
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	_, err = f.WriteString(output_string) // write content to file
// 	if err != nil {
// 		panic(err)
// 	}
// }

// type Book struct {
// 	Category         string
// 	No               int
// 	Title            string
// 	Author           string
// 	Discount_percent int
// 	Discount_price   int
// 	Sub_url          string
// 	Publication_date string
// }
