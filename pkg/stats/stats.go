package stats

import "github.com/Tursunkhuja/bank/v2/pkg/types"

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
