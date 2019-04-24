package broker

import (
	"testing"
	"time"
	"sync"
)

func TestBroker(t *testing.T) {
	testBasic(t)
	testCocurrent(t)
}

func testBasic(t *testing.T) {
	broker := NewBroker()

	var n int

	sub1, err := broker.Subscribe("t1", func(publication Publication) error {
		t.Log("\tsub1 receive msg: " + string(publication.Msg))
		n++
		return nil
	})
	if err != nil {
		t.Fatal("fail to subscribe with error: ", err)
	}

	sub2, err := broker.Subscribe("t1", func(publication Publication) error {
		t.Log("\tsub2 receive msg: " + string(publication.Msg))
		n++
		return nil
	})
	if err != nil {
		t.Fatal("fail to subscribe with error: ", err)
	}

	sub3, err := broker.Subscribe("t1", func(publication Publication) error {
		t.Log("\tsub3 receive msg: " + string(publication.Msg))
		n++
		return nil
	})
	if err != nil {
		t.Fatal("fail to subscribe with error: ", err)
	}

	err = broker.Publish(Publication{"hello world", "t1"})
	if err != nil {
		t.Fatal("fail to publish with error: ", err)
	}
	if n != 3 {
		t.Fatal("not all subscribers received message")
	}

	// ubsubscribe
	sub2.Unsubscribe()

	<- time.After(1 * time.Second)

	n = 0
	err = broker.Publish(Publication{"hello world", "t1"})
	if err != nil {
		t.Fatal("fail to publish with error: ", err)
	}
	if n != 2 {
		t.Fatalf("fail to unsubscribe with n=%d", n)
	}

	sub1.Unsubscribe()
	sub3.Unsubscribe()
}

func testCocurrent(t *testing.T) {
	broker := NewBroker()

	var wg sync.WaitGroup
	n := 1000
	wg.Add(2 * n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			_, err := broker.Subscribe("t1", func(publication Publication) error {
				return nil
			})
			if err != nil {
				t.Fatal("fail to subscribe with error: ", err)
			}
		}(i)
	}
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			err := broker.Publish(Publication{"hello world", "t1"})
			if err != nil {
				t.Fatal("fail to publish with error: ", err)
			}
		}(i)
	}
	wg.Wait()
}
