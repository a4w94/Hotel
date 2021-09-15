package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

type AllRoomTypeURL struct {
	Url string `yaml:"URL_RoomType"`
}

func main() {
	HttpRoomTypeGet()

}

func HttpRoomTypeGet() {

	var allRoomTypeUrl AllRoomTypeURL
	allRoomTypeUrl.GetUrl()

	resp, err := http.Get(allRoomTypeUrl.Url)

	Error(err)

	body, err := ioutil.ReadAll(resp.Body)
	Error(err)
	fmt.Println(string(body))

}

func (roomtype *AllRoomTypeURL) GetUrl() {

	url, err := ioutil.ReadFile("./url.yaml")
	Error(err)

	//yaml文件内容影射到结构体中
	err1 := yaml.Unmarshal(url, &roomtype)
	Error(err1)

}

func Error(err error) {
	if err != nil {
		panic(err)
	}
}
