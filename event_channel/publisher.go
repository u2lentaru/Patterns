package main

import (
	"fmt"
)

type Channels map[string]*Channel

type Publisher struct {
	channels Channels
}

func NewPublisher() *Publisher {
	return &Publisher{
		channels: Channels{},
	}
}

func (p *Publisher) AddChannel(name string, channel *Channel) {
	p.channels[name] = channel
}

//TODO: Удаление канала и получение списка каналов
func (p *Publisher) RemoveChannel(chname string) (Channels, error) {
	if _, ok := p.channels[chname]; ok {
		delete(p.channels, chname)
		return p.channels, nil
	}

	return nil, fmt.Errorf("channel %s not found", chname)

}

//TODO: Если 0 имён каналов - отправлять всем
func (p *Publisher) Send(msg string, channels ...string) error {
	if len(channels) == 0 {
		for _, channel := range p.channels {
			channel.Send(msg)
		}
		return nil
	}

	for _, ch := range channels {
		channel, ok := p.channels[ch]
		if !ok {
			return fmt.Errorf("channel %s can't be found", ch)
		}
		channel.Send(msg)
	}

	return nil
}
