package stats

import (
	"github.com/Shahlojon/bank/pkg/types"
	"fmt"
)


func ExampleAvg(){
	payments := []types.Payment{
	  {
		ID:2,
		Amount:53_00,
		Category: "Cat",
	  },
	  {
		ID:1,
		Amount:53_00,
		Category: "Cat",
	  },
	  {
		ID:3,
		Amount:55_00,
		Category: "Cat",
	  },
	}
  
	fmt.Println(Avg(payments))
  
	//Output: 5366
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
	  {
		ID:2,
		Amount:53_00,
		Category: "Cafe",
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
	  },
	}
  
	fmt.Println(TotalInCategory(payments, "Cafe"))
  
	//Output: 5300
}