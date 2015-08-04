# Make AppEngine Segfault

GC and/or compiler bug to do with pinning a reference to a closure via another
closure within a for loop.

Only causes AppEngine to crash. Regular Go binary seems fine.

`dev_appserver.py ./aedir` will start the AppEngine example.

`go run ./gorunnable/main.go` will demonstrate the Go binary doesn't have this issue.