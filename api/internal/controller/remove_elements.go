package controller

func RemoveElements(originalSlice []string, stringsToDelete []string) []string {
	result := originalSlice[:0]

	for _, str := range originalSlice {
		if !contains(stringsToDelete, str) {
			result = append(result, str)
		}
	}
	return result
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
