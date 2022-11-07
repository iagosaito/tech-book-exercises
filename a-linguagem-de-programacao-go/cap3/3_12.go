package main

import "fmt"

func main() {

	fmt.Println(isAnagrams("amor", "roma"))
	fmt.Println(isAnagrams("amsor", "roma"))
	fmt.Println(isAnagrams("yara", "raay"))
	fmt.Println(isAnagrams("ana", "anana"))
}

func isAnagrams(s1 string, s2 string) bool {

	f1 := make(map[string]int)
	f2 := make(map[string]int)

	for _, v := range s1 {
		f1[string(v)]++
	}

	for _, v := range s2 {
		f2[string(v)]++
	}

	for k, v := range f1 {
		v2 := f2[k]

		if v2 != v {
			return false
		}
	}

	return true
}
