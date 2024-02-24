//go:build goexperiment.rangefunc

package goiter

import (
	"math"
	"slices"
	"testing"
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

	// overflowing test 2: incremental iteration, stepSize causes addition overflow
	actualInt8 = make([]int8, 0)
	for each := range RangeStep(int8(120), int8(127), 10) {
		actualInt8 = append(actualInt8, each)
	}
	expectInt8 = []int8{120}
	if !slices.Equal(expectInt8, actualInt8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectInt8, actualInt8)
	}

	// overflowing test 3: decremental iteration, stepSize causes subtraction overflow
	actualInt8 = make([]int8, 0)
	for each := range RangeStep(int8(-120), int8(-128), 10) {
		actualInt8 = append(actualInt8, each)
	}
	expectInt8 = []int8{-120}
	if !slices.Equal(expectInt8, actualInt8) {
		t.Fatalf("test RangeStep failed, expect %d, got %v", expectInt8, actualInt8)
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
