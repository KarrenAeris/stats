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

func TestPeriodsDynamic_positive(t *testing.T) {
	first := map[types.Category]types.Money{
		"cafe": 20,
		"auto": 14,
	}
	second := map[types.Category]types.Money{
		"cafe": 35,
		"auto": 17,
	}
	changes := map[types.Category]types.Money{
		"cafe": 15,
		"auto": 3,
	}

	got := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(changes, got) {
		t.Errorf("\n got > %v \n want > %v", got, changes)
	}

}

func TestPeriodsDynamic_notEqualMap(t *testing.T) {
	first := map[types.Category]types.Money{
		"cafe": 20,
		"auto": 14,
	}
	second := map[types.Category]types.Money{
		"cafe": 35,
	}
	changes := map[types.Category]types.Money{
		"cafe": 15,
		"auto": -14,
	}

	got := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(changes, got) {
		t.Errorf("\n got > %v \n want > %v", got, changes)
	}

}

func TestPeriodsDynamic_OneMoreElem(t *testing.T) {
	first := map[types.Category]types.Money{
		"cafe": 20,
		"auto": 14,
	}
	second := map[types.Category]types.Money{
		"cafe":   35,
		"auto":   17,
		"mobile": 17,
	}
	changes := map[types.Category]types.Money{
		"cafe":   15,
		"auto":   3,
		"mobile": 17,
	}

	got := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(changes, got) {
		t.Errorf("\n got > %v \n want > %v", got, changes)
	}
}