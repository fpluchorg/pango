package settingmanagement

import (
	"github.com/fpluchorg/pango/version"
)

type testCase struct {
	desc    string
	version version.Number
	conf    Config
}

func getTests() []testCase {
	trueVal := true
	return []testCase{
		{"v1 with raw", version.Number{6, 1, 0, ""}, Config{
			EnableLogHighDpLoad: &trueVal,
			//EnableHighSpeedLogForwarding: &trueVal,
			SupportUtf8ForLogOutput: &trueVal,
			TrafficStopOnLogdbFull:  &trueVal,
		}},
	}
}
