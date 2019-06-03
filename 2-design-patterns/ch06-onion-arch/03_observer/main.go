package main

import . "github.com/YoungsoonLee/funtional-programming-in-go/2-design-patterns/ch06-onion-arch/03_observer/observer"

func main() {
	subject := Subject{}
	oa := Observable{Name: "A"}
	ob := Observable{Name: "B"}
	subject.AddObserver(&Observer{})
	subject.NotifyObservers(oa, ob)

	oc := Observable{Name: "C"}
	subject.NotifyObservers(oa, ob, oc)

	subject.DeleteObserver(&Observer{})
	subject.NotifyObservers(oa, ob, oc)

	od := Observable{Name: "D"}
	subject.NotifyObservers(oa, ob, oc, od)

}
