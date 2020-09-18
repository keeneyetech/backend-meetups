package main

import (
	"fmt"
	"net/http"
	"sayhello"
	"sayhello/clock"
	"sayhello/db"
	"sayhello/ipstack"
	"strings"

	"github.com/go-redis/redis"
)

func main() {
	// connect to redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("no redis available: %+v", err))
	}

	sh := sayhello.New(db.New(rdb), ipstack.New(), clock.New())
	err = http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s, err := sh.SayHello(getIP(r))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, s)
	}))
	panic(err)
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}
