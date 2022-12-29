package vlans

import (
	"strconv"
	"strings"
)

func GroupVlans(vlans []int) string {
	if len(vlans) == 0 {
		return ""
	}

	var result strings.Builder
	result.WriteString(strconv.Itoa(vlans[0]))

	for i := 1; i < len(vlans); i++ {
		if vlans[i] == vlans[i-1]+1 {
			result.WriteString("-")
			for ; i < len(vlans) && vlans[i] == vlans[i-1]+1; i++ {
			}
			i--
			result.WriteString(strconv.Itoa(vlans[i]))
		} else {
			result.WriteString("," + strconv.Itoa(vlans[i]))
		}
	}

	return result.String()
}
