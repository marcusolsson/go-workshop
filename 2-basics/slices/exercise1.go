// +build OMIT

package main

import "fmt"

func printKeyValues(kvs []string) {
	for i := 0; i < len(kvs); i += 2 {
		printKeyValue(kvs[i : i+2])
	}
}

func printKeyValue(kv []string) {
	fmt.Println(kv[0], "=", kv[1])
}

func main() {
	keyvals := []string{
		"id", "ABC123",
		"first_name", "Bruce",
		"last_name", "Wayne",
	}

	printKeyValues(keyvals)
}
