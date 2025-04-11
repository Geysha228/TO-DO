package main

import (
	logicfortxt "TO-DO/logicForTxt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)


func main() {
	//New instance of Echo
	e := echo.New()

	e.GET("/getAllTasks", getAllTasks)
	e.Logger.Fatal(e.Start(":3000"))
}




//getAllTasks open file read it
//and return full tasks in it
func getAllTasks(c echo.Context) error{
	file, err := os.Open("D:\\GoProjects\\TO-DO\\listOfTasks.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	return c.JSON(http.StatusOK, logicfortxt.ReadFileF(*file)) 

}