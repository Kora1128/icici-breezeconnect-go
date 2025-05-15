package models

// CustomerDetails represents the response from customer details API
type CustomerDetails struct {
	Success struct {
		ExgTradeDate struct {
			NSE string `json:"NSE"`
			BSE string `json:"BSE"`
			FNO string `json:"FNO"`
			NDX string `json:"NDX"`
		} `json:"exg_trade_date"`
		ExgStatus struct {
			NSE string `json:"NSE"`
			BSE string `json:"BSE"`
			FNO string `json:"FNO"`
			NDX string `json:"NDX"`
		} `json:"exg_status"`
		SegmentsAllowed struct {
			Trading     string `json:"Trading"`
			Equity      string `json:"Equity"`
			Derivatives string `json:"Derivatives"`
			Currency    string `json:"Currency"`
		} `json:"segments_allowed"`
		IDirectUserID        string `json:"idirect_userid"`
		IDirectUserName      string `json:"idirect_user_name"`
		IDirectORDTYP        string `json:"idirect_ORD_TYP"`
		IDirectLastLoginTime string `json:"idirect_lastlogin_time"`
		SessionToken         string `json:"session_token"`
	} `json:"Success"`
	Status int    `json:"Status"`
	Error  string `json:"Error"`
}
