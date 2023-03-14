package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func HttpLogin() []byte {
	url := "http://127.0.0.1:8080/login"
	var jsonStr = []byte(`{"account":"xionzhi","password":"123456"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	// 判断Http Code
	if err != nil && resp.StatusCode != 200 {
		panic(err)
	}

	//关闭连接
	defer resp.Body.Close()

	//读取报文中所有内容
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	//输出内容
	return body
}
