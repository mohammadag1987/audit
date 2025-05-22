package models

type CloudMachine struct {
	ID             int    `json:"id"`
	ServerName     string `json:"server_name"`
	SiteName       string `json:"site_name"`
	CustomerTitle  string `json:"customer_title"`
	CustomerDLCode string `json:"customer_dl_code"`
}

func (cm *CloudMachine) New(id int, serverName, siteName string, customerTitle string, customerDLCode string) *CloudMachine {
	return &CloudMachine{
		ID:             id,
		ServerName:     serverName,
		SiteName:       siteName,
		CustomerTitle:  customerTitle,
		CustomerDLCode: customerDLCode,
	}
}
