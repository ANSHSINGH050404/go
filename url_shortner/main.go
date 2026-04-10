package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID string  `json:"id"`
	OriginalURL string  `json:"original_url"`
	ShortURL string  `json:"short_url"`
	CreationDate time.Time  `json:"creation_date"`
}


var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL)) // it conerts the string to a byte slice and writes it to the hasher
	hash := hasher.Sum(nil)
	data := hex.EncodeToString(hash)
	shortURL := data[:9]
    fmt.Println("shorturl",shortURL)
	return shortURL
	
}

func createURL(OriginalURL string) URL {
	shortURL := generateShortURL(OriginalURL)
	id := shortURL

	urlDB[id] = URL{
		ID: id,
		OriginalURL: OriginalURL,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}
	return urlDB[id]

}

func getURL(id string) (URL, error) {
   url,ok := urlDB[id]
   if !ok {
   	return URL{}, fmt.Errorf("URL not found")
   }

   return url, nil

}

func main() {
	fmt.Println("Starting URL Shortner")

	originalURL := "https://x.com/AnshSin18258375"
	url := createURL(originalURL)
	fmt.Printf("Original URL: %s\nShort URL: %s\n", url.OriginalURL, url.ShortURL)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}

}