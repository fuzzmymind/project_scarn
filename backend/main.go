package main

import (
	b64 "encoding/base64"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func epochToHuman(s string) (time.Time, error) {
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
}

func main() {
	home := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w,
			`A collection of tools that are kinda helpful you know!`)
	}

	epochConverter := func(w http.ResponseWriter, req *http.Request) {
		param := req.URL.Query().Get("param")
		result, err := epochToHuman(param)
		if err != nil {
			io.WriteString(w, "Failed to convert")
		} else {
			io.WriteString(w, result.String())
		}
	}

	base64encode := func(w http.ResponseWriter, req *http.Request) {
		param := req.URL.Query().Get("param")
		result := b64.StdEncoding.EncodeToString([]byte(param))
		io.WriteString(w, result)
	}

	base64decode := func(w http.ResponseWriter, req *http.Request) {
		param := req.URL.Query().Get("param")
		result, err := b64.StdEncoding.DecodeString(param)
		if err != nil {
			io.WriteString(w, "Failed to decode")
		} else {
			io.WriteString(w, string(result))
		}
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/epoch", epochConverter)
	http.HandleFunc("/base64encode", base64encode)
	http.HandleFunc("/base64decode", base64decode)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
