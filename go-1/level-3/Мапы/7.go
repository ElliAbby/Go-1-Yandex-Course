package main

import "fmt"

func SwapKeysAndValues(m map[string]string) map[string]string {
	newMap := make(map[string]string)

	for key, val := range m {
		newMap[val] = key
	}
	return newMap
}

func main() {
	m := map[string]string{
		"Яндекс": "+74957397000",
		"Музей Яндекса": "+74991101133",
	}
	fmt.Println(SwapKeysAndValues(m)) // map[+74957397000:Яндекс +74991101133:Музей Яндекса]
}
