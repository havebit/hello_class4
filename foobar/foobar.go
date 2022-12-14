package foobar

import "strconv"

func SayAny(n any) {
	var a int
	switch v := n.(type) {
	case int:
		a = v
	case string:
		a, _ = strconv.Atoi(v)
	default:
		return
	}

	Say(a)
}

func Say(n int) string {

	if n == 5 {
		return "Bar"
	}
	if n == 6 {
		return "Foo"
	}
	if n == 3 {
		return "Foo"
	}

	return strconv.Itoa(n)
}
