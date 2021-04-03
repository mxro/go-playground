package hello

import (
	"mxro.de/go/hello/greetings"

	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}

func MyHello() string {
	return greetings.English()
}
