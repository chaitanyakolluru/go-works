// The observer pattern is a behavioural design pattern that allows you to define a one-to-many dependency between objects, such that when one object changes state,
// all of its dependents are notified and updated automatically. This can be useful when you need to update multiple objects in response to a change in another object.
// In Go, you can implement the observer pattern using channels to communicate state changes between objects, like this:

package main

import "fmt"

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
	Named() string
}

type Observer interface {
	Update(Subject)
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(s Subject) {
	o.name = s.Named()
}

type ConcreteSubject struct {
	name      string
	observers []Observer
}

func (s *ConcreteSubject) Named() string { return s.name }

func (s *ConcreteSubject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *ConcreteSubject) Detach(o Observer) {
	for i, oo := range s.observers {
		if oo == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			return
		}
	}
}

func (s *ConcreteSubject) Notify() {
	for _, oo := range s.observers {
		oo.Update(s)
	}
}

func main() {
	fmt.Println("observer-pattern")
	subj := &ConcreteSubject{name: "cjhai"}
	obs := &ConcreteObserver{name: "default name"}
	subj.Attach(obs)
	fmt.Println(obs.name)
	subj.Notify()
	fmt.Println(obs.name)
}
