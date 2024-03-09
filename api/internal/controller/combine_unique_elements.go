package controller

// CombineUnique extract a list of unique elements in two lists
func CombineUnique(list1, list2 []string) []string {
	combinedMap := make(map[string]bool)
	var combinedSlice []string

	for _, str := range list1 {
		combinedMap[str] = true
	}

	for _, str := range list2 {
		if !combinedMap[str] {
			combinedMap[str] = true
		}
	}

	for key := range combinedMap {
		combinedSlice = append(combinedSlice, key)
	}

	return combinedSlice
}
