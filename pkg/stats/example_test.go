package stats

import (
	"fmt"

	"github.com/KarrenAeris/bank/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   53_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       2,
			Amount:   51_00,
			Category: "Cat",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cat",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "Cafe",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   51_00,
			Category: "Cafe",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Restaurant",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(TotalInCategory(payments, "Cafe"))

	//Output: 10400
}
