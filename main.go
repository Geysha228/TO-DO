package main

import "github.com/labstack/echo/v4"

type task struct{
	id int32
	text string
}


func main() {
	//New instance of Echo
	e := echo.New()


}

func getAllTasks() []task{

}