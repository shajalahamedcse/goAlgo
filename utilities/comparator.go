package utilities

/*
-----------------------------------------------------------

IntegerComparator will make integer type assertion,
which will panic if a or b are not of the integer type.

-----------------------------------------------------------

Will Return:
	0 : if a == b
	1 : if a > b
   -1 : if a < b
------------------------------------------------------------

*/

func IntegerComparator(a, b interface{}) int {
	aInt := a.(int)
	bInt := b.(int)

	switch {
	case aInt > bInt:
		return 1
	case aInt < bInt:
		return -1
	default:
		return 0
	}
}

/*
-----------------------------------------------------------

StringComparator will make string type assertion,
which will panic if a or b are not of the string type.

-----------------------------------------------------------

Will Return:


------------------------------------------------------------

*/

func StringComparator(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)

	min := len(s2)

	if len(s1) < len(s2) {
		min = len(s1)
	}

	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}

	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	return diff
}
