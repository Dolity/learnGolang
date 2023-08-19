package main

import "fmt"

type employee struct {
	EID    string
	Ename  string
	Ephone string
}

func main() {
	// employee1 := employee{
	// 	EID:    "E001",
	// 	Ename:  "John",
	// 	Ephone: "08123456789",
	// }
	// fmt.Println("Employee =", employee1)
	// fmt.Printf("Employee ID: %v\n", employee1.EID)

	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	EID:    "E002",
	// 	Ename:  "Aoi",
	// 	Ephone: "081233456789",
	// }
	// employeeList[1] = employee{
	// 	EID:    "E003",
	// 	Ename:  "Aoi",
	// 	Ephone: "081233456789",
	// }
	// employeeList[2] = employee{
	// 	EID:    "E004",
	// 	Ename:  "toy",
	// 	Ephone: "0812334489",
	// }
	// fmt.Println("Employee List =", employeeList)

	employeeList := []employee{}
	employee1 := employee{
		EID:    "E001",
		Ename:  "John",
		Ephone: "08123456789",
	}
	employeeList = append(employeeList, employee1)
	fmt.Println("Employee =", employeeList)

}
