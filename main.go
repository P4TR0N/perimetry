package main

import (
	"fmt"
	"perimetry/v1/services"
)


func main()  {
	fmt.Println("Попытка запустить http-сервер")
	services.HttpServer()
}