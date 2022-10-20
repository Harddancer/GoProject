package workers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// open file nioss make map of data from nioss ip:[element name,....]
func Readernioss(f string) (m map[string][]string) {
	s := make(map[string][]string)
	nioss, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(nioss)
	for scanner.Scan() {
		re := regexp.MustCompile("^\\d+.\\d+.\\d+.\\d+")
		match_ip := strings.Join(re.FindAllString(strings.Split(scanner.Text(), ";")[4], 1), "")
		match_ip = strings.Replace(match_ip, ":", "@", -1)
		s[match_ip] = []string{strings.Split(scanner.Text(), ";")[0]}

	}
	return s

}

// open file sbc
func Createdata(s string, m map[string][]string) (gooddata map[string][]string, baddata map[string][]string) {
	m_bad := make(map[string][]string)
	sbc, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	scanner2 := bufio.NewScanner(sbc)
	for scanner2.Scan() {
		// search ip host from sbc
		ip_sbc := strings.Split(scanner2.Text(), ";")[4]
		// look for number from sbc
		re := regexp.MustCompile("[:]\\d+")
		number_client := strings.Join(re.FindAllString(strings.Split(scanner2.Text(), ";")[1], 1), "")
		number_client = strings.Replace(number_client, ":", "", -1)
		// add a number to map names m
		if _, ok := m[ip_sbc]; ok {
			m[ip_sbc] = append(m[ip_sbc], number_client)

		} else {
			//key does not exist -add number to map names m_bad
			m_bad[ip_sbc] = append(m_bad[ip_sbc], fmt.Sprintf("Данного хоста нет в NIOSS, номер абонента"), number_client)

		}

	}
	return m_bad, m
}
