package stats

import (
	"reflect"
	"testing"
	"github.com/Shahlojon/bank/v2/pkg/types"
	"fmt"
)


func ExampleAvg(){
	payments := []types.Payment{
	  {
		ID:2,
		Amount:53_00,
		Category: "Box",
		Status: types.StatusFail,
	  },
	  {
		ID:1,
		Amount:53_00,
		Category: "Box",
		Status: types.StatusOk,
	  },
	  {
		ID:3,
		Amount:55_00,
		Category: "Cat",
		Status: types.StatusOk,
	  },
	}
  
	fmt.Println(Avg(payments))
  
	//Output: 5400
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
	  {
		ID:2,
		Amount:53_00,
		Category: "Shop",
	  },
	  {
		ID:1,
		Amount:51_00,
		Category: "Restaurant",
	  },
	  {
		ID:3,
		Amount:52_00,
		Category: "Shop",
		Status: types.StatusFail,
	  },
	}
  
	fmt.Println(TotalInCategory(payments, "Shop"))
  
	//Output: 5300
}


func TestCategoriesAvgUser(t *testing.T) {
	payments := []types.Payment {
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 2, Category: "auto", Amount: 1_000_00},
		{ID: 3, Category: "food", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money {
		"auto": 2_000_00,
		"food": 3_000_00,
		"fun": 5_000_00,
	}
    
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected,result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}