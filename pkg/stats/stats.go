package stats

import (
	"github.com/KarrenAeris/bank/v2/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	n := types.Money(0)
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
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
		if payment.Status == types.StatusFail {
			continue
		}
		sum += payment.Amount
	}
	return sum
}

//CategoriesAvg рассчитывает среднюю сумму платежа
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	quantity := map[types.Category]types.Money{}
	result := map[types.Category]types.Money{}

	for _, payment := range payments {
		result[payment.Category] += payment.Amount
		quantity[payment.Category]++
	}

	for key := range result {
		result[key] /= quantity[key]
	}

	return result
}

//PeriodsDynamic сравнивает расходы по категориям за два периода
func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	
	changes := map[types.Category]types.Money{}

	for c := range second {
		changes[c] += second[c]
	}

	for c := range first {
		changes[c] -= first[c]
	}

	return changes
}