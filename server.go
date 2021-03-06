package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	type Tweet struct {
		Username string `json:"username"`
		Tweet    string `json:"tweet"`
	}
	tweets := []Tweet{
		{Username: "@GuyEndoreKaiser", Tweet: "If you live to be 100, you should make up some fake reason why, just to mess with people... like claim you ate a pinecone every single day."},
		{Username: "@mikeleffingwell", Tweet: "STOP TELLING ME YOUR NEWBORN'S WEIGHT AND LENGTH I DON'T KNOW WHAT TO DO WITH THAT INFORMATION."},
	}
	tweet_buf, _ := json.Marshal(tweets)

	type Status struct {
		Name   string `json:"name"`
		Status string `json:"status"`
	}
	statuses := []Status{
		{Name: "Some Friend", Status: "Here's some photos of my holiday. Look how much more fun I'm having than you are!"},
		{Name: "Drama Pig", Status: "I am in a hospital. I will not tell you anything about why I am here."},
	}
	fb_buf, _ := json.Marshal(statuses)

	type Post struct {
		Username string `json:"username"`
		Picture  string `json:"picture"`
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
		time.Sleep(time.Duration(delay) * time.Second)
		if delay >= 4 {
			w.WriteHeader(500)
			w.Write(err_buf)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(tweet_buf)
		return
	})

	http.HandleFunc("/facebook", func(w http.ResponseWriter, r *http.Request) {
		delay := rand.Intn(5)
		time.Sleep(time.Duration(delay) * time.Second)
		if delay >= 3 {
			w.WriteHeader(500)
			w.Write(err_buf)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(fb_buf)
		return
	})

	http.HandleFunc("/instagram", func(w http.ResponseWriter, r *http.Request) {
		delay := rand.Intn(5)
		time.Sleep(time.Duration(delay) * time.Second)
		if delay >= 2 {
			w.WriteHeader(500)
			w.Write(err_buf)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(insta_buf)
		return
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
