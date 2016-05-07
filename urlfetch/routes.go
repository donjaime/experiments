package app

import (
	"net/http"
	"net/url"
	"strings"

	"io"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/storage/v1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/api/fetchFile", fetchFromGcs)
}

func fetchFromGcs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	gcsFile := r.Form.Get("gcsFile")
	if gcsFile == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctx := appengine.NewContext(r)
	urlStr := "https://storage.googleapis.com/{gcsFile}?alt=media"

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Errorf(ctx, "Failed to make new request: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// you have to do this crazy dance because (by default) net/url will normalize URL escaping, which means
	// that will undo any "/" -> "%2f" escaping that we need to do to escape the slashes in the filepath
	req.URL.Path = strings.Replace(req.URL.Path, "{gcsFile}", url.QueryEscape(gcsFile), 1)
	req.URL.Opaque = "//" + req.URL.Host + req.URL.Path

	client, err := google.DefaultClient(ctx, storage.DevstorageReadOnlyScope)
	if err != nil {
		log.Errorf(ctx, "Failed to make new Google storage client: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Large files that transfer more than 30MB will error here with urlfetch: UNSPECIFIED_ERROR.
	// They are supposed to fail with urlfetch: truncated body.
	rsp, err := client.Do(req)
	if err != nil {
		log.Errorf(ctx, "Failed to Do: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(w, rsp.Body); err != nil {
		log.Errorf(ctx, "Failed to Copy: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	rsp.Body.Close()
}
