package concurrency

// WebsiteChecker checks a url, returning a bool.
type (
	WebsiteChecker func(string) bool
	// As we don't need either value to be named, each of them is anonymous within the struct; this can be useful in when it's hard to know what to name a value.
	result struct {
		string
		bool
	}
)

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	// Now when we iterate over the urls, instead of writing to the map directly we're sending a result struct for each call to wc to the resultChannel with a send statement.
	for _, url := range urls {
		go func() {
			// By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time.
			resultChannel <- result{url, wc(url)}
		}()
	}
	// iterates once for each of the urls. Inside we're using a receive expression, which assigns a value received from a channel to a variable.
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
