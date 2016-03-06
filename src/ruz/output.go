package ruz

import (
	"encoding/json"
	"fmt"
)

type lesson struct {
	BeginLesson     string
	EndLesson       string
	DayOfWeekString string
	Discipline      string
}

// PrintWeek красиво распечатывает ответ сервера РУЗ
func PrintWeek(response []byte) {
	var lessons []lesson
	json.Unmarshal(response, &lessons)

	if len(lessons) == 0 {
		fmt.Println("Пар нет.")
		return
	}

	currentDate := ""
	for _, lesson := range lessons {
		if currentDate != lesson.DayOfWeekString {
			fmt.Printf("\n%v:\n", lesson.DayOfWeekString)
			currentDate = lesson.DayOfWeekString
		}

		fmt.Printf("%v - %v: %v\n", lesson.BeginLesson, lesson.EndLesson, lesson.Discipline)
	}
}
