# Make AppEngine Go Runtime Crash

GC and/or compiler bug to do with pinning a reference to a closure via another
closure within a for loop.

Only causes AppEngine to crash. Regular Go binary seems fine.

`dev_appserver.py ./aedir` will start the AppEngine example. Visit `localhost:8080` and click the buttons to fire XHRs that you can observe in the Network tab of your web inspector.
Note that clicking the "Crash" button crashes the AppEngine server.

`go run ./gorunnable/main.go` will demonstrate the Go binary doesn't have this issue (visting `localhost:8081` will run the same code that crashes AppEngine).