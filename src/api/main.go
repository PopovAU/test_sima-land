package main

import (
	"api/db"
	"api/myhttp"
	"api/setting"
)

func init() {
	setting.OpenEnv()

	db.Init()
}

func main() {
	myhttp.Start()
}
