package logicfortxt

import (
	"TO-DO/models"
	"os"
)


func openAndReadFile(file os.File) []models.Task{
	os.Open("../listOfTasks.txt")
}