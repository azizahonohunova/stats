package stats

import (
	"reflect"
	"testing"

	"github.com/azizahonohunova/bank/v2/pkg/types"
)

func TestPeriodsDynamic_Nil(t *testing.T) {
	first := map[types.Category]types.Money{}
	second := map[types.Category]types.Money{}
	result := PeriodsDynamic(first, second)
	if len(result) > 0 {
		t.Error("first && second is Nil but result is not empty")
	}
}

func TestPeriodsDynamic_Empty(t *testing.T) {
	first := map[types.Category]types.Money{}
	second := map[types.Category]types.Money{}
	result := PeriodsDynamic(first, second)
	if len(result) > 0 {
		t.Error("first && second is empty but result is not empty")
	}
}
func TestPeriodsDynamic_FoundOne(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
	}
	second := map[types.Category]types.Money{
		"food": 20,
	}
	expected := map[types.Category]types.Money{
		"auto": -10,
		"food": 20,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\ninvalid output\ngot: %v\nwont: %v\n\n", result, expected)
	}
}
func TestPeriodsDynamic_FoundMul(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map[types.Category]types.Money{
		"food":   25,
		"auto":   10,
		"mobile": 5,
	}
	expected := map[types.Category]types.Money{
		"auto":   0,
		"food":   5,
		"mobile": 5,
	}
	result := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\ninvalid output\ngot: %v\nwont: %v\n\n", result, expected)
	}
}

//check with payments == nil
func TestCategoriesAvg_Nil(t *testing.T) {
	var payments []types.Payment
	result := CategoriesAvg(payments)
	if len(result) > 0 {
		t.Error("len of map with payments nil > 0")
	}
}

//check with payments == empty
func TestCategoriesAvg_Empty(t *testing.T) {
	payments := []types.Payment{}
	result := CategoriesAvg(payments)
	if len(result) > 0 {
		t.Error("len of map with payments empty > 0")
	}
}

func TestCategoriesAvg_NotFound(t *testing.T) {
	payments := []types.Payment{
		{Category: "Cafe", Amount: 10_0, Status: "FAIL"},
		{Category: "Cafe", Amount: 20_0, Status: "FAIL"},
		{Category: "internet", Amount: 50_0, Status: "FAIL"},
	}
	expected := map[types.Category]types.Money{}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\ninvalid output\nget: %v,\nwont: %v\n", result, expected)
	}
}

func TestCategoriesAvg_FoundOne(t *testing.T) {
	payments := []types.Payment{
		{Category: "Cafe", Amount: 10_0, Status: "FAIL"},
		{Category: "Cafe", Amount: 10_0, Status: "OK"},
		{Category: "internet", Amount: 20_0, Status: "FAIL"},
	}
	expected := map[types.Category]types.Money{
		"Cafe": 10_0,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\ninvalid output\nget: %v,\nwont: %v\n", result, expected)
	}
}

func TestCategoriesAvg_FoundMul(t *testing.T) {
	payments := []types.Payment{
		{Category: "Cafe", Amount: 10_0},
		{Category: "Cafe", Amount: 10_0},
		{Category: "internet", Amount: 20_0},
	}
	expected := map[types.Category]types.Money{
		"Cafe":     10_0,
		"internet": 20_0,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("\ninvalid output\nget: %v,\nwont: %v\n", result, expected)
	}
}
