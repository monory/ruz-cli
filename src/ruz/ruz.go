package ruz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// PrettifyResponse делает ответ сервера РУЗ с отступами
func PrettifyResponse(response []byte) string {
	var result bytes.Buffer
	json.Indent(&result, response, "", "\t")

	return result.String()
}

// MakeRequest посылает запрос серверу РУЗ и возвращает ответ
func MakeRequest(email, from, to string) ([]byte, error) {
	resp, err := http.Get(constructURL(email, from, to))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func constructURL(email, from, to string) string {
	u := url.URL{
		Scheme: "http",
		Host:   "ruz.hse.ru",
		Path:   "RUZService.svc/personlessons"}

	v := url.Values{}
	v.Add("email", email)
	v.Add("fromdate", from)
	v.Add("todate", to)

	u.RawQuery = v.Encode()

	fmt.Println(u.String())

	return u.String()
}
