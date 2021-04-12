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

	v = newVal("u", "u", "j")

	expect(t, do(v, "unu", "li", "usawi", "e", "jan"))
	expect(t, do(v, "uta", "unu", "li", "jami"))

	NEXPCT(t, do(v, "ike", "nasa"))

	v = newVal("l", "p", "i")

	expect(t, do(v, "li", "poka", "ipi"))

	v = newVal("l", "o", "a")

	expect(t, do(v, "li", "o", "a"))
}
