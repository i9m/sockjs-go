package main

import (
	"log"
	"net/http"

	"github.com/i9m/sockjs-go/v3/sockjs"
)

func main() {
	http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, echoHandler))
	http.Handle("/", http.FileServer(http.Dir("web/")))
	log.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func echoHandler(session sockjs.Session) {
	var chat Publisher
	log.Println("new sockjs session established")
	var closedSession = make(chan struct{})
	chat.Publish("[info] new participant joined chat")
	defer chat.Publish("[info] participant left chat")
	go func() {
		reader, _ := chat.SubChannel(nil)
		for {
			select {
			case <-closedSession:
				return
			case msg := <-reader:
				if err := session.Send(msg.(string)); err != nil {
					return
				}
			}

		}
	}()
	for {
		if msg, err := session.Recv(); err == nil {
			chat.Publish(msg)
			continue
		}
		break
	}
	close(closedSession)
	log.Println("sockjs session closed")
}
