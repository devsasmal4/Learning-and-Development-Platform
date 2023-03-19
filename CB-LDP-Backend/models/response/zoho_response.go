package response

type ZohoResponse []struct {
	EmailID    string `json:"EmailID"`
	LastName   string `json:"LastName"`
	EmployeeID string `json:"EmployeeID"`
	ZohoID     int64  `json:"Zoho_ID"`
	FirstName  string `json:"FirstName"`
}
