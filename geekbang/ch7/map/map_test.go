package my_map

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 4, 3: 5}
	t.Log(m1[2])
	t.Logf("m1 len=%d", len(m1))

	m2 := map[int]int{}
	t.Logf("m2 len=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("m3 len=%d", len(m3))
}

func TestAccessMap(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Logf("key s is not existing.")
	}
}

func TestRangeMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m {
		t.Log(k, v)
	}

}
