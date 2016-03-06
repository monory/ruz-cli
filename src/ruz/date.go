package ruz

import (
	"fmt"
	"time"
)

// ParseDate пытается распарсить дату по нескольким обычным шаблонам
func ParseDate(arg string) (time.Time, error) {
	const template = "2.1.06"
	const template2 = "2.1.2006"

	result, err := time.Parse(template, arg)
	if err != nil {
		result, err = time.Parse(template2, arg)
	}

	return result, err
}

// FindWeek возвращает две даты - понедельник и субботу на той неделе,
// которая в неё передаётся
func FindWeek(day time.Time) (time.Time, time.Time) {
	return findMonday(day), findSaturday(day)
}

// FormatDate возвращает дату в формате, в котором ожидает сервер РУЗ
func FormatDate(day time.Time) string {
	return fmt.Sprintf("%d.%d.%d", day.Year(), day.Month(), day.Day())
}

func findMonday(day time.Time) time.Time {
	return day.AddDate(0, 0, -(int(day.Weekday()) - 1))
}

func findSaturday(day time.Time) time.Time {
	return day.AddDate(0, 0, 6-int(day.Weekday()))
}
