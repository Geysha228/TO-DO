package logicfortxt

import (
	models "TO-DO/Models"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)


func ReadFileF(file os.File) []models.Task{
	allText, err := io.ReadAll(&file)
	if err != nil {
		log.Fatal(err)
	}

	var bufForIndex strings.Builder
	var bufForText strings.Builder
	indexPoint := 0

	for i, val := range allText {
		if val == '.' {
			indexPoint = i + 1
			break
		}
		bufForIndex.WriteString(string(val))
	}

	for i := indexPoint; i < len(allText); i++ {
		if allText[i] == '\n' {
			break
		}
		bufForText.WriteString(string(allText[i]))
	}

	var task models.Task
	task.ID, err = strconv.ParseInt(bufForIndex.String(), 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	task.Text = bufForText.String()

	var tasks []models.Task
	tasks = append(tasks, task)
	return tasks
}