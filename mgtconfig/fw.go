package mgtconfig

import (
	"github.com/fpluchorg/pango/mgtconfig/passwordcomplexity"
	"github.com/fpluchorg/pango/mgtconfig/user"
	"github.com/fpluchorg/pango/util"
)

// Firewall is the client.MGTConfig namespace.
type Firewall struct {
	PasswordComplexity *passwordcomplexity.Firewall
	User *user.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		PasswordComplexity: passwordcomplexity.FirewallNamespace(x),
		User: user.FirewallNamespace(x),
	}
}
