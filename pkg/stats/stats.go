package stats

import (
	"github.com/KarrenAeris/bank/pkg/types"
)

//Avg расчитовает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	n := types.Money(0)
	for _, payment := range payments {
		sum += payment.Amount
		n++
	}
	return sum / n
}

// TotalInCategory находит 	сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category != category {
			continue
		}
		sum += payment.Amount
	}
	return sum
}
