package main

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/fatih/color"
	"sort"
)

//cmd版本 根据中文查找日语意思
//qt版本 根据中文或者截图查找日语意思
const qurl = "https://dict.hjenglish.com/jp/jc/%s" //后面+单词


type meaning struct {
	hiragana string
	alphabet string
	means map[string][]string
}
type QueryWord struct {
	name string
	means []meaning
}

func (q *QueryWord)Show() {
	color.Blue("%s",q.name)
	for _, v := range q.means {
		color.Green("%s %s",v.hiragana,v.alphabet)
		 var keys []string
		for k := range v.means {
			keys = append(keys,k)
		}
		sort.Strings(keys)
		for _, value := range keys {
			color.Yellow("%s ",value)
			for _, mean := range v.means[value] {
				color.Red("%s",mean)
			}
		}
	}
}
func getDocument(word string)*goquery.Document{
	resp, err := http.Get(fmt.Sprintf(qurl, word))
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	document, e := goquery.NewDocumentFromReader(resp.Body)
	if e != nil {
		log.Fatal(e)
		return nil
	}
	return document
}

func getResult(document *goquery.Document)*QueryWord{
	contain := document.Find(".word-details-pane-header")
	if contain.Length() == 0 {
		suggstion := document.Find(".word-suggestions")
		if suggstion.Length() != 0 {
			fmt.Println("查找结果不存在,你想查的是不是: ")
			var means []string
			suggstion.Find("ul li").Each(func(i int, item *goquery.Selection) {
				means = append(means,item.Text())
			})
			fmt.Printf("%s\n",strings.Join(means, " "))
		}
		return nil
	}

	wordSec := contain.Find(".word-text h2")
	means := make( []meaning,wordSec.Length())
	name := wordSec.First().Text()
	 contain.Find(".pronounces>span:nth-child(1)").Each(func(i int, item *goquery.Selection) {
		 means[i].hiragana = item.Text()
	 })
	contain.Find(".pronounces>span:nth-child(2)").Each(func(i int, item *goquery.Selection) {
		means[i].alphabet = item.Text()
	})


	document.Find(".simple").Each(func(i int, item *goquery.Selection) {
		mean := make(map[string][]string)
		selection := item.Find("h2")
		part := mean[selection.Text()]
		selection.Siblings().ChildrenFiltered("li").Each(func(i int, selection *goquery.Selection) {
			part = append(part,selection.Text())
		})
		mean[selection.Text()] = part
		means[i].means = mean
	})

	return &QueryWord{
		name:name,
		means:means,
	}
}
