//package g_web_simple_ex
package main

import (
	"fmt"
	"g-web-simple-ex/rmq"
	"log"
	"net/http"
)

func main() {

	rmq.ReceiveMsgs() // receive RabbitMQ mwssages from 'queue.one' queue

	// HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello from GO!"
		fmt.Fprintf(w, msg)
		err := rmq.SenMsg(msg)
		if err != nil {
			log.Fatalln(err)
		}
	})

	err2 := http.ListenAndServe(":4000", nil)
	if err2 != nil {
		log.Fatalln(err2)
	}
}
