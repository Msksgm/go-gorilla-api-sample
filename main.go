package main

// go modulesで管理すると絶対パスでのimportが必須です
import "github.com/msksgm/go-gorilla-api-sample/controllers"

func main() {
	controllers.StartWebServer()
}
