package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"strconv"
)

type GinConfig struct {
        Port int `yaml:"port"`
}

type ReqData struct {
	Item int `json:"item"`
}

func main() {
	databytes, _ := ioutil.ReadFile("gin_port.yaml")
	var gin_config GinConfig
	yaml.Unmarshal(databytes, &gin_config)
	r := gin.Default()
	cors_config := cors.DefaultConfig()
	cors_config.AllowAllOrigins = true
	r.Use(cors.New(cors_config))
	run_port := ":" + strconv.Itoa(gin_config.Port)
	r.POST("/api/v1/mbti_quizzes", fetch)
	r.Run(run_port)
}

func fetch(context *gin.Context) {
	var req_data ReqData
	context.BindJSON(&req_data)
	item := req_data.Item
	res := fetch_data(item)
	context.JSON(http.StatusOK, gin.H{
		"item":  res["item"],
		"Q": res["Q"],
		"A": res["A"],
		"B": res["B"],
	})
}
