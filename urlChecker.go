// package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var errRequestFailed = errors.New("Fail to Rquest url")

// func main() {
// 	var results = make(map[string]string)
// 	urls := []string{
// 		"https://www.airbnb.com/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://www.reddit.com/",
// 		"https://soundcloud.com/",
// 		"https://www.instagram.com/",
// 		"https://academy.nomadcoders.co/",
// 	}

// 	for _, url := range urls {
// 		result := "OK"
// 		err := hitURL(url)

// 		if err != nil {
// 			result = "FAIL"
// 		}
// 		results[url] = result
// 	}

// 	for url, result := range results {
// 		fmt.Println(url, result)
// 	}

// }

// func hitURL(url string) error {
// 	fmt.Println("Check Url : ", url)
// 	resp, err := http.Get(url)

// 	if err != nil || resp.StatusCode >= 400 {
// 		fmt.Println(err, resp.StatusCode)
// 		return errRequestFailed
// 	}

// 	return nil
// }
