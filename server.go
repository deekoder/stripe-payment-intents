package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

	// start logger
	f, err := os.OpenFile("stripe_log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Application Started")
	// serves the home page
	http.Handle("/", http.FileServer(http.Dir("./static")))

	//used to create the payment intent
	http.HandleFunc("/create-payment-intent", func(w http.ResponseWriter, r *http.Request) {
		log.Println("In create payment handler")
		//MUST DO: Drive this via ENV
		stripe.Key = "sk_test_51HuHfCIbY8JtcKY78g9naTM9v4EsVkced5C6zZG5XHOEE6LsF8qAGIQItzpEaRsKIMVJ9LXCDTr8QiRTu4fTYh9k00wcjgZF0g"

		params := &stripe.PaymentIntentParams{
			Amount:   stripe.Int64(1099), //inr since i am in India for now
			Currency: stripe.String(string(stripe.CurrencyINR)),
		}

		// Verify your integration in this guide by including this parameter
		// This can be used to do tagging and other useful things with this operations.
		params.AddMetadata("integration_check", "accept_a_payment")
		//Let's log that we added the metadata for this intent.
		log.Println("Added Payment Metadata")

		// the actual call here to create the PaymentIntent
		// TODO: use err to handle any errors later.
		intent, err := paymentintent.New(params)
		if err != nil {
			//lets log the error. Use ID to go dig in the stripe dashboard logs
			log.Println("Error in creating payment intent : " + intent.ID)

		} else {
			//lets log success. Print ID to go look up in the stripe dashboard
			log.Println("Successfully created paymentIntent : " + intent.ID)
		}

		// Serve the publishable key also via this data structure
		// TODO: Store keys in keys.toml later. It's not a good idea to check in the keys in git.
		// TODO: In Prod, use ENV variables to set key values. For now this is ok.
		data := CheckoutData{
			ClientSecret:   intent.ClientSecret,
			PublishableKey: "pk_test_51HuHfCIbY8JtcKY7QvTYZXS9ENfP1ZQGLOc22x370kva6qyVpQmC97NbMXhTniivae74k5S0KEwOujeoMuDr25sF00oqBP3qh0",
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
				log.Println("Something went wrong with payment succeeded event : " + err.Error())

				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("PaymentIntent was successful!")
			log.Println("Success with payment succeeded event " + paymentIntent.ID)

		case "payment_intent.payment_failed":
			var paymentIntent stripe.PaymentIntent
			err := json.Unmarshal(event.Data.Raw, &paymentIntent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				log.Println("Something went wrong with payment succeeded event : " + err.Error())

				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("PaymentIntent was NOT successful!")
			log.Println("Failure with payment failed event " + paymentIntent.ID)

		case "payment_method.attached":
			var paymentMethod stripe.PaymentMethod
			err := json.Unmarshal(event.Data.Raw, &paymentMethod)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				log.Println("Something went wrong with payment method attached event : " + err.Error())
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("PaymentMethod was attached to a Customer!")
			log.Println("Success with payment method attached : " + paymentMethod.ID)

		// ... handle other event types
		case "charge.succeeded":
			log.Println("This is a charge suceeded webhook")

		case "payment_intent.created":
			var paymentIntent stripe.PaymentIntent
			err := json.Unmarshal(event.Data.Raw, &paymentIntent)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				log.Println("Something went wrong with payment succeeded event : " + err.Error())

				w.WriteHeader(http.StatusBadRequest)
				return
			}
			fmt.Println("PaymentIntent was created!")
			log.Println("Success with payment created event. Submit with card details to make a charge. " + paymentIntent.ID)

		case "payment_intent.requires_action":
			log.Println("Complete authorization to proceed.")

		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
			log.Println("Something Unexpected happened" + event.Type)

		}

		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":3000", nil)
}
