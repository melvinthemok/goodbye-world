package mockapi

import "rsc.io/quote"

func Request() string {
	return quote.Hello()
}
