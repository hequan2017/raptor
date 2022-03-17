package task

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 钉钉报警

func DingToInfo(accessToken,info string) bool {
	content, data := make(map[string]string), make(map[string]interface{})
	content["content"] = info
	data["msgtype"] = "text"
	data["text"] = content
	b, _ := json.Marshal(data)

	resp, err := http.Post("https://oapi.dingtalk.com/robot/send?access_token="+accessToken,
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
