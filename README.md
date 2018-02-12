# ginprometheus
 [![GoDoc](https://godoc.org/github.com/FortnoxAB/ginprometheus?status.svg)](https://godoc.org/github.com/fortnoxab/ginprometheus)

Send gin request metrics to prometheus


## installation
```
go get github.com/fortnoxab/ginprometheus
```


## usage example
```go
package main

import (
	"flag"
	"log"
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
	"github.com/fortnoxab/ginprometheus"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	g := gin.Default()

	m := ginprometheus.New("http")
	m.Use(g)

	g.GET("/one/:id/", func(c *gin.Context) {
		c.String(200, "one")
	})
	g.GET("/two/:id/", func(c *gin.Context) {
		c.String(200, "two")
	})
	log.Fatal(http.ListenAndServe(*addr, g))
}
```
