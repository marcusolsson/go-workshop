// +build OMIT

func main() {
	
	var customers map[string]int

	customers = make(map[string]int)

	customers = map[string]int{
		"a": 34,
		"b": 12,
		"c": 2,
	}

	v := customers["b"]

	v, ok := customers["b"]
	if !ok {
		fmt.Println("did not find value")
	}
}