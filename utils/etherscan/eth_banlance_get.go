package etherscan

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"math/big"
	"net/http"
)

type GetETHBalanceRes struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func GetETHBalance(addr string, decimals int) (decimal.Decimal, error) {
	var (
		price decimal.Decimal
		err   error
	)

	getETHBalanceUrl := global.GVA_CONFIG.Etherscan.EtherscanApi.GetEthBalanceUrl

	url := fmt.Sprintf(getETHBalanceUrl, addr, global.GVA_CONFIG.Etherscan.Apikey)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return price, err
	}
	res, err := client.Do(req)
	if err != nil {
		return price, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return price, err
	}
	var resData GetETHBalanceRes
	err = json.Unmarshal(body, &resData)
	if err != nil {
		return price, err
	}
	if resData.Status != "1" {
		return price, fmt.Errorf("get eth balance api is error", resData.Status, resData.Message)
	}
	price = ToDecimal(resData.Result, decimals)

	return price, nil
}

// 将Wei转换为小数
func ToDecimal(ivalue interface{}, decimals int) decimal.Decimal {
	value := new(big.Int)
	switch v := ivalue.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	result := num.Div(mul)

	return result
}
