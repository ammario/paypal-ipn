package ipn

import (
	"net/url"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// Notification is sent from PayPal to our application.
// See https://developer.paypal.com/docs/classic/ipn/integration-guide/IPNandPDTVariables for more info
type Notification struct {
	TxnType          string `schema:"txn_type"`
	TxnID            string `schema:"txn_id"`
	Business         string `schema:"business"`
	Custom           string `schema:"custom"`
	ParentTxnID      string `schema:"parent_txn_id"`
	ReceiptID        string `schema:"receipt_id"`
	RecieverEmail    string `schema:"receiver_email"`
	RecieverID       string `schema:"receiver_id"`
	Resend           bool   `schema:"resend"`
	ResidenceCountry string `schema:"residence_country"`
	TestIPN          bool   `schema:"test_ipn"`

	//Buyer address information
	AddressCountry     string `schema:"address_country"`
	AddressCity        string `schema:"address_city"`
	AddressCountryCode string `schema:"address_country_code"`
	AddressName        string `schema:"address_name"`
	AddressState       string `schema:"address_state"`
	AddressStatus      string `schema:"address_status"`
	AddressStreet      string `schema:"address_street"`
	AddressZip         string `schema:"address_zip"`

	//Misc buyer info
	ContactPhone      string `schema:"contact_phone"`
	FirstName         string `schema:"first_name"`
	LastName          string `schema:"last_name"`
	PayerBusinessName string `schema:"payer_business_name"`
	PayerEmail        string `schema:"payer_email"`
	PayerID           string `schema:"payer_id"`
	PayerStatus       string `schema:"payer_status"`

	AuthAmount string `schema:"auth_amount"`
	AuthExpire string `schema:"auth_exp"`
	AuthIfD    string `schema:"auth_id"`
	AuthStatus string `schema:"auth_status"`
	Invoice    string `schema:"invoice"`

	//Payment amount
	Currency string  `schema:"mc_currency"`
	Fee      float64 `schema:"mc_fee"`
	Gross    float64 `schema:"mc_gross"`

	Memo string `schema:"memo"`
}

//ReadNotification reads a notification from an //IPN request
func ReadNotification(vals url.Values) (*Notification, error) {
	n := &Notification{}
	return n, decoder.Decode(n, vals)
}
