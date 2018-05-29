package main

import "testing"

func Test_diff(t *testing.T) {
	orig := []string{"a"}
	new := []string{"b", "a"}
	expected := []string{"b"}

	actual := getDiff(orig, new)

	if len(actual) != len(expected) {
		t.Fatalf("length: %d != %d", len(expected), len(actual))
	}

	if actual[0] != expected[0] {
		t.Fatalf("%v != %v", actual, expected)
	}

}
