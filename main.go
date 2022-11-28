package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

type result struct {
	url   string
	count int
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://youtube.com",
		"https://twitter.com",
		"https://instagram.com",
		"https://baidu.com",
		"https://wikipedia.org",
		"https://yandex.ru",
		"https://yahoo.com",
		"https://whatsapp.com",
		"https://amazon.com",
		"https://netflix.com",
		"https://live.com",
		"https://pornhub.com",
		"https://reddit.com",
		"https://office.com",
		"https://xhamster.com",
		"https://vk.com",
		"https://linkedin.com",
		"https://discord.com",
		"https://tiktok.com",
		"https://naver.com",
		"https://twitch.tv",
		"https://pinterest.com",
		"https://bing.com",
		"https://roblox.com",
		"https://qq.com",
		"https://duckduckgo.com",
		"https://samsung.com",
		"https://globo.com",
		"https://msn.com",
		"https://microsoft.com",
		"https://ebay.com",
		"https://accuweather.com",
		"https://weather.com",
		"https://bongacams.com",
		"https://indeed.com",
		"https://cnn.com",
		"https://paypal.com",
		"https://quora.com",
	}
	var wg sync.WaitGroup
	ch := make(chan result, 5)
	ch2 := make(chan int)
	go func() {
		total := 0
		for res := range ch {
			total += res.count
			fmt.Printf("Count for %v: %v\n", res.url, res.count)
		}
		ch2 <- total
	}()

	for _, url := range urls {
		wg.Add(1)
		go worker(url, ch, &wg)
	}
	wg.Wait()
	close(ch)
	fmt.Println("Total: ", <-ch2)
}

func worker(url string, ch chan result, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		return
	}
	count := strings.Count(string(respData), "Go")
	ch <- result{
		url:   url,
		count: count,
	}
}