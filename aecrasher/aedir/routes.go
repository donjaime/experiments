package app

import (
	"appengine"
	"math/rand"
	"net/http"
	"runtime"
)

func init() {
	http.HandleFunc("/api/crashServer", GCBugHandler)
}

func GCBugHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cmd := r.Form.Get("cmd")
	if cmd != "crash" {
		return
	}

	// Abandon all hope ye who enter here...

	ctx := appengine.NewContext(r)

	// Allocate lots of shit.
	lotsOfShit := make([]uint32, 10000000) // 40 MB
	for i := 0; i < len(lotsOfShit); i++ {
		lotsOfShit[i] = rand.Uint32()
	}

	thingThatLeaks := func() error {
		ctx.Debugf("Thing!")
		return nil
	}

	var thingThatIsSupposedToPin func()

	for i := 0; i < 3; i++ {
		thingThatIsSupposedToPin = func() {
			thingThatLeaks()
		}
	}

	ctx.Debugf("%d", lotsOfShit[23]) // pin lots of shit until we get here
	runtime.GC()

	if thingThatIsSupposedToPin != nil {
		thingThatIsSupposedToPin()
		return
	}

	ctx.Debugf("Reaches here if not thing.")
}