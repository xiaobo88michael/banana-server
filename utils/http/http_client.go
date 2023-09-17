package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// 签名信息
type SignReq struct {
	OperationID string `json:"OperationID"`
	SignData    string `json:"signData"`
}
type Response struct {
	ErrCode int         `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	Data    interface{} `json:"data,omitempty"`
}

type CommResp struct {
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}
type CommDataResp struct {
	CommResp
	Data []map[string]interface{} `json:"data"`
}
type CommDataRespOne struct {
	CommResp
	Data map[string]interface{} `json:"data"`
}

func Get(url string) (response []byte, err error) {
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// application/json; charset=utf-8
func Post(url string, data interface{}, timeOutSecond int) (content []byte, err error) {

	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Add("content-type", "application/json; charset=utf-8")

	client := &http.Client{Timeout: time.Duration(timeOutSecond) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// application/json; charset=utf-8
func PostAndHeader(url string, headerMap map[string]string, data interface{}, timeOutSecond int) (content []byte, err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}
	req.Close = true
	req.Header.Add("content-type", "application/json; charset=utf-8")

	for k, v := range headerMap {
		req.Header.Add(k, v)
	}

	client := &http.Client{Timeout: time.Duration(timeOutSecond) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func PostReturn(url string, input, output interface{}, timeOut int) error {
	b, err := Post(url, input, timeOut)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, output); err != nil {
		return err
	}
	return nil
}
