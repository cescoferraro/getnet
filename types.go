package getnet

// AuthResponse sdfkjn
type AuthFullResponse struct {
	Code     int          `json:"code"`
	Response AuthResponse `json:"payload"`
}

// AuthResponse sdfkjn
type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIN   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

// Bearer sdkfj
func (auth AuthResponse) Bearer() string {
	return "Bearer " + auth.AccessToken
}

// OrderItem asdfkjn
type OrderItem struct {
	Amount      int    `json:"amount"`
	Currency    string `json:"currency"`
	ID          string `json:"id"`
	Description string `json:"description"`
	TaxPercent  string `json:"tax_percent"`
	TaxAmount   string `json:"tax_amount"`
}

// MaketPlaceSubSellerPayments sdfjk
type MaketPlaceSubSellerPayments struct {
	SubsellerSalesAmount int         "subseller_sales_amount"
	SubsellerID          int         "subseller_sales_amount"
	OrderItems           []OrderItem "order_items"
}

// BuyOrder sdfkjb
type BuyOrder struct {
	OrderID     string `json:"order_id"`
	SalesTax    int    `json:"sales_tax"`
	ProductType string `json:"product_type"`
}

// Address sdfkj
type Address struct {
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	District   string `json:"district"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

// Customer sdfkj
type Customer struct {
	CustomerID     string  `json:"customer_id"`
	FirstName      string  `json:"first_name"`
	LastName       string  `json:"last_name"`
	Name           string  `json:"name"`
	Email          string  `json:"email"`
	DocumentType   string  `json:"document_type"`
	PhoneNumber    string  `json:"phone_number"`
	BillingAddress Address `json:"billing_address"`
}

// Card sdkjsdf
type Card struct {
	NumberToken     string `json:"number_token"`
	Bin             string `json:"bin"`
	SecurityCode    string `json:"security_code"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	CardholderName  string `json:"cardholder_name"`
}

// BuyCredit sdkjsdf
type BuyCredit struct {
	Delayed            bool   `json:"delayed"`
	SaveCardData       bool   `json:"save_card_data"`
	TransactionType    string `json:"transaction_type"`
	NumberInstallments int    `json:"number_installments"`
	SoftDescriptor     string `json:"soft_descriptor"`
	DynamicMcc         int    `json:"dynamic_mcc"`
	Card               Card   `json:"card"`
}
