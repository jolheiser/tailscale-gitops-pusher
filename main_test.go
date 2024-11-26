// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause
package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEmbeddedTypeUnmarshal(t *testing.T) {
	var gitopsErr ACLGitopsTestError
	gitopsErr.Message = "gitops response error"
	gitopsErr.Data = []ACLTestFailureSummary{
		{
			User:   "GitopsError",
			Errors: []string{"this was initially created as a gitops error"},
		},
	}

	var aclTestErr ACLTestError
	aclTestErr.Message = "native ACL response error"
	aclTestErr.Data = []ACLTestFailureSummary{
		{
			User:   "ACLError",
			Errors: []string{"this was initially created as an ACL error"},
		},
	}

	t.Run("unmarshal gitops type from acl type", func(t *testing.T) {
		b, _ := json.Marshal(aclTestErr)
		var e ACLGitopsTestError
		err := json.Unmarshal(b, &e)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(e.Error(), "For user ACLError") { // the gitops error prints out the user, the acl error doesn't
			t.Fatalf("user heading for 'ACLError' not found in gitops error: %v", e.Error())
		}
	})
	t.Run("unmarshal acl type from gitops type", func(t *testing.T) {
		b, _ := json.Marshal(gitopsErr)
		var e ACLTestError
		err := json.Unmarshal(b, &e)
		if err != nil {
			t.Fatal(err)
		}
		expectedErr := `Status: 0, Message: "gitops response error", Data: [{User:GitopsError Errors:[this was initially created as a gitops error] Warnings:[]}]`
		if e.Error() != expectedErr {
			t.Fatalf("got %v\n, expected %v", e.Error(), expectedErr)
		}
	})
}
