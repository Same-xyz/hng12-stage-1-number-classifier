# hng12-stage-1-number-classifier
A simple api for classifying number
# Number Classification API

## Overview
The **Number Classification API** is an API built with Golang that classifies numbers based on mathematical properties and fun facts. It can determine whether a number is prime, even, odd, perfect, and more.

## Features
- Checks if a number is **even** or **odd**.
- Determine if a number is **prime**.
- Check if a numbers is perfect.
- Determines the sum of the digits
- Gets fun facts about numbers.

## Installation & Setup
### Prerequisites
- Install [Go](https://go.dev/dl/) (version 1.18 or later recommended)

### Clone the Repository
```sh
git clone https://github.com:Same-xyz/hng12-stage-1-number-classifier.git
cd number-classifier
```
### Run the API
```sh
go run main.go
or run main.exe
```
The API will start at `http://localhost:8080` by default.

## API Endpoints
### 1. Classify a Number
**Endpoint:**
```http
GET /classify/{number}
```
**Response Example:**
```json
{
	  "number": 307,
	    "is_prime": true,
	      "is_perfect": false,
	        "properties": [
	        	    "odd"
	        	      ],
	        	        "digit_sum": 10,
	        	          "fun_fact": "307 is a non-palindrome with a palindromic square."
	        	          }
           	        ]
	              }
	              ```

### 2. Get Fun Facts About a Number
**Endpoint:**
```http
GET /fun-fact/{number}
```
**Response Example:**
```json
{
"number": 7,
 "fact": "7 is the number of colors in the rainbow."
 }
 ```

 ## Example Usage
 ### Using Curl
 ```sh
 curl http://localhost:8080/classify/28
 ```

 ### Using Postman
 1. Open Postman and create a **GET** request to `http://localhost:8080/classify/28`.
 2. Click **Send** to see the response.

 ## Deployment
 You can deploy the API using Docker or host it on a cloud service like AWS, Heroku, or DigitalOcean.

 ### Run with Docker
 1. Build the Docker image:
 ```sh
 docker build -t number-classifier-api .
 ```
 2. Run the container:
 ```sh
 docker run -p 8080:8080 number-classifier-api
 ```

 ## Contributing
 1. Fork the repository.
 2. Create a new branch (`feature-new-classification`).
 3. Commit your changes and push.
 4. Open a pull request.
