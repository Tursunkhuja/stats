package stats

import (
	"fmt"

	"github.com/tursunkhuja/bank/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{}
	payments = append(payments, types.Payment{ID: 1, Amount: 10})
	payments = append(payments, types.Payment{ID: 2, Amount: 8})

	avg := Avg(payments)
	fmt.Println(avg)
	// Output: 9
}
func ExampleTotalInCategory() {
	payments := []types.Payment{}
	payments = append(payments, types.Payment{ID: 1, Amount: 10, Category: "A"})
	payments = append(payments, types.Payment{ID: 2, Amount: 8, Category: "B"})
	payments = append(payments, types.Payment{ID: 3, Amount: 5, Category: "A"})

	result := TotalInCategory(payments, "A")
	fmt.Println(result)
	// Output: 15
}
