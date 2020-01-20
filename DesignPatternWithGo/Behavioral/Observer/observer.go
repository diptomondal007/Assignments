package Observer

import "fmt"

type Observer interface {
	Notify(s string)
}

type Publisher struct {
	ObserversList []Observer
}

func(p *Publisher) AddObserver(o Observer){
	p.ObserversList = append(p.ObserversList, o)
}

func(p *Publisher) RemoveObserver(o Observer){
	var indexToRemove int
	for i, observer := range p.ObserversList { if observer == o {
		indexToRemove = i
		break }
	}
	p.ObserversList = append(p.ObserversList[:indexToRemove], p.ObserversList[indexToRemove+1:]...)
}

func (p *Publisher) NotifyObservers(s string){
	fmt.Printf("Publisher received message '%s' to notify observers\n", s)
	for _, observer := range p.ObserversList {
		observer.Notify(s)
	}
}

type TestObserver struct {
	ID int
	Message string
}

func (t *TestObserver) Notify(s string){
	fmt.Printf("Observer %d: message '%s' received \n", t.ID, s)
	t.Message = s
}