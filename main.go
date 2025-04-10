package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
)


func main() {
	//New instance of Echo
	e := echo.New()


}

//getAllTasks open file read it
//and return full tasks in it
func getAllTasks() []Task{
	file, err := os.Open("../listOfTasks")
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()



}