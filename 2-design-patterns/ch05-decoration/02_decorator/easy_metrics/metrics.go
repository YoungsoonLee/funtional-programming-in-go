package decorator

import (
	"math/rand"
	"time"

	. "github.com/YoungsoonLee/funtional-programming-in-go/2-design-patterns/ch05-decoration/02_decorator/decorator"
)

func work() {
	randInt := rand.Intn(5000)
	decorator.Debug.Printf("- randInt: %v", randInt)
	workTime := time.Duration(randInt) * time.Millisecond
	time.Sleep(workTime)
}
