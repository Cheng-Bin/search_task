package main

import (
	"log"
	"os"

	"github.com/chapin/search_task/search"
	_ "github.com/chapin/search_task/search"
)

func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	//使用特定项做搜索
	search.Run("president")
}
