package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	// задаем флаги для url и о (скачать весь сайт целиком)
	url := flag.String("url", "", "url")
	entireSite := flag.Bool("o", false, "download entire site")

	flag.Parse()
	// узнаем название конечной страницы
	pathArr := strings.Split(*url, "/")
	outputPath := "testSite/" + pathArr[len(pathArr)-1]
	fmt.Println(outputPath)
	// Логика на выбор - скачиваем файл или сайт целиком
	if !*entireSite {
		err := downloadFile(outputPath, *url)
		if err != nil {
			log.Fatalln(err)
		}
		return
	}

	// сохраняем все страницы с сайта
	DownloadAllFiles(*url)
}

//LinkParser получаем все ссылки с сайта
func LinkParser(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var links []string

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		links = append(links, link)
	})
	return links
}

//
//// DownloadAllFiles записываем все в html файлы
func DownloadAllFiles(url string) {
	links := LinkParser(url)

	for _, l := range links {
		fmt.Println(l)
		pathArr := strings.Split(l, "/")
		outputPath := "testSite/" + pathArr[len(pathArr)-1]
		if len(pathArr) > 2 {
			continue
		}
	
		resp, err := http.Get(url + l)
		if err != nil {
			fmt.Println("failed")

		}
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()
		
		f, err := os.Create(outputPath)
		if err != nil {
			fmt.Println("creating file failed")
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}()

		_, err = io.Copy(f, resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
