package main

import (
	"github.com/go-resty/resty"
)

var urlKong = "http://localhost:8001/services"

type postServiceBody struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type postRouteBody struct {
	Hosts     []string `json:"hosts"`
	Path      []string `json:"paths"`
	StripPath bool     `json:"strip_path"`
	Methods   []string `json:"methods"`
}

func postKongAsService(name string, path string) (isSuccess bool, error error) {
	client := resty.New()
	client.SetDebug(true)
	jsonBody := postServiceBody{
		Name: name,
		Url:  path,
	}
	// POST JSON string
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonBody).
		Post(urlKong)
	if err != nil {
		return false, err
	}

	isSucessHitToKong := resp.IsSuccess()
	return isSucessHitToKong, nil
}

func postKongAsRoute(name string, hosts string, method string, path string) (isSuccess bool, error error) {
	client := resty.New()
	client.SetDebug(true)

	hostsFix := []string{hosts}
	pathFix := []string{path}
	methodFix := []string{method}
	urlFix := urlKong + "/" + name + "/routes"

	jsonBody := postRouteBody{
		Hosts:     hostsFix,
		Path:      pathFix,
		StripPath: false,
		Methods:   methodFix,
	}

	// POST JSON string
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(jsonBody).
		Post(urlFix)

	if err != nil {
		return false, err
	}
	isSucessHitToKong := resp.IsSuccess()
	return isSucessHitToKong, nil
}
