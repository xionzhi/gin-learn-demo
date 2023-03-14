package test

import (
	"io/ioutil"
	"net/http"
)

func Sum(a int, b int) int {
	return a + b
}

func HttpIndex() []byte {
	resp, err := http.Get("http://127.0.0.1:8080/")
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

/*
go test -v ./test

go test -v ./test -run="none" -bench="Bench*"

go test -v ./test -bench="Bench*"
*/
