package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResolveTwitchClip(t *testing.T) {
	router := makeRouter("")
	handleLinkResolver(router)
	ts := httptest.NewServer(router)
	defer ts.Close()
	fmt.Println(ts.URL)
	const url = `https%3A%2F%2Fclips.twitch.tv%2FGorgeousAntsyPizzaSaltBae`
	res, err := http.Get(ts.URL + "/link_resolver/" + url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var jsonResponse LinkResolverResponse
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonResponse)
	// go listen("127.0.0.1:5902", router)
	//
	// 	r := httptest.NewRequest("GET", "http://Pepega/link_resolver/http%3A%2F%2Fpajlada.com", nil)
	// 	w := httptest.NewRecorder()
	//
	// 	linkResolver(w, r)
	//
	// 	resp := w.Result()
	// 	body, _ := ioutil.ReadAll(resp.Body)
	//
	// 	fmt.Println(string(body))
	//
	// 	fmt.Println("lol")
}

func TestResolve1M(t *testing.T) {
	// var resp *http.Response
	// var err error

	// fmt.Println("test resolve 1M")
	// resp, err = makeRequest("http://speedtest.tele2.net/100MB.zip")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// resp, err = makeRequest("http://httpbin.org/redirect/5")
	// fmt.Println(resp)
	// fmt.Println(resp.Request.URL)
	// resp, err = makeRequest("http://httpbin.org/image")
	// fmt.Println(resp)
	// resp, err = makeRequest("http://speedtest.tele2.net/1MB.zip")
	// fmt.Println(resp)
}

func TestDoRequest(t *testing.T) {
	// var err error

	// fmt.Println("test resolve 1M")
	// _, err = doRequest("http://speedtest.tele2.net/100MB.zip")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
}
