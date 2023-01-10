package gross

import "fmt"

func main() {
	units := Units()
	fmt.Println(units)

	bill := NewBill()
	fmt.Println(bill)

	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)

	ok1 := RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok1)

	bill1 := map[string]int{"carrot": 12, "grapes": 3}
	qty, ok2 := GetItem(bill1, "carrot")
	fmt.Println(qty)
	// Output: 12
	fmt.Println(ok2)
	// Output: true
}

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if newValue, newUnit := units[unit]; newUnit {
		bill[item] += newValue
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	newValue, newUnit := units[unit]
	if newUnit {
		if bill[item] >= newValue {
			bill[item] -= newValue
			if bill[item] == 0 {
				delete(bill, item)
			}
			return true
		}
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, newItem := bill[item]
	return value, newItem
}
