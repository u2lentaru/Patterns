package main

import (
	"log"

	"patterns/myeventchannel"
)

func main() {

	user1 := myeventchannel.NewUser("jkulvich")
	user2 := myeventchannel.NewUser("vasya")

	ch1 := myeventchannel.NewChannel()
	ch1.Subscribe(user1)
	ch1.Subscribe(user2)

	ch2 := myeventchannel.NewChannel()
	ch2.Subscribe(user2)

	pub := myeventchannel.NewPublisher()
	pub.AddChannel("test", ch1)
	pub.AddChannel("test2", ch2)

	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO ALL!"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

}
