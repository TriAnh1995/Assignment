package controller

func ExtractCommonElements(list1 []string, list2 []string) (commonList []string) {
	elementsInCommon := make(map[string]bool)

	for _, element := range list1 {
		elementsInCommon[element] = true
	}

	for _, element := range list2 {
		if elementsInCommon[element] {
			commonList = append(commonList, element)
		}
	}
	return commonList
}
