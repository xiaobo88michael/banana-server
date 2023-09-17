package etherscan

import (
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
)

func GetERC20TokenBalance(contractAddr, addr string, decimals int) (decimal.Decimal, error) {
	var (
		price decimal.Decimal
		err   error
	)
	getERC20BalanceUrl := global.GVA_CONFIG.Etherscan.EtherscanApi.GetERC20BalanceUrl

	url := fmt.Sprintf(getERC20BalanceUrl, contractAddr, addr, global.GVA_CONFIG.Etherscan.Apikey)
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
