package client

import (
	"fmt"
	"github.com/meua/gin-test-project/geekbang/ch15/series"
	"testing"
)

func init() {
	fmt.Println("init3")
}

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonaciSeries(10))
	//t.Log(series.square(5))
}
