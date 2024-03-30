package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil" //简化io的包
	"log"
	"net/http"
	"os"
)

/*
api-key:21d71799
使用开放电影数据库的JSON服务接口，允许你检索和下载 https://omdbapi.com/
上电影的名字和对应的海报图像。编写一个poster工具，通过命令行输入的电影名字，下载对应的海报
*/
//获取电影数据
func getMovieData(movieName string) ([]byte, error) {
	url := "https://omdbapi.com/?t=" + movieName + "&apikey=21d71799"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//延迟关闭 HTTP 响应的主体
	defer resp.Body.Close()
	//fmt.Println("%+v", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// get poster url from movies json
func getPosterUrl(movieJson []byte) (string, error) {
	var movie struct {
		Poster string `json:"Poster"`
	}
	//json.Unmarshal 函数用于将 JSON 数据解析到结构体中
	err := json.Unmarshal(movieJson, &movie)
	//fmt.Println("%+v", movie)
	if err != nil {
		return "", err
	}
	return movie.Poster, nil
}
func main() {
	//命令行读入电影名字
	movieName := os.Args[1]
	movieJson, err := getMovieData(movieName)
	if err != nil {
		log.Fatal(err)
	}
	posterUrl, err := getPosterUrl(movieJson)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(posterUrl)
	resp, err := http.Get(posterUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(movieName+".jpg", body, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(movieName + ".jpg downloaded")
}
