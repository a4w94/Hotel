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

const token = "bwRMIGnrqqLsvMyDlM8OxT89yTSREwUxoNWxPIZ1wuVZvT3FgYtLGgsmIRIb"

func main() {
	HttpRoomTypeGet()

}

func HttpRoomTypeGet() {

	var allRoomTypeUrl AllRoomTypeURL
	allRoomTypeUrl.GetUrl()

	client := &http.Client{}

	req, err := http.NewRequest("GET", allRoomTypeUrl.Url, nil)
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/json")

	Error(err)
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	Error(err)
	fmt.Println(string(body))
	defer resp.Body.Close()

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
