package main

import "fmt"

type Patient struct {
	Name string `json:"name"`
}

type Doctor struct {
	Name string `json:"name"`
	// Patients []Patient `json:"patients"`
	Schedule map[int]Patient `json:"schedule"`
}

func createSchedule() (doc *Doctor) {
	doc = &Doctor{Name: "Doc", Schedule: map[int]Patient{}}
	for i := 0; i < 16; i++ {
		doc.Schedule[i] = Patient{Name: fmt.Sprintf("patient no %d", i)}
	}
	return
}

func main() {
	doctorWithSchedule := createSchedule()
	fmt.Println(doctorWithSchedule)
}
