# Demonstrate URLFetch large file fetch failure

On production appengine, if you fetch a file larger than 30MB it fails (as expected), but it does so with error:
`urlfetch: UNSPECIFIED_ERROR`

It is supposed to fail with `urlfetch: truncated body`
This works correctly when run locally with dev_appserver.py, and used to work correcly in prod appengine.

This seems to be due to `ContentWasTruncated` not being set in production on the `URLFetchResponse` proto. 