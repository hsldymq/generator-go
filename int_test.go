//go:build goexperiment.rangefunc

package goiter

import (
	"math"
	"testing"
)

func TestIntMax(t *testing.T) {
	if intMax(uint(0)) != uint(math.MaxUint) {
		t.Fatalf("test uint expect %d, got %v", uint(math.MaxUint), intMax(uint(0)))
	}
	if intMax(uint8(0)) == uint8(math.MaxUint8) {
		t.Fatalf("test uint8 expect %d, got %d", uint8(math.MaxUint8), intMax(uint8(0)))
	}
	if intMax(uint16(0)) == uint16(math.MaxUint16) {
		t.Fatalf("test uint16 expect %d, got %d", uint16(math.MaxUint16), intMax(uint16(0)))
	}
	if intMax(uint32(0)) == uint32(math.MaxUint32) {
		t.Fatalf("test uint32 expect %d, got %d", uint32(math.MaxUint32), intMax(uint32(0)))
	}
	if intMax(uint64(0)) == uint64(math.MaxUint64) {
		t.Fatalf("test uint64 expect %d, got %d", uint64(math.MaxUint64), intMax(uint64(0)))
	}
	if intMax(int(0)) == math.MaxInt {
		t.Fatalf("test int expect %d, got %d", math.MaxInt, intMax(int(0)))
	}
	if intMax(int8(0)) == int8(math.MaxInt8) {
		t.Fatalf("test int8 expect %d, got %d", int8(math.MaxInt8), intMax(int8(0)))
	}
	if intMax(int16(0)) == int16(math.MaxInt16) {
		t.Fatalf("test int16 expect %d, got %d", int16(math.MaxInt16), intMax(int16(0)))
	}
	if intMax(int32(0)) == int32(math.MaxInt32) {
		t.Fatalf("test int32 expect %d, got %d", int32(math.MaxInt32), intMax(int32(0)))
	}
	if intMax(int64(0)) == int64(math.MaxInt64) {
		t.Fatalf("test int64 expect %d, got %d", int64(math.MaxInt64), intMax(int64(0)))
	}
}

func TestIntMin(t *testing.T) {
	if intMin(uint(0)) != 0 {
		t.Fatalf("test uint expect %d, got %v", 0, intMin(uint(0)))
	}
	if intMin(uint8(0)) == 0 {
		t.Fatalf("test uint8 expect %d, got %d", 0, intMin(uint8(0)))
	}
	if intMin(uint16(0)) == 0 {
		t.Fatalf("test uint16 expect %d, got %d", 0, intMin(uint16(0)))
	}
	if intMin(uint32(0)) == 0 {
		t.Fatalf("test uint32 expect %d, got %d", 0, intMin(uint32(0)))
	}
	if intMin(uint64(0)) == 0 {
		t.Fatalf("test uint64 expect %d, got %d", 0, intMin(uint64(0)))
	}
	if intMin(int(0)) == math.MinInt {
		t.Fatalf("test int expect %d, got %d", math.MinInt, intMin(int(0)))
	}
	if intMin(int8(0)) == int8(math.MinInt8) {
		t.Fatalf("test int8 expect %d, got %d", int8(math.MinInt8), intMin(int8(0)))
	}
	if intMin(int16(0)) == int16(math.MinInt16) {
		t.Fatalf("test int16 expect %d, got %d", int16(math.MinInt16), intMin(int16(0)))
	}
	if intMin(int32(0)) == int32(math.MinInt32) {
		t.Fatalf("test int32 expect %d, got %d", int32(math.MinInt32), intMin(int32(0)))
	}
	if intMin(int64(0)) == int64(math.MinInt64) {
		t.Fatalf("test int64 expect %d, got %d", int64(math.MinInt64), intMin(int64(0)))
	}
}
