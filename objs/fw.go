package objs

import (
	"github.com/fpluchorg/pango/util"

	"github.com/fpluchorg/pango/objs/addr"
	"github.com/fpluchorg/pango/objs/addrgrp"
	"github.com/fpluchorg/pango/objs/app"
	appgrp "github.com/fpluchorg/pango/objs/app/group"
	"github.com/fpluchorg/pango/objs/app/signature"
	"github.com/fpluchorg/pango/objs/app/signature/andcond"
	"github.com/fpluchorg/pango/objs/app/signature/orcond"
	datapat "github.com/fpluchorg/pango/objs/custom/data"
	cusspy "github.com/fpluchorg/pango/objs/custom/spyware"
	cusurl "github.com/fpluchorg/pango/objs/custom/url"
	cusvuln "github.com/fpluchorg/pango/objs/custom/vulnerability"
	"github.com/fpluchorg/pango/objs/dug"
	"github.com/fpluchorg/pango/objs/edl"
	"github.com/fpluchorg/pango/objs/profile/logfwd"
	"github.com/fpluchorg/pango/objs/profile/logfwd/matchlist"
	"github.com/fpluchorg/pango/objs/profile/logfwd/matchlist/action"
	dfsp "github.com/fpluchorg/pango/objs/profile/security/data"
	dpsp "github.com/fpluchorg/pango/objs/profile/security/dos"
	fprof "github.com/fpluchorg/pango/objs/profile/security/file"
	spg "github.com/fpluchorg/pango/objs/profile/security/group"
	"github.com/fpluchorg/pango/objs/profile/security/spyware"
	ufsp "github.com/fpluchorg/pango/objs/profile/security/url"
	"github.com/fpluchorg/pango/objs/profile/security/virus"
	"github.com/fpluchorg/pango/objs/profile/security/vulnerability"
	wfasp "github.com/fpluchorg/pango/objs/profile/security/wildfire"
	"github.com/fpluchorg/pango/objs/srvc"
	"github.com/fpluchorg/pango/objs/srvcgrp"
	"github.com/fpluchorg/pango/objs/tags"
)

// FwObjs is the client.Objects namespace.
type FwObjs struct {
	Address                             *addr.Firewall
	AddressGroup                        *addrgrp.Firewall
	AntiSpywareProfile                  *spyware.Firewall
	AntivirusProfile                    *virus.Firewall
	Application                         *app.Firewall
	AppGroup                            *appgrp.FwGroup
	AppSignature                        *signature.FwSignature
	AppSigAndCond                       *andcond.FwAndCond
	AppSigOrCond                        *orcond.FwOrCond
	CustomSpyware                       *cusspy.Firewall
	CustomUrlCategory                   *cusurl.Firewall
	CustomVulnerability                 *cusvuln.Firewall
	DataPattern                         *datapat.Firewall
	DataFilteringProfile                *dfsp.Firewall
	DosProtectionProfile                *dpsp.Firewall
	DynamicUserGroup                    *dug.Firewall
	Edl                                 *edl.Firewall
	FileBlockingProfile                 *fprof.Firewall
	LogForwardingProfile                *logfwd.Firewall
	LogForwardingProfileMatchList       *matchlist.FwMatchList
	LogForwardingProfileMatchListAction *action.FwAction
	SecurityProfileGroup                *spg.Firewall
	Services                            *srvc.Firewall
	ServiceGroup                        *srvcgrp.Firewall
	Tags                                *tags.Firewall
	UrlFilteringProfile                 *ufsp.Firewall
	VulnerabilityProfile                *vulnerability.Firewall
	WildfireAnalysisProfile             *wfasp.Firewall
}

// Initialize is invoked on client.Initialize().
func (c *FwObjs) Initialize(i util.XapiClient) {
	c.Address = addr.FirewallNamespace(i)
	c.AddressGroup = addrgrp.FirewallNamespace(i)
	c.AntiSpywareProfile = spyware.FirewallNamespace(i)
	c.AntivirusProfile = virus.FirewallNamespace(i)
	c.Application = app.FirewallNamespace(i)

	c.AppGroup = &appgrp.FwGroup{}
	c.AppGroup.Initialize(i)

	c.AppSignature = &signature.FwSignature{}
	c.AppSignature.Initialize(i)

	c.AppSigAndCond = &andcond.FwAndCond{}
	c.AppSigAndCond.Initialize(i)

	c.AppSigOrCond = &orcond.FwOrCond{}
	c.AppSigOrCond.Initialize(i)

	c.CustomSpyware = cusspy.FirewallNamespace(i)
	c.CustomUrlCategory = cusurl.FirewallNamespace(i)
	c.CustomVulnerability = cusvuln.FirewallNamespace(i)
	c.DataFilteringProfile = dfsp.FirewallNamespace(i)
	c.DataPattern = datapat.FirewallNamespace(i)
	c.DosProtectionProfile = dpsp.FirewallNamespace(i)
	c.DynamicUserGroup = dug.FirewallNamespace(i)
	c.Edl = edl.FirewallNamespace(i)
	c.FileBlockingProfile = fprof.FirewallNamespace(i)
	c.LogForwardingProfile = logfwd.FirewallNamespace(i)

	c.LogForwardingProfileMatchList = &matchlist.FwMatchList{}
	c.LogForwardingProfileMatchList.Initialize(i)

	c.LogForwardingProfileMatchListAction = &action.FwAction{}
	c.LogForwardingProfileMatchListAction.Initialize(i)

	c.SecurityProfileGroup = spg.FirewallNamespace(i)
	c.Services = srvc.FirewallNamespace(i)
	c.ServiceGroup = srvcgrp.FirewallNamespace(i)
	c.Tags = tags.FirewallNamespace(i)
	c.UrlFilteringProfile = ufsp.FirewallNamespace(i)
	c.VulnerabilityProfile = vulnerability.FirewallNamespace(i)
	c.WildfireAnalysisProfile = wfasp.FirewallNamespace(i)
}
