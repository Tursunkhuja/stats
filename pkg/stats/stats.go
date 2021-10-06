package stats

import (
	"github.com/tursunkhuja/bank/pkg/types"
)

//Avg calculate avarage sum of payments
func Avg(payments []types.Payment) types.Money {
	sum := int64(0)
	i := 0
	for _, payment := range payments {
		sum += int64(payment.Amount)
		i += 1
	}

	result := sum / int64(i)
	return types.Money(result)
}

//TotalInCategory returns sum of payments in specific category
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := int64(0)
	for _, payment := range payments {
		if payment.Category == category {
			sum += int64(payment.Amount)
		}
	}

	return types.Money(sum)
}
