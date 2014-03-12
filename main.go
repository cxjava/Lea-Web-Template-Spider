package main

import (
	"fmt"
	"time"
)

func main() {
	ReadConfig()
	start := time.Now()
	fmt.Println("start...")

	lea := NewLeaSpider()
	lea.Fetch(conf.ThemesUrl, conf.SaveFolder)

	fmt.Printf("time cost %v\n", time.Now().Sub(start))
}
