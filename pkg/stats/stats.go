package stats

import "github.com/Tursunkhuja/bank/v2/pkg/types"

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	results := map[types.Category]types.Money{}

	for category, _ := range first {
		_, ok := second[category]
		if ok {
		} else {
			second[category] = 0
		}
	}

	for category, amount := range second {
		results[category] = amount - first[category]
	}

	return results
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {

	categoriesTotalAmount := map[types.Category]types.Money{}
	categoriesCount := map[types.Category]types.Money{}
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categoriesTotalAmount[payment.Category] += payment.Amount
		categoriesCount[payment.Category] += 1
	}

	for category, amount := range categoriesTotalAmount {
		categories[category] = amount / categoriesCount[category]
	}

	return categories
}
