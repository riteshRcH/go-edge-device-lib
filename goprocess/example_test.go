package goprocess_test

import (
	"fmt"
	"time"

	"github.com/riteshRcH/go-edge-device-lib/goprocess"
)

func ExampleGo() {
	p := goprocess.Go(func(p goprocess.Process) {
		ticker := time.Tick(200 * time.Millisecond)
		for {
			select {
			case <-ticker:
				fmt.Println("tick")
			case <-p.Closing():
				fmt.Println("closing")
				return
			}
		}
	})

	<-time.After(1100 * time.Millisecond)
	p.Close()
	fmt.Println("closed")
	<-time.After(100 * time.Millisecond)

	// Output:
	// tick
	// tick
	// tick
	// tick
	// tick
	// closing
	// closed
}
