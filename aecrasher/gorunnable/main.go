package main
import (
	"net/http"
	"math/rand"
	"runtime"
	"log"
	"fmt"
)

// So you can test with go run.
func main() {
	log.Println("Visit: localhost:8081 to show that this does not crash")

	http.HandleFunc("/", GCBugHandler)

	if err := (&http.Server{Addr: ":8081"}).ListenAndServe(); err != nil {
		panic(err)
	}
}

func GCBugHandler(w http.ResponseWriter, r *http.Request) {
	// Allocate lots of shit.
	lotsOfShit := make([]uint32, 10000000) // 40 MB
	for i := 0; i < len(lotsOfShit); i++ {
		lotsOfShit[i] = rand.Uint32()
	}

	thingThatLeaks := func() error {
		log.Println("Thing!")
		return nil
	}

	var thingThatIsSupposedToPin func()

	for i := 0; i < 3; i++ {
		thingThatIsSupposedToPin = func() {
			thingThatLeaks()
		}
	}

	log.Println(fmt.Sprintf("%d", lotsOfShit[23])) // pin lots of shit until we get here
	runtime.GC()

	if thingThatIsSupposedToPin != nil {
		thingThatIsSupposedToPin()
		return
	}

	log.Println("Reaches here if not thing.")
}