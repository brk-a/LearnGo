package main

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	newBalance := balance - newItem.GetCost()
	if newBalance<0.0 {
		return nil, 0.0, errors.New("insufficient funds")
	}

	oldItems = append(oldItems, newItem)

	return oldItems, newBalance, nil
}

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	switch s.interval {
	case "monthly":
		return 25.00
	case "yearly":
		return 250.00
	default:
		return 0.00
	}
}

type oneTimeUsagePlan struct {
	userEmail string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one-time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

func main() {
	test(subscription{
		userEmail: "goatmatata@example.co.ke",
		startDate: time.Now().UTC(),
		interval: "yearly",
	},
		[]subscription{},
		1000.00,
	)

	test(subscription{
		userEmail: "kakambwamwitu@example.co.ke",
		startDate: time.Now().UTC(),
		interval: "monthly",
	},
		[]subscription{
			{
				userEmail: "dadangombe@example.co.ke",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval: "monthly",
			},
			{
				userEmail: "kakatai@example.co.ke",
				startDate: time.Now().UTC().Add(-time.Hour * 24 * 7),
				interval: "yearly",
			},
		},
		666.67,
	)

	test(oneTimeUsagePlan{
		userEmail: "bibuibui@example.co.ke",
		numEmailsAllowed: 5000,
	},
		[]oneTimeUsagePlan{},
		314.15,
	)

	test(oneTimeUsagePlan{
		userEmail: "pakathecat@example.co.ke",
		numEmailsAllowed: 7000,
	},
		[]oneTimeUsagePlan{
			{
				userEmail: "papatheshark@example.co.ke",
				numEmailsAllowed: 10000,
			},
			{
				userEmail: "mzeekobe@example.co.ke",
				numEmailsAllowed: 1000,
			},
		},
		1618.29,
	)
}

func test[T lineItem](newItem T, oldItems []T, balance float64) {
	fmt.Println("===============================================")
	fmt.Printf("Charging customer for a '%s'. Current balance is %v", newItem.GetName(), balance)
	newItems, newBalance, err := chargeForLineItem(newItem, oldItems, balance)
	if err!=nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}
	fmt.Printf("New balance is %v. Total number of line items is now %v", newBalance, len(newItems))
}