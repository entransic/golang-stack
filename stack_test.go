package stack_test

import (
	"math"
	"testing"
	"unicode/utf8"

	adt "entransic.com/stack"
)

func TestIsEmpty(t *testing.T) {

	stack := &adt.Stack[string]{}
	new := stack.NewStack()
	got := new.IsEmpty()
	want := true

	if got != want {
		t.Errorf("got %t, want true", got)
	}

}

func TestNotEmpty(t *testing.T) {

	elem := "Bob"

	new := &adt.Stack[string]{}
	stack := new.NewStack()
	_ = stack.Push(elem)
	got := stack.IsEmpty()
	want := false

	if got != want {
		t.Errorf("got %t, want false", got)
	}
}

func TestStackSizeZero(t *testing.T) {
	new := &adt.Stack[string]{}
	stack := new.NewStack()
	got := stack.Size()
	want := 0
	if got != want {
		t.Errorf("got %d, want  %d", got, want)
	}
}

func TestStackSizeOne(t *testing.T) {
	elem := "Bob"
	new := &adt.Stack[string]{}
	stack := new.NewStack()
	stack.Push(elem)
	got := stack.Size()
	want := 1

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestStackSizeThree(t *testing.T) {
	elem1 := "Bob"
	elem2 := "The"
	elem3 := "Fish"

	new := &adt.Stack[string]{}
	stack := new.NewStack()
	stack.Push(elem1)
	stack.Push(elem2)
	stack.Push(elem3)

	got := stack.Size()
	want := 3

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPopOne(t *testing.T) {
	elem := "Bob"
	new := &adt.Stack[string]{}
	stack := new.NewStack()
	stack.Push(elem)
	got, _ := stack.Pop()
	want := elem

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	if stack.IsEmpty() != true {
		t.Errorf("Stack not empty")
	}
}

func TestPushThreePopThree(t *testing.T) {

	elem1 := "Bob"
	elem2 := "The"
	elem3 := "Fish"

	new := &adt.Stack[string]{}
	stack := new.NewStack()
	stack.Push(elem1)
	stack.Push(elem2)
	stack.Push(elem3)

	got, _ := stack.Pop()
	want := elem3

	if got != want {
		t.Errorf("got1 %s, want %s", got, want)
	}

	got2, _ := stack.Pop()
	want2 := elem2

	if got2 != want2 {
		t.Errorf("got2 %s, want %s", got2, want2)
	}

	got3, _ := stack.Pop()
	want3 := elem1

	if got3 != want3 {
		t.Errorf("got3 %s, want %s", got3, want3)
	}

}

func TestStackFull(t *testing.T) {

	elem := "Bob"

	new := &adt.Stack[string]{}
	stack := new.NewStack()

	for range 4 {
		_ = stack.Push(elem)
	}

	got := stack.Push(elem)

	if got == nil {
		t.Errorf("Did not receive error")
	}

}

func TestPopEmpty(t *testing.T) {
	new := &adt.Stack[string]{}
	stack := new.NewStack()

	got, err := stack.Pop()
	want := ""

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestEmpty(t *testing.T) {

	new := &adt.Stack[string]{}
	stack := new.NewStack()

	stack.Push("Bob")
	stack.Empty()

	got := stack.Size()
	want := 0

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	_, err := stack.Pop()

	if err == nil {
		t.Error("Didn't get an error popping empty stack")
	}
}

func TestIntegerStack(t *testing.T) {

	new := &adt.Stack[int]{}
	stack := new.NewStack()
	want := 5

	stack.Push(want)
	got, err := stack.Pop()

	if err != nil {
		t.Error("error received")
	}
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}

}

func TestFloat32Stack(t *testing.T) {
	new := &adt.Stack[float32]{}
	stack := new.NewStack()
	var want float32 = math.Pi

	stack.Push(want)
	got, err := stack.Pop()

	if err != nil {
		t.Error("error received")
	}
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

func TestFloat64Stack(t *testing.T) {
	new := &adt.Stack[float64]{}
	stack := new.NewStack()
	var want float64 = math.Pi

	stack.Push(want)
	got, err := stack.Pop()

	if err != nil {
		t.Error("error received")
	}
	if got != want {
		t.Errorf("got %f, want %f", got, want)
	}
}

// BENACHMARK TESTS

func BenchmarkNewPushPop(b *testing.B) {

	for b.Loop() {
		new := &adt.Stack[string]{}
		stack := new.NewStack()
		element := "Bob"
		_ = stack.Push(element)
		_, _ = stack.Pop()
	}
}

func BenchmarkIsEmpty(b *testing.B) {

	for b.Loop() {
		new := &adt.Stack[string]{}
		stack := new.NewStack()

		stack.IsEmpty()
	}
}

func BenchmarkEmpty(b *testing.B) {
	for b.Loop() {
		new := &adt.Stack[string]{}
		stack := new.NewStack()
		stack.Empty()
	}
}

// FUZZ TESTS

func FuzzPush(f *testing.F) {
	testcases := []string{"Element", "", "0", "!@#$%^&*()"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, s string) {
		new := &adt.Stack[string]{}
		stack := new.NewStack()
		got := stack.Push(s)
		if got != nil {
			t.Errorf("got an error")
		}

	})
}

func FuzzPop(f *testing.F) {
	testcases := []string{"Element", "", "0", "!@#$%^&*()"}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, want string) {
		new := &adt.Stack[string]{}
		stack := new.NewStack()
		_ = stack.Push(want)
		got, _ := stack.Pop()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

		if utf8.ValidString(want) && !utf8.ValidString(got) {
			t.Errorf("Pop produced invalid UTF-8 string %q", got)
		}
	})
}
