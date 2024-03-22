//go:build goexperiment.rangefunc

package goiter

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"testing"
	"time"
)

func TestRangeStep(t *testing.T) {
	var actual []int
	// equivalent to RangeStep(0, 3, 1)
	for each := range Range(0, 3) {
		actual = append(actual, each)
	}
	expect := []int{0, 1, 2}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expect, actual)
	}

	actual = make([]int, 0)
	for each := range RangeStep(0, 8, 2) {
		actual = append(actual, each)
	}
	expect = []int{0, 2, 4, 6}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expect, actual)
	}

	actual = make([]int, 0)
	i := 0
	for each := range RangeStep(0, 8, 2) {
		actual = append(actual, each)
		i++
		if i >= 3 {
			break
		}
	}
	expect = []int{0, 2, 4}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expect, actual)
	}

	actualUint8 := make([]uint8, 0)
	for each := range RangeStep(uint8(100), uint8(251), 50) {
		actualUint8 = append(actualUint8, each)
	}
	expectUint8 := []uint8{100, 150, 200, 250}
	if !slices.Equal(expectUint8, actualUint8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectUint8, actualUint8)
	}

	actualUint8 = make([]uint8, 0)
	for each := range RangeStep(uint8(0), uint8(251), 50) {
		actualUint8 = append(actualUint8, each)
	}
	expectUint8 = []uint8{0, 50, 100, 150, 200, 250}
	if !slices.Equal(expectUint8, actualUint8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectUint8, actualUint8)
	}

	actual = make([]int, 0)
	for each := range Range(3, -2) {
		actual = append(actual, each)
	}
	expect = []int{3, 2, 1, 0, -1}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expect, actual)
	}

	actual = make([]int, 0)
	for each := range RangeStep(8, -4, 2) {
		actual = append(actual, each)
	}
	expect = []int{8, 6, 4, 2, 0, -2}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expect, actual)
	}

	actualUint8 = make([]uint8, 0)
	for each := range RangeStep(uint8(201), uint8(0), 50) {
		actualUint8 = append(actualUint8, each)
	}
	expectUint8 = []uint8{201, 151, 101, 51, 1}
	if !slices.Equal(expectUint8, actualUint8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectUint8, actualUint8)
	}

	actualUint8 = make([]uint8, 0)
	for each := range RangeStep(uint8(255), uint8(0), 50) {
		actualUint8 = append(actualUint8, each)
	}
	expectUint8 = []uint8{255, 205, 155, 105, 55, 5}
	if !slices.Equal(expectUint8, actualUint8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectUint8, actualUint8)
	}

	// zero stepSize will lead to infinite loops, so it will not iterate
	actual = make([]int, 0)
	for each := range RangeStep(0, 5, 0) {
		actual = append(actual, each)
	}
	expect = []int{}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expect, actual)
	}

	// RangeStep does not accept negative stepSize, so it will not iterate
	actual = make([]int, 0)
	for each := range RangeStep(0, 5, -1) {
		actual = append(actual, each)
	}
	expect = []int{}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expect, actual)
	}

	// overflowing test 1: stepSize has value beyond the value range of T
	actualInt8 := make([]int8, 0)
	for each := range RangeStep(int8(0), int8(5), 256) {
		actualInt8 = append(actualInt8, each)
	}
	expectInt8 := []int8{0}
	if !slices.Equal(expectInt8, actualInt8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectInt8, actualInt8)
	}

	// overflowing test 2: produces increasing sequence
	actualInt8 = make([]int8, 0)
	for each := range RangeStep(int8(120), int8(127), 10) {
		actualInt8 = append(actualInt8, each)
	}
	expectInt8 = []int8{120}
	if !slices.Equal(expectInt8, actualInt8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectInt8, actualInt8)
	}

	// overflowing test 3: produces decreasing sequence
	actualInt8 = make([]int8, 0)
	for each := range RangeStep(int8(-120), int8(-128), 10) {
		actualInt8 = append(actualInt8, each)
	}
	expectInt8 = []int8{-120}
	if !slices.Equal(expectInt8, actualInt8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectInt8, actualInt8)
	}
}

