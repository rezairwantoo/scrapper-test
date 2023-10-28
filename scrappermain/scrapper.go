package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"reza/scrapper-test/config"
	"reza/scrapper-test/model"
	"reza/scrapper-test/usecase"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	// initializing the slice of structs that will contain the scraped data
	var products []model.CreateRequest

	// the first pagination URL to scrape
	pageToScrape := "https://www.tokopedia.com/p/handphone-tablet/handphone?page=1"

	// initializing a Colly instance
	c := colly.NewCollector()
	// setting a valid User-Agent header
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"
	// scraping the product data
	c.OnHTML("[data-testid=lstCL2ProductList]", func(e *colly.HTMLElement) {
		product := model.CreateRequest{}
		linkDesc := "for detail please visit " + e.ChildAttr("a[data-testid=lnkProductContainer]", "href")
		e.ForEach("[data-testid=divProductWrapper]", func(i int, h *colly.HTMLElement) {
			h.DOM.Each(func(j int, s *goquery.Selection) {
				image, _ := s.First().Find("img").Attr("src")
				test := s.Last().Find("span")
				test.Each(func(k int, l *goquery.Selection) {
					switch k {
					case 0:
						product.Name = l.Text()

					case 1:
						product.Price = l.Text()

					case 3:
						product.MerchantName = l.Text()

					case 4:
						product.ImageLink = image
						product.Description = linkDesc
						products = append(products, product)
						product = model.CreateRequest{}
					}
				})
			})
		})
		fmt.Printf("%+v", products)
	})

	// visiting the first page
	c.Visit(pageToScrape)

	ctx := context.Background()
	config.LoadConfigFile(ctx)
	settings, err := config.NewSettings(ctx)
	if err != nil {
		errWrap := errors.New("initialize settings, err: " + err.Error())
		log.Fatalln("initialize settings error", errWrap)
	}

	settings.Load(settings.SetPostgresRepo(settings))
	usecaseProducts := usecase.NewProductUsecase(settings.PostgresSQLProvider)

	usecaseProducts.CreateScrapper(ctx, products)
}
