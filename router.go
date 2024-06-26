package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveHTTP() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.GET("/", HTTPAPIServerIndex)
	router.GET("/stream/all/:pt", MultiPlayer)
	router.GET("/stream/floor/:uuid", HTTPAPIServerFloor)
	router.GET("/stream/player/:uuid", HTTPAPIServerStreamPlayer)
	router.GET("/stream/thumbnail/:uuid", HTTPAPIServerThumbnail)
	router.POST("/stream/receiver/:uuid", HTTPAPIServerStreamWebRTC)
	router.GET("/stream/codec/:uuid", HTTPAPIServerStreamCodec)

	router.StaticFS("/static", http.Dir("web/static"))
	err := router.Run(Config.Server.HTTPPort)
	if err != nil {
		log.Fatalln("Start HTTP Server error", err)
	}
}
