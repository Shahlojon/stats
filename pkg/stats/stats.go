package stats

import (
	"github.com/Shahlojon/bank/v2/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) (avg types.Money) {
	n := types.Money(0)
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			n++
			avg += (payment.Amount)
		}
	}
	avg = avg / n
	return avg
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			if payment.Category == category {
				sum += payment.Amount
			}
		}
	}
	return sum
}

//FilterByCategory
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
		for _, payment := range payments {
			if payment.Category == category {
				filtered = append(filtered, payment)
			}
		}
		return filtered
}

//CategoriesAvg, которая считает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment)  map[types.Category]types.Money {
	n := make(map[types.Category]types.Money)
	// categories := map[types.Category]types.Money{}
	for _, payment := range payments {
		if _, er := n[payment.Category]; er {
			continue
		}
		filtered := FilterByCategory(payments, payment.Category)
		n[payment.Category]=Avg(filtered)
	}
	return n 
}
