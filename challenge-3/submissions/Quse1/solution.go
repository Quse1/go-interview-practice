package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
	m.Employees = append(m.Employees, e)
	// TODO: Implement this method
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
	delidx := -1
	for i, field := range m.Employees {
		if field.ID == id {
			delidx = i
			break
		}
	}
	if delidx != -1 {
		m.Employees = append(m.Employees[:delidx], m.Employees[delidx+1:]...)
	}
	// TODO: Implement this method
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	// TODO: Implement this method
	if len(m.Employees) == 0 {
		return 0
	}
	var sum float64

	for _, field := range m.Employees {
		sum += field.Salary
	}
	return sum / float64(len(m.Employees))
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
	for _, field := range m.Employees {
		if field.ID == id {
			return &Employee{
				ID:     field.ID,
				Name:   field.Name,
				Age:    field.Age,
				Salary: field.Salary,
			}
		}
	}
	// TODO: Implement this method
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
