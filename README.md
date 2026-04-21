# 1. A Simple Web Server
2. 
Let me preface this by saying, if you stumbled across this repository while browsing my profile, this is me practicing GoLang!

In this repo, I have built a simple http web server. It has been made using http library, and it can simply take a request, and return a response. 

The server has 3 routes - 
- ```/ - root ```  takes you to "Static Website." page
- ```/form.html```  will show the form.html page, where you can add your name and address (this won't be stored, ofcourse) and then it returns a response, printing the name and address you gave it.
- ```/hello```  prints a simple "Hello, Gophers!" message.

# 2. A Simple Crud Application
3. 
![CRUD FlowChart](pictures/crud.png)

- In this application, I have simulated a database using Structs. It is a database of movies. Now, I have built 4 REST APIs which can do the following:
1. Create a new movie entry
2. Updpate an existing entry
3. Delete an existing entry
4. GET all the movies
5. GET a movie by ID.
6. This is done without using a web-framework, but with http. I bind different functions to different API routes, and also set specific GET / PUT / PATCH / DELETE methods. 
7. This was my FIRST proper REST API using Go!

# 3. Boring

This is an example of using fan-in and fan-out using channels!
Channels BLOCK. We can use the channels to basically delegate a long, time-consuming task to multiple worker functions.
The functions will then write to a single channel.

```
Fan-out : When we have a large, time-consuming task, we can give it to multiple functions and launch multiple Go-Routines.
This will reduce the time taken.

Fan-in : We can have multiple channels from multiple go-routines write to a single channel, and read the final output from that channel.
```

Boring is an interesting, documented example given by Rob Pike (the GOAT). I am simply trying to implement it without looking.

# 4. AWS Lambda using Go-lang

AWS lambda is a serverless offering of AWS, which allows us to basically write functions and expose them via aws to other users, and we are billed for the number of times that function is called by the users. 
It is a serverless technology, meaning our lambda function can scale very rapidly (upto 1000 times concurrently) according to the demand. We also don't have to maintain our own servers to do this, so it is cheap for MVP products.

In my example, I have created a Lambda function that, given a text passage, will compute:
1. The word count
2. The number of each character
3. The most frequent word
4. The approximate read time (standard human reading rate is 200 wpm)

I've also practiced using channels! Basically, 4 go-routines are launched individually and each go-routine carries out one computation.
Each go-routine then writes to a channel. At the end, I then read from each channel to compute the answer struct. 
This is also a nice example of fan-in and fan-out structure! I fan out to compute everything concurrently, and then fan-in to compose my answer struct. 

To connect with a Lambda Function, I had to go through the following 4 steps in AWS CLI:
1. Create a trust-policy.json:
```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```
2. Create an IAM Role:
```
aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
```
3. Attach the policy to the role:
```
aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```
Then, we want to build a binary for linux, and zip it:
```
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```
```
zip function.zip bootstrap
```
Then, we want to basically attach the go-function to that lambda instance:
```
aws lambda create-function --function-name go-text-analyzer --runtime provided.al2 --handler bootstrap --role arn:<YOUR_ARN_HERE>:role/lambda-ex --zip-file fileb://function.zip
```

After this, we are ready to invoke - 
```
aws lambda invoke \
  --function-name go-text-analyzer \
  --payload '{
    "text": "Go is often praised for its simplicity, but that simplicity can be deceptive. Developers coming from languages like Java or Python sometimes underestimate how much thought has gone into Go’s design. The language avoids unnecessary abstractions, encourages composition over inheritance, and provides built-in support for concurrency through goroutines and channels. Go is not just fast in execution; it is fast to write, fast to read, and fast to maintain. In modern cloud-native systems, where services need to scale horizontally and handle thousands of requests per second, Go has become a natural choice. Many engineers choose Go not because it has every feature, but because it has the right features, implemented cleanly and predictably."
  }' response.json
```

### Expected Response:
```
{
    "StatusCode": 200,
    "ExecutedVersion": "$LATEST"
}

```
and, in response.json:
```
{
   "word_count":113,
   "char_count":{
      "a":48,
      "b":9,
      "c":28,
      "d":21,
      "e":73,
      "f":11,
      "g":18,
      "h":23,
      "i":45,
      "j":2,
      "k":1,
      "l":18,
      "m":16,
      "n":48,
      "o":49,
      "p":13,
      "q":1,
      "r":33,
      "s":50,
      "t":54,
      "u":24,
      "v":9,
      "w":3,
      "x":1,
      "y":11,
      "z":1
   },
   "most_frequent_word":"and",
   "reading_time_minutes":1
}
```
