package common

// Balance 余额结构体
type Balance struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Remain    string `json:"remain"`
	Account   string `json:"account"`
	Index     string `json:"index"`
	Status    stat   `json:"status"`
}

// GetBalance 获取阿里云余额
func GetBalance() Balance {
	return Balance{
		Name:      "",
		EndOfTime: "",
		Type:      "",
		Remain:    "",
	}
}
