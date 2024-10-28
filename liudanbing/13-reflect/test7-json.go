package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string `json:"title"` // 如果需要将Movie这个结构体转换为json格式，那么Title这个字段在json中会以title的形式展示
	Year int `json:"year"`
	Price int `json:"dollar"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"周星驰", "莫文蔚"}}

	// 编码的过程：将Go中的结构体转换为json格式
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr: %s\n", jsonStr) // jsonStr: {"title":"喜剧之王","year":2000,"dollar":10,"actors":["周星驰","莫文蔚"]}

	// 解码的过程：将json格式转换为Go中的结构体
	// jsonStr: {"title":"喜剧之王","year":2000,"dollar":10,"actors":["周星驰","莫文蔚"]}
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Printf("my_movie: %v\n", my_movie) // my_movie: {喜剧之王 2000 10 [周星驰 莫文蔚]}
}

