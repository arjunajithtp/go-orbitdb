package orbitdb

import (
	"fmt"
	"io"
	"time"

	"github.com/keks/go-ipfs-colog"
)

func ExampleEventStore() {
	s1, err := NewOrbitDB(topic)
	assert(err == nil, err)

	time.Sleep(5 * time.Millisecond)

	s2, err := NewOrbitDB(topic)
	assert(err == nil, err)

	ev1 := NewEventStore(s1)
	ev2 := NewEventStore(s2)

	fmt.Println("ev1: add foo")
	_, err = ev1.Add("foo")
	assert(err == nil, err)

	time.Sleep(5 * time.Millisecond)

	fmt.Println("ev2: add bar")
	_, err = ev2.Add("bar")
	assert(err == nil, err)

	res := ev2.Query(colog.Query{})

	var (
		p Event
	)
	for err == nil {
		p, err = res()
		assert(err == nil || err == io.EOF, err)

		fmt.Println(p.GetString())
	}

	// Output:
	// ev1: add foo
	// ev2: add bar
	// foo
	// bar
}
