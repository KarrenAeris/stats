package stats

import (
	"github.com/KarrenAeris/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	n := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
		n++
	}
	return sum / n
}

func TotalInCategory (payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if category == payment.Category {
			sum += payment.Amount
		}
	}
	return sum
}