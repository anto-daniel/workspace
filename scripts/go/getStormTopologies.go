package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
  "reflect"
)

func main() {
	// read api json result from uri
	// parse json result into map
	// print map
  uri := "http://127.0.0.1:5000/home/5/6"

	resp, err := http.Get(uri)
	if err != nil {
		fmt.Println("Error while fetching data from uri", err)
		return
	}
	defer resp.Body.Close()
	// read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading response body", err)
		return
	}
	fmt.Println(reflect.TypeOf(body))
	fmt.Println(string(body))
	// parse json result into map
	var mymap map[string]int
	json.Unmarshal(body, &mymap)
	fmt.Println("Print Type")
	fmt.Println("================")
	fmt.Println(reflect.TypeOf(mymap))
	fmt.Println("================")
	for _, ele := range mymap {
		fmt.Println(ele)
	}
}
