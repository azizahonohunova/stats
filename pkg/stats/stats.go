package stats

import (
	"github.com/azizahonohunova/bank/v2/pkg/types"
)

func PeriodsDynamic(
	first map[types.Category]types.Money,
	second map[types.Category]types.Money,
) map[types.Category]types.Money {

	//equate lengths
	for key := range second {
		_, ok := first[key]
		if ok == false {
			first[key] = 0
		}
	}
	for key := range first {
		_, ok := second[key]
		if ok == false {
			second[key] = 0
		}
	}
	if len(first) != len(second) {
		panic("Something is WRONG")
	}

	res := map[types.Category]types.Money{}
	for key := range second {
		res[key] = second[key] - first[key]
	}
	return res
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	cnt := map[types.Category]types.Money{}
	res := map[types.Category]types.Money{}

	for _, x := range payments {
		if x.Status == "FAIL" {
			continue
		}
		res[x.Category] += x.Amount
		cnt[x.Category]++
	}
	for key := range res {
		res[key] /= cnt[key]
	}
	return res
}
