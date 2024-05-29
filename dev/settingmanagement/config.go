package settingmanagement

import (
	"encoding/xml"
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
	"strconv"
)

// Config is a normalized, version independent representation of a device's
// setting management.
type Config struct {
	Template                     string
	EnableLogHighDpLoad          *bool
	EnableHighSpeedLogForwarding *bool
	SupportUtf8ForLogOutput      *bool
	TrafficStopOnLogdbFull       *bool
	HostnameTypeInSyslog         string
	FailedAttempts               *int
	LockoutTime                  *int
	MaxSessionCount              *int
	MaxSessionTime               *int
	IdleTimeout                  *int
	Internal                     bool // flag to know if it's a panorama resource
}

// Merge copies non connectivity variables from source Config `s` to this
// object.
func (o *Config) Merge(s Config) {
	if s.Template != util.EmptyString {
		o.Template = s.Template
	}

	o.EnableLogHighDpLoad = s.EnableLogHighDpLoad
	o.EnableHighSpeedLogForwarding = s.EnableHighSpeedLogForwarding
	o.SupportUtf8ForLogOutput = s.SupportUtf8ForLogOutput
	o.TrafficStopOnLogdbFull = s.TrafficStopOnLogdbFull
	if s.HostnameTypeInSyslog != util.EmptyString {
		o.HostnameTypeInSyslog = s.HostnameTypeInSyslog
	}
	o.FailedAttempts = s.FailedAttempts
	o.LockoutTime = s.LockoutTime
	o.MaxSessionCount = s.MaxSessionCount
	o.MaxSessionTime = s.MaxSessionTime
	o.IdleTimeout = s.IdleTimeout
	o.Internal = s.Internal
}

/** Structs / functions for this namespace. **/

func (o Config) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning()
	return util.EmptyString, fn(o)
}

type normalizer interface {
	Normalize() []Config
	Names() []string
}

type container_v1 struct {
	Answer []config_device `xml:"management"`
}

func (o *container_v1) Names() []string {
	return nil
}

func (o *container_v1) Normalize() []Config {
	ans := make([]Config, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *config_device) normalize() Config {
	ans := Config{
		EnableLogHighDpLoad:          util.AsBoolEmpty(o.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.AsBoolEmpty(o.EnableLogHighDpLoad),
		SupportUtf8ForLogOutput:      util.AsBoolEmpty(o.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.AsBoolEmpty(o.TrafficStopOnLogdbFull),
		HostnameTypeInSyslog:         o.HostnameTypeInSyslog,
	}

	if o.AdminLockout != nil {
		if o.AdminLockout.FailedAttempts != util.EmptyString {
			failedAttempts, err := strconv.Atoi(o.AdminLockout.FailedAttempts)
			if err != nil {
				panic(err)
			}
			ans.FailedAttempts = &failedAttempts
		}

		if o.AdminLockout.LockoutTime != util.EmptyString {
			lockoutTime, err := strconv.Atoi(o.AdminLockout.LockoutTime)
			if err != nil {
				panic(err)
			}
			ans.LockoutTime = &lockoutTime
		}
	}

	if o.AdminSession != nil {
		if o.AdminSession.MaxSessionCount != util.EmptyString {
			maxSessionCount, err := strconv.Atoi(o.AdminSession.MaxSessionCount)
			if err != nil {
				panic(err)
			}
			ans.MaxSessionCount = &maxSessionCount
		}

		if o.AdminSession.MaxSessionTime != util.EmptyString {
			maxSessionTime, err := strconv.Atoi(o.AdminSession.MaxSessionTime)
			if err != nil {
				panic(err)
			}
			ans.MaxSessionTime = &maxSessionTime
		}
	}

	if o.IdleTimeout != util.EmptyString {
		idleTimeout, err := strconv.Atoi(o.IdleTimeout)
		if err != nil {
			panic(err)
		}
		ans.IdleTimeout = &idleTimeout
	}

	return ans
}

type config_device struct {
	XMLName                      xml.Name      `xml:"management"`
	EnableLogHighDpLoad          string        `xml:"enable-log-high-dp-load,omitempty"`
	EnableHighSpeedLogForwarding string        `xml:"enable-high-speed-log-forwarding,omitempty"`
	SupportUtf8ForLogOutput      string        `xml:"support-utf8-for-log-output,omitempty"`
	TrafficStopOnLogdbFull       string        `xml:"traffic-stop-on-logdb-full,omitempty"`
	HostnameTypeInSyslog         string        `xml:"hostname-type-in-syslog,omitempty"`
	IdleTimeout                  string        `xml:"idle-timeout,omitempty"`
	AdminLockout                 *adminLockout `xml:"admin-lockout"`
	AdminSession                 *adminSession `xml:"admin-session"`
}

type config_panos struct {
	config_device
	StoragePartition *storagePartition `xml:"storage-partition"`
}

type adminLockout struct {
	FailedAttempts string `xml:"failed-attempts,omitempty"`
	LockoutTime    string `xml:"lockout-time,omitempty"`
}

type adminSession struct {
	MaxSessionCount string `xml:"max-session-count,omitempty"`
	MaxSessionTime  string `xml:"max-session-time,omitempty"`
}

type storagePartition struct {
	Internal string `xml:"internal"`
}

// specify_v1 based on the flag Internal we will know the resource if it's panos or device
// config_panos is used for panorama as config_device is used for the device
// this separation logic is done as panos api will not work without having storage-partition param
func specify_v1(c Config) interface{} {
	if c.Internal {
		ans := config_panos{}
		ans.HostnameTypeInSyslog = c.HostnameTypeInSyslog
		if c.FailedAttempts != nil || c.LockoutTime != nil {
			ans.AdminLockout = &adminLockout{}
			if c.FailedAttempts != nil {
				ans.AdminLockout.FailedAttempts = strconv.Itoa(*c.FailedAttempts)
			}
			if c.LockoutTime != nil {
				ans.AdminLockout.LockoutTime = strconv.Itoa(*c.LockoutTime)
			}
		}

		if c.MaxSessionTime != nil || c.MaxSessionCount != nil {
			ans.AdminSession = &adminSession{}
			if c.MaxSessionTime != nil {
				ans.AdminSession.MaxSessionTime = strconv.Itoa(*c.MaxSessionTime)

			}
			if c.MaxSessionCount != nil {
				ans.AdminSession.MaxSessionCount = strconv.Itoa(*c.MaxSessionCount)
			}
		}

		if c.IdleTimeout != nil {
			ans.IdleTimeout = strconv.Itoa(*c.IdleTimeout)
		}

		ans.StoragePartition = &storagePartition{}

		return ans
	}
	ans := config_device{
		EnableLogHighDpLoad:          util.YesNoEmpty(c.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.YesNoEmpty(c.EnableHighSpeedLogForwarding),
		SupportUtf8ForLogOutput:      util.YesNoEmpty(c.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.YesNoEmpty(c.TrafficStopOnLogdbFull),
	}

	return ans
}
