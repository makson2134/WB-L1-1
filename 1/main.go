package main

import (
	"bufio"
	"fmt"
	"os"
)

type Human struct{
	Age uint8
	Name string
	HeightCM uint8
	WeightKG uint16 //Ну мало ли кто больше 255 весит, всякое бывает
}

func (h *Human) SayHello() {
	fmt.Printf("Hello, my name is %s, I'm %d years old.\n", h.Name, h.Age)
}

func (h *Human) CalculateBM()  {
	heightM := float32(h.HeightCM) / 100
	fmt.Printf("BMI is: %.2f\n", float32(h.WeightKG) / (heightM * heightM))
}

func NewHuman(name string, age uint8, heightCM uint8, weightKG uint16) *Human {
	return &Human{
		Name: name,
		Age: age,
		HeightCM: heightCM,
		WeightKG: weightKG,
	}
}
	
type Action struct {
	JobTitle string
	Human
}

func NewAction(name string, age uint8, heightCM uint8, weightKG uint16, Job string) *Action {
	return &Action{
		JobTitle: Job,
		Human: Human{
			Name: name,
			Age: age,
			HeightCM: heightCM,
			WeightKG: weightKG,
		},
	}	
}

func (a *Action) PerformAction()  {
	fmt.Println("Doing Job: " + a.JobTitle)
}

func main() {
	var name string
	var age uint8
	var heightCM uint8
	var weightKG uint16
	var job string
	

	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter Name, Age, Height (in cm) and Weight (in kg) for Human(with spaces without commas): ")
	

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%s %d %d %d", &name, &age, &heightCM, &weightKG)

	h := NewHuman(name, age, heightCM, weightKG)

	h.SayHello()
	h.CalculateBM()

	fmt.Println("Enter Name, Age, Height (in cm), Weight (in kg) and Job for Action (with spaces without commas): ")
	

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%s %d %d %d %s", &name, &age, &heightCM, &weightKG, &job)

	a := NewAction(name, age, heightCM, weightKG, job)

	a.SayHello()
	a.CalculateBM()
	a.PerformAction()


}