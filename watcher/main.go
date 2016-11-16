package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var (
	addr = "http://127.0.0.1:8000/session/"
)

func main() {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.LoadHTMLGlob("tpl/*")
	r.GET("/", index)
	r.GET("/size/:id/:height/:width", size)
	r.GET("/shot/:id", screenshot)
	r.Run(":8080")
}

func size(c *gin.Context) {
	var buf = new(bytes.Buffer)
	var size struct {
		Width  string `json:"width"`
		Height string `json:"height"`
	}

	size.Height = c.Param("height")
	size.Width = c.Param("width")
	id := c.Param("id")

	json.NewEncoder(buf).Encode(size)

	url := addr + id + "/window/current/size"

	req, _ := http.NewRequest("POST", url, buf)

	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"server": c.Param("s")})
}

func screenshot(c *gin.Context) {
	id := c.Param("id")
	img := shot(id)
	c.HTML(http.StatusOK, "screenshot.html", gin.H{
		"title": "screen of the Session" + id,
		"img":   img,
	})
}

func shot(id string) string {
	var v struct {
		Value string `json:"value"`
	}
	url := addr + id + "/screenshot"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &v)
	return v.Value
}
