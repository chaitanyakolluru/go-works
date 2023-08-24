package main

import "fmt"

type simpleStructure struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Profession string   `json:"profession"`
	Interests  []string `json:"interests"`
	Happiness  float64  `json:"happiness"`
	Fulfilment float64  `json:"fulfilment"`
}

func (s *simpleStructure) GetDetails() {
	fmt.Println("details of the user are provided hereunder")
	fmt.Printf("name: %s\n", s.Name)
	fmt.Printf("age: %d\n", s.Age)
	fmt.Printf("profession: %s\n", s.Profession)
	fmt.Printf("interests: %v\n", s.Interests)
	fmt.Printf("happiness: %v\n", s.Happiness)
	fmt.Printf("fulfilment: %v\n", s.Fulfilment)
}

func simeplFun() {
	chaitanya := simpleStructure{
		Name:       "chaitanya",
		Age:        34,
		Profession: "software engineer",
		Interests:  []string{"tech", "kettlebells", "spirituality", "games", "code"},
		Happiness:  0.6,
		Fulfilment: 0.4,
	}

	chaitanya.GetDetails()

}

func main() {
	simeplFun()
}
