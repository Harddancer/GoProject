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
		match_ip := strings.Join(re.FindAllString(strings.Split(scanner.Text(), ";")[1], 1), "")
		// match_ip = strings.Replace(match_ip, ":", "@", -1)
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
		slice_sbc := strings.Fields(scanner2.Text())
		ip_sbc := slice_sbc[len(slice_sbc)-1]
		// ip_sbc := strings.Split(scanner2.Text(), ";")[4]
		// look for number from sbc
		re := regexp.MustCompile("[:]\\d+@")
		number_client := strings.Join(re.FindAllString(slice_sbc[0]+slice_sbc[1], 1), "")
		number_client = strings.Replace(number_client, ":", "", -1)
		number_client = strings.Replace(number_client, "@", "", -1)
		var number_client2 string
		if len(number_client) >= 10 {
			number_client2 = fmt.Sprintf("(%s)%s %s|", number_client[0:3], number_client[3:6], number_client[6:10])
		} else {
			number_client2 = number_client
		}
		// number_client := strings.Join(re.FindAllString(strings.Split(scanner2.Text(), ";")[1], 1), "")
		// number_client = strings.Replace(number_client, ":", "", -1)
		// add a number to map names m
		if _, ok := m[ip_sbc]; ok {
			m[ip_sbc] = append(m[ip_sbc], number_client2)

		} else {
			//key does not exist -add number to map names m_bad
			m_bad[ip_sbc] = append(m_bad[ip_sbc], number_client2)

		}

	}
	return m, m_bad


	func ReaderMRnioss(f string) (m map[string][]string) {
		s := make(map[string][]string)
		nioss, err := os.Open(f)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(nioss)
		for scanner.Scan() {
			re := regexp.MustCompile("^\\d+.\\d+.\\d+.\\d+")
			match_ip := strings.Join(re.FindAllString(strings.Split(scanner.Text(), ";")[1], 1), "")
			// match_ip = strings.Replace(match_ip, ":", "@", -1)
			s[match_ip] = []string{strings.Split(scanner.Text(), ";")[0]}
	
		}
		return s
	
	}
}
