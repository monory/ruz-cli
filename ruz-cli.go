package main

import (
	"./ruz"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseArgs(args []string) (string, time.Time) {
	email := args[1]

	if !strings.Contains(email, "@") {
		email += "@edu.hse.ru"
	}

	date, _ := ruz.ParseDate(args[2])

	return email, date
}

func main() {
	email, date := parseArgs(os.Args)

	from, to := ruz.FindWeek(date)
	resp, _ := ruz.MakeRequest(email, ruz.FormatDate(from), ruz.FormatDate(to))

	fmt.Printf("Показываем расписание для %v с %v до %v.", email, from.Format("2.1.2006"), to.Format("2.1.2006"))
	ruz.PrintWeek(resp)
}
