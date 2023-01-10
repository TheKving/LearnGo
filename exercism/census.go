// Package census simulates a system used to collect census data.
package census

import "fmt"

func main() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	// => &{Matthew Sanabria 29 map[street:Main St.]}
	fmt.Println(NewResident(name, age, address))

	name11 := ""
	age11 := 0
	address11 := make(map[string]string)
	resident11 := NewResident(name11, age11, address11)
	fmt.Println(resident11.HasRequiredInfo()) // => false

	resident := NewResident(name, age, address)
	fmt.Println(resident) // Output: &{Matthew Sanabria 29 map[street:Main St.]}
	resident.Delete()
	fmt.Println(resident) // Output: &{ 0 map[]}

	name1 := "Matthew Sanabria"
	age1 := 29
	address1 := map[string]string{"street": "Main St."}

	resident1 := NewResident(name1, age1, address1)

	name2 := "Rob Pike"
	age2 := 0
	address2 := make(map[string]string)

	resident2 := NewResident(name2, age2, address2)

	residents := []*Resident{resident1, resident2}

	fmt.Println(Count(residents))
}

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{name, age, address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	return r.Name != "" && r.Address["street"] != ""
	/*if r.Name == "" {
		return false
	} else if value, exists := r.Address["street"]; exists {
		if value == "" {
			return false
		}
		fmt.Println(exists)
		return true
	}
	return false*/
	//return (r.Name != "" || _, exists := r.Address["street"]; exists)
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) (countResidents int) {
	for v := range residents {
		if residents[v].HasRequiredInfo() {
			countResidents++
		}
	}
	return countResidents
}
