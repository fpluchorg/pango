package dev

import (
	"github.com/fpluchorg/pango/dev/settingmanagement"
	"github.com/fpluchorg/pango/util"

	cert "github.com/fpluchorg/pango/dev/certificate"
	"github.com/fpluchorg/pango/dev/general"
	"github.com/fpluchorg/pango/dev/ha"
	halink "github.com/fpluchorg/pango/dev/ha/monitor/link"
	hapath "github.com/fpluchorg/pango/dev/ha/monitor/path"
	"github.com/fpluchorg/pango/dev/localuserdb/group"
	"github.com/fpluchorg/pango/dev/localuserdb/user"
	"github.com/fpluchorg/pango/dev/profile/authentication"
	"github.com/fpluchorg/pango/dev/profile/certificate"
	"github.com/fpluchorg/pango/dev/profile/email"
	"github.com/fpluchorg/pango/dev/profile/http"
	"github.com/fpluchorg/pango/dev/profile/kerberos"
	"github.com/fpluchorg/pango/dev/profile/ldap"
	"github.com/fpluchorg/pango/dev/profile/radius"
	"github.com/fpluchorg/pango/dev/profile/saml"
	"github.com/fpluchorg/pango/dev/profile/snmp"
	"github.com/fpluchorg/pango/dev/profile/ssltls"
	"github.com/fpluchorg/pango/dev/profile/syslog"
	"github.com/fpluchorg/pango/dev/profile/tacplus"
	"github.com/fpluchorg/pango/dev/ssldecrypt"
	"github.com/fpluchorg/pango/dev/vminfosource"
)

// Panorama is the client.Device namespace.
type Panorama struct {
	AuthenticationProfile *authentication.Panorama
	Certificate           *cert.Panorama
	CertificateProfile    *certificate.Panorama
	EmailServerProfile    *email.Panorama
	GeneralSettings       *general.Panorama
	HaConfig              *ha.Panorama
	HaLinkMonitorGroup    *halink.Panorama
	HaPathMonitorGroup    *hapath.Panorama
	HttpServerProfile     *http.Panorama
	KerberosProfile       *kerberos.Panorama
	LdapProfile           *ldap.Panorama
	LocalUserDbGroup      *group.Panorama
	LocalUserDbUser       *user.Panorama
	RadiusProfile         *radius.Panorama
	SamlProfile           *saml.Panorama
	SettingManagement     *settingmanagement.Panorama
	SslTlsServiceProfile  *ssltls.Panorama
	SslDecrypt            *ssldecrypt.Panorama
	SnmpServerProfile     *snmp.Panorama
	SyslogServerProfile   *syslog.Panorama
	TacacsPlusProfile     *tacplus.Panorama
	VmInfoSource          *vminfosource.Panorama
}

// Initialize is invoked on client.Initialize().
func PanoramaNamespace(x util.XapiClient) *Panorama {
	return &Panorama{
		AuthenticationProfile: authentication.PanoramaNamespace(x),
		Certificate:           cert.PanoramaNamespace(x),
		CertificateProfile:    certificate.PanoramaNamespace(x),
		EmailServerProfile:    email.PanoramaNamespace(x),
		GeneralSettings:       general.PanoramaNamespace(x),
		HaConfig:              ha.PanoramaNamespace(x),
		HaLinkMonitorGroup:    halink.PanoramaNamespace(x),
		HaPathMonitorGroup:    hapath.PanoramaNamespace(x),
		HttpServerProfile:     http.PanoramaNamespace(x),
		KerberosProfile:       kerberos.PanoramaNamespace(x),
		LdapProfile:           ldap.PanoramaNamespace(x),
		LocalUserDbGroup:      group.PanoramaNamespace(x),
		LocalUserDbUser:       user.PanoramaNamespace(x),
		RadiusProfile:         radius.PanoramaNamespace(x),
		SamlProfile:           saml.PanoramaNamespace(x),
		SettingManagement:     settingmanagement.PanoramaNamespace(x),
		SslTlsServiceProfile:  ssltls.PanoramaNamespace(x),
		SslDecrypt:            ssldecrypt.PanoramaNamespace(x),
		SnmpServerProfile:     snmp.PanoramaNamespace(x),
		SyslogServerProfile:   syslog.PanoramaNamespace(x),
		TacacsPlusProfile:     tacplus.PanoramaNamespace(x),
		VmInfoSource:          vminfosource.PanoramaNamespace(x),
	}
}
