package sublist

type Relation string

const (
	RelationEqual     Relation = "equal"
	RelationSublist   Relation = "sublist"
	RelationSuperlist Relation = "superlist"
	RelationUnequal   Relation = "unequal"
)

func findRelation(smaller []int, bigger []int, expected Relation) Relation {
	if len(smaller) == 0 {
		return expected
	}

	var indexes []int

	for index, element := range bigger {
		if element == smaller[0] {
			indexes = append(indexes, index)
		}
	}

	if len(indexes) == 0 {
		return RelationUnequal
	}

	for _, element := range indexes {
		var equalElements []int

		for i := 0; i < len(smaller); i++ {
			if (element+i) < len(bigger) && smaller[i] == bigger[element+i] {
				equalElements = append(equalElements, smaller[i])
			} else {
				break
			}
		}

		if len(equalElements) == len(smaller) {
			return expected
		}
	}

	return RelationUnequal
}

func Sublist(listA []int, listB []int) Relation {
	if len(listA) == len(listB) {
		for index, e := range listA {
			if e != listB[index] {
				return RelationUnequal
			}
		}

		return RelationEqual
	}

	// Check if Sublist
	if len(listA) < len(listB) {
		return findRelation(listA, listB, RelationSublist)
	}

	// Check if Superlist
	return findRelation(listB, listA, RelationSuperlist)
}
