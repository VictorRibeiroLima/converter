package converter_test

import (
	"testing"

	"github.com/VictorRibeiroLima/converter"
)

func TestEmbeddedArray(t *testing.T) {
	type FromSecondLayer struct {
		SName string
		SID   uint
	}
	type FromFirstLayer struct {
		Name   string
		ID     uint
		Second []FromSecondLayer
	}
	type From struct {
		First []FromFirstLayer
	}

	type ToSecondLayer struct {
		SName string
		SID   uint
	}
	type ToFirstLayer struct {
		Name   string
		ID     uint
		Second []ToSecondLayer
	}
	type To struct {
		First []ToFirstLayer
	}

	fs1 := FromSecondLayer{
		SName: "1",
		SID:   1,
	}
	fs2 := FromSecondLayer{
		SName: "2",
		SID:   2,
	}

	ff1 := FromFirstLayer{
		Name: "1",
		ID:   1,
	}

	ff2 := FromFirstLayer{
		Name:   "2",
		ID:     2,
		Second: []FromSecondLayer{fs1, fs2},
	}

	ff3 := FromFirstLayer{
		Name: "3",
		ID:   3,
	}

	f := From{
		First: []FromFirstLayer{ff1, ff2, ff3},
	}

	var to To
	converter.Convert(&to, f)

	if to.First[0].Second != nil {
		t.Error("Expected First[0].Second to be nil. instead got a value \n")
	}
	if to.First[1].Second == nil {
		t.Error("Expected First[1].Second to have a value. instead got nil \n")
	}
	if len(to.First[1].Second) != 2 {
		t.Errorf("Expected First[1].Second to have a len of 2. instead got %d \n", len(to.First[1].Second))
	}
	if to.First[2].Second != nil {
		t.Error("Expected First[2].Second to be nil. instead got a value \n")
	}
}

func TestSimpleTypeConversion(t *testing.T) {
	type To struct {
		A string
		B int
		C string
		P *string
		e string
	}

	type From struct {
		A string
		B int
		D string
		P *string
		e string
	}
	p := "asdasd"
	from := From{
		A: "test",
		B: 1,
		D: "test",
		P: &p,
		e: "test",
	}

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	if to.A != from.A {
		t.Errorf("Property 'A' expected to be %s. instead got %s", from.A, to.A)
	}

	if to.B != from.B {
		t.Errorf("Property 'B' expected to be %d. instead got %d", from.B, to.B)
	}

	if *to.P != *from.P {
		t.Errorf("Property 'P' expected to be %s. instead got %s", *from.P, *to.P)
	}

	if to.C != "" {
		t.Errorf("Property 'C' should be empty. instead got %s", to.C)
	}

	if to.e != "" {
		t.Errorf("Property 'e' should be empty. instead got %s", to.e)
	}
}

