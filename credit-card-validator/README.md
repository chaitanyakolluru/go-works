Credit cards will often use the Luhn algorithm to confirm the validity of a credit card number. First, implement the algorithm as a microservice and then expose the functionality with a JSON API.

This project is a web-enabled micro service. It accepts a credit card number in an HTTP request before returning a response. The response indicates whether the number is valid according to the Luhn algorithm.

Implementing this project requires a series of steps that looks something like this:

Implement the Luhn algorithm
Create an HTTP server
Configure the server to respond to GET requests having a JSON payload
Accept valid JSON requests and proceed to step 5, whilst rejecting invalid requests using an HTTP 400 status code
Extract the credit card number from the JSON payload
Run the Luhn algorithm on the number
Wrap the result into a JSON response payload
Return the payload back to the user through the HTTP server
