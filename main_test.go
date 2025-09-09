package main

import (
	"testing"
)

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	got := generateRandomElements(0)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got len=%d", len(got))
	}
}

func TestGenerateRandomElements_PositiveSize(t *testing.T) {
	const n = 10_000
	got := generateRandomElements(n)
	if len(got) != n {
		t.Fatalf("expected len=%d, got %d", n, len(got))
	}

	allZero := true
	for _, v := range got {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Fatalf("expected non-zero random values, got all zeros")
	}
}

func TestMaximum_Empty(t *testing.T) {
	if got := maximum(nil); got != 0 {
		t.Fatalf("expected 0 for empty slice, got %d", got)
	}
	if got := maximum([]int{}); got != 0 {
		t.Fatalf("expected 0 for empty slice, got %d", got)
	}
}

func TestMaximum_Single(t *testing.T) {
	if got := maximum([]int{42}); got != 42 {
		t.Fatalf("expected 42, got %d", got)
	}
}

func TestMaximum_Many(t *testing.T) {
	data := []int{1, 7, 3, 9, 2, 9, 4}
	if got := maximum(data); got != 9 {
		t.Fatalf("expected 9, got %d", got)
	}
}

func TestMaxChunks_Correct(t *testing.T) {
	data := []int{1, 7, 3, 9, 2, 9, 4, 15, 8, 10}
	want := maximum(data)
	got := maxChunks(data)
	if got != want {
		t.Fatalf("maxChunks=%d, want=%d", got, want)
	}
}

func TestMaxChunks_Empty(t *testing.T) {
	if got := maxChunks(nil); got != 0 {
		t.Fatalf("expected 0 for empty slice, got %d", got)
	}
}

func TestMaxChunks_ShorterThanChunks(t *testing.T) {
	data := []int{5, 1, 3}
	want := maximum(data)
	got := maxChunks(data)
	if got != want {
		t.Fatalf("maxChunks=%d, want=%d", got, want)
	}
}
