package week2

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const APIKey = "193ef3a"

type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
}

func SendGetRequest(url string) (string, error) {
	resp, err := http.Get(url) //1.요청에 대한 응답을 담아낼 문자열, 2. 요청 혹은 응답 파싱 과정에서 발생할 수 있는 오류 혹인 nil을 담아낼 error type
	if err != nil {
		return "", err
	}

	defer resp.Body.Close() //이 함수가 나중에 반환되기 전에 응답에서 본문 입력 스트림을 반드시 종료하게 함
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status) //200이 아닐경우 오류 생성하여 반환
	}
	return string(body), nil
}

func SearchByName(name string) (*MovieInfo, error) { //error가 nil만 반환하기 위해서는 movieinfo도 포인터로 선언해야함
	parms := url.Values{}
	parms.Set("apikey", APIKey)
	parms.Set("t", name)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	body, err := SendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody:" + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)

}

func SearchByID(id string) (*MovieInfo, error) {
	parms := url.Values{}
	parms.Set("apikey", APIKey)
	parms.Set("i", id)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	body, err := SendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody:" + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)

}
