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

	for scanner.Scan(){
		line := scanner.Text()

		var bufForIndex strings.Builder
		var bufForText strings.Builder
		indexPoint := 0


		//Read Index
		for i, val := range line {
		if val == '.' {
			indexPoint = i + 1
			break
			}
		bufForIndex.WriteString(string(val))
		}

		//Read Text
		for i := indexPoint; i < len(line); i++ {
		if line[i] == '\n' {
			break
			}
		bufForText.WriteString(string(line[i]))
		}

		var task models.Task
		var err error

		task.ID, err = strconv.ParseInt(bufForIndex.String(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		task.Text = bufForText.String()


		tasks = append(tasks, task)
	}

	return tasks
}


func CountOfLinesInFile(file *os.File) int32 {
	scanner := bufio.NewScanner(file)
	res := 0

	for scanner.Scan(){
		res++
	}
	return int32(res)
}