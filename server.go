//    Copyright 2018 storyicon@foxmail.com
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/storyicon/graphquery"
	"github.com/storyicon/graphquery-playground/examples"
)

// Port is the listening port number.
// After the service starts, you can access Playground through 127.0.0.1:port.
const Port = ":8558"

// Response defines the API response format.
type Response struct {
	// Data is the carrier for returning data.
	Data interface{}
	// Error records the errors in this request.
	Error interface{}
	// TimeCost recorded the time wastage of the request.
	TimeCost int64
}

// HTTPRequest is used to download the text of the specified URL using the Get request.
func HTTPRequest(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	return string(body), err
}

// Get the absolute path of the index.html file.
// * In order to fix the file path error caused by cross compilation
func getIndexPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(dir, "index.html")
}

// Start is used to start the http server
func Start() {
	router := gin.Default()
	router.LoadHTMLFiles(getIndexPath())
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title":    "GraphQuery PlayGround",
			"examples": examples.GetExamples(),
		})
	})
	router.GET("/download", func(c *gin.Context) {
		url := c.Query("url")
		timeStart := time.Now().UnixNano()
		body, err := HTTPRequest(url)
		c.JSON(http.StatusOK, Response{
			Data:     body,
			TimeCost: (time.Now().UnixNano() - timeStart),
			Error:    err,
		})
	})
	router.POST("/analyze", func(c *gin.Context) {
		document := c.PostForm("document")
		expr := c.PostForm("expr")
		timeStart := time.Now().UnixNano()
		conseq := graphquery.ParseFromString(document, expr)
		var err interface{}
		if len(conseq.Errors) > 0 {
			err = conseq.Errors[0]
		}
		c.JSON(http.StatusOK, Response{
			Data:     conseq.Data,
			TimeCost: (time.Now().UnixNano() - timeStart),
			Error:    err,
		})
	})
	router.Run(Port)
}

func main() {
	Start()
}
