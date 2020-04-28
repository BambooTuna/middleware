package main

import (
	"fmt"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	serverPort := config.GetEnvString("PORT", "9090")
	r := gin.Default()

	target := config.GetEnvString("API_SERVER_ENDPOINT", "http://localhost:18080")
	r.Use(reverseProxy("/v1", target))

	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	log.Fatal(r.Run(fmt.Sprintf(":%s", serverPort)))
}

func reverseProxy(urlPrefix string, target string) gin.HandlerFunc {
	url, err := url.Parse(target)
	if err != nil {
		log.Println("Reverse Proxy target url could not be parsed:", err)
		return nil
	}

	http.Handle("/", httputil.NewSingleHostReverseProxy(url))
	proxy := httputil.NewSingleHostReverseProxy(url)
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, urlPrefix) {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, urlPrefix, "/", 1)
			proxy.ServeHTTP(c.Writer, c.Request)
		}
	}
}
