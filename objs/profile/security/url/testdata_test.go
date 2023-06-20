package url

import (
	"github.com/fpluchorg/pango/version"
)

type tc struct {
	desc    string
	version version.Number
	conf    Entry
}

func getTests() []tc {
	return []tc{
		{"v1 minimal", version.Number{7, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "block",
		}},
		{"v1 all bools true", version.Number{7, 1, 0, ""}, Entry{
			Name:                   "t1",
			Description:            "foobar",
			BlockListAction:        "continue",
			DynamicUrl:             true,
			ExpiredLicenseAction:   true,
			TrackContainerPage:     true,
			LogContainerPageOnly:   true,
			SafeSearchEnforcement:  true,
			LogHttpHeaderXff:       true,
			LogHttpHeaderUserAgent: true,
			LogHttpHeaderReferer:   true,
		}},
		{"v1 block list", version.Number{7, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			BlockList:       []string{"www.hotmail.com", "www.cnn.com/news"},
		}},
		{"v1 allow list", version.Number{7, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AllowList:       []string{"www.hotmail.com", "www.cnn.com/news"},
		}},
		{"v1 allow categories", version.Number{7, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v1 alert categories", version.Number{7, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v1 block categories", version.Number{7, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			BlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v1 continue categories", version.Number{7, 1, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			ContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v1 override categories", version.Number{7, 1, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			OverrideCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 minimal", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "block",
		}},
		{"v2 all bools true", version.Number{8, 0, 0, ""}, Entry{
			Name:                   "t1",
			Description:            "foobar",
			BlockListAction:        "continue",
			DynamicUrl:             true,
			ExpiredLicenseAction:   true,
			TrackContainerPage:     true,
			LogContainerPageOnly:   true,
			SafeSearchEnforcement:  true,
			LogHttpHeaderXff:       true,
			LogHttpHeaderUserAgent: true,
			LogHttpHeaderReferer:   true,
		}},
		{"v2 block list", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			BlockList:       []string{"www.hotmail.com", "www.cnn.com/news"},
		}},
		{"v2 allow list", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AllowList:       []string{"www.hotmail.com", "www.cnn.com/news"},
		}},
		{"v2 allow categories", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 alert categories", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 block categories", version.Number{8, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			BlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 continue categories", version.Number{8, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			ContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 override categories", version.Number{8, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			OverrideCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 with credential enforcement settings mode disabled", version.Number{8, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			UcdMode:            UcdModeDisabled,
			UcdAllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 with credential enforcement settings mode user ip user mappings", version.Number{8, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			UcdMode:            UcdModeIpUser,
			UcdLogSeverity:     "informational",
			UcdAlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 with credential enforcement settings mode domain credentials", version.Number{8, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			UcdMode:            UcdModeDomainCredentials,
			UcdLogSeverity:     "informational",
			UcdBlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v2 with credential enforcement settings mode group mapping", version.Number{8, 0, 0, ""}, Entry{
			Name:                  "t1",
			Description:           "foobar",
			BlockListAction:       "alert",
			UcdMode:               UcdModeGroupMapping,
			UcdModeGroupMapping:   "myGroup",
			UcdLogSeverity:        "informational",
			UcdContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 minimal", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "block",
		}},
		{"v3 all bools true", version.Number{8, 1, 0, ""}, Entry{
			Name:                   "t1",
			Description:            "foobar",
			BlockListAction:        "continue",
			DynamicUrl:             true,
			ExpiredLicenseAction:   true,
			TrackContainerPage:     true,
			LogContainerPageOnly:   true,
			SafeSearchEnforcement:  true,
			LogHttpHeaderXff:       true,
			LogHttpHeaderUserAgent: true,
			LogHttpHeaderReferer:   true,
		}},
		{"v3 block list", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			BlockList:       []string{"www.hotmail.com", "www.cnn.com/news"},
		}},
		{"v3 allow list", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AllowList:       []string{"www.hotmail.com", "www.cnn.com/news"},
		}},
		{"v3 allow categories", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 alert categories", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			AlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 block categories", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			BlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 continue categories", version.Number{8, 1, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			ContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 override categories", version.Number{8, 1, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			OverrideCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 with credential enforcement settings mode disabled", version.Number{8, 1, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			UcdMode:            UcdModeDisabled,
			UcdAllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 with credential enforcement settings mode user ip user mappings", version.Number{8, 1, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			UcdMode:            UcdModeIpUser,
			UcdLogSeverity:     "informational",
			UcdAlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 with credential enforcement settings mode domain credentials", version.Number{8, 1, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			BlockListAction:    "alert",
			UcdMode:            UcdModeDomainCredentials,
			UcdLogSeverity:     "informational",
			UcdBlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 with credential enforcement settings mode group mapping", version.Number{8, 1, 0, ""}, Entry{
			Name:                  "t1",
			Description:           "foobar",
			BlockListAction:       "alert",
			UcdMode:               UcdModeGroupMapping,
			UcdModeGroupMapping:   "myGroup",
			UcdLogSeverity:        "informational",
			UcdContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v3 with header insertion", version.Number{8, 1, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockListAction: "alert",
			HttpHeaderInsertions: []HttpHeaderInsertion{
				HttpHeaderInsertion{
					Name:    "h1",
					Type:    "Custom",
					Domains: []string{"domain1.com", "domain2.com"},
					HttpHeaders: []HttpHeader{
						HttpHeader{
							Name:   "header-0",
							Header: "X-Panw",
							Value:  "foo",
							Log:    true,
						},
						HttpHeader{
							Name:   "header-1",
							Header: "X-Api",
							Value:  "bar",
							Log:    false,
						},
					},
				},
			},
		}},
		{"v4 minimal", version.Number{9, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
		}},
		{"v4 all bools true", version.Number{9, 0, 0, ""}, Entry{
			Name:                   "t1",
			Description:            "foobar",
			TrackContainerPage:     true,
			LogContainerPageOnly:   true,
			SafeSearchEnforcement:  true,
			LogHttpHeaderXff:       true,
			LogHttpHeaderUserAgent: true,
			LogHttpHeaderReferer:   true,
		}},
		{"v4 allow categories", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			AllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 alert categories", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			AlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 block categories", version.Number{9, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 continue categories", version.Number{9, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			ContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 override categories", version.Number{9, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			OverrideCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 with credential enforcement settings mode disabled", version.Number{9, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			UcdMode:            UcdModeDisabled,
			UcdAllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 with credential enforcement settings mode user ip user mappings", version.Number{9, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			UcdMode:            UcdModeIpUser,
			UcdLogSeverity:     "informational",
			UcdAlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 with credential enforcement settings mode domain credentials", version.Number{9, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			UcdMode:            UcdModeDomainCredentials,
			UcdLogSeverity:     "informational",
			UcdBlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 with credential enforcement settings mode group mapping", version.Number{9, 0, 0, ""}, Entry{
			Name:                  "t1",
			Description:           "foobar",
			UcdMode:               UcdModeGroupMapping,
			UcdModeGroupMapping:   "myGroup",
			UcdLogSeverity:        "informational",
			UcdContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v4 with header insertion", version.Number{9, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			HttpHeaderInsertions: []HttpHeaderInsertion{
				HttpHeaderInsertion{
					Name:    "h1",
					Type:    "Custom",
					Domains: []string{"domain1.com", "domain2.com"},
					HttpHeaders: []HttpHeader{
						HttpHeader{
							Name:   "header-0",
							Header: "X-Panw",
							Value:  "foo",
							Log:    true,
						},
						HttpHeader{
							Name:   "header-1",
							Header: "X-Api",
							Value:  "bar",
							Log:    false,
						},
					},
				},
			},
		}},
		{"v5 minimal", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
		}},
		{"v5 all bools true", version.Number{10, 0, 0, ""}, Entry{
			Name:                   "t1",
			Description:            "foobar",
			TrackContainerPage:     true,
			LogContainerPageOnly:   true,
			SafeSearchEnforcement:  true,
			LogHttpHeaderXff:       true,
			LogHttpHeaderUserAgent: true,
			LogHttpHeaderReferer:   true,
		}},
		{"v5 allow categories", version.Number{10, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			AllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 alert categories", version.Number{10, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			AlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 block categories", version.Number{10, 0, 0, ""}, Entry{
			Name:            "t1",
			Description:     "foobar",
			BlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 continue categories", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			ContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 override categories", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			OverrideCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 with credential enforcement settings mode disabled", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			UcdMode:            UcdModeDisabled,
			UcdAllowCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 with credential enforcement settings mode user ip user mappings", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			UcdMode:            UcdModeIpUser,
			UcdLogSeverity:     "informational",
			UcdAlertCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 with credential enforcement settings mode domain credentials", version.Number{10, 0, 0, ""}, Entry{
			Name:               "t1",
			Description:        "foobar",
			UcdMode:            UcdModeDomainCredentials,
			UcdLogSeverity:     "informational",
			UcdBlockCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 with credential enforcement settings mode group mapping", version.Number{10, 0, 0, ""}, Entry{
			Name:                  "t1",
			Description:           "foobar",
			UcdMode:               UcdModeGroupMapping,
			UcdModeGroupMapping:   "myGroup",
			UcdLogSeverity:        "informational",
			UcdContinueCategories: []string{"hacking", "dating", "malware", "phishing"},
		}},
		{"v5 with header insertion", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			HttpHeaderInsertions: []HttpHeaderInsertion{
				HttpHeaderInsertion{
					Name:    "h1",
					Type:    "Custom",
					Domains: []string{"domain1.com", "domain2.com"},
					HttpHeaders: []HttpHeader{
						HttpHeader{
							Name:   "header-0",
							Header: "X-Panw",
							Value:  "foo",
							Log:    true,
						},
						HttpHeader{
							Name:   "header-1",
							Header: "X-Api",
							Value:  "bar",
							Log:    false,
						},
					},
				},
			},
		}},
		{"v5 ml models", version.Number{10, 0, 0, ""}, Entry{
			Name:        "t1",
			Description: "foobar",
			MachineLearningModels: []MachineLearningModel{
				MachineLearningModel{
					Model:  "model1",
					Action: "allow",
				},
				MachineLearningModel{
					Model:  "model2",
					Action: "alert",
				},
				MachineLearningModel{
					Model:  "model3",
					Action: "block",
				},
			},
		}},
		{"v5 ml exceptions", version.Number{10, 0, 0, ""}, Entry{
			Name:                      "t1",
			Description:               "foobar",
			MachineLearningExceptions: []string{"ex1", "ex2"},
		}},
	}
}
