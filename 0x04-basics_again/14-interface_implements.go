package main

import "fmt"


type employee interface {
	getName() string
	getSalary() float64
}

type contractor struct {
	name string
	hourlyPay float64
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() float64 {
	return c.hourlyPay * float64(c.hoursPerYear)
}

type fullTime struct {
	name string
	salary float64
}

func (ft fullTime) getSalary() float64 {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func testEmp(e employee) {
	fmt.Println(e.getName(), e.getSalary())
	fmt.Println("===============================================================")
}

func main() {
	testEmp(fullTime{
		name: "Kaka Mbweha",
		salary: 100000.00,
	})
	testEmp(contractor{
		name: "Kaka Tai",
		hourlyPay: 70.25,
		hoursPerYear: 2400,
	})
}