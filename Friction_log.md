## What went well
* Copious documentation. 
* Server side examples were great and easy to copy paste
* The flow diagram to show what we were about to do really helped understand the context of how the APIs will help me make a charge.
* I could submit an intent to charge, send the details of the card with the client_secret and submit a charge from the middleware in about 30 mins total. 
* The prototyping process held up pretty well for me even through non contiguous time spent. I had to switch and return back to this task. But the documentation really helped me proceed smoothly. 
* The tests ran and helped me uncover all the testing that should be done such as additional authorizations.
* The webhooks worked beautifully to show that event processing happens and helps imagine the possibilities.


## What could be better
* Documentation: I didn't know what i wanted, pre-built or custom flow. And there was an immersive option. I felt that the as an average developer I had to learn what to learn and where to learn from before even I got started. The number of choices was comforting but it affected my focus a bit. Anyway, i decided to skip the pre-built and immersive options and proceed down the steps.

* I scanned the docs and decided to start with the server side. Copying the server side code as is didn't compile for me immediately in go. So I had to fix it up for about 10 mins. I think this is ok because the developer should know how to handle this part.

* Next, client side code led to a Github link which didnt land on the exact file with the client code. So i had to navigate a bit. I thought it could have had more attention to detail for the developer's experience to minimize distractions.

*  I created my own simple client at first and then realized that certain classes had to be present on the html so I switched back to the client provided by stripe. 

* Presenting a simple piece of client code would have been easier for me. To try out the flow, we just need to present a simple client. The image of an attractive looking client on the page without any code was a surprise. Then i found a github link to go get the client side code made me realize that the client side code was not available on the property itself. By this time i had too many stripe documentation tabs open. 

* On clicking the link, the very first step took me to a github project. The Github link goes to Stripe's general home page on Github rather than the intended file or directory. I had to spend a few minutes orienting myself on where I wanted to go. Finally I navigated to the correct project and found a client folder and went into it to find the HTML or react code. I thought this portion, significantly affected my onboarding experience of the APIs because I had only a 20 min window available before I had to go to some other meeting. 

* Using the correct keys, I realized there is a logged in and a non logged in experience when I hovered over the key where a popover indicated, that the documentation would be customized with my own keys. So I set out to create an account. I wonder aloud if getting my keys could be a step of its own. 

*  The flow diagram was very helpful on what we are accomplishing. I got a bit confused about why the stripe token was called the client secret. I then realized that it was not my API secret key, which should never be sent from the middleware to the UX. Perhaps this is my own confusion because in my language I think of this client_secret as a token. It was not clear to me that the user cannot submit a payment with this client_secret and card details without the API secret key. 

* I had a couple of go related errors in my middleware code around the CheckoutData in my func main(). I later realized the documentation page may have been focusing more on understanding the steps rather than take and bake into a prototype. I adjusted my perception that it would work as soon as I paste the code and run.

* Everything ran finally, I got a payment success and I was able to make a charge. I could see it on my dashboard. The use of the APIs became clear. Next steps was a bit hazy. So documentation support for this may have helped. Error codes, and what each error means documentation could be helpful. Eg. Network interruptions, token invalidated, API Secret reset during transaction.

* Could not comprehend the pairing phrase when adding stripe CLI. I suppose it's a better log in experience. I compare this to heroku CLI and I can see how that required me to do login at the CLI itself. I reckon the pairing phrase is simply like a textual QR code to establish session?

Recommendations
 - The documentation is the UX for API products. So obsesss about how the archetype dev uses it. As a PM, I would even do a user study to see how they onboard. Presenting options is nice but being opinionated is also good. 

  - Help users navigate what APIs they need for what kind of product they are building. Twilio has improved in this regard a great deal. They have a logged in experience which clearly helps users navigate its sprawling APIs. Some of them overlap. They have done clever branding around APIs to help the user pick the right APIs for the right task.

 - Improve Client Side documentation for the onboarding experience could be more self contained inside the onboarding page.

 - Notification when payment success from the Stripe Dashboard. However this may not be loved by all devs. Should validate and then decide if this is a good idea to enable it in test mode while prototyping.   

 - Need to have "What's Next" kind of guidance in the documentation.
 
 - A community of developers help each other. Establish a store front say via a slack channel to answer questions.
 



