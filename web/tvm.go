package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func TvmProxyHandler(c *gin.Context) {
	ts := c.Query("ts")
	id := c.Query("id")
	if ts == "" { //请求m3u8文件
		id = strings.TrimSpace(id)
		if id == "" {
			id = "CCTV1HD"
		}
		bitrate := "1500000"
		if strings.Contains(id, "HD") {
			bitrate = "3000000"
		}
		url := "https://live-bdxcx.mtq.tvmmedia.cn/baidu/live_" + id + ".m3u8"
		log.Println(url)
		client := http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 8.0.0; BLN-AL40 Build/HONORBLN-AL40; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/63.0.3239.83 Mobile Safari/537.36 T7/11.11 swan/2.8.0 swan-baiduboxapp/11.11.0.12 baiduboxapp/11.11.0.12 (Baidu; P1 8.0.0) dumedia/6.14.2.8")

		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if resp.StatusCode != http.StatusOK {
			log.Println(resp.StatusCode)
			c.AbortWithStatus(resp.StatusCode)
			return
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		bodyString := string(bodyBytes)
		if !strings.Contains(bodyString, "#EXTM3U") {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		bodyString = strings.ReplaceAll(bodyString, "http://live-bdxcx.mtq.tvmmedia.cn/", "/tvm.m3u8?ts=")
		bodyString = strings.ReplaceAll(bodyString, "1500000", bitrate)
		bodyString = strings.ReplaceAll(bodyString, "\n", "\r\n")

		c.Data(http.StatusOK, "application/vnd.apple.mpegurl", []byte(bodyString))

	} else { //请求数据流
		url := "http://live-bdxcx.mtq.tvmmedia.cn/" + ts
		client := http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 8.0.0; BLN-AL40 Build/HONORBLN-AL40; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/63.0.3239.83 Mobile Safari/537.36 T7/11.11 swan/2.8.0 swan-baiduboxapp/11.11.0.12 baiduboxapp/11.11.0.12 (Baidu; P1 8.0.0) dumedia/6.14.2.8")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if resp.StatusCode != http.StatusOK {
			c.AbortWithStatus(resp.StatusCode)
			return
		}
		defer resp.Body.Close()
		c.DataFromReader(http.StatusOK, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
	}
}
