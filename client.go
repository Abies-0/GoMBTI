package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"strconv"
)

type Group struct {
	Item int    `json:"item"`
	Q    string `json:"Q"`
	A    string `json:"A"`
	B    string `json:"B"`
}

type ReqData struct {
	Item int `json: "item"`
}

type GinConfig struct {
        Port int `yaml:"port"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	item := 0
	group := next(item)
	fmt.Println(group.Q)
	fmt.Println("A:", group.A)
	fmt.Println("B:", group.B)

	for item < 31 {
		ans := answer(reader)
		item = check(item, ans)
		group = next(item)
		fmt.Println(group.Q)
		if item < 31 {
			fmt.Println("A:", group.A)
			fmt.Println("B:", group.B)
		} else {
			fmt.Printf("測試結果：你的MBTI是「%s」，代表角色「%s」。", group.A, group.B)
		}
	}
}

func next(item int) Group {
	var gin_config GinConfig
	databytes, _ := ioutil.ReadFile("gin_port.yaml")
        yaml.Unmarshal(databytes, &gin_config)
	data := make(map[string]int)
	data["item"] = item
	jsonData, err := json.Marshal(data)
	url := "http://localhost:" + strconv.Itoa(gin_config.Port) + "/api/v1/mbti_quizzes"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("connect to API error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	var group Group
	if err := json.NewDecoder(resp.Body).Decode(&group); err != nil {
		fmt.Println("parse API error:", err)
		os.Exit(1)
	}

	return group
}

func answer(r *bufio.Reader) string {
	for {
		fmt.Print("請輸入您的選擇 (A/B): ")
		ans, _ := r.ReadString('\n')
		ans = strings.TrimSpace(ans)
		if strings.ToUpper(ans) == "A" || strings.ToUpper(ans) == "B" {
			return strings.ToUpper(ans)
		}
		fmt.Println("請輸入有效的選擇 (A/B)")
	}
}

func check(item int, ans string) int {
	if ans == "A" {
		return 2*item + 1
	} else {
		return 2*item + 2
	}
}

