package handler

import (
	"io"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter,r *http.Request)  {
	if r.Method=="GET"{
		//返回上传html
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			return io.WriteString(w,"inter error")
		}

	}else if r.Method=="POST" {
		//接收文件流以及存储到本地


	}
	
}