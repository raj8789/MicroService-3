package handle

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	hello.l.Println("Good Morning WelCome Again to Coding Of Go-Lang")
	d, _ := io.ReadAll(r.Body)
	fmt.Printf("Data =%s", d)
}
