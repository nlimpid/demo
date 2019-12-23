package main

// if one of the channel is ready
// then the or channel will be ready

import (
	"fmt"
	"time"

	"github.com/miekg/dns"
	"github.com/sirupsen/logrus"
)

func main() {

	msg := &dns.Msg{}
	localc := &dns.Client{
		ReadTimeout: 5 * time.Second,
	}
	localc.Exchange(msg, "ingtube.com")
	logrus.Infof("msg: %v", msg.Question)
	var or func(chs ...chan interface{}) chan interface{}
	or = func(chs ...chan interface{}) chan interface{} {
		if len(chs) == 0 {
			return nil
		}

		if len(chs) == 1 {
			return chs[0]
		}

		orDone := make(chan interface{})

		go func() {
			defer close(orDone)
			switch len(chs) {
			case 2:
				select {
				case <-chs[0]:
				case <-chs[1]:
				}
			default: // len(chs) > 2
				select {
				case <-chs[0]:
				case <-chs[1]:
				case <-chs[2]:
				case <-or(append(chs[3:], orDone)...):
				}
			}
		}()

		return orDone
	}

	start := time.Now()

	sig := func(after time.Duration) chan interface{} {
		ch := make(chan interface{})

		go func() {
			defer close(ch)
			time.Sleep(after)
		}()

		return ch
	}
	<-or(
		sig(time.Second),
		sig(time.Second*2),
		sig(time.Minute),
	)
	fmt.Println(time.Since(start))
}
