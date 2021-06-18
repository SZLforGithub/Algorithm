package hashtable

func Example() map[string]float64 {
	book := make(map[string]float64)

	book["apple"] = 30.0
	book["milk"] = 45.0
	book["avocado"] = 60.5

	return book
}
