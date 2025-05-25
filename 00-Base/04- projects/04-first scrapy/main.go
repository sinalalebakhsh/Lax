package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/gocolly/colly"
)

// product structure to keep the scraped data
type Product struct {
    Url, Image, Name, Price string
}

func main() {

    // instantiate a new collector object
    c := colly.NewCollector(
        colly.AllowedDomains("www.scrapingcourse.com"),
    )

    // initialize the slice of structs that will contain the scraped data
    var products []Product

    // define a sync to filter visited URLs
    var visitedUrls sync.Map

    // OnHTML callback for scraping product information
    c.OnHTML("li.product", func(e *colly.HTMLElement) {

        // initialize a new Product instance
        product := Product{}

        // scrape the target data
        product.Url = e.ChildAttr("a", "href")
        product.Image = e.ChildAttr("img", "src")
        product.Name = e.ChildText(".product-name")
        product.Price = e.ChildText(".price")

        // add the product instance with scraped data to the list of products
        products = append(products, product)
    })

    // OnHTML callback for handling pagination
    c.OnHTML("a.next", func(e *colly.HTMLElement) {

        // extract the next page URL from the next button
        nextPage := e.Attr("href")

        // check if the nextPage URL has been visited
        if _, found := visitedUrls.Load(nextPage); !found {
            fmt.Println("scraping:", nextPage)
            // mark the URL as visited
            visitedUrls.Store(nextPage, struct{}{})
            // visit the next page
            e.Request.Visit(nextPage)
        }
    })

    // store the data to a CSV after extraction
    c.OnScraped(func(r *colly.Response) {

        // open the CSV file
        file, err := os.Create("products.csv")
        if err != nil {
            log.Fatalln("Failed to create output CSV file", err)
        }
        defer file.Close()

        // initialize a file writer
        writer := csv.NewWriter(file)

        // write the CSV headers
        headers := []string{
            "Url",
            "Image",
            "Name",
            "Price",
        }
        writer.Write(headers)

        // write each product as a CSV row
        for _, product := range products {
            // convert a Product to an array of strings
            record := []string{
                product.Url,
                product.Image,
                product.Name,
                product.Price,
            }

            // add a CSV record to the output file
            writer.Write(record)
        }
        writer.Flush()
    })

    // open the target URL
    c.Visit("https://www.scrapingcourse.com/ecommerce")
}
