//package g_web_simple_ex
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//rmq.ReceiveMsgs() // receive RabbitMQ mwssages from 'queue.one' queue

	// HTTP
	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		msg := "Hello from GO!"
		fmt.Fprintf(wr, msg)
		//err := rmq.SendMsg(msg)
		//if err != nil {
		//	log.Fatalln(err)
		//}
	})

	err2 := http.ListenAndServe(":4001", nil)
	if err2 != nil {
		log.Fatalln(err2)
	}
}
