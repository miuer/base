package aqueue_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/miuer/base/structure/queue/aqueue"
)

func TestAQueue(t *testing.T) {
	aq := aqueue.NewAqueue()

	aq.EnQueue("enqueue")

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
}
