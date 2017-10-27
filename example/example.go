package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/fortnoxab/ginprometheus"
	gin "gopkg.in/gin-gonic/gin.v1"
)

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	g := gin.Default()

	m := ginprometheus.New("")
	m.Use(g)

	g.GET("/one/:id/", func(c *gin.Context) {
		c.String(200, "one")
	})
	g.GET("/two/:id/", func(c *gin.Context) {
		c.String(200, "two")
	})
	log.Fatal(http.ListenAndServe(*addr, g))
}
