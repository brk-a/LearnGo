package main

import "fmt"

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.00
	if ub.Plan=="pro" {
		amount = 100.00
	}

	return bill{
		Customer: u,
		Amount: amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
} 

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.00
	if ob.Plan=="pro" {
		amount = 3000.00
	}

	return bill {
		Customer: o,
		Amount: amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func main() {
	test[user](
		userBiller{Plan: "basic"},
		user{UserEmail: "goatmatata@example.co.ke"},
	)

	test[user](
		userBiller{Plan: "pro"},
		user{UserEmail: "kakasungura@example.co.ke"},
	)

	test[org](
		orgBiller{Plan: "basic"},
		org{Admin: user{UserEmail: "dadangombe@example.co.ke"}},
	)

	test[org](
		orgBiller{Plan: "pro"},
		org{Admin: user{UserEmail: "bibuibui@example.co.ke"}},
	)
}

func test[C customer](b biller[C], c C) {
	fmt.Printf("Using '%s' to create a bill for '%s'\n", b.Name(), c.GetBillingEmail())
	bill := b.Charge(c)
	fmt.Printf("Bill created. KES %v\n", bill.Amount)
	fmt.Println("=================================================================")
}