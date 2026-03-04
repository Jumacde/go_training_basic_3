package main

/*
this file is written by go-structure
say greeting to each current time
*/
import (
	"time"
)

type greet interface {
	GetGreeting() (string, int)
	showGreeting() (string, int)
}

type Greeting struct {
	greeting_call string
	present_time  int
}

func (g *Greeting) SetGreeting(greeting_call string, present_time int) {
	g.greeting_call = greeting_call
	g.present_time = present_time
	g.showGreeting()
}

func (g *Greeting) GetGreeting() (string, int) {
	return g.greeting_call, g.present_time
}

func (g *Greeting) showGreeting() (string, int) {
	g.present_time = time.Now().Hour()

	switch {
	case g.present_time > 7 || g.present_time < 4:
		g.greeting_call = "good night"
	case g.present_time < 10:
		g.greeting_call = "good moring"
	default:
		g.greeting_call = "good afternoon"
	}
	return g.greeting_call, g.present_time

}
