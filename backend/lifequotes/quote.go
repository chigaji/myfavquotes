package lifequotes

import (
	"math/rand"
	"time"
)

func Quote() string {

	quotes := []string{
		"Formal Education will make you aliving, Self-education will make you a 'Fortune'",
		"To be greate at all, One must be greate Here and Now.",
		"Never neglect to do the little things.",
		"What's Easy to do is also Easy Not to do.",
	}
	rand.Seed(time.Now().UnixNano()) //initial a random generator

	return quotes[rand.Intn(len(quotes))]
}
