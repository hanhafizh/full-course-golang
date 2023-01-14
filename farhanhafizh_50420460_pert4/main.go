package main

import (
	"farhanhafizh_50420460_pert4/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/", handler.API)

	//Ganti 2 digit akhir port dengan 2 digit akhir NPM anda
	log.Println("localhost : 8060")
	http.ListenAndServe(":8060", nil)
}
