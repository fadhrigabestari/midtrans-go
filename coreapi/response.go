package coreapi

type ResponseWithMap map[string]interface{}

// VANumber : bank virtual account number
type VANumber struct {
	Bank     string `json:"bank"`
	VANumber string `json:"va_number"`
}

// Action represents response action
type Action struct {
	Name   string   `json:"name"`
	Method string   `json:"method"`
	URL    string   `json:"url"`
	Fields []string `json:"fields"`
}

type PaymentAmount struct {
	PaidAt string `json:"paid_at"`
	Amount string `json:"amount"`
}

// Response after calling the API
type Response struct {
	StatusCode           string          `json:"status_code"`
	StatusMessage        string          `json:"status_message"`
	PermataVaNumber      string          `json:"permata_va_number"`
	SignKey              string          `json:"signature_key"`
	CardToken            string          `json:"token_id"`
	SavedCardTokenID     string          `json:"saved_token_id"`
	SavedTokenExpAt      string          `json:"saved_token_id_expired_at"`
	SecureToken          bool            `json:"secure_token"`
	Bank                 string          `json:"bank"`
	BillerCode           string          `json:"biller_code"`
	BillKey              string          `json:"bill_key"`
	XlTunaiOrderID       string          `json:"xl_tunai_order_id"`
	BIIVaNumber          string          `json:"bii_va_number"`
	ReURL                string          `json:"redirect_url"`
	ECI                  string          `json:"eci"`
	ValMessages          []string        `json:"validation_messages"`
	Page                 int             `json:"page"`
	TotalPage            int             `json:"total_page"`
	TotalRecord          int             `json:"total_record"`
	FraudStatus          string          `json:"fraud_status"`
	PaymentType          string          `json:"payment_type"`
	OrderID              string          `json:"order_id"`
	TransactionID        string          `json:"transaction_id"`
	TransactionTime      string          `json:"transaction_time"`
	TransactionStatus    string          `json:"transaction_status"`
	GrossAmount          string          `json:"gross_amount"`
	VANumbers            []VANumber      `json:"va_numbers"`
	PaymentAmounts       []PaymentAmount `json:"payment_amounts"`
	PaymentCode          string          `json:"payment_code"`
	Store                string          `json:"store"`
	MerchantID           string          `json:"merchant_id"`
	MaskedCard           string          `json:"masked_card"`
	Currency             string          `json:"currency"`
	CardType             string          `json:"card_type"`
	Actions              []Action        `json:"actions"`
	RefundChargebackID   int             `json:"refund_chargeback_id"`
	RefundAmount         string          `json:"refund_amount"`
	RefundKey            string          `json:"refund_key"`
	Refunds              []RefundDetails `json:"refunds"`
	ChannelResponseCode  string          `json:"channel_response_code"`
	ChannelStatusMessage string          `json:"channel_status_message"`
}

type TransactionStatusResponse struct {
	TransactionTime        string          `json:"transaction_time"`
	GrossAmount            string          `json:"gross_amount"`
	Currency               string          `json:"currency"`
	OrderID                string          `json:"order_id"`
	PaymentType            string          `json:"payment_type"`
	SignatureKey           string          `json:"signature_key"`
	StatusCode             string          `json:"status_code"`
	TransactionID          string          `json:"transaction_id"`
	TransactionStatus      string          `json:"transaction_status"`
	FraudStatus            string          `json:"fraud_status"`
	SettlementTime         string          `json:"settlement_time"`
	StatusMessage          string          `json:"status_message"`
	MerchantID             string          `json:"merchant_id"`
	PermataVaNumber        string          `json:"permata_va_number"`
	VaNumbers              []VANumber      `json:"va_numbers"`
	PaymentAmounts         []PaymentAmount `json:"payment_amounts"`
	ID                     string          `json:"id"`
	PaymentCode            string          `json:"payment_code"`
	Store                  string          `json:"store"`
	MaskedCard             string          `json:"masked_card"`
	Bank                   string          `json:"bank"`
	ApprovalCode           string          `json:"approval_code"`
	Eci                    string          `json:"eci"`
	ChannelResponseCode    string          `json:"channel_response_code"`
	ChannelResponseMessage string          `json:"channel_response_message"`
	CardType               string          `json:"card_type"`
	Refunds                []RefundDetails `json:"refunds"`
	RefundAmount           string          `json:"refund_amount"`
	BillKey                string          `json:"bill_key"`
	BillerCode             string          `json:"biller_code"`
	TransactionType        string          `json:"transaction_type"`
	Issuer                 string          `json:"issuer"`
	Acquirer               string          `json:"acquirer"`
	CustomField1           string          `json:"custom_field1"`
	CustomField2           string          `json:"custom_field2"`
	CustomField3           string          `json:"custom_field3"`
}

