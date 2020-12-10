## What went well
* Copious documentation. 

* Server side examples were great and easy to copy paste

* The flow diagram to show what we were about to do really helped understand the context of how the APIs will help me make a charge.

* I could submit an intent to charge, send the details of the card with the client_secret and submit a charge from the middleware in about 30 mins total. It was exciting to get quick results.

* The prototyping process held up pretty well for me even though I had some distracted engagement where non contiguous time was spent. I had to switch and return back to this task. But the documentation really helped me proceed smoothly. 

* The tests ran and helped me uncover all the testing that should be done such as additional authorizations.

<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/tested.png?raw=true" alt="tests" width="500"/>



* I performed some additional tests with international cards and expired tokens
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/detected_expired_key.png?raw=true" alt="tests" width="500"/>


* The webhooks worked beautifully to show that event processing happens and helps imagine the possibilities.I used this to discover all the event types and understand the API better.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/post_events.png?raw=true" alt="tests" width="500"/>

* The log file setup and the webhooks together aided discovery of the Payment Intent APIs further.I ended up enhancing the webhook handler further with charge.succeeded, payment_intent.requires_action and other events.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/rolling_log.png?raw=true" alt="tests" width="500"/>

* Complete environment looked like this.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/dev_setup.png?raw=true" alt="tests" width="500"/>


## What could be better
* Documentation: I didn't know what i wanted, pre-built or custom flow. And there was an immersive option. I felt that as an average developer I had to learn what to learn and where to learn from before even I got started. The number of choices was comforting but it affected my focus a little bit. Anyway, I decided to skip the pre-built and immersive options and proceed down the steps.

* I scanned the docs and decided to start with the server side. Copying the server side code as is didn't compile for me immediately in go. So I had to fix it up for about 10 mins. I had some errors in the CheckOutData struct. I think this is ok because the developer should know how to handle this part.

* I created my own simple form client at first and then realized that certain classes had to be present on the html so I switched back to the client provided by stripe. 

* Presenting a simple piece of client code would have been easier for me.  The image of an attractive looking client on the page without any code was a surprise. Then I found a github link to go get the client side code made me realize that the client side code was not available on the property itself. By this time I had too many stripe documentation tabs open. 

* On clicking the github [link](https://github.com/stripe-samples) the very first step took me to a github project. The Github link goes to Stripe's general home page on Github rather than the intended file or directory. I had to spend a few seconds orienting myself on where I wanted to go. Finally I navigated to the correct project and found a client folder and went into it to find the HTML or react code. I thought this portion, significantly affected my onboarding experience of the APIs due to the number of clicks involved. I was painfully aware that this was all pre-setup to explore the APIs and clicking around looking for client code.
<img src="https://github.com/deekoder/stripe-payment-intents/blob/master/static/screenshots/Github_Generic.png?raw=true" alt="github link" width="500"/>

* Using the correct keys, I realized there is a logged in and a non-logged in experience when I hovered over the key where a popover indicated, that the documentation would be customized with my own keys. So I set out to create an account. I wondered aloud if getting my keys could be a step of its own. 

* The flow diagram was very helpful on what we are accomplishing. It helped me understand that the payment intent is created even before the client card details are collected in my example. I got a bit confused about why the stripe token was called the client secret. I then realized that it was not my API secret key, which should never be sent from the middleware to the UX. Perhaps this is my own confusion.   

* I had a couple of go related errors in my middleware code around the CheckoutData in my func main(). I later realized the documentation page may have been focusing more on understanding the steps rather than take and bake into a prototype. I adjusted my perception that it would work as soon as I paste the code and run. I explored the PaymentIntent API on whether I could return the keys as when Payment Intent succeeded. I peeked into the struct and found all the members. One of the things that I discovered was the Charges API and had some difficulty disambiguating the differences with Payment Intent APIs. 

* Everything ran finally, I got a payment success and I was able to make a charge. I could see it on my dashboard. I thought the Next Steps section could have used some improvements in recommendations.

* Could not comprehend the pairing phrase when adding stripe CLI. I suppose it's a better log in experience. I compare this to heroku CLI and I can see how that required me to do login at the CLI itself. I reckon the pairing phrase is simply like a textual QR code to establish session?

* The right hand side deep links in the documentation was very useful to jump sections in the page.

## General Recommendations
 * The documentation is the UX for APIs. So improve navigation and present multiple paths to exploring in a streamlined manner (with recommended paths). 

 * Reduce onboarding time. I love how Sentry.io has a 5 min onboarding process. Client code is dispensed on a single button click for the choice of client. Focus is quickly returned back to server side exploration of APIs. 

 * Brand the APIs better to help users select which API product they should use for their needs. For the logged in experience - enhance the developer tab in the dashboard to get specific code that can be taken and run immediately. Twilio has recently improved its sprawling set of overlapping APIs very well with branding.

 *  Notification when payment success from the Stripe Dashboard. However this may not be loved by all devs. Should validate and then decide if this is a good idea to enable it in test mode while prototyping.   
 
 * A community of developers help each other. Establish a store front say via a slack channel to answer questions.
 
## Payment Intent API Recommendations
* Full Example Application per server side language with the same client would be helpful. 
* Add expire/cancel token with time setting capabilities with corresponding webhooks for these operations,  if not already available (in case i missed this).
* Schedule a payment intent for usecases where a future charge can be performed. Auto expire if not used by certain time horizon.
* Allow user to configure settings for advanced usage. Eg. Only a single operation with a give token or with retry option. Eg. usecase is a fill or kill ordering system in trading context. Add capability to auto generate new token say after n amount of retries. Make the API behaviour configurable for the power user if this is a need of the user (validate).
* A language specific SDK should do more than wrap the rest API in that language.
* I really appreciated the automated tests after completing the initial prototype. Add extended payment intent API testing before go live.
* Offer a Stripe Enterprise Payment Intent Token Server for enterprise customers (validate need). It's a server that can be customized, configured, managed, monitored, audited for all payment intents created in their org. This could be a SAAS service hosted. As a dev you simply enter custom code in a function as a service like UX via stripe dashboard. Removes the friction of server creation. Enables client side devs to uptake stripe without learning much server side programming.



