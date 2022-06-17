package main

type WebsiteChecker func(url string) bool

type result struct {
	string
	bool
}

func CheckWebsite(w WebsiteChecker, urls []string) map[string]bool {
	results := map[string]bool{}
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Sending a result struct for each call to w
			// to the resultChannel with the send statement.
			// This uses the <- opertor (send operator),
			// taking a channel on the left & a value on the right
			// Send statement
			resultChannel <- result{u, w(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive experssion
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}

// goroutines: the basic unit of concurrency in GO
// anonymous functions: which we used to start each of the concurrent processes that check websites
// channels: to help organize & control the communication between the different processes
// the race detector which helped us debug problems with concurrent code
