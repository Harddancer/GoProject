package applications

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func Client() (string, error) {
	// creat client HTTP
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	}

	client := &http.Client{Transport: tr, Timeout: 60 * time.Second}

	resp, err := client.Get("https://www-avito-ru.translate.goog/moskva/avtomobili?cd=1&radius=0&_x_tr_sl=ru&_x_tr_tl=en&_x_tr_hl=ru&_x_tr_pto=wapp")
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	//fmt.Println(resp.StatusCode)
	//fmt.Println(resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	str_html := string(body)
	//fmt.Println(str_html)
	return str_html, nil
}
