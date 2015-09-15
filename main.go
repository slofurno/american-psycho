package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type americanResponse struct {
	Start      int
	Count      int
	TotalLines int
	Error      string
	Lines      []string
}

func parseQuerystring(count string, n string) (int, int, error) {

	sentenceNumber := -1
	sentenceCount := -1
	var err error

	if count != "" {
		sentenceCount, err = strconv.Atoi(count)

		if err != nil {
			return -1, -1, err
		}
	} else {
		sentenceCount = 1
	}

	if n != "" {
		sentenceNumber, err = strconv.Atoi(n)

		if err != nil {
			return -1, -1, err
		}
	} else {

		max := len(lines) - sentenceCount

		sentenceNumber = rand.Intn(max + 1)
	}

	return sentenceNumber, sentenceCount, nil

}

func handleRequest(w http.ResponseWriter, req *http.Request) {

	qs := req.URL.Query()
	count := qs.Get("count")
	start := qs.Get("n")

	n, c, err := parseQuerystring(count, start)

	r := &americanResponse{}

	if err != nil {
		r.Error = err.Error()
	}

	r.Lines = lines[n : n+c]
	r.Count = c
	r.Start = n
	r.TotalLines = len(lines)

	j, err := json.Marshal(r)

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json ; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(j)

}

var lines []string

func init() {
	lines = []string{}
}

func main() {

	file, err := os.Open("./ap.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(len(lines))

	http.HandleFunc("/api", handleRequest)
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.ListenAndServe(":912", nil)
}
