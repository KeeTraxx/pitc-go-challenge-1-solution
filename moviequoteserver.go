package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// MovieQuote represents a quote from a movie
type MovieQuote struct {
	Movie     string `json:"movie"`
	Quote     string `json:"quote"`
	Character string `json:"character"`
	Actor     string `json:"actor"`
	Year      uint   `json:"year"`
}

func main() {
	var movieQuotes []MovieQuote

	// Initialize command line parameters
	verbose := flag.Bool("v", false, "Be verbose")
	help := flag.Bool("h", false, "Display usage")
	flag.Parse()

	// Print usage text if binary is called with -h
	if *help {
		flag.Usage()
		os.Exit(0)
	}

	initLogging(*verbose)

	Logger.DEBUG.Println("Initializing random number generator...")
	rand.Seed(time.Now().Unix())

	Logger.DEBUG.Println("Adding basic set of movie quotes from moviequotes.json...")
	/*movieQuotes = append(movieQuotes, MovieQuote{
		Movie:     "Star Wars: Episode V - The Empire Strikes Back",
		Quote:     "Luke, I am your father!",
		Character: "Darth Vader",
		Actor:     "David Prowse",
		Year:      1980,
	})*/

	data, err := ioutil.ReadFile("moviequotes.json")

	if err != nil {
		Logger.ERROR.Println(err)
		os.Exit(1)
	}

	err = json.Unmarshal(data, &movieQuotes)

	if err != nil {
		Logger.ERROR.Println(err)
		os.Exit(1)
	}

	http.HandleFunc("/v1/moviequotes", func(resp http.ResponseWriter, req *http.Request) {

		// Switch according to HTTP Request Method
		switch req.Method {
		case "GET":
			Logger.INFO.Println("GET /v1/moviequotes")
			Logger.DEBUG.Printf("HTTP headers: %+v\n", req.Header)

			// Set MIME type application/json to HTTP Response
			resp.Header().Add("Content-Type", "application/json")

			// Convert movieQuotes slice to json string
			jsonData, err := json.Marshal(movieQuotes)
			if err != nil {
				// if movieQuotes can't be converted to json somehow?
				Logger.ERROR.Println(err)
				resp.WriteHeader(http.StatusInternalServerError)
				return
			}

			// Write json string to response body
			resp.Write(jsonData)
		case "POST":
			/*
				Test the POST method with the following curl command
				curl -v \
				  -H "Content-Type: application/json" \
				  -d '{"movie":"Terminator 2: Judgment Day","quote":"Hasta La Vista, Baby","actor":"Arnold Schwarzenegger","character":"The Terminator T-101", "year": 1991 }' \
				  http://localhost:1323/v1/moviequotes
			*/

			// Read HTTP POST Body
			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				Logger.ERROR.Println(err)
				resp.WriteHeader(http.StatusUnprocessableEntity)
				return
			}
			Logger.DEBUG.Printf("HTTP Body %v\n", body)

			// Initialize MovieQuote variable
			var moviequote MovieQuote
			json.Unmarshal(body, &moviequote)
			movieQuotes = append(movieQuotes, moviequote)

			Logger.INFO.Printf("Added new movie quote: %+v\n", moviequote)
			// Send back HTTP Status 204
			resp.WriteHeader(http.StatusNoContent)
		default:
			// Return HTTP Status 405 method allowed code if the user tries other http verbs like PATCH,PUT,DELETE
			Logger.WARN.Printf("Denied %v request to %v", req.Method, req.RequestURI)
			resp.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/v1/moviequotes/random", func(resp http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			// Get a random movie quote
			resp.Header().Add("Content-Type", "application/json")
			jsonData, err := json.Marshal(movieQuotes[rand.Intn(len(movieQuotes))])
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
			}
			resp.Write(jsonData)
		default:
			resp.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	Logger.INFO.Println("net/http listening at http://localhost:1323")
	Logger.ERROR.Fatalln(http.ListenAndServe(":1323", nil))
}

// Logger is a basic logging facility
var Logger struct {
	INFO  *log.Logger
	ERROR *log.Logger
	WARN  *log.Logger
	DEBUG *log.Logger
}

func initLogging(verbose bool) {
	Logger.INFO = log.New(os.Stdout, "[INFO] - ", log.LstdFlags)
	Logger.ERROR = log.New(os.Stderr, "[ERROR] - ", log.LstdFlags)
	Logger.WARN = log.New(os.Stderr, "[WARN] - ", log.LstdFlags)
	if verbose {
		Logger.DEBUG = log.New(os.Stdout, "[DEBUG] - ", log.LstdFlags)
	} else {
		Logger.DEBUG = log.New(ioutil.Discard, "[DEBUG] - ", log.LstdFlags)
	}
}