func TestNestedStructTypeConversion(t *testing.T) {

	type NestedTo struct {
		A string
		B int
		C string
	}

	type To struct {
		A      string
		B      int
		C      string
		e      string
		Nested NestedTo
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type From struct {
		A      string
		B      int
		D      string
		e      string
		Nested NestedFrom
	}
	from := From{
		A: "test",
		B: 1,
		D: "test",
		e: "test",
		Nested: NestedFrom{
			A: "test",
			B: 1,
			D: "Test",
		},
	}

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	if to.A != from.A {
		t.Errorf("Property 'A' expected to be %s. instead got %s", from.A, to.A)
	}

	if to.B != from.B {
		t.Errorf("Property 'B' expected to be %d. instead got %d", from.B, to.B)
	}

	if to.C != "" {
		t.Errorf("Property 'C' should be empty. instead got %s", to.C)
	}

	if to.e != "" {
		t.Errorf("Property 'e' should be empty. instead got %s", to.e)
	}

	if to.Nested.A != from.Nested.A {
		t.Errorf("Property 'A' expected to be %s. instead got %s", from.Nested.A, to.Nested.A)
	}

	if to.Nested.B != from.Nested.B {
		t.Errorf("Property 'B' expected to be %d. instead got %d", from.Nested.B, to.Nested.B)
	}

	if to.Nested.C != "" {
		t.Errorf("Property 'C' should be empty. instead got %s", to.Nested.C)
	}
}

func TestArrayStructTypeConversion(t *testing.T) {

	type NestedTo struct {
		A string
		B int
		C string
	}

	type To struct {
		A            string
		B            int
		C            string
		e            string
		Nesteds      []NestedTo
		Array        []int
		NotMachArray []string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type From struct {
		A            string
		B            int
		D            string
		e            string
		Nesteds      []NestedFrom
		Array        []int
		NotMachArray []int
	}
	from := From{
		A:            "test",
		B:            1,
		D:            "test",
		e:            "test",
		Array:        []int{1, 2, 3},
		NotMachArray: []int{4, 5, 6},
	}

	from.Nesteds = append(from.Nesteds, NestedFrom{
		A: "test",
		B: 1,
		D: "test",
	}, NestedFrom{
		A: "test",
		B: 2,
		D: "test",
	})
	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	if to.A != from.A {
		t.Errorf("Property 'A' expected to be %s. instead got %s", from.A, to.A)
	}

	if to.B != from.B {
		t.Errorf("Property 'B' expected to be %d. instead got %d", from.B, to.B)
	}

	if to.C != "" {
		t.Errorf("Property 'C' should be empty. instead got %s", to.C)
	}

	if to.e != "" {
		t.Errorf("Property 'e' should be empty. instead got %s", to.e)
	}

	for i := range from.Nesteds {
		nested := to.Nesteds[i]
		fromNested := from.Nesteds[i]
		if nested.A != fromNested.A {
			t.Errorf("Property 'A' expected to be %s. instead got %s", fromNested.A, nested.A)
		}

		if nested.B != fromNested.B {
			t.Errorf("Property 'B' expected to be %d. instead got %d", fromNested.B, nested.B)
		}

		if nested.C != "" {
			t.Errorf("Property 'C' should be empty. instead got %s", nested.C)
		}
	}

	for i := range from.Array {
		value := to.Array[i]
		fromValue := from.Array[i]
		if value != fromValue {
			t.Errorf("Property 'A' expected to be %d. instead got %d", value, fromValue)
		}
	}

	if len(to.NotMachArray) > 0 {
		t.Errorf("Property 'NotMachArray' should be empty. instead got %d", len(to.NotMachArray))
	}
}

func TestValueToPointerTypeConversion(t *testing.T) {
	type NestedTo struct {
		A string
		B int
		C string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type To struct {
		P      *string
		Nested *NestedTo
		Hp     *int
		e      string
	}

	type From struct {
		P      string
		Nested NestedFrom
		Hp     string
		e      string
	}
	from := From{

		P: "asdasd",
		Nested: NestedFrom{
			A: "TEST",
			B: 1,
			D: "FASFAs",
		},
		e: "test",
	}

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	if *to.P != from.P {
		t.Errorf("Property 'P' expected to be %s. instead got %s", from.P, *to.P)
	}

	if to.e != "" {
		t.Errorf("Property 'e' should be empty. instead got %s", to.e)
	}

	if (*to.Nested).A != from.Nested.A {
		t.Errorf("Property 'A' expected to be %s. instead got %s", from.Nested.A, to.Nested.A)
	}

	if (*to.Nested).B != from.Nested.B {
		t.Errorf("Property 'B' expected to be %d. instead got %d", from.Nested.B, to.Nested.B)
	}

	if (*to.Nested).C != "" {
		t.Errorf("Property 'C' should be empty. instead got %s", to.Nested.C)
	}

	if to.Hp != nil {
		t.Errorf("Property 'Hp' shold be nil. instead got %d", *to.Hp)
	}
}

func TestValueToPointerFromPointerTypeConversion(t *testing.T) {
	type NestedTo struct {
		A string
		B int
		C string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type To struct {
		P      *string
		Nested *NestedTo
		Hp     *int
		e      string
	}

	type From struct {
		P      string
		Nested *NestedFrom
		Hp     string
		e      string
	}

	nested := NestedFrom{
		A: "TEST",
		B: 1,
		D: "FASFAs",
	}

	from := From{

		P:      "asdasd",
		Nested: &nested,
		e:      "test",
	}

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	if *to.P != from.P {
		t.Errorf("Property 'P' expected to be %s. instead got %s", from.P, *to.P)
	}

	if to.e != "" {
		t.Errorf("Property 'e' should be empty. instead got %s", to.e)
	}

	if (*to.Nested).A != from.Nested.A {
		t.Errorf("Property 'A' expected to be %s. instead got %s", from.Nested.A, to.Nested.A)
	}

	if (*to.Nested).B != from.Nested.B {
		t.Errorf("Property 'B' expected to be %d. instead got %d", from.Nested.B, to.Nested.B)
	}

	if (*to.Nested).C != "" {
		t.Errorf("Property 'C' should be empty. instead got %s", to.Nested.C)
	}

	if to.Hp != nil {
		t.Errorf("Property 'Hp' shold be nil. instead got %d", *to.Hp)
	}
}

func TestValueToPointerFromPointerNillTypeConversion(t *testing.T) {
	type NestedTo struct {
		A string
		B int
		C string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type To struct {
		P      *string
		Nested *NestedTo
		Hp     *int
		e      string
	}

	type From struct {
		P      string
		Nested *NestedFrom
		Hp     string
		e      string
	}

	from := From{

		P:      "asdasd",
		Nested: nil,
		e:      "test",
	}

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	if *to.P != from.P {
		t.Errorf("Property 'P' expected to be %s. instead got %s", from.P, *to.P)
	}

	if to.e != "" {
		t.Errorf("Property 'e' should be empty. instead got %s", to.e)
	}

	if to.Nested != nil {
		t.Errorf("Property 'Nested' shold be nil")
	}

	if to.Hp != nil {
		t.Errorf("Property 'Hp' shold be nil. instead got %d", *to.Hp)
	}
}

func TestValueToPointerArrayTypeConversion(t *testing.T) {
	type NestedTo struct {
		A string
		B int
		C string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type To struct {
		Nested       []*NestedTo
		NotMachArray []*string
	}

	type From struct {
		Nested       []NestedFrom
		NotMachArray []int
	}
	from := From{
		NotMachArray: []int{1, 2},
	}

	from.Nested = append(from.Nested, NestedFrom{
		A: "TEST",
		B: 1,
		D: "FASFAs",
	}, NestedFrom{
		A: "TEST",
		B: 2,
		D: "FASFAs",
	})

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}
	for i := range from.Nested {
		nested := to.Nested[i]
		fromNested := from.Nested[i]
		if (*nested).A != fromNested.A {
			t.Errorf("Property 'A' expected to be %s. instead got %s", fromNested.A, (*nested).A)
		}

		if (*nested).B != fromNested.B {
			t.Errorf("Property 'B' expected to be %d. instead got %d", fromNested.B, (*nested).B)
		}

		if (*nested).C != "" {
			t.Errorf("Property 'C' should be empty. instead got %s", (*nested).C)
		}
	}
	if len(to.NotMachArray) > 0 {
		t.Errorf("Property 'NotMachArray' should be empty. instead got %d", len(to.NotMachArray))
	}
}

func TestValueToPointerArrayPointerTypeConversiont(t *testing.T) {
	type NestedTo struct {
		A string
		B int
		C string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type To struct {
		Nested []*NestedTo
	}

	type From struct {
		Nested []*NestedFrom
	}
	from := From{}

	from.Nested = append(from.Nested, &NestedFrom{
		A: "TEST",
		B: 1,
		D: "FASFAs",
	}, &NestedFrom{
		A: "TEST",
		B: 2,
		D: "FASFAs",
	})

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	for i := range from.Nested {
		fromValue := *from.Nested[i]
		toValue := *to.Nested[i]

		if toValue.A != fromValue.A {
			t.Errorf("Property 'A' expected to be %s. instead got %s", fromValue.A, toValue.A)
		}
		if toValue.B != fromValue.B {
			t.Errorf("Property 'B' expected to be %d. instead got %d", fromValue.B, toValue.B)
		}
		if toValue.C != "" {
			t.Errorf("Property 'C' expected to be empty. instead got %s", toValue.C)
		}
	}
}

func TestPointerToValueTypeConversiont(t *testing.T) {
	type NestedTo struct {
		A string
		B int
		C string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type To struct {
		P      string
		Nested NestedTo
		e      string
	}

	type From struct {
		P      *string
		Nested *NestedFrom
		e      string
	}
	p := "afasfas"
	n := &NestedFrom{
		A: "Test",
		B: 1,
		D: "asfasfas",
	}
	from := From{
		P:      &p,
		Nested: n,
		e:      "test",
	}

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}

	if to.P != *from.P {
		t.Errorf("Property 'P' expected to be %s. instead got %s", to.P, *from.P)
	}
	if to.Nested.A != (*from.Nested).A {
		t.Errorf("Property 'A' expected to be %s. instead got %s", (*from.Nested).A, to.Nested.A)
	}

	if to.Nested.B != (*from.Nested).B {
		t.Errorf("Property 'B' expected to be %d. instead got %d", (*from.Nested).B, to.Nested.B)
	}

	if to.Nested.C != "" {
		t.Errorf("Property 'C' should be empty. instead got %s", to.Nested.C)
	}
}

func TestPoiterToValueArrayTypeConversion(t *testing.T) {
	type NestedTo struct {
		A string
		B int
		C string
	}

	type NestedFrom struct {
		A string
		B int
		D string
	}

	type To struct {
		Nested       []NestedTo
		NotMachArray []string
	}

	type From struct {
		Nested       []*NestedFrom
		NotMachArray []*int
	}
	from := From{}

	from.Nested = append(from.Nested, &NestedFrom{
		A: "TEST",
		B: 1,
		D: "FASFAs",
	}, &NestedFrom{
		A: "TEST",
		B: 2,
		D: "FASFAs",
	})

	n1 := 1
	n2 := 2
	from.NotMachArray = append(from.NotMachArray, &n1, &n2)

	var to To

	err := converter.Convert(&to, from)

	if err != nil {
		t.Error("Simple conversion error")
	}
	for i := range from.Nested {
		nested := to.Nested[i]
		fromNested := *from.Nested[i]
		if nested.A != fromNested.A {
			t.Errorf("Property 'A' expected to be %s. instead got %s", fromNested.A, nested.A)
		}

		if nested.B != fromNested.B {
			t.Errorf("Property 'B' expected to be %d. instead got %d", fromNested.B, nested.B)
		}

		if nested.C != "" {
			t.Errorf("Property 'C' should be empty. instead got %s", nested.C)
		}
	}
	if len(to.NotMachArray) > 0 {
		t.Errorf("Property 'NotMachArray' should be empty. instead got %d", len(to.NotMachArray))
	}
}

func TestArrayConversion(t *testing.T) {
	type FromSecondLayer struct {
		SName string
		SID   uint
	}
	type FromFirstLayer struct {
		Name   string
		ID     uint
		Second []FromSecondLayer
	}
	type ToSecondLayer struct {
		SName string
		SID   uint
	}
	type ToFirstLayer struct {
		Name   string
		ID     uint
		Second []ToSecondLayer
	}

	fs1 := FromSecondLayer{
		SName: "1",
		SID:   1,
	}
	fs2 := FromSecondLayer{
		SName: "2",
		SID:   2,
	}

	ff1 := FromFirstLayer{
		Name: "1",
		ID:   1,
	}

	ff2 := FromFirstLayer{
		Name:   "2",
		ID:     2,
		Second: []FromSecondLayer{fs1, fs2},
	}

	ff3 := FromFirstLayer{
		Name: "3",
		ID:   3,
	}
	from := []FromFirstLayer{ff1, ff2, ff3}
	var to []ToFirstLayer
	err := converter.Convert(&to, from)
	if err != nil {
		t.Errorf("Expected err to be nil. instead got %s", err)
	}
	for i := range to {
		fromValue := from[i]
		toValue := to[i]
		if toValue.ID != fromValue.ID {
			t.Errorf("Expected ID to equal %d. instead got %d", fromValue.ID, toValue.ID)
		}
		if toValue.Name != fromValue.Name {
			t.Errorf("Expected ID to equal %s. instead got %s", fromValue.Name, toValue.Name)
		}
	}
	if to[0].Second != nil {
		t.Error("Expected to[0].Second to be nil. instead got a value \n")
	}
	if to[1].Second == nil {
		t.Error("Expected to[1].Second to have a value. instead got nil \n")
	}
	if len(to[1].Second) != 2 {
		t.Errorf("Expected to[1].Second to have a len of 2. instead got %d \n", len(to[1].Second))
	}
	if to[2].Second != nil {
		t.Error("Expected to[2].Second to be nil. instead got a value \n")
	}
}

func TestIncompatibleType(t *testing.T) {
	type A struct {
		Name string
	}
	type B struct {
		Name string
	}
	t.Run("to array", func(t *testing.T) {
		a := A{Name: "1"}
		var b []B
		err := converter.Convert(&b, a)
		if err == nil {
			t.Fatal("Expected err to have a value. instead got nil")
		}

		if err.Error() != "incompatible types" {
			t.Errorf("Expected err to equal 'incompatible types'. instead got '%s'", err)
		}

	})
	t.Run("to struct", func(t *testing.T) {
		a := []A{{Name: "1"}}
		var b B
		err := converter.Convert(&b, a)
		if err == nil {
			t.Fatal("Expected err to have a value. instead got nil")
		}

		if err.Error() != "incompatible types" {
			t.Errorf("Expected err to equal 'incompatible types'. instead got '%s'", err)
		}
	})
}
