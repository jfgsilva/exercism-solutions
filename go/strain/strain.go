package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var slcInt Ints
	for _, val := range i {
		if filter(val) {
			slcInt = append(slcInt, val)
		}
	}
	return slcInt
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var slcInt Ints
	for _, val := range i {
		if !filter(val) {
			slcInt = append(slcInt, val)
		}
	}
	return slcInt
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var newLists Lists
	for _, innerList := range l {
		if filter(innerList) {
			newLists = append(newLists, innerList)
		}
	}
	return newLists
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var str Strings
	for _, innerStr := range s {
		if filter(innerStr) {
			str = append(str, innerStr)
		}
	}
	return str
}
