package stats

import (
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