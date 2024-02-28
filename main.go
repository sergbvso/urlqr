package main

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/skip2/go-qrcode"
)

var html_start string = `<!doctype html>
<html>
<head>
<title>URL QR</title>
<meta name="description" content="URL QR">
</head>
<body>
<img src='data:image/png;base64, `
var html_end string = `' />
</body>
</html>
`

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: proto
	reqUrl := "https://" + req.Host + req.RequestURI
	png, _ := qrcode.Encode(reqUrl, qrcode.Medium, 512)
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(png)))
	base64.StdEncoding.Encode(dst, png)
	htmlFull := html_start + string(dst) + html_end
	_, err := w.Write([]byte(string(htmlFull)))
	if err != nil {
		log.Println(err)
	}
	log.Println(reqUrl)
}

func main() {
	log.Println(http.ListenAndServe("127.0.0.1:8080", &Handler{}))
}
