package main

import "fmt"

// type Org interface {
// 	Salary
// 	Emp
// }
type Emp interface {
	Getter()
	Setter(emp Employee)
}

type Salary interface {
	CalculateSalary() float64
}

func (e Employee) Getter() {
	fmt.Printf("Employee Details\n")
	fmt.Printf("Id: %v\nName: %s\nAddress:\nCity: %s State: %s Pincode: %v\n PhoneNumber: %d", e.Id, e.Name, e.Address.City, e.Address.State, e.Address.Pincode, e.Address.PhoneNumber)
}
func (e *Employee) Setter(emp Employee) {
	e.Name = emp.Name
	e.Address.City = emp.Address.City
	e.Address.State = emp.Address.State
	e.Address.Pincode = emp.Address.Pincode
	e.Address.PhoneNumber = emp.Address.PhoneNumber
}

func (c Contractual) CalculateSalary() float64 {
	return c.CostPerHr * c.Hours
}
func (i Intern) CalculateSalary() float64 {
	return i.Stipend + i.Bonus
}
func (f Fte) CalculateSalary() float64 {
	return f.Base + f.Bonus + f.Pf
}

type Employee struct {
	Id      int
	Name    string
	Address Address
}
type Address struct {
	City        string
	State       string
	Pincode     int
	PhoneNumber int64
}
type Contractual struct {
	CostPerHr float64
	Hours     float64
}
type Intern struct {
	Stipend float64
	Bonus   float64
}
type Fte struct {
	Base  float64
	Bonus float64
	Pf    float64
}

func main() {
	emp1 := Employee{1, "Shivani", Address{"Guna", "MP", 473287, 7373737373}}
	emp2 := Employee{2, "Sharif", Address{"Delhi", "Delhi", 473247, 7374737373}}
	// emp3 := Employee{2, "Richesh", Address{"Lucknow", "UP", 923287, 7373730173}}
	intern1 := Intern{30000, 10000}
	cntrct1 := Contractual{100, 50}
	fte1 := Fte{100000, 78000, 27829.80}
	var x, y, z Salary
	var x1 Emp
	// x = emp1
	x = fte1
	y = cntrct1
	z = intern1
	// x.Getter()
	x1 = &emp1
	x1.Getter()
	x1.Setter(emp2)
	x1.Getter()
	fmt.Printf("Total Salary %v", x.CalculateSalary()+y.CalculateSalary()+z.CalculateSalary())

}
