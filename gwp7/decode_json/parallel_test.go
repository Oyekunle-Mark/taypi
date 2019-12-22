package main

import (
	"testing"
	"time"
)

func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}
