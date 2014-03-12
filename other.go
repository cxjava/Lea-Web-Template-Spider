package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/BurntSushi/toml"
)

type config struct {
	Proxy        bool
	ProxyURL     string
	ThemesUrl    string
	SaveFolder   string
	GoroutineNum int
}

var (
	conf   config
	client = &http.Client{}
)

//读取配置文件
func ReadConfig() {
	//读取配置文件
	if _, err := toml.DecodeFile("config.ini", &conf); err != nil {
		fmt.Println("配置文件错误:", err)
		return
	}
	//如果有代理，设置代理
	if conf.Proxy {
		pr, err := url.Parse(conf.ProxyURL)
		if err != nil {
			fmt.Println("url.Parse(conf.ProxyURL):", err)
			return
		}
		client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(pr),
			},
		}
	}

	conf.SaveFolder = conf.SaveFolder + "/"
}

//添加消息头
func AddReqestHeader(request *http.Request) {
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Accept-Charset", "utf-8;q=0.7,*;q=0.3")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.72 Safari/537.36")
}

//获取响应
func GetResponse(targetUrl string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Println("GetResponse:", targetUrl, err)
		return
	}
	AddReqestHeader(req)
	return client.Do(req)
}

//获取内容
func GetResponseBody(resp *http.Response) string {
	var body []byte
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
		}
		defer reader.Close()
		bodyByte, err := ioutil.ReadAll(reader)
		if err != nil {
		}
		body = bodyByte
	default:
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		}
		body = bodyByte
	}
	return string(body)
}
