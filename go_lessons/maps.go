package main

import (
	"fmt"
)

func main() {
	s := []string{"хлеб", "буженина", "сыр", "огурцы"}
	maps_task_1(s)

}
func maps_task_1(z []string) {
	// func should print all price, which is low 500
	Price := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	for k, v := range Price {
		if v < 500 {
			fmt.Printf("Price %s low 500\n", k)
		}
	}

	summ := 0

	for _, k := range z {
		v, ok := Price[k]
		if ok {
			summ += v
		}

	}
	fmt.Printf("Summ: %d\n", summ)
}
