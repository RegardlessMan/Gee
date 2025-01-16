/**
 * @Author QG
 * @Date  2024/12/30 22:45
 * @description
**/

package main

import (
	gee_cache "Gee/gee-cache"
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	http.HandleFunc("/first", index)
	http.HandleFunc("/second", second)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func second(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(resp, "URL Path is %s", req.URL.Path)
}

func index(resp http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(resp, "Header[%v] == %v\n", k, v)
	}
}

func TestCache() {
	gee_cache.NewGroup("scores", 2<<10, gee_cache.GetterFunc(func(key string) ([]byte, error) {
		log.Println("[SlowDB] search key", key)
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, fmt.Errorf("%s not exist", key)
	}))

	addr := "localhost:9999"
	pool := gee_cache.NewHTTPPool(addr)
	log.Println("gee cache is running at", addr)
	http.ListenAndServe(addr, pool)
}
