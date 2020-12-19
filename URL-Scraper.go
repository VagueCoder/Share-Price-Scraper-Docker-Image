package main

import (
	"fmt"
    "log"
	"time"
	"path"
	"regexp"
	"strings"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func get_URLs() []string {
	client := &http.Client{
        Timeout: 30 * time.Second,
    }

	url := "https://www.moneycontrol.com/india/stockpricequote"
    response, err := client.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    document, err := goquery.NewDocumentFromReader(response.Body)
    if err != nil {
        log.Fatal("Error loading HTTP response body. ", err)
    }

	var hrefs []string
	document.Find("div.alph_pagn").Each(func(c_index int, container *goquery.Selection) {
		container.Find("a").Each(func(e_index1 int, element *goquery.Selection) {
			href, exists := element.Attr("href")
			if exists {
				hrefs = append(hrefs, "https://www.moneycontrol.com"+href)
			}
		})
	})
	hrefs = hrefs[1:] // Pop first element as site returns base URL

	return page_urls(hrefs)
}

func page_urls(urls []string) []string {
	client := &http.Client{
        Timeout: 30 * time.Second,
    }

	var result []string
	var output string
	common := "https://priceapi.moneycontrol.com/pricefeed/bse/equitycash/"

	for i, url := range urls {
		output = fmt.Sprintf("Page Checking Progress: (%d/%d)", i+1, len(urls))
		fmt.Printf(output)
		response, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body. ", err)
		}

		document.Find("a.bl_12").Each(func(index int, element *goquery.Selection) {
				href, exists := element.Attr("href")
				if exists {
					temp := path.Base(href)
					match, _:= regexp.Match("[A-Z]+[0-9]*", []byte(temp))
					if match {
						result = append(result, common + temp)
					}
				}
		})
		fmt.Printf("\r%s\r", strings.Repeat(" ", len(output)))
	}

	return result
}