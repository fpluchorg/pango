package mgtconfig

import (
	"github.com/fpluchorg/pango/mgtconfig/user"
	"github.com/fpluchorg/pango/util"
)

// Firewall is the client.MGTConfig namespace.
type Firewall struct {
	User *user.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		User: user.FirewallNamespace(x),
	}
}
