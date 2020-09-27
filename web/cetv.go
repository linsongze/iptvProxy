package web

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var re, _ = regexp.Compile("<source src=\"(.*?)\"")

func cetvProxyHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		id = "451"
	}
	url := "http://app.cetv.cn/video/videojs/index?site_id=10001&id=" + id
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	bodyString := string(bodyBytes)
	sub := re.FindSubmatch([]byte(bodyString))
	println(sub)
	log.Println(bodyString)
	defer resp.Body.Close()
	c.Redirect(301, string(sub[1]))

}
