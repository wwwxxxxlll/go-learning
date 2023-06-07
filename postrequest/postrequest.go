package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	jq "github.com/savaki/jq"
)

type Response struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {
	postUrl := "http://localhost:8502/test1"
	body := []byte(`{
		"title": "Post title",
		"body": "Post description",
		"userId": 1
	}`)
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, reserr := client.Do(req)
	if reserr != nil {
		panic(reserr)
	}
	data, derr := ioutil.ReadAll(res.Body)
	if derr != nil {
		panic(derr)
	}
	defer res.Body.Close()
	op, _ := jq.Parse(".title")
	ans, _ := op.Apply(data)
	fmt.Println((string)(ans))
}
