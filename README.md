# Pre-Requisites
 * Working Go environment set up. Setup here if you don't have one.
 * Install Strip CLI. If you need this install it from here.

# Running the Server
`go run server.go`
 
# Accessing the Site

* Visit http://localhost:3000
!(Web App)[https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/home.png=200x200)


* Enter a Test Card. Test cards are available here.
* Start on the terminal for events : 
    `stripe login`

    `stripe listen --forward-to http://localhost:3000/webhook`
* (Friction Log)[./Friction_Log.md]

