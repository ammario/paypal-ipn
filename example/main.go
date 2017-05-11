package main

import (
	"log"
	"net/http"

	ipn "github.com/ammario/paypal-ipn"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/paypal-ipn", ipn.Listener(func(err error, n *ipn.Notification) {
		if err != nil {
			log.Printf("IPN error: %v", err)
			return
		}

		//It's critical you verify the payment was sent to the correct business in the correct currency
		const (
			BusinessEmail = "ammar@ammar.io"
			Currency      = "USD"
		)
		if n.Business != BusinessEmail {
			log.Printf("Payment sent to wrong business: %v", err)
			return
		}
		if n.Currency != Currency {
			log.Printf("Payment in wrong currency: %v", n.Currency)
			return
		}

		log.Printf("%v sent me %v!", n.PayerEmail, n.Gross)
	}))
	log.Fatalf("failed to run http server: %v", http.ListenAndServe(":80", mux))
}
