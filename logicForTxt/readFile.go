package logicfortxt

import (
	models "TO-DO/Models"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)


func ReadFileF(file os.File) []models.Task{

	countLines := CountOfLinesInFile(&file)

	tasks := make([]models.Task, 0, countLines)

	scanner := bufio.NewScanner(&file)

	file.Seek(0,0)

	for scanner.Scan(){
		line := scanner.Text()

		var bufForIndex string
		var bufForText string
		indexPoint := 0



		bufForIndex, indexPoint = ReadIndex(line)  

		bufForText = ReadText(line, indexPoint)

		var task models.Task
		var err error

		task.ID, err = strconv.ParseInt(bufForIndex, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		task.Text = bufForText


		tasks = append(tasks, task)
	}

	return tasks
}

func ReadIndex(line string) (id string, indexPoint int){
	var indPoint int 
	var bfIndex strings.Builder
	
	for i, val := range line {
		if val == '.' {
			indPoint = i
			break
			}
		bfIndex.WriteString(string(val))
	}
	return bfIndex.String(), indPoint
}

func ReadText(line string, indPoint int) string{
	var bfText strings.Builder

	for i := indPoint + 1; i < len(line); i++ {
		if line[i] == '\n' {
			break
			}
		bfText.WriteString(string(line[i]))
		}
	return bfText.String()
}

func CountOfLinesInFile(file *os.File) int32 {
	scanner := bufio.NewScanner(file)
	res := 0

	for scanner.Scan(){
		res++
	}
	return int32(res)
}

