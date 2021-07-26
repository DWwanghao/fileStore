package main

import (
	"fileStore/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/filestore/upload",handler.UploadHandler)
	http.ListenAndServe(":8080",nil)


}
