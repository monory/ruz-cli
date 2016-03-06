package main

import (
	"./ruz"
)

func main() {
	resp, _ := ruz.MakeRequest("npchudinov@edu.hse.ru", "2016.2.22", "2016.2.25")

	ruz.PrintWeek(resp)
}
