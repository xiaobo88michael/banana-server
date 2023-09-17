package config

type Etherscan struct {
	Apikey       string       `json:"apikey" yaml:"apikey"`
	EtherscanApi EtherscanApi `json:"etherscanApi" yaml:"etherscanApi"`
}

type EtherscanApi struct {
	GetEthBalanceUrl   string `json:"getEthBalanceUrl" yaml:"getEthBalanceUrl"`
	GetERC20BalanceUrl string `json:"getERC20BalanceUrl" yaml:"getERC20BalanceUrl"`
}
