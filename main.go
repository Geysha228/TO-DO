package main

import (
	models "TO-DO/Models"
	logicfortxt "TO-DO/logicForTxt"
	"fmt"
	"log"
	"os"
)


func main() {
	//New instance of Echo
	// e := echo.New()

	res := getAllTasks()
	fmt.Println(res)
}

//getAllTasks open file read it
//and return full tasks in it
func getAllTasks() []models.Task{
	file, err := os.Open("D:\\GoProjects\\TO-DO\\listOfTasks.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	defer file.Close()

	return logicfortxt.ReadFileF(*file)

}