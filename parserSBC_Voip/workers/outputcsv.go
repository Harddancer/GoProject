package workers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Writegoodvoice(mg map[string][]string) (s string) {
	file, err := os.Create("voice_map.csv")
	if err != nil {
		fmt.Println("Cannot create CSV file:", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	header := []string{"VoipGW_IP", "NAME", "Client_number"}
	e := writer.Write(header)
	if e != nil {
		fmt.Print("ошибка заголовка")
	}
	// map[string][]string conver to CSV

	for key, value := range mg {

		r := make([]string, 0, 1+len(value))
		el := value[0]
		r = append(r, key)
		r = append(r, el)
		phonesLine := strings.Join(value[1:], "")
		a := []rune(phonesLine)
		l := len(a) - 1
		s := a
		if l > 2 {
			s = a[:l]
		}
		// fmt.Println(string(s))
		r = append(r, string(s))

		err := writer.Write(r)
		if err != nil {
			fmt.Print("ошибка")
		}
		// fmt.Print(r)
	}

	writer.Flush()
	return fmt.Sprintf("Файл voice_map.csv записан!!!")
}
func Writebadvoice(mb map[string][]string) (s string) {
	file2, err := os.Create("voice_map_bad.csv")
	if err != nil {
		fmt.Println("Cannot create CSV file:", err)
	}
	defer file2.Close()

	writer2 := csv.NewWriter(file2)
	header2 := []string{"VoipGW_IP", "Number"}
	e2 := writer2.Write(header2)
	if e2 != nil {
		fmt.Print("ошибка заголовка")
	}
	// map[string][]string conver to CSV
	for key, value := range mb {
		r := make([]string, 0, 1+len(value))
		r = append(r, key)
		r = append(r, value...)
		err := writer2.Write(r)

		if err != nil {
			fmt.Print("ошибка")
		}
	}
	writer2.Flush()
	return fmt.Sprintf("Файл voice_map_bad.csv записан!!!")
}
