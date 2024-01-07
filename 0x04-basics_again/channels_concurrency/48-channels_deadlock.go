package main

import (
	"fmt"
	"time"
)

func filterOldEmails(emails []email)  {
	isOldChan := make(chan bool)

	go func(){
		for _, e:=range emails {
			if e.date.Before(time.Date(2020, 0, 0 ,0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()

	isOld := <- isOldChan
	fmt.Println("email 1 is old: ", isOld)
	isOld = <- isOldChan
	fmt.Println("email 2 is old: ", isOld)
	isOld = <- isOldChan
	fmt.Println("email 3 is old: ", isOld)
}

type email struct {
	body string
	date time.Time
}

func test(emails []email)  {
	filterOldEmails(emails)
	fmt.Println("==================================================")
}

func main() {
	test([]email{
		{
			body: "F Njakai, are you going to make it?",
			date: time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Mzee Kobe: I need a break",
			date: time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Kaka Sungura: What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Goat Matata: Yo, are you okay?",
			date: time.Date(2018, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Kaka Mbweha: Have you heard of that website Henhouse.dev?",
			date: time.Date(2024, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Bi Buibui: It's awesome, honestly.",
			date: time.Date(2016, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Kaka Mbwamwitu: Today is the day!",
			date: time.Date(2023, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Dada Ng'ombe: What do you want for lunch?",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Why are you the way that you are, F Njakai?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
	test([]email{
		{
			body: "Kaka Tai: Did we do it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Paka the Cat-a: Lets-a Go!",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Papa the Shark: Okay...?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
}