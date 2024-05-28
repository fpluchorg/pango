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
	o.IdleTimeout = s.IdleTimeout
}

/** Structs / functions for this namespace. **/

func (o Config) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return util.EmptyString, fn(o)
}

type normalizer interface {
	Normalize() []Config
	Names() []string
}

// 6.1+
type container_v1 struct {
	Answer []config_v1 `xml:"management"`
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

func (o *config_v1) normalize() Config {
	ans := Config{
		EnableLogHighDpLoad:          util.AsBoolEmpty(o.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.AsBoolEmpty(o.EnableLogHighDpLoad),
		SupportUtf8ForLogOutput:      util.AsBoolEmpty(o.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.AsBoolEmpty(o.TrafficStopOnLogdbFull),
		HostnameTypeInSyslog:         o.HostnameTypeInSyslog,
	}

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

	return ans
}

type config_v1 struct {
	XMLName                      xml.Name      `xml:"management"`
	EnableLogHighDpLoad          string        `xml:"enable-log-high-dp-load,omitempty"`
	EnableHighSpeedLogForwarding string        `xml:"enable-high-speed-log-forwarding,omitempty"`
	SupportUtf8ForLogOutput      string        `xml:"support-utf8-for-log-output,omitempty"`
	TrafficStopOnLogdbFull       string        `xml:"traffic-stop-on-logdb-full,omitempty"`
	HostnameTypeInSyslog         string        `xml:"hostname-type-in-syslog,omitempty"`
	AdminLockout                 *adminLockout `xml:"admin-lockout"`
	AdminSession                 *adminSession `xml:"admin-session"`
}

type adminLockout struct {
	FailedAttempts string `xml:"failed-attempts,omitempty"`
	LockoutTime    string `xml:"lockout-time,omitempty"`
}

type adminSession struct {
	MaxSessionCount string `xml:"max-session-count,omitempty"`
	MaxSessionTime  string `xml:"max-session-time,omitempty"`
}

func specify_v1(c Config) interface{} {
	ans := config_v1{
		EnableLogHighDpLoad:          util.YesNoEmpty(c.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.YesNoEmpty(c.EnableHighSpeedLogForwarding),
		SupportUtf8ForLogOutput:      util.YesNoEmpty(c.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.YesNoEmpty(c.TrafficStopOnLogdbFull),
	}

	failedAttempts := *c.FailedAttempts
	lockoutTime := *c.LockoutTime
	ans.AdminLockout = &adminLockout{
		FailedAttempts: strconv.Itoa(failedAttempts),
		LockoutTime:    strconv.Itoa(lockoutTime),
	}

	maxSessionTime := *c.MaxSessionTime
	maxSessionCount := *c.MaxSessionCount
	ans.AdminSession = &adminSession{
		MaxSessionTime:  strconv.Itoa(maxSessionTime),
		MaxSessionCount: strconv.Itoa(maxSessionCount),
	}
	return ans
}

// 9.0
type container_v2 struct {
	Answer []config_v2 `xml:"management"`
}

func (o *container_v2) Names() []string {
	return nil
}

func (o *container_v2) Normalize() []Config {
	ans := make([]Config, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *config_v2) normalize() Config {
	ans := Config{
		EnableLogHighDpLoad:          util.AsBoolEmpty(o.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.AsBoolEmpty(o.EnableLogHighDpLoad),
		SupportUtf8ForLogOutput:      util.AsBoolEmpty(o.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.AsBoolEmpty(o.TrafficStopOnLogdbFull),
		HostnameTypeInSyslog:         o.HostnameTypeInSyslog,
	}

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

	return ans
}

type config_v2 struct {
	XMLName                      xml.Name      `xml:"management"`
	EnableLogHighDpLoad          string        `xml:"enable-log-high-dp-load,omitempty"`
	EnableHighSpeedLogForwarding string        `xml:"enable-high-speed-log-forwarding,omitempty"`
	SupportUtf8ForLogOutput      string        `xml:"support-utf8-for-log-output,omitempty"`
	TrafficStopOnLogdbFull       string        `xml:"traffic-stop-on-logdb-full,omitempty"`
	HostnameTypeInSyslog         string        `xml:"hostname-type-in-syslog,omitempty"`
	AdminLockout                 *adminLockout `xml:"admin-lockout"`
	AdminSession                 *adminSession `xml:"admin-session"`
}

func specify_v2(c Config) interface{} {
	ans := config_v1{
		EnableLogHighDpLoad:          util.YesNoEmpty(c.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.YesNoEmpty(c.EnableHighSpeedLogForwarding),
		SupportUtf8ForLogOutput:      util.YesNoEmpty(c.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.YesNoEmpty(c.TrafficStopOnLogdbFull),
	}

	failedAttempts := *c.FailedAttempts
	lockoutTime := *c.LockoutTime
	ans.AdminLockout = &adminLockout{
		FailedAttempts: strconv.Itoa(failedAttempts),
		LockoutTime:    strconv.Itoa(lockoutTime),
	}

	maxSessionTime := *c.MaxSessionTime
	maxSessionCount := *c.MaxSessionCount
	ans.AdminSession = &adminSession{
		MaxSessionTime:  strconv.Itoa(maxSessionTime),
		MaxSessionCount: strconv.Itoa(maxSessionCount),
	}
	return ans
}

// 10.0
type container_v3 struct {
	Answer []config_v3 `xml:"management"`
}

func (o *container_v3) Names() []string {
	return nil
}

func (o *container_v3) Normalize() []Config {
	ans := make([]Config, 0, len(o.Answer))
	for i := range o.Answer {
		ans = append(ans, o.Answer[i].normalize())
	}

	return ans
}

func (o *config_v3) normalize() Config {
	ans := Config{
		EnableLogHighDpLoad:          util.AsBoolEmpty(o.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.AsBoolEmpty(o.EnableLogHighDpLoad),
		SupportUtf8ForLogOutput:      util.AsBoolEmpty(o.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.AsBoolEmpty(o.TrafficStopOnLogdbFull),
		HostnameTypeInSyslog:         o.HostnameTypeInSyslog,
	}

	if o.FailedAttempts != util.EmptyString {
		failedAttempts, err := strconv.Atoi(o.FailedAttempts)
		if err != nil {
			panic(err)
		}
		ans.FailedAttempts = &failedAttempts
	}

	if o.LockoutTime != util.EmptyString {
		lockoutTime, err := strconv.Atoi(o.LockoutTime)
		if err != nil {
			panic(err)
		}
		ans.LockoutTime = &lockoutTime
	}

	if o.MaxSessionCount != util.EmptyString {
		maxSessionCount, err := strconv.Atoi(o.MaxSessionCount)
		if err != nil {
			panic(err)
		}
		ans.MaxSessionCount = &maxSessionCount
	}

	if o.MaxSessionTime != util.EmptyString {
		maxSessionTime, err := strconv.Atoi(o.MaxSessionTime)
		if err != nil {
			panic(err)
		}
		ans.MaxSessionTime = &maxSessionTime
	}

	return ans
}

type config_v3 struct {
	XMLName                      xml.Name `xml:"management"`
	EnableLogHighDpLoad          string   `xml:"enable-log-high-dp-load,omitempty"`
	EnableHighSpeedLogForwarding string   `xml:"enable-high-speed-log-forwarding,omitempty"`
	SupportUtf8ForLogOutput      string   `xml:"support-utf8-for-log-output,omitempty"`
	TrafficStopOnLogdbFull       string   `xml:"traffic-stop-on-logdb-full,omitempty"`
	HostnameTypeInSyslog         string   `xml:"hostname-type-in-syslog,omitempty"`
	//AdminLockout                 *adminLockout `xml:"admin-lockout"`
	//AdminSession                 *adminSession `xml:"admin-session"`
	FailedAttempts  string `xml:"admin-lockout>failed-attempts,omitempty"`
	LockoutTime     string `xml:"admin-lockout>lockout-time,omitempty"`
	MaxSessionCount string `xml:"admin-session>max-session-count,omitempty"`
	MaxSessionTime  string `xml:"admin-session>max-session-time,omitempty"`
}

func specify_v3(c Config) interface{} {
	ans := config_v1{
		EnableLogHighDpLoad:          util.YesNoEmpty(c.EnableLogHighDpLoad),
		EnableHighSpeedLogForwarding: util.YesNoEmpty(c.EnableHighSpeedLogForwarding),
		SupportUtf8ForLogOutput:      util.YesNoEmpty(c.SupportUtf8ForLogOutput),
		TrafficStopOnLogdbFull:       util.YesNoEmpty(c.TrafficStopOnLogdbFull),
	}

	failedAttempts := *c.FailedAttempts
	lockoutTime := *c.LockoutTime
	ans.AdminLockout = &adminLockout{
		FailedAttempts: strconv.Itoa(failedAttempts),
		LockoutTime:    strconv.Itoa(lockoutTime),
	}

	maxSessionTime := *c.MaxSessionTime
	maxSessionCount := *c.MaxSessionCount
	ans.AdminSession = &adminSession{
		MaxSessionTime:  strconv.Itoa(maxSessionTime),
		MaxSessionCount: strconv.Itoa(maxSessionCount),
	}
	return ans
}