func TestRangeTime(t *testing.T) {
	from, _ := time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ := time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual := make([]string, 0)
	for t := range RangeTime(from, to, 24*time.Hour) {
		actual = append(actual, t.Format(time.DateOnly))
	}
	expect := []string{"2024-01-01", "2024-01-02", "2024-01-03", "2024-01-04", "2024-01-05"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:10:00")
	to, _ = time.Parse(time.DateTime, "2024-01-01 03:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{"2024-01-01 00:10:00", "2024-01-01 01:10:00", "2024-01-01 02:10:00"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ = time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual = make([]string, 0)
	i := 0
	for t := range RangeTime(from, to, 24*time.Hour) {
		actual = append(actual, t.Format(time.DateOnly))
		i++
		if i >= 3 {
			break
		}
	}
	expect = []string{"2024-01-01", "2024-01-02", "2024-01-03"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 03:10:00")
	to, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{"2024-01-01 03:10:00", "2024-01-01 02:10:00", "2024-01-01 01:10:00", "2024-01-01 00:10:00"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 03:10:00")
	to, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	actual = make([]string, 0)
	i = 0
	for t := range RangeTime(from, to, time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
		i++
		if i >= 3 {
			break
		}
	}
	expect = []string{"2024-01-01 03:10:00", "2024-01-01 02:10:00", "2024-01-01 01:10:00"}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ = time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, time.Duration(0)) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}

	from, _ = time.Parse(time.DateTime, "2024-01-01 00:00:00")
	to, _ = time.Parse(time.DateTime, "2024-01-05 00:00:00")
	actual = make([]string, 0)
	for t := range RangeTime(from, to, -time.Hour) {
		actual = append(actual, t.Format(time.DateTime))
	}
	expect = []string{}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test RangeTime failed, expect %v, got %v", expect, actual)
	}
}

func TestSequence(t *testing.T) {
	// case 1
	genFib := func() GeneratorFunc[int] {
		a, b := 0, 1
		return func() (int, bool) {
			a, b = b, a+b
			return a, true
		}
	}
	actual := make([]int, 0, 10)
	i := 0
	for v := range Sequence(genFib()) {
		actual = append(actual, v)
		i++
		if i >= 10 {
			break
		}
	}
	expect := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test Sequence failed, expect %v, got %v", expect, actual)
	}

	// case 2
	genAlphabet := func(until rune) GeneratorFunc[string] {
		c := 'a'
		return func() (string, bool) {
			if c > until {
				return "", false
			}
			c++
			return string(c - 1), true
		}
	}
	actual2 := make([]string, 0, 10)
	for v := range Sequence(genAlphabet('g')) {
		actual2 = append(actual2, v)
	}
	expect2 := []string{"a", "b", "c", "d", "e", "f", "g"}
	if !slices.Equal(expect2, actual2) {
		t.Fatalf("test Sequence failed, expect %v, got %v", expect2, actual2)
	}
}

func TestSequence2(t *testing.T) {
	// case 1
	genFib := func() GeneratorFunc2[int, int] {
		n := 0
		a, b := 0, 1
		return func() (int, int, bool) {
			a, b = b, a+b
			n++
			return n, a, true
		}
	}

	actual := make([]int, 0, 10)
	for n, v := range Sequence2(genFib()) {
		actual = append(actual, v)
		if n >= 10 {
			break
		}
	}
	expect := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	if !slices.Equal(expect, actual) {
		t.Fatalf("test Sequence failed, expect %v, got %v", expect, actual)
	}

	// case 2
	genAlphabet := func(until rune) GeneratorFunc2[int, string] {
		n := 0
		c := 'a'
		return func() (int, string, bool) {
			if c > until {
				return 0, "", false
			}
			n++
			c++
			return n, string(c - 1), true
		}
	}
	actual2 := make([]string, 0, 10)
	for _, v := range Sequence2(genAlphabet('g')) {
		actual2 = append(actual2, v)
	}
	expect2 := []string{"a", "b", "c", "d", "e", "f", "g"}
	if !slices.Equal(expect2, actual2) {
		t.Fatalf("test Sequence failed, expect %v, got %v", expect2, actual2)
	}
}

func TestIntMax(t *testing.T) {
	if intMax(uint(0)) != uint(math.MaxUint) {
		t.Fatalf("test uint expect %d, got %v", uint(math.MaxUint), intMax(uint(0)))
	}
	if intMax(uint8(0)) != uint8(math.MaxUint8) {
		t.Fatalf("test uint8 expect %d, got %d", uint8(math.MaxUint8), intMax(uint8(0)))
	}
	if intMax(uint16(0)) != uint16(math.MaxUint16) {
		t.Fatalf("test uint16 expect %d, got %d", uint16(math.MaxUint16), intMax(uint16(0)))
	}
	if intMax(uint32(0)) != uint32(math.MaxUint32) {
		t.Fatalf("test uint32 expect %d, got %d", uint32(math.MaxUint32), intMax(uint32(0)))
	}
	if intMax(uint64(0)) != uint64(math.MaxUint64) {
		t.Fatalf("test uint64 expect %d, got %d", uint64(math.MaxUint64), intMax(uint64(0)))
	}
	if intMax(int(0)) != math.MaxInt {
		t.Fatalf("test int expect %d, got %d", math.MaxInt, intMax(int(0)))
	}
	if intMax(int8(0)) != int8(math.MaxInt8) {
		t.Fatalf("test int8 expect %d, got %d", int8(math.MaxInt8), intMax(int8(0)))
	}
	if intMax(int16(0)) != int16(math.MaxInt16) {
		t.Fatalf("test int16 expect %d, got %d", int16(math.MaxInt16), intMax(int16(0)))
	}
	if intMax(int32(0)) != int32(math.MaxInt32) {
		t.Fatalf("test int32 expect %d, got %d", int32(math.MaxInt32), intMax(int32(0)))
	}
	if intMax(int64(0)) != int64(math.MaxInt64) {
		t.Fatalf("test int64 expect %d, got %d", int64(math.MaxInt64), intMax(int64(0)))
	}
}

