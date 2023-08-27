package main

import "fmt"

type workLife struct {
	Happiness float64 `json:"happiness"`
	Success   float64 `json:"success"`
	Growth    float64 `json:"growth"`
}

type simpleStructure struct {
	Name       string   `json:"name"`
	Age        int      `json:"age"`
	Profession string   `json:"profession"`
	Interests  []string `json:"interests"`
	Happiness  float64  `json:"happiness"`
	Fulfilment float64  `json:"fulfilment"`
	workLife
}

func (s *simpleStructure) GetDetails() {
	fmt.Println("details of the user are provided hereunder")
	fmt.Printf("name: %s\n", s.Name)
	fmt.Printf("age: %d\n", s.Age)
	fmt.Printf("profession: %s\n", s.Profession)
	fmt.Printf("interests: %v\n", s.Interests)
	fmt.Printf("happiness: %v\n", s.Happiness)
	fmt.Printf("fulfilment: %v\n", s.Fulfilment)
	fmt.Printf("happiness: %v, success: %v, growth: %v\n", s.workLife.Happiness, s.workLife.Success, s.workLife.Growth)
}

func simeplFun() {
	chaitanya := simpleStructure{
		Name:       "chaitanya",
		Age:        34,
		Profession: "software engineer",
		Interests:  []string{"tech", "kettlebells", "spirituality", "games", "code"},
		Happiness:  0.6,
		Fulfilment: 0.4,
		workLife: workLife{Happiness: 0.5,
			Success: 0.4,
			Growth:  0.5,
		},
	}

	chaitanya.GetDetails()

}

func main() {
	simeplFun()
}