type TransactionStatusB2bResponse struct {
	StatusCode    string                      `json:"status_code"`
	StatusMessage string                      `json:"status_message"`
	ID            string                      `json:"id"`
	Transactions  []TransactionStatusResponse `json:"transactions"`
}

// RefundDetails Details
type RefundDetails struct {
	RefundChargebackID   int    `json:"refund_chargeback_id"`
	RefundChargebackUUID string `json:"refund_chargeback_uuid"`
	RefundAmount         string `json:"refund_amount"`
	Reason               string `json:"reason"`
	RefundKey            string `json:"refund_key"`
	RefundMethod         string `json:"refund_method"`
	BankConfirmedAt      string `json:"bank_confirmed_at"`
	CreatedAt            string `json:"created_at"`
}

type RefundResponse struct {
	StatusCode           string `json:"status_code"`
	StatusMessage        string `json:"status_message"`
	ID                   string `json:"id"`
	TransactionID        string `json:"transaction_id"`
	OrderID              string `json:"order_id"`
	GrossAmount          string `json:"gross_amount"`
	Currency             string `json:"currency"`
	MerchantID           string `json:"merchant_id"`
	PaymentType          string `json:"payment_type"`
	TransactionTime      string `json:"transaction_time"`
	TransactionStatus    string `json:"transaction_status"`
	SettlementTime       string `json:"settlement_time"`
	FraudStatus          string `json:"fraud_status"`
	RefundChargebackID   int    `json:"refund_chargeback_id"`
	RefundChargebackUUID string `json:"refund_chargeback_uuid"`
	RefundAmount         string `json:"refund_amount"`
	RefundKey            string `json:"refund_key"`
}

type CardTokenResponse struct {
	StatusCode        string   `json:"status_code"`
	StatusMessage     string   `json:"status_message"`
	ValidationMessage []string `json:"validation_messages"`
	Id                string   `json:"id"`
	TokenID           string   `json:"token_id"`
	Hash              string   `json:"hash"`
	RedirectURL       string   `json:"redirect_url"`
	Bank              string   `json:"bank"`
}

type CardRegisterResponse struct {
	StatusCode        string   `json:"status_code"`
	StatusMessage     string   `json:"status_message"`
	ValidationMessage []string `json:"validation_messages"`
	Id                string   `json:"id"`
	SavedTokenID      string   `json:"saved_token_id"`
	TransactionID     string   `json:"transaction_id"`
	MaskCard          string   `json:"masked_card"`
}

type BinResponse struct {
	Data struct {
		RegistrationRequired interface{} `json:"registration_required"`
		CountryName          interface{} `json:"country_name"`
		CountryCode          string      `json:"country_code"`
		Channel              string      `json:"channel"`
		Brand                string      `json:"brand"`
		BinType              interface{} `json:"bin_type"`
		BinClass             interface{} `json:"bin_class"`
		Bin                  string      `json:"bin"`
		BankCode             interface{} `json:"bank_code"`
		Bank                 string      `json:"bank"`
	} `json:"data"`
}
