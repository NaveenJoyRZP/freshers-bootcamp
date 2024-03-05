package main

import "fmt"

type payroll interface {
  calculateSalary() float32
}

type Employee struct {
  name string
  wage float32
  employeeType EmployeeType
}

type EmployeeType struct {
  typeOfEmployee string
  wageDistribution string
}

const fullTimeEmployee = "Full Time Employee"
const contractor = "Contractor"
const freeLancer = "FreeLancer"
const dailyWageDistribution = "Daily"
const hourlyWageDistribution = "Hourly"

var fullTimeEmpType = EmployeeType{fullTimeEmployee, dailyWageDistribution}
var contractorEmpType = EmployeeType{contractor, dailyWageDistribution}
var freeLancerEmpType = EmployeeType{contractor, dailyWageDistribution}

func (emp *Employee) calculateSalary() float32 {
  switch empType := emp.employeeType.typeOfEmployee; empType {
  case fullTimeEmployee:
    return emp.wage * 30
  case contractor:
    return emp.wage * 20
  default:
    return 0

  }
}

func main() {
  emp1 := Employee{name: "Naveen", employeeType: fullTimeEmpType, wage: 1000}
  fmt.Println("emp1 salary : ", emp1.calculateSalary())

  emp2 := Employee{name: "Geethu", employeeType: contractorEmpType, wage: 120}
  fmt.Println("emp2 salary : ", emp2.calculateSalary())

}
