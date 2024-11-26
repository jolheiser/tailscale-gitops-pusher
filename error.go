package main

import "fmt"

// ACLTestError is ErrResponse but with an extra field to account for ACLTestFailureSummary.
// https://github.com/tailscale/tailscale/blob/26de518413277e0869b815c373f694f6b5d18562/client/tailscale/acl.go#L176-L184
type ACLTestError struct {
	ErrResponse
	Data []ACLTestFailureSummary `json:"data"`
}

func (e ACLTestError) Error() string {
	return fmt.Sprintf("%s, Data: %+v", e.ErrResponse.Error(), e.Data)
}

// ACLTestFailureSummary specifies a user for which ACL tests
// failed and the related user-friendly error messages.
//
// ACLTestFailureSummary specifies the JSON format sent to the
// JavaScript client to be rendered in the HTML.
// https://github.com/tailscale/tailscale/blob/26de518413277e0869b815c373f694f6b5d18562/client/tailscale/acl.go#L160-L174
type ACLTestFailureSummary struct {
	// User is the source ("src") value of the ACL test that failed.
	// The name "user" is a legacy holdover from the original naming and
	// is kept for compatibility but it may also contain any value
	// that's valid in a ACL test "src" field.
	User string `json:"user,omitempty"`

	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// ErrResponse is the HTTP error returned by the Tailscale server.
// https://github.com/tailscale/tailscale/blob/26de518413277e0869b815c373f694f6b5d18562/client/tailscale/tailscale.go#L141-L149
type ErrResponse struct {
	Status  int
	Message string
}

func (e ErrResponse) Error() string {
	return fmt.Sprintf("Status: %d, Message: %q", e.Status, e.Message)
}
