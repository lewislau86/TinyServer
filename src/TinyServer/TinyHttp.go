package TinyServer

import (
	"fmt"
	"net/http"
)

func StartHttpServer() {
	fmt.Println("Start Tiny Http Server.")
	http.Handle("/", http.FileServer(http.Dir("./")))
	go http.ListenAndServe(":8123", nil)
}
