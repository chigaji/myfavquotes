package financialquotes

import (
	"math/rand"
	"time"
)

func Quotes() string {
	quotes := []string{
		"Income alone can never make you wealthy.",
		"Money can't solve all your problems, but neither can poverty",
		"Opportunity is missed my most people because it is dressed in overalls and looks like work.",
		"It is not the man who has too little, but the man who craves more, that is poor.",
		"By when everyone else is selling and hold untill everyone else is selling. It's the essences of successfull investing.",
		"Money is a terrible Master but an Excellent servant.",
		"You must gain control over your Money or the lack of it will forever control you.",
		"Everytime you borrow money, you are robbing your future self.",
		"It's not how much money you make, but how much you keep, how hard it works for you and how many generations you keep it for.",
	}
	rand.Seed(time.Now().UnixNano()) //initial a random generator
	return quotes[rand.Intn(len((quotes)))]

}
