package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string // her må du skrive riktig type for input
		want  int    // her må du skrive riktig type for returverdien
	}

	var tests = []test{ // Her må du legge inn korrekte testverdier
		{Sun, 0}, //tests := []test{
		{1, 1},   //  {input: , want: },
		{2, 1},   //}
		{3, 2},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got) // tc.input,
		}
	}
}
