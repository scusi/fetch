// go prog to fetch a file from a given URL

package main

import (
    "fmt"
    "os"
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
)

func checkErr(err error) {
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(2)
    }
}

func getContent(url string) (b []byte) {
    resp, err := http.Get(url)
    checkErr(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    checkErr(err)
    return body
}

func writeFile(b []byte, filename string) (i int) {
    f, err := os.Create(filename)
    checkErr(err)
    defer f.Close()
    i, err = f.Write(b)
    checkErr(err)
    return i
}

func filenameFromPath(path string) (filename string) {
    pathElements := strings.Split(path, "/")
    numOfElements := len(pathElements)
    filenameElement := pathElements[numOfElements-1:numOfElements]
    filename = filenameElement[0]
    return filename
}

func fetchFromUrl(uri string) (filename string) {
	body := getContent(uri)
	l, err := url.Parse(uri)
        checkErr(err)
	filename = filenameFromPath(l.Path)
	i := writeFile(body, filename)
	fmt.Printf("%d bytes written to '%s'\n", i, filename)
	return
}

func main() {
    urlList := os.Args[1:]
    for u := range urlList {
        uri := urlList[u]
	fetchFromUrl(uri)
    }
}
