package main

import "testing"

func expect(t *testing.T, b bool) {
	if !b {
		panic("failed")
	}
}

func NEXPCT(t *testing.T, b bool) {
	expect(t, !b)
}

func TestComamnds(t *testing.T) {
	do := func(v validations, s ...string) bool {
		v = v.copy()
		return checksOut(s, v)
	}

	v := newVal("m", "o")

	expect(t, do(v, "mi", "open"))
	expect(t, do(v, "mi", "open", "o"))
	expect(t, do(v, "mi", "o", "open"))
	expect(t, do(v, "mi", "o"))

	NEXPCT(t, do(v, "mi", "open", "ike"))
	NEXPCT(t, do(v, "mi", "ike"))
	NEXPCT(t, do(v, "jasima"))

	v = newVal("m", "o", "o")

	expect(t, do(v, "mi", "open", "o"))
	expect(t, do(v, "mi", "open", "o", "o", "o"))
	expect(t, do(v, "mi", "open", "olin", "o"))

}
