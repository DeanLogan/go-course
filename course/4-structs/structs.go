package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
	user user // nested struct
}

type user struct {
	firstName string
	lastName string
}

type messageRecieced struct {
	messageToSend // embedded struct
	sentFrom string
}

type rect struct {
	width int
	height int
}

// area has a receiver of (r rect)
func (r rect) area() int {
	return r.width * r.height
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v who is %s %s\n", m.message, m.phoneNumber, m.user.firstName, m.user.lastName)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
		user: user{"Dean", "Logan"},
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
		user: user{"John", "McClane"},
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
		user: user{"Peter", "Parker"},
	})

	messageR := messageRecieced{
		messageToSend: messageToSend{ 
			phoneNumber: 148255510984,
			message:     "How are you?",
			user: user{"Bruce", "Banner"},
		},
		sentFrom: "Tony Stark",
	}
	fmt.Printf("Recieving message: '%s' to: %v who is %s %s from %s \n", messageR.message, messageR.phoneNumber, messageR.user.firstName, messageR.user.lastName, messageR.sentFrom)
	fmt.Println("")
	r := rect{
		width: 5,
		height: 10,
	}
	
	fmt.Println(r.area())
}