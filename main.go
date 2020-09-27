package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"io"
	"iptvProxy/web"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	logFile, err := os.OpenFile(os.Getenv("LIVETV_DATADIR")+"/iptvproxy.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.SetOutput(io.MultiWriter(os.Stderr, os.Stdout))
		log.Println(err)
	} else {
		log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	web.Register(router)
	log.Println("Server listen", os.Getenv("LIVETV_LISTEN"))

	srv := &http.Server{
		Addr:              os.Getenv("LIVETV_LISTEN"),
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		Handler:           router,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panicf("listen: %s\n", err)
	}
}
