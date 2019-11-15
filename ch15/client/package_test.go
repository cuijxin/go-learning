package client

import (
	"go-actual/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(10))
	t.Log(series.Square(10))
}
