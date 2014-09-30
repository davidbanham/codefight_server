package main

import (
  "net/http"
  "log"
  "os"
  "encoding/json"
  "time"
  "math/rand"
)

func main() {
  type Tweet struct {
    Username string `json:"username"`
    Tweet string `json:"tweet"`
  }
  tweets := []Tweet{
    {Username: "@Bill_Nye_tho", Tweet: "I STRAIGHT UP HAVE NO IDEA HOW PORCUPINES FUCK EACH OTHER"},
    {Username: "@WillFerrell", Tweet: "I want to have 3 kids and name them Ctrl, Alt and Delete. Then if they fuck up I will just hit them all at once."},
    {Username: "@shitmydadsays", Tweet: "We ain't a sharp species. We kill each other over arguments about what happens when you die, then fail to see the fucking irony in that."},
  }
  tweet_buf, _ := json.Marshal(tweets)

  type Status struct {
    Name string `json:"name"`
    Status string `json:"status"`
  }
  statuses := []Status{
    {Name: "Some Friend", Status: "Here's some photos of my holiday. Look how much more fun I'm having than you are!"},
    {Name: "Drama Pig", Status: "I am in a hospital. I will not tell you anything about why I am here."},
  }
  fb_buf, _ := json.Marshal(statuses)

  type Post struct {
    Username string `json:"username"`
    Picture string `json:"picture"`
  }
  posts := []Post{
    {Username: "hipster1", Picture: "food"},
    {Username: "hipster2", Picture: "coffee"},
    {Username: "hipster3", Picture: "coffee"},
    {Username: "hipster4", Picture: "food"},
    {Username: "hipster5", Picture: "this one is of a cat"},
  }
  insta_buf, _ := json.Marshal(posts)

  err_buf := []byte("I am trapped in a social media factory send help")

  http.HandleFunc("/twitter", func(w http.ResponseWriter, r *http.Request) {
    delay := rand.Intn(5)
    time.Sleep(time.Duration(delay)*time.Second)
    if (delay >= 4) {
      w.WriteHeader(500)
      w.Write(err_buf)
      return
    }
    w.Write(tweet_buf)
    return
  })

  http.HandleFunc("/facebook", func(w http.ResponseWriter, r *http.Request) {
    delay := rand.Intn(5)
    time.Sleep(time.Duration(delay)*time.Second)
    if (delay >= 4) {
      w.WriteHeader(500)
      w.Write(err_buf)
      return
    }
    w.Write(fb_buf)
    return
  })

  http.HandleFunc("/instagram", func(w http.ResponseWriter, r *http.Request) {
    delay := rand.Intn(5)
    time.Sleep(time.Duration(delay)*time.Second)
    if (delay >= 4) {
      w.WriteHeader(500)
      w.Write(err_buf)
      return
    }
    w.Write(insta_buf)
    return
  })

  log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
