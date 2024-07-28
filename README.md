# Fetch Receipt Processor Challenge Solution

This is the solution for the [receipt processor code challenge](https://github.com/fetch-rewards/receipt-processor-challenge/tree/main) given to me by Fetch.

This solution is written entirely in Go and can be run locally on your machine.

## Instructions to run application 
1. Clone this repository from Github to your machine.
2. In your preferred command line, change your directory to the location you cloned this repository to.
3. Change your directory to "/cmd" and then change to "/webserver".
4. Verify that you are now in `<CLONED_REPOSITORY_LOCATION>/cmd/webserver`
5. In your command line, run the command `go run main.go`
   - There may be a pop up asking to allow the service to make a connection. Click Allow.
6. If there are no errors, the service should be running successfully on your Localhost.
7. Now you can interact with the service:
   - There are many different tools that API calls can be made to interact with the service. I personally made use of the cURL tool which can be downloaded from [here](https://curl.se/download.html). Some assistance on successfully installing cURL can be found [here](https://developer.zendesk.com/documentation/api-basics/getting-started/installing-and-using-curl/#installing-curl)
   - To save a bit of time, I have included example cURL commands that I put together using the example data that was provided in the code challenge repository.
### POST receipts/process
Example 1:
```
curl --header "Content-Type: application/json" --request POST --data '{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}' http://localhost:8080/receipts/process
```

Example 2:
```
curl --header "Content-Type: application/json" --request POST --data '{         "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}' http://localhost:8080/receipts/process
```
Provide different values for the JSON object after the `--data` tag to send different payloads to the endpoint.

### GET receipts/{id}/points
Example 1:
```
curl --request GET http://localhost:8080/receipts/090c7564-12bc-465a-bb35-976a8c882ce5/points
```
You will have to use the POST command first to load a few receipts into the system and receive their IDs to put in the `{id}` value of the URL. In this example, `090c7564-12bc-465a-bb35-976a8c882ce5` is the test ID that you would replace with IDs that are generated from posting receipts.

## Questions?
Did I miss anything? If there is anything at all that may need some clarification, feel free to ask! I would be very glad to talk more about this solution I put together and also the challenge in general. 

## Reflection
This was a great exercise! I would be very glad to reflect on it as there are a few changes that I would make in the future if this was an application that I would be deploying on a cloud solution. I also would be very glad to hear any feedback! I don't think I've gotten an actual code challenge in a very long time and would be very glad to hear any sort of feedback of things I may be able to do better to ultimately continue improving as an engineer. Thank you!
