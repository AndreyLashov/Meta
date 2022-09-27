package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/", fileServer)

	http.HandleFunc("/templates", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Сервак не работает!\n")
	})

	log.Println("Запускаю сервак...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
