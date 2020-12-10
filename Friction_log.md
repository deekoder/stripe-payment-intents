## What went well
* Availabilitiy of copious documentation. 

* Server side examples were great and easy to copy paste.

* The flow diagram to show what we were about to do really helped understand the context of how the APIs will help me make a charge.

* I could submit an intent to charge, send the details of the card with the client_secret and submit a charge from the middleware in about 30 mins total. It was exciting to get quick results.

* The prototyping process held up pretty well through distractions and interruptions. 

* The tests ran and helped uncover all the testing that should be done such as additional authorizations.

<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/tested.png?raw=true" alt="tests" width="500"/>



* The testing inspired me to think of additional test cases such as international cards and expired tokens.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/detected_expired_key.png?raw=true" alt="tests" width="500"/>


* After the basic flow started working, I put in the webhook endpoint and that worked smoothly for me.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/post_events.png?raw=true" alt="tests" width="500"/>

* The log file setup and the webhooks together aided discovery of the Payment Intent APIs further. I ended up enhancing the webhook handler further with charge.succeeded, payment_intent.requires_action and other events. So it aided discoverability.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/rolling_log.png?raw=true" alt="tests" width="500"/>

* Here's the complete environment:
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/dev_setup.png?raw=true" alt="tests" width="500"/>


## What could be better

* Documentation: I didn't know what I wanted, pre-built or custom flow. There was also an immersive option. I felt that as an average developer, I had to learn what to learn and where to learn from before even I got started. While the number of choices was comforting, it affected my focus. Move decision making outside of the steps in the documentation. Ask questions and land the user in one of the options, based on their answers.

* Server Side: I scanned the docs and decided to start with the server side. Copying the server side code as is didn't compile for me immediately in go. So I had to fix it up for about 10 mins. I had some errors in the CheckOutData struct. I decided to put the publishable key in that struct as well. I thought the client UX needed to send card details, publishable key, and client_secret back to the middleware. 

* Client Side: I created my own simple form client at first and then realized that certain classes had to be present on the html for form injection, so I decided to switch to Stripe provided client code.

* Presenting a simple piece of client code would have been easier for me.  The image of an attractive looking client on the page without any code was a surprise to me. A github link to go get the client side code made me realize I could go fetch it from Github and clicked on it. By now I had several browser tabs open for stripe.

* On clicking [link](https://github.com/stripe-samples) the link goes to Stripe's general home page on Github rather than the intended file or directory. I had to spend a few seconds orienting myself on where I wanted to go. Finally I navigated to the correct project and found a client folder and went into it to find the HTML or react code. I thought this portion, significantly affected my onboarding experience of the Payment Intent APIs due to the number of clicks involved.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/Github_Generic.png?raw=true" alt="github link" width="500"/>

* Keys: I realized there is a logged in and a non-logged in experience when I hovered over the key where a popover indicated, that the documentation would be customized with my own keys. So I set out to create an account. I wondered aloud if getting my keys could be a step of its own. 

* So far, I had all the setup to start using the Payment Intent APIs. The flow diagram was very helpful on what we are accomplishing. The payment intent is created even before the client card details are collected in the example. So I could comprehend that the word "Payment Intent" indicates the intent to process a charge. A token is dispensed to the client to fulfill this intention.

* I got a bit confused about why the stripe token was called the client secret. I then realized that it was not my API secret key, which should never be sent from the middleware to the UX. 

* During the server side coding, I had a couple of go related errors in my middleware code around the CheckoutData in my func main(). I concluded  that the documentation page may have been focusing more on understanding the steps rather than take and bake into a prototype. I adjusted my perception that it would work as soon as I paste the code and run. I explored the PaymentIntent API on whether I could return the publishable keys as when Payment Intent succeeded. I peeked into the struct and found all the members. One of the things that I discovered, was the Charges API and had some difficulty disambiguating the differences with Payment Intent APIs. I decided to table thi and pursue the thread of Charges vs Payment Intent APIs later. 

* Everything ran finally, I got a payment success and I was able to make a charge. I could see it on my dashboard. I thought the Next Steps section could have used some improvements in recommendations.

* I could not comprehend the reason for the pairing phrase when adding stripe CLI. I suppose it's a better log in experience. I compare this to heroku CLI and I can see how that required me to do login at the CLI itself. I reckon the pairing phrase is simply like a textual QR code to establish session?


## General Recommendations
 * The documentation is the UX for APIs. The time taken to setup affects the exploration of specific APIs. Customize documentation to each API product.

 * Reduce onboarding time. I love how Sentry.io has a 5 min onboarding process. Client code is dispensed on a single button click for the choice of client. Focus is quickly returned back to server side exploration of APIs. The experience is only logged in. You can see results quickly on the dashboard.

 * Brand the APIs better to help users select which API product they should use for their needs. For the logged in experience - enhance the developer tab in the dashboard to get specific code that can be taken and run immediately. Twilio has recently improved its sprawling set of overlapping APIs very well with branding.

 * Notify when payment success occurs, on the Stripe Dashboard. However this may not be loved by all devs. PM should validate and then decide if it is a good idea to enable it in test mode while prototyping.   
 
 * A community of developers can help each other. Establish a store front say via a slack channel to answer questions.
 
## Payment Intent API Recommendations

* Full example application per server side language with the same client would be helpful. Cloning and perusing a full application, is another way to learn the APIs. 

* Add expire/cancel token with time capabilities. Add corresponding webhooks for these operations. This can power several usecases. Eg. Buy this ticket with a timer option. Or Lightning deals on shopping sites.

* Schedule a payment intent for usecases where a future charge can be performed. Auto expire if not used by certain time horizon. This can power several usecases. Eg. Any busines that needs to hold a deposit down can use this feature. Webhooks can be used to do analytics and email campaigns to remind the user. Booking.com is a site that takes no card to reserve a hotel room. You can pay when you use. If you don't use it in 3 days, they send email reminders.

* Allow the power user to configure settings for advanced usage of the Payment Intent APIs. Eg. Only a single operation with a give token or with retry option. Eg. usecase is a fill or kill ordering system in a trading context. Add capability to auto generate new token say after n amount of retries. If paying rent is missed after 3 retries, notify a building manager to call this tenant. PM should validate the strength of these new features with customers before they are developed.

* A language specific SDK should do more than wrap the rest API in that language. Understand the types of customers who use different SDKs and see how to make the SDK more than a wrapper of REST APIs.

* I really appreciated the automated tests after completing the initial prototype. Add extended payment intent API test cases before, go live. We can charge customers for full suite of tests. Customers are willing to pay for peace of mind.

* Offer a Stripe Enterprise Payment Intent Token Server for enterprise customers (validate need). It's a server that can be customized, configured, managed, monitored, audited for all payment intents created in their org. Missed intents can show new opportunities to develop revenue for a customer. Missed payment intents can help fraud detection perhaps.  

* Lambda Option for Growth : As a dev you simply enter custom code in a function as a service like UX via stripe dashboard. Removes the friction of server creation. Enables client side devs to uptake stripe without learning much server side programming.



