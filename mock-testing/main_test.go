package main

import "testing"

func TestGetEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee fullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeByID = func(id int) (employee, error) {
					return employee{
						id:       1,
						position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (person, error) {
					return person{
						name: "Nombre",
						age:  23,
						DNI:  "123",
					}, nil
				}
			},
		},
	}
}
