package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "net/http"
)

type Post struct {
  Id:     `json:"id"`
  Title:  `json:"title"`
  Body:   `json:"body"`
  UserId: `json:"userId"`
}

func main() {

  newPost := []byte(`{
    "userId": "john",
    "title": "hi",
    "body": "hello world"
  }`)

  postUrl := "http://json.com/posts"

  client := &http.Client{}

  r, err := http.NewHttpRequest("POST", postUrl, bytes.NewBuffer(newPost) 
  if err != nil {
    panic(err)
  }

  r.Headers.Add("Content-Type", "application/json")

  res, err := client.Do(r)
  if err != nil {
    panic(err)
  }

  defer res.Body.Close()

  if res.StatusCode != http.StatusCreated {
    panic(res.StatusCode)
  }

  post := &Post{}
  err1 := json.NewDecoder(res.Body).Decode(post)
  if err1 != nil {
    panic(err1)
  }

	fmt.Println("Id:", post.Id)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
	fmt.Println("UserId:", post.UserId)
}





