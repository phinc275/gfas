package gfas

import "github.com/ory/ladon"

var DefaultPolicies = []ladon.Policy{
	&ladon.DefaultPolicy{
		ID:          "",
		Description: `Allow authenticated users to create`,
		Subjects:    []string{"<.+>"},
		Resources:   []string{"<.+>"},
		Actions:     []string{"<.+>"},
		Effect:      ladon.AllowAccess,
	},
}
