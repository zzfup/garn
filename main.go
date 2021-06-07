package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func crawl(q string) {
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://pkg.go.dev/search?q=" + q
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + q + "失败")
		return
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取"+q+"失败", err)
		return
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + q + "失败")
		return
	}

	allData := []map[string]string{}
	document.Find(".SearchResults .LegacySearchSnippet").Each(func(i int, selection *goquery.Selection) {
		pkgName := selection.Find("h2").Text()
		desc := selection.Find("p").Text()
		allData = append(allData, map[string]string{
			"pkg":  strings.TrimSpace(pkgName),
			"desc": desc,
		})
	})

	first := allData[0]
	for _, v := range allData {
		if strings.Contains(v["pkg"], q) {
			first = v
			break
		}
	}

	fmt.Println("找到包：", first["pkg"])
	fmt.Println(first["desc"])

	path, _ := os.Getwd()
	fmt.Println("工作目录：", path)

	cmd := exec.Command("go", "get", "-u", "-v", first["pkg"])
	fmt.Println("开始执行：", cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out), err)
		return
	}
	fmt.Println(string(out))
}

func main() {
	args := os.Args

	if args[1] != "add" {
		fmt.Println("gopm add xxx")
	}
	q := args[2]

	fmt.Println("开始抓取：", q)
	crawl(q)
}
