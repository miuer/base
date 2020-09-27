package lqueue_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/miuer/base/structure/queue/lqueue"
)

func TestLQueue(t *testing.T) {
	aq := lqueue.NewLqueue()

	aq.EnQueue("1")
	aq.EnQueue("2")

	q, aQ, err := aq.DeQueue()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(q)

	q, _, err = aQ.DeQueue()

	fmt.Println(q)

	if err != nil {
		log.Fatalln(err)
	}

	q, _, err = aQ.DeQueue()

	fmt.Println(q)

	if err != nil {
		log.Fatalln(err)
	}
}
