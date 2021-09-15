package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type GetAllRoomType struct {
	Url string `yaml:"URL_RoomType"`
}

func main() {
	HttpRoomTypeGet()

}

func HttpRoomTypeGet() {

	var roomtype GetAllRoomType
	roomtype.GetAllRoomTypeAPI()

}

func (roomtype *GetAllRoomType) GetAllRoomTypeAPI() {

	url, err := ioutil.ReadFile("./url.yaml")
	Error(err)
	fmt.Println(url)

	//yaml文件内容影射到结构体中
	err1 := yaml.Unmarshal(url, &roomtype)
	Error(err1)
	fmt.Println(roomtype.Url)

}

func Error(err error) {
	if err != nil {
		panic(err)
	}
}
