package main
import (
	"fmt"
	"net/http"
	
)

func checkStatus(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status for %s: %d\n", url, resp.StatusCode)
}


func main() {

	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://anshdev-tau.vercel.app",
	}

	for _, url := range urls {
		checkStatus(url)
	}

}

