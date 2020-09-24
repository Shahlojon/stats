package stats

import (
	"github.com/Shahlojon/bank/pkg/types"
)

//Avg рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) (avg types.Money) {
	n := types.Money(len(payments))
	for _, payment := range payments {
		avg += (payment.Amount)
	}
	avg = avg/n
	return avg
}

//TotalInCategory находит сумму покупок в определённой категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment := range payments{
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}