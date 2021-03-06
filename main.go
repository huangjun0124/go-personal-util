package main

import (
	"flag"
	_ "github.com/joho/godotenv/autoload"
	"goutils/funcs"
	_024 "goutils/funcs/1024"
	"goutils/utils"
)

func main() {
	funcName := flag.String("cmd", "mtloop", "输入功能（mtloop）")
	flag.Parse()
	utils.ConsolePl("Cmd is ", *funcName, " login task scheduing")
	switch *funcName {
	case "mtloop":
		funcs.ScheduleMtLogin()
		//阻塞主线程停止
		select {}
	case "1024":
		_024.CrawDagaierFlagsPictures()
	}
}