func TestIntMin(t *testing.T) {
	if intMin(uint(0)) != 0 {
		t.Fatalf("test uint expect %d, got %v", 0, intMin(uint(0)))
	}
	if intMin(uint8(0)) != 0 {
		t.Fatalf("test uint8 expect %d, got %d", 0, intMin(uint8(0)))
	}
	if intMin(uint16(0)) != 0 {
		t.Fatalf("test uint16 expect %d, got %d", 0, intMin(uint16(0)))
	}
	if intMin(uint32(0)) != 0 {
		t.Fatalf("test uint32 expect %d, got %d", 0, intMin(uint32(0)))
	}
	if intMin(uint64(0)) != 0 {
		t.Fatalf("test uint64 expect %d, got %d", 0, intMin(uint64(0)))
	}
	if intMin(int(0)) != math.MinInt {
		t.Fatalf("test int expect %d, got %d", math.MinInt, intMin(int(0)))
	}
	if intMin(int8(0)) != int8(math.MinInt8) {
		t.Fatalf("test int8 expect %d, got %d", int8(math.MinInt8), intMin(int8(0)))
	}
	if intMin(int16(0)) != int16(math.MinInt16) {
		t.Fatalf("test int16 expect %d, got %d", int16(math.MinInt16), intMin(int16(0)))
	}
	if intMin(int32(0)) != int32(math.MinInt32) {
		t.Fatalf("test int32 expect %d, got %d", int32(math.MinInt32), intMin(int32(0)))
	}
	if intMin(int64(0)) != int64(math.MinInt64) {
		t.Fatalf("test int64 expect %d, got %d", int64(math.MinInt64), intMin(int64(0)))
	}
}

func TestConcat(t *testing.T) {
	c1 := []int{1, 2, 3}
	c2 := []int{4, 5, 6}
	actual := make([]int, 0, 6)
	for v := range Concat(SliceElem(c1), SliceElem(c2)) {
		actual = append(actual, v)
	}
	expect := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	c1 = []int{4, 5, 6}
	c2 = []int{7, 8, 9}
	actual = make([]int, 0, 6)
	for v := range Concat(SliceElem(c1), SliceElem(c2)) {
		if v == 8 {
			break
		}
		actual = append(actual, v)
	}
	expect = []int{4, 5, 6, 7}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestConcat2(t *testing.T) {
	type person struct {
		name string
		age  int
	}
	p1 := []person{{"john", 25}, {"jane", 20}}
	i1 := Transform12(SliceElem(p1), func(p person) (string, int) {
		return p.name, p.age
	})
	p2 := []person{{"joe", 35}, {"ann", 30}, {"josh", 15}}
	i2 := Transform12(SliceElem(p2), func(p person) (string, int) {
		return p.name, p.age
	})

	actual := make(map[string]int)
	for name, age := range Concat2(i1, i2) {
		actual[name] = age
	}
	expect := map[string]int{"john": 25, "jane": 20, "joe": 35, "ann": 30, "josh": 15}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make(map[string]int)
	for name, age := range Concat2(i1, i2) {
		if name == "ann" {
			break
		}
		actual[name] = age
	}
	expect = map[string]int{"john": 25, "jane": 20, "joe": 35}
	if !maps.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	actual := make([]int, 0, 5)
	for v := range Reverse(SliceElem(input)) {
		actual = append(actual, v)
	}
	expect := []int{5, 4, 3, 2, 1}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make([]int, 0, 3)
	for v := range Reverse(SliceElem(input)) {
		if v < 3 {
			break
		}
		actual = append(actual, v)
	}
	expect = []int{5, 4, 3}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}

func TestReverse2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	actual := make([]Tuple[int, int], 0, 5)
	for idx, v := range Reverse2(Slice(input)) {
		actual = append(actual, Tuple[int, int]{V1: idx, V2: v})
	}
	expect := []Tuple[int, int]{
		{V1: 4, V2: 5},
		{V1: 3, V2: 4},
		{V1: 2, V2: 3},
		{V1: 1, V2: 2},
		{V1: 0, V2: 1},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}

	actual = make([]Tuple[int, int], 0, 3)
	for idx, v := range Reverse2(Slice(input)) {
		if v < 3 {
			break
		}
		actual = append(actual, Tuple[int, int]{V1: idx, V2: v})
	}
	expect = []Tuple[int, int]{
		{V1: 4, V2: 5},
		{V1: 3, V2: 4},
		{V1: 2, V2: 3},
	}
	if !slices.Equal(expect, actual) {
		t.Fatal(fmt.Sprintf("expect: %v, actual: %v", expect, actual))
	}
}
