package main

import "testing"

func TestNormal(t *testing.T) {
    result := TopKFrequent(
        "aa bb cc aa cc cc cc aa ab ac bb",
        3,
    )

    if result != "cc aa bb" {
        t.Errorf("Expected cc aa bb, got %s", result)
    }
}

func TestEmpty(t *testing.T) {
    result := TopKFrequent("", 3)

    if result != "" {
        t.Errorf("Expected empty string, got %s", result)
    }
}

func TestLargeK(t *testing.T) {
    result := TopKFrequent("cat dog cat", 10)

    if result != "cat dog" {
        t.Errorf("Expected cat dog, got %s", result)
    }
}