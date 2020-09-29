package stats

import (
	"reflect"
	"testing"
	"github.com/KarrenAeris/bank/v2/pkg/types"
)

func TestCategoriesAvg_User(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 3_000_00},
		{ID: 1, Category: "food", Amount: 5_000_00},
		{ID: 1, Category: "auto", Amount: 6_000_00},
		{ID: 1, Category: "auto", Amount: 9_000_00},
		{ID: 1, Category: "clothes", Amount: 4_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 6_000_00,
		"food": 5_000_00,
		"clothes":  4_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}
}

