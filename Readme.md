<h1>Go Load Testing CLI Application</h1>
This Go application performs load testing on a web service by sending HTTP requests and generating a report based on the results. The user specifies the URL of the service, the total number of requests, and the level of concurrency.

### Features
- Configurable URL for load testing
- Adjustable total number of requests
- Adjustable level of concurrency

### Generates a detailed report with:
- Total execution time
- Total number of requests made
- Number of requests with HTTP 200 status
- Distribution of other HTTP status codes

### Requirements
- Docker
- Go 1.22.3 or later (if running locally without Docker)

### Usage
You can use this application via Docker. Example command:
````
docker run renatafborges/stress-test-golang:v1.0.1 --url=http://google.com --requests=20 --concurrency=10   
````

### Installation

1. Clone the repository:

    ````
    git clone https://github.com/renatafborges/stress-test-golang
    cd stress-test-golang
    ````


2. Build the Docker image:

    ````
    docker build -t stresser .
    ````

<h1>Running the Application</h1>

### Run the Docker container with the required parameters:

````
docker run stresser --url=http://example.com --requests=1000 --concurrency=10
````

<h1>Implementation Details</h1>

The application creates a number of workers based on the concurrency level specified by the user. These workers send HTTP GET requests to the specified URL and collect the status codes of the responses. Once all requests are completed, a report is generated summarizing the execution time, total number of requests, number of successful requests (HTTP 200), and the distribution of other HTTP status codes.

### Code Structure:
- cmd/stresser/main.go: The main application code
- internal/processor/stress.go: Where we receive the values from CLI and make the stress test
- Dockerfile: Docker configuration for containerizing the application

### Example Output
````
Total time spent on execution: 17.233724877s
Total number of requests: 100
Number of requests with HTTP status 200: 100
````

### Conclusion
This application provides a simple yet powerful way to perform load testing on a web service. By adjusting the number of requests and the concurrency level, users can simulate different load conditions and observe how their service performs under stress.
