package broker

import (
	"sync"
	"github.com/pkg/errors"
	"github.com/google/uuid"
)

type Topic string

var InvalidTopicError = errors.New("无效的topic")

// 订阅者
type Subscriber struct {
	Id   string
	T    Topic
	H    Handler
	Broker *Broker
}

func (s *Subscriber) Unsubscribe() {
	s.Broker.Lock()
	defer s.Broker.Unlock()
	var newSubscribers []*Subscriber
	// 移除订阅者
	for _, subscriber := range s.Broker.Subs[s.T] {
		if subscriber.Id == s.Id {
			continue
		}
		newSubscribers = append(newSubscribers, subscriber)
	}
	s.Broker.Subs[s.T] = newSubscribers
}

func (s *Subscriber) Topic() Topic {
	return s.T
}

// Broker简单处理同进程内的消息
type Broker struct {
	Subs map[Topic][]*Subscriber
	sync.RWMutex
}

func (b *Broker) Subscribe(topic Topic, handler Handler) (ISubscriber, error) {
	// 支持并发写map
	b.Lock()
	defer b.Unlock()

	if len(topic) == 0 {
		return nil, InvalidTopicError
	}
	sub := &Subscriber{
		Id:   uuid.New().String(),
		T:    topic,
		H:    handler,
		Broker: b,
	}
	b.Subs[topic] = append(b.Subs[topic], sub)

	return sub, nil
}

func (b *Broker) Publish(publication Publication) error {
	// 支持并发读map
	b.RLock()
	defer b.RUnlock()
	topic := publication.Topic
	for _, subscriber := range b.Subs[topic] {
		err := subscriber.H(publication)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBroker() IBroker {
	return &Broker{
		Subs: make(map[Topic][]*Subscriber),
	}
}