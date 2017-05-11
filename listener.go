package ipn

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"net/url"

	"github.com/pkg/errors"
)

//LiveIPNEndpoint contains the notification verification URL
const LiveIPNEndpoint = "https://www.paypal.com/cgi-bin/webscr"

var Debug = false

//Listener creates a PayPal listener.
//if err is set in cb, PayPal will resend the notification at some future point.
func Listener(cb func(err error, n *Notification)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			cb(errors.Wrap(err, "failed to read body"), nil)
		}

		form, err := url.ParseQuery(string(body))
		if err != nil {
			cb(errors.Wrap(err, "failed to parse query"), nil)
			return
		}

		notification := ReadNotification(form)

		if Debug {
			fmt.Printf("paypal: form: %s, parsed: %+v\n", body, notification)
		}

		body = append([]byte("cmd=_notify-validate&"), body...)

		resp, err := http.Post(LiveIPNEndpoint, r.Header.Get("Content-Type"), bytes.NewReader(body))
		if err != nil {
			cb(errors.Wrap(err, "failed to create post verification req"), nil)
			return
		}
		defer resp.Body.Close()

		verifyStatus, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			cb(errors.Wrap(err, "failed to read verification response"), nil)
			return
		}
		if string(verifyStatus) != "VERIFIED" {
			cb(errors.Errorf("unexpected verify status %q", verifyStatus), nil)
			return
		}

		// notification confirmed here

		// tell PayPal to not send more notificatins
		w.WriteHeader(http.StatusOK)
		cb(nil, notification)
	}
}
