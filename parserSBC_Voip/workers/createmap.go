package workers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// open file nioss make map of data from nioss ip:[element name,....]
// parsing name_voipgw depend of ip host
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
		s[match_ip] = []string{strings.Split(scanner.Text(), ";")[0]}

	}
	return s

}

// open file sbc
func Createdata(s string, m map[string][]string) (gooddata map[string][]string, baddata map[string][]string) {
	// dict code_regions for short numbers mrc
	codenumber_region := map[string]string{
		"ACCESS_P_Tambov":           "4752",
		"ACCESS_P_TVER":             "4822",
		"ACCESS_P_SML":              "4812",
		"ACCESS_P_ORL_COMSTAR":      "4862",
		"ACCESS_P_OREL_MP":          "4862",
		"ACCESS_P_ORL_MTS":          "4862",
		"ACCESS_P_KURSK_MP":         "4712",
		"ACCESS_P_IVANOVO":          "4932",
		"ACCESS_P_BELG_COMSTAR":     "4722",
		"ACCESS_P_IVN_MP":           "4932",
		"ACCESS_P_Ryazan_MP":        "4912",
		"ACCESS_P_TVER_OZON_REZERV": "4822",
		"ACCESS_P_YAR":              "4852",
		"ACCESS_P_YAR_MP":           "4852",
	}

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

		// look for number from sbc
		// parsing number
		re := regexp.MustCompile("[:]\\d+@")
		number_new := re.FindAllString(slice_sbc[0]+slice_sbc[1], 1)
		str_number := strings.Join(number_new, "")
		number_client := strings.Replace(str_number, ":", "", -1)
		number_client = strings.Replace(number_client, "@", "", -1)

		var number_client2 string
		if len(number_client) == 10 {
			number_client2 = fmt.Sprintf("(%s)%s %s|", number_client[0:3], number_client[3:6], number_client[6:10])
		} else if len(number_client) == 6 {
			// parsing name_realm for detection code region
			realm_string := strings.Join(slice_sbc, "")
			//fmt.Print(realm_string)
			re := regexp.MustCompile(`^([A-Za-z_\s]+)sip:`)
			num := re.FindAllStringSubmatch(realm_string, -1)[0][1]
			// fmt.Print(num)
			// fmt.Print("NEXT")
			code := codenumber_region[num]
			new_number := code + number_client
			//fmt.Print(new_number)
			number_client2 = fmt.Sprintf("(%s)%s %s|", new_number[0:3], new_number[3:6], new_number[6:10])

		}

		if _, ok := m[ip_sbc]; ok {
			m[ip_sbc] = append(m[ip_sbc], number_client2)

		} else {
			//key does not exist -add number to map names m_bad
			m_bad[ip_sbc] = append(m_bad[ip_sbc], number_client2)

		}

	}
	return m, m_bad

}
