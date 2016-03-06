package main

import (
	"./ruz"
	"fmt"
)

func main() {
	resp, _ := ruz.MakeRequest("npchudinov@edu.hse.ru", "2016.2.22", "2016.2.25")
	fmt.Println(ruz.PrettifyResponse(resp))

	date, _ := ruz.ParseDate("6.3.2016")
	fmt.Println(date)
	mon, sat := ruz.FindWeek(date)
	fmt.Println(ruz.FormatDate(mon), ruz.FormatDate(sat))
}
