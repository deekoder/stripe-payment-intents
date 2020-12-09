package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/paymentintent"
)

// CheckoutData is the struct that holds our client secret.
type CheckoutData struct {
	ClientSecret   string `json:"client_secret"`
	PublishableKey string `json:"publishable_key"`
}


func main() {

	// serves the home page
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//used to create the payment intent
	http.HandleFunc("/create-payment-intent", func(w http.ResponseWriter, r *http.Request) {
		stripe.Key = "sk_test_51HuHfCIbY8JtcKY78g9naTM9v4EsVkced5C6zZG5XHOEE6LsF8qAGIQItzpEaRsKIMVJ9LXCDTr8QiRTu4fTYh9k00wcjgZF0g"

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(1099),
			Currency: stripe.String(string(stripe.CurrencyINR)),
		}
		// Verify your integration in this guide by including this parameter
		// This can be used to do tagging and other useful things with this operations.
		params.AddMetadata("integration_check", "accept_a_payment")

		// the actual call here to create the PaymentIntent
		// TODO: use err to handle any errors later.
		intent, _ := paymentintent.New(params)

		// Serve the publishable key also via this data structure
		// TODO: Store keys in keys.toml later. It's not a good idea to check in the keys in git.
		// TODO: In Prod, use ENV variables to set key values. For now this is ok.
		data := CheckoutData{
			ClientSecret:   intent.ClientSecret,
			PublishableKey: "pk_test_51HuHfCIbY8JtcKY7jShQWboVEXHwOg7MbAgey5cSn6DUTM3zpiv48NjX19rcgGDpbBdvLJ4g3zi79QJpMNFSTEdg008jUkrkaN",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	})

	// This handler helps us react to events and we can do tasks we want to do upon such events.
	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		event := stripe.Event{}

		if err := json.Unmarshal(payload, &event); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Unmarshal the event data into an appropriate struct depending on its Type
		switch event.Type {
		case "payment_intent.succeeded":
			var paymentIntent stripe.PaymentIntent
			err := json.Unmarshal(event.Data.Raw, &paymentIntent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("PaymentIntent was successful!")
		case "payment_method.attached":
			var paymentMethod stripe.PaymentMethod
			err := json.Unmarshal(event.Data.Raw, &paymentMethod)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("PaymentMethod was attached to a Customer!")
		// ... handle other event types
		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":3000", nil)
}
