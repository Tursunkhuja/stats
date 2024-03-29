package stats

import (
	"reflect"
	"testing"

	"github.com/Tursunkhuja/bank/v2/pkg/types"
)

func Test_PeriodsDynamic(t *testing.T) {
	firstP := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	secondP := map[types.Category]types.Money{
		"food": 3,
	}

	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": -17,
	}

	result := PeriodsDynamic(firstP, secondP)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}

func Test_CategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 2_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 2_000_00},
		{ID: 6, Category: "fun", Amount: 2_000_00},
	}

	expected := map[types.Category]types.Money{
		"auto": 3_000_00,
		"food": 2_000_00,
		"fun":  2_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}
