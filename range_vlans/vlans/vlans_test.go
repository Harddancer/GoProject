package vlans

import (
	"testing"
)

type testvlans struct {
	got  []int
	want string
}

var test = []testvlans{
	testvlans{[]int{1, 2, 3, 4, 5, 10}, "1-5,10"},
	testvlans{[]int{1, 2, 3, 6}, "1-3,6"},
	testvlans{[]int{1, 2, 3, 6, 7, 8, 11}, "1-3,6-8,11"},
	testvlans{[]int{1, 2, 3, 6, 7, 8}, "1-3,6-8"},
}

func TestGroupvlans(t *testing.T) {
	for _, val := range test {
		if output := GroupVlans(val.got); output != val.want {
			t.Errorf("Output %q not equal to expected %q", output, val.want)
		}

	}

}
