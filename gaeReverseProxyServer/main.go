package main

import (
	"fmt"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/gin-gonic/gin"
	"log"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	apiVersion := "v1"
	serverPort := config.GetEnvString("PORT", "8080")
	r := gin.Default()

	r.Use(
		reverseProxy(
			"/"+apiVersion,
			&url.URL{Scheme: "http", Host: config.GetEnvString("API_SERVER_HOST", "10.128.0.2:18080")},
			&url.URL{Scheme: "http", Host: config.GetEnvString("FRONT_SERVER_HOST", "10.128.0.2:9090")},
		),
	)

	log.Fatal(r.Run(fmt.Sprintf(":%s", serverPort)))
}

func reverseProxy(urlPrefix string, target *url.URL, otherTarget *url.URL) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.FlushInterval = -1

	otherProxy := httputil.NewSingleHostReverseProxy(otherTarget)
	otherProxy.FlushInterval = -1

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, urlPrefix) {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, urlPrefix, "", 1)
			proxy.ServeHTTP(c.Writer, c.Request)
		} else {
			otherProxy.ServeHTTP(c.Writer, c.Request)
		}
	}
}
