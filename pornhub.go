package pornhub

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// apiURL base for API
const apiURL = "http://www.pornhub.com/webmasters/"

// APITimeout in seconds
const APITimeout = 10

// SearchVideosByTerm function
func SearchVideosByTerm(search string, page int) SearchResult {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	// http://www.pornhub.com/webmasters/search?search=Hooker&page=1
	resp, _ := client.Get(fmt.Sprintf(apiURL+"search?search=%s&page=%d", url.QueryEscape(search), page))
	b, _ := ioutil.ReadAll(resp.Body)
	var result SearchResult
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// SearchVideosByCategory function
func SearchVideosByCategory(search string, page int) SearchResult {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	// https://www.pornhub.com/webmasters/search?category=2d&page=1
	resp, _ := client.Get(fmt.Sprintf(apiURL+"search?category=%s&page=%d", url.QueryEscape(search), page))
	b, _ := ioutil.ReadAll(resp.Body)
	var result SearchResult
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// ListCategories function
func ListCategories() CategoriesList {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL + "categories"))
	b, _ := ioutil.ReadAll(resp.Body)
	var result CategoriesList
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// ListStars function
func ListStars() StarsList {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL + "stars"))
	b, _ := ioutil.ReadAll(resp.Body)
	var result StarsList
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// NewestVideos function
func NewestVideos(page int) SearchResult {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"search?ordering=newest&period=daily&page=%d", page))
	b, _ := ioutil.ReadAll(resp.Body)
	var result SearchResult
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}

// GetVideoByID function
func GetVideoByID(ID string) SingleVideo {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"video_by_id?id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result SingleVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result

}

// GetVideoEmbedCode function
func GetVideoEmbedCode(ID string) EmbedCode {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"video_embed_code?id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result EmbedCode
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}
