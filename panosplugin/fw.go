package panosplugin

import (
	"github.com/fpluchorg/pango/panosplugin/cloudwatch"
	"github.com/fpluchorg/pango/util"
)

type Firewall struct {
	AwsCloudWatch *cloudwatch.Firewall
}

func FirewallNamespace(x util.XapiClient) *Firewall {
	return &Firewall{
		AwsCloudWatch: cloudwatch.FirewallNamespace(x),
	}
}
