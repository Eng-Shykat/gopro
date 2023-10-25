package main

import ( 
	"fmt"
	"net/url"
)

// const myurl string = "https://innovatpark.com"
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=1234&price=50"

func main() {
	fmt.Println("Welcome to handling urls in Go")
	fmt.Println("The url is:", myurl)

	//parshing the url
	result, _ := url.Parse(myurl)
	// fmt.Println("result,scheme:", result.Scheme) // returns the scheme of the url e.g. http, https, ftp etc
	// fmt.Println("result,host:", result.Host) // returns the host name and port number if any in the url
	// fmt.Println("result,path:", result.Path) // returns the path if any in the url else returns an empty string
	// fmt.Println("result,query:", result.RawQuery) // returns the query string if any in the url else returns an empty string
	// fmt.Println("result,Port:", result.Port()) // returns the port number if any in the url else returns an empty string

	qprams := result.Query() // returns a map of the query string parameters
	fmt.Printf("The type of qprams is: %T\n", qprams)
	fmt.Println("qprams:", qprams["coursename"]) // returns the value of the query string parameter with the key "coursename"

	for _, val := range qprams { // loop through the map of query string parameters and print the key and value
		fmt.Println("Params Is:", val)
	}

	partsOfUrl := &url.URL{ // create a url object using the url package and set the scheme, host and path
		Scheme: "https",
		Host:   "lco.dev",
		Path:   "/tutcss",
		RawQuery: "user=shykat&password=1234",
	}
	anotherUrl := partsOfUrl.String() // convert the url object to a string and store in anotherUrl
	fmt.Println("anotherUrl:", anotherUrl)

	pOfUrl:= &url.URL{
		Scheme: "https",
		Host: "innovatpark.com",
		RawQuery: "user=shykat&password=1234", // this is the query string part of the url
	}
	anotherUrl2 := pOfUrl.String()
	fmt.Println("anotherUrl2:", anotherUrl2)


}
