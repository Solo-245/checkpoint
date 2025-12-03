//package main

import "fmt"

type Employee struct {
	name     string
	age      int
	isRemote bool
}

func main() {
	employee1 := Employee{
		name:     "Alice",
		age:      30,
		isRemote: true,
	}
	fmt.Println("Employee name:", employee1.name)
	fmt.Println("Employee name:", employee1.age)
	fmt.Println("Employee name:", employee1.isRemote)

	job := struct {
		title  string
		salary int
	}{
		title:  "Software Engineer",
		salary: 2000,
	}
	fmt.Println("job:", job.title)
	fmt.Println("Salary:", job.salary)

	employeePtr := &employee1
	employeePtr.age = 31

}
