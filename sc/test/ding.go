package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const webhook_url = "https://oapi.dingtalk.com/robot/send?access_token=721c5fc3c75602776331f778933bcad40405b8ac58db00726942e2e459764f82"

func dingToInfo(s string) bool {
	content, data := make(map[string]string), make(map[string]interface{})
	content["content"] = s
	data["msgtype"] = "text"
	data["text"] = content
	b, _ := json.Marshal(data)

	resp, err := http.Post(webhook_url,
		"application/json",
		bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return true
}

func main()  {
	dingToInfo("wxexecutor报警")
}