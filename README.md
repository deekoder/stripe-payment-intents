# Pre-Requisites
 * Working Go 1.14 and up environment set up. Setup [here](https://golang.org/dl/) if you don't have one.
 * Install Stripe CLI. If you need this install it from [here](https://stripe.com/docs/stripe-cli).

# Running the Server
`go run server.go`

# Stripe CLI 
`stripe login` complete the login on your browser.
` stripe listen --forward-to http://localhost:3000/webhook` to observe events from the go server running at localhost:3000
 
# Accessing the Site

* Visit http://localhost:3000. You should see this below.

<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/home.png?raw=true" alt="tests" width="500"/>

* Enter various Test Cards. Test cards are available [here](https://stripe.com/docs/testing).

# Friction Log
* [The Friction Log](./Friction_Log.md)

