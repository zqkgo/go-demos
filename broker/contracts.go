package broker

// 订阅者的处理Handler
type Handler func(Publication) error

// 发送的消息
type Publication struct {
	Msg string
	Topic Topic
}

// 订阅者接口
type ISubscriber interface {
	// 取消订阅
	Unsubscribe()
	// 返回订阅者topic
	Topic() Topic
}

// Broker接口
type IBroker interface {
	Subscribe(topic Topic, handler Handler) (ISubscriber, error)
	Publish(publication Publication) error
}
