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
			expectedEmployee: fullTimeEmployee{
				person: person{
					name: "Nombre",
					age:  23,
					DNI:  "123",
				},
				employee: employee{
					id:       1,
					position: "CEO",
				},
			},
		},
	}
	originalGetEmployeeById := GetEmployeeByID
	originalGetPersonByDni := GetPersonByDNI
	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmloyeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("error when getting employee")
		}
		if ft.age != test.expectedEmployee.age {
			t.Errorf("got %d but expected %d", ft.age, test.expectedEmployee.age)
		}
	}
	GetEmployeeByID = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni
}
