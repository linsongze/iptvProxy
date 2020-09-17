package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"iptvProxy/web"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())
	web.Register(router)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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
