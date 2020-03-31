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

	ch3 := myeventchannel.NewChannel()

	pub := myeventchannel.NewPublisher()
	pub.AddChannel("test", ch1)
	pub.AddChannel("test2", ch2)
	pub.AddChannel("test3", ch3)

	if err := pub.Send("HELLO!", "test"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO FROM CH2", "test2"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	if err := pub.Send("HELLO ALL!"); err != nil {
		log.Fatalf("can't send: %s", err)
	}

	rc := "test2"
	cl, err := pub.RemoveChannel(rc)
	if err != nil {
		log.Fatalf("can't send: %s", err)
	}
	log.Printf("Channel %v removed!", rc)
	log.Println("Remains channels:")
	for chname := range cl {
		log.Println(chname)
	}
}
