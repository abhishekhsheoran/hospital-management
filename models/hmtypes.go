package models

import "time"

type Doctor struct {
	Name      string
	Age       int
	Contact   int
	Address   []Address
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Address struct {
	State string
	City  string
	Pin   int
}

type Patient struct {
	Name    string
	Age     int
	Contact int
	Address Address
}

type Employees struct {
	Name    string
	Age     int
	Contact int
	Address Address
}
