// Company Manager - Create an hierarchy of classes - abstract class Employee and subclasses HourlyEmployee,
// SalariedEmployee, Manager and Executive. Every one's pay is calculated differently, research a bit about it.
// After you've established an employee hierarchy, create a Company class that allows you to manage the employees.
//You should be able to hire, fire and raise employees.

package main

import (
	"fmt"
	"reflect"
)

type Company struct {
	Employees []CompanyEmployee
}

func (c *Company) HireEmployee(e CompanyEmployee) {
	c.Employees = append(c.Employees, e)
}

func (c *Company) FireEmployee(e CompanyEmployee) {
	for i, ce := range c.Employees {
		if ce == e {
			c.Employees = append(c.Employees[:i], c.Employees[i+1:]...)
			break
		}
	}
}

func (c *Company) RaiseEmployee(to CompanyEmployee, e CompanyEmployee, added float32) {
	emp, pay := e.GetDetails()
	switch to.GetType() {
	case "Manager":
		if len(pay) != 2 {
			break
		}
		for i, ce := range c.Employees {
			if ce == e {
				newM := &Manager{Employee: emp, Base: pay[0], Bonus: pay[1], Stock: added}
				c.Employees[i] = newM
				break
			}
		}
	case "Executive":
		if len(pay) != 3 {
			break
		}
		for i, ce := range c.Employees {
			if ce == e {
				newM := &Executive{Employee: emp, Base: pay[0], Bonus: pay[1], Stock: pay[2], Additional: added}
				c.Employees[i] = newM
				break
			}
		}
	}
}

type CompanyEmployee interface {
	IsEmployee() bool
	GetDetails() (*Employee, []float32)
	GetType() string
}

type Employee struct {
	HireDate string
}

type Salaried struct {
	*Employee
	Base  float32
	Bonus float32
}

func (s *Salaried) IsEmployee() bool { return true }

func (s *Salaried) GetDetails() (*Employee, []float32) {
	return s.Employee, []float32{s.Base, s.Bonus}
}

func (s *Salaried) GetType() string {
	return reflect.TypeOf(*s).Name()
}

type Manager struct {
	*Employee
	Base  float32
	Bonus float32
	Stock float32
}

func (m *Manager) IsEmployee() bool { return true }

func (m *Manager) GetDetails() (*Employee, []float32) {
	return m.Employee, []float32{m.Base, m.Bonus, m.Stock}
}

func (m *Manager) GetType() string {
	return reflect.TypeOf(*m).Name()
}

type Executive struct {
	*Employee
	Base       float32
	Bonus      float32
	Stock      float32
	Additional float32
}

func (e *Executive) IsEmployee() bool { return true }

func (e *Executive) GetDetails() (*Employee, []float32) {
	return e.Employee, []float32{e.Base, e.Bonus, e.Stock, e.Additional}
}

func (e *Executive) GetType() string {
	return reflect.TypeOf(*e).Name()
}

type withEmployeeOption func(e *Employee)

func withHiredate(o string) withEmployeeOption {
	return func(e *Employee) {
		e.HireDate = o
	}
}

func NewEmployee(opts ...withEmployeeOption) *Employee {
	e := &Employee{}

	for _, f := range opts {
		f(e)
	}
	return e
}

func main() {
	salaried := &Salaried{Employee: NewEmployee(withHiredate("01-01-2023")), Base: 100000.00, Bonus: 10000.00}
	salaried2 := &Salaried{Employee: NewEmployee(withHiredate("01-01-2013")), Base: 100000.00, Bonus: 10000.00}
	manager := &Manager{Employee: NewEmployee(withHiredate("01-01-2020")), Base: 300000.00, Bonus: 20000.00, Stock: 4000.00}
	executive := &Executive{Employee: NewEmployee(withHiredate("01-01-2010")), Base: 500000.00, Bonus: 40000.00, Stock: 8000.00, Additional: 10000.00}

	cc := &Company{}
	cc.HireEmployee(salaried)
	cc.HireEmployee(salaried2)
	cc.HireEmployee(manager)
	cc.HireEmployee(executive)

	cc.FireEmployee(salaried2)

	cc.RaiseEmployee(&Manager{}, salaried, 4000.00)
	cc.RaiseEmployee(&Executive{}, manager, 8000.00)

	for _, ce := range cc.Employees {
		fmt.Println(ce.GetType())
		fmt.Println(ce.GetDetails())
		fmt.Println("##########")
	}
}
