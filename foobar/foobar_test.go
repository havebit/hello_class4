package foobar_test

import (
	"testing"

	"github.com/pallat/hello_api_class4/foobar"
)

func TestGivenOneWantOne(t *testing.T) {
	given := 1
	want := "1"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q; want %s", given, result, want)
	}
}
func TestGivenTwoWantTwo(t *testing.T) {
	given := 2
	want := "2"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q; want %s", given, result, want)
	}
}
func TestGivenThreeWantFoo(t *testing.T) {
	given := 3
	want := "Foo"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q; want %s", given, result, want)
	}
}
func TestGivenFourWantFour(t *testing.T) {
	given := 4
	want := "4"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q; want %s", given, result, want)
	}
}
func TestGivenFiveWantBar(t *testing.T) {
	given := 5
	want := "Bar"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q; want %s", given, result, want)
	}
}
func TestGivenSixWantFoo(t *testing.T) {
	given := 6
	want := "Foo"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q; want %s", given, result, want)
	}
}
func TestGivenSevenWantSeven(t *testing.T) {
	given := 7
	want := "7"

	result := foobar.Say(given)

	if result != want {
		t.Errorf("Say(%d) = %q; want %s", given, result, want)
	}
}
