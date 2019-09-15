package fn__test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFn(t *testing.T) {
	i1, _ := returnMutilValues()
	t.Log(i1)
	tsFn := timeSpent(SlowFn)
	t.Log(tsFn(10))
}

func returnMutilValues() (int, int) {
	return rand.Intn(20), rand.Intn(10)
}

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		ret := inner(op)
		fmt.Println("spent time:", time.Since(start))
		return ret
	}
}

func SlowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestVarparam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4))
}

func Sum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret += v
	}
	return ret
}

func clear() {
	fmt.Println("clear resources")
}

func TestDefer(t *testing.T) {
	defer clear()
	t.Log("Starting...")
	panic("Fatal error")
	t.Log("end")
}
