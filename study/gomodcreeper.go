package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

//新版本爬虫
func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://pkg.go.dev/search?q=goquery")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	// class选择器
	// 查找标签为a,属性为href
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		//获取属性为href的Value
		href, exists := s.Attr("href")
		fmt.Println(href, exists)

		// For each item found, get the title
		//title := s.Find("a").Text()
		//fmt.Printf("Review %d: %s\n", i, title)
	})
	fmt.Println("---------------------------------------------------")

	//需要找共同特征,才能找到想要的属性
	//如果是找class的值,那么就需要前面加个"."
	// class选择器
	// 在SearchSnippet下获取标签为a所有的超链接
	doc.Find(".SearchSnippet").Find("a").Each(func(i int, s *goquery.Selection) {
		//获取属性为href的Value
		href, exists := s.Attr("href")
		fmt.Println(href, exists)

		// For each item found, get the title
		//title := s.Find("a").Text()
		//fmt.Printf("Review %d: %s\n", i, title)
	})
	fmt.Println("---------------------------------------------------")
	//id选择器 格式: "#idName"
	//只获取id为jump-to-modal,class的值
	fmt.Println(doc.Find("#jump-to-modal").Attr("class"))
	//获取其中html代码
	fmt.Println(doc.Find("#jump-to-modal").Html())
	//只获取其中文本
	fmt.Println(doc.Find("#jump-to-modal").Text())

	//符合选择器
	//tag+class
	// <div></div><div class="nav"></div><span class="nav"></span>
	//tag.class

	//子选择器
	//子孙选择器 格式为:"select01 select02 select03..."
	fmt.Println("---------------------子选择器示例1------------------------------")
	doc.Find(".SearchSnippet a").Each(func(i int, s *goquery.Selection) {
		//获取属性为href的Value
		href, exists := s.Attr("href")
		fmt.Println(href, exists)
	})
	//筛选parent这个父元素下，符合child这个条件的最直接（一级）的子元素。
	//注意是一级!!!
	//用不来妈的
	fmt.Println("---------------------子选择器示例2------------------------------")
	doc.Find(".SearchSnippet>.SearchSnippet-headerContainer>h2>a").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		//获取属性为href的Value
		//href, exists := s.Attr("href")
		//fmt.Println(href, exists)
	})
}

func main() {
	ExampleScrape()
}
