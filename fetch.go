// go prog to fetch a file from a given URL

package main

// import needed modules
import (
    "fmt"
    "os"
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
)

// checkErr - a function to check the error return value of another function.
func checkErr(err error) {
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(2)
    }
}

// getContent - takes an url as argument and returns the body of the document
// from that url
func getContent(url string) (b []byte) {
    resp, err := http.Get(url)
    checkErr(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    checkErr(err)
    return body
}

// writeFile - takes an array of bytes, and a filename as argument, writes 
// those bytes into a file with the supplied name (filename), returns the number of bytes written
func writeFile(b []byte, filename string) (i int) {
    f, err := os.Create(filename)
    checkErr(err)
    defer f.Close()
    i, err = f.Write(b)
    checkErr(err)
    return i
}

// filenameFromPath - takes a path (as a string) as argument,
// and returns the filename from that path
func filenameFromPath(path string) (filename string) {
    pathElements := strings.Split(path, "/")
    numOfElements := len(pathElements)
    filenameElement := pathElements[numOfElements-1:numOfElements]
    filename = filenameElement[0]
    return filename
}
// fetchFromUrl - takes an url as argument (as string),
// retrieves the body (by calling getContent) from that url,
// parses the URL, extracts the filename from the parsed URL 
// by calling filenameFromPath,
// writes the body to a file with the extracted filename,
// and prints howmany bytes have been written to the file (with the name 
// extracted previous) to standard output.
func fetchFromUrl(uri string) (string) {
	l, err := url.Parse(uri)
    checkErr(err)
	var filename string
	if filenameFromPath(l.Path) == "" {
		filename = "outfile"
	} else {
		filename = filenameFromPath(l.Path)
	}
	body := getContent(uri)
	i := writeFile(body, filename)
	fmt.Printf("%d bytes written to '%s'\n", i, filename)
	return filename
}

// gets the arguments 'fetch' was called with,
// iterates over arguments (urls) and calls fetchFromUrl for each of it.
func main() {
    urlList := os.Args[1:]
    for u := range urlList {
        uri := urlList[u]
	fetchFromUrl(uri)
    }
}
