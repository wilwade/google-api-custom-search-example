# google-api-custom-search-example

## Introduction

This is an example of using the Google Custom Search Engine (CSE) API in Go.

This will search the CSE for keywords and search for a string in the domain and report back the position and response.

## Installation

Before you can use the CSE API you will need to create a new Custom Search Engine (CSE). You can create one at [https://cse.google.com](https://cse.google.com). 
After it is created make sure to copy your "Search Engine ID" from the Basics tab.
 
Now you will also need to create an API key to access your CSE via the API. 
 1 Visit the Google Developers Console at [https://console.developers.google.com](https://console.developers.google.com) 
 2 Create a new project.
 3 Enable and manage APIs
 4 Find and enable the Custom Search API
 5 Go to Credentials
 6 Add Credentials -> Service Account
 7 JSON and Save file over search-key.json in the project folder
 
You are now ready to run ```go install && go build``` and run ```./google-api-custom-search-example```

## Using the Command

```./google-api-custom-search-example``` will show usage information.

```./google-api-custom-search-example -id [Search Engine ID] -q [Query String] -d [Domain/path to search for]```

Remember that with the free CSE version you only get 100 queries a day!

## OAuth Instead?

You want to see an OAuth version? Here is the [Google API Library in Go Getting Started Guide](https://github.com/google/google-api-go-client/blob/master/GettingStarted.md#oauth-http-client) which should give you a good starting point.

## More Information

 - [Custom Search Engine](https://cse.google.com)
 - [Custom Search Engine API Docs](https://developers.google.com/custom-search/json-api/v1/reference/)
 - [Google Developers Console](https://console.developers.google.com)
 - [Google API Libraries in Go](https://github.com/google/google-api-go-client)


## License

MIT
