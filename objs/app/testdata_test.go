package app

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
		{"v1 no default", version.Number{7, 1, 0, ""}, Entry{
			Name:                                 "v1",
			DefaultType:                          DefaultTypeNone,
			Category:                             "category",
			Subcategory:                          "subcat",
			Technology:                           "tech",
			Description:                          "my desc",
			Timeout:                              5,
			TcpTimeout:                           6,
			UdpTimeout:                           7,
			TcpHalfClosedTimeout:                 8,
			TcpTimeWaitTimeout:                   9,
			Risk:                                 2,
			AbleToFileTransfer:                   true,
			ExcessiveBandwidth:                   true,
			TunnelsOtherApplications:             true,
			HasKnownVulnerability:                true,
			UsedByMalware:                        true,
			EvasiveBehavior:                      true,
			PervasiveUse:                         true,
			ProneToMisuse:                        true,
			ContinueScanningForOtherApplications: true,
			FileTypeIdent:                        true,
			VirusIdent:                           true,
			DataIdent:                            true,
			AlgDisableCapability:                 "alg disable",
			ParentApp:                            "parent app",
			raw: map[string]string{
				"sigs": "a bunch of sig data",
			},
		}},
		{"v1 default ports", version.Number{7, 1, 0, ""}, Entry{
			Name:                 "v1",
			DefaultType:          DefaultTypePort,
			DefaultPorts:         []string{"tcp/dynamic", "udp/32"},
			Category:             "category",
			Subcategory:          "subcat",
			Technology:           "tech",
			Description:          "my desc",
			Timeout:              5,
			TcpTimeout:           6,
			UdpTimeout:           7,
			TcpHalfClosedTimeout: 8,
			TcpTimeWaitTimeout:   9,
			Risk:                 2,
			FileTypeIdent:        true,
			VirusIdent:           true,
			DataIdent:            true,
			AlgDisableCapability: "alg disable",
			ParentApp:            "parent app",
		}},
		{"v1 default ip protocol", version.Number{7, 1, 0, ""}, Entry{
			Name:                                 "v1",
			DefaultType:                          DefaultTypeIpProtocol,
			DefaultIpProtocol:                    "101",
			Category:                             "category",
			Subcategory:                          "subcat",
			Technology:                           "tech",
			Description:                          "my desc",
			Timeout:                              5,
			TcpTimeout:                           6,
			UdpTimeout:                           7,
			TcpHalfClosedTimeout:                 8,
			TcpTimeWaitTimeout:                   9,
			Risk:                                 2,
			AbleToFileTransfer:                   true,
			ExcessiveBandwidth:                   true,
			TunnelsOtherApplications:             true,
			HasKnownVulnerability:                true,
			UsedByMalware:                        true,
			EvasiveBehavior:                      true,
			PervasiveUse:                         true,
			ProneToMisuse:                        true,
			ContinueScanningForOtherApplications: true,
			AlgDisableCapability:                 "alg disable",
			ParentApp:                            "parent app",
		}},
		{"v1 default icmp", version.Number{7, 1, 0, ""}, Entry{
			Name:                  "v1",
			DefaultType:           DefaultTypeIcmp,
			DefaultIcmpType:       55,
			DefaultIcmpCode:       44,
			Category:              "category",
			Subcategory:           "subcat",
			Technology:            "tech",
			Description:           "my desc",
			Timeout:               5,
			TcpTimeout:            6,
			UdpTimeout:            7,
			TcpHalfClosedTimeout:  8,
			TcpTimeWaitTimeout:    9,
			Risk:                  2,
			ExcessiveBandwidth:    true,
			HasKnownVulnerability: true,
			EvasiveBehavior:       true,
			ProneToMisuse:         true,
			FileTypeIdent:         true,
			DataIdent:             true,
			AlgDisableCapability:  "alg disable",
			ParentApp:             "parent app",
		}},
		{"v1 no default", version.Number{7, 1, 0, ""}, Entry{
			Name:                                 "v1",
			DefaultType:                          DefaultTypeIcmp6,
			DefaultIcmpType:                      33,
			DefaultIcmpCode:                      22,
			Category:                             "category",
			Subcategory:                          "subcat",
			Technology:                           "tech",
			Description:                          "my desc",
			Timeout:                              5,
			TcpTimeout:                           6,
			UdpTimeout:                           7,
			TcpHalfClosedTimeout:                 8,
			TcpTimeWaitTimeout:                   9,
			Risk:                                 2,
			AbleToFileTransfer:                   true,
			TunnelsOtherApplications:             true,
			UsedByMalware:                        true,
			PervasiveUse:                         true,
			ContinueScanningForOtherApplications: true,
			VirusIdent:                           true,
			AlgDisableCapability:                 "alg disable",
			ParentApp:                            "parent app",
		}},
		{"v2 no default", version.Number{8, 1, 0, ""}, Entry{
			Name:                                 "v2",
			DefaultType:                          DefaultTypeNone,
			Category:                             "category",
			Subcategory:                          "subcat",
			Technology:                           "tech",
			Description:                          "my desc",
			Timeout:                              5,
			TcpTimeout:                           6,
			UdpTimeout:                           7,
			TcpHalfClosedTimeout:                 8,
			TcpTimeWaitTimeout:                   9,
			Risk:                                 2,
			AbleToFileTransfer:                   true,
			ExcessiveBandwidth:                   true,
			TunnelsOtherApplications:             true,
			HasKnownVulnerability:                true,
			UsedByMalware:                        true,
			EvasiveBehavior:                      true,
			PervasiveUse:                         true,
			ProneToMisuse:                        true,
			ContinueScanningForOtherApplications: true,
			FileTypeIdent:                        true,
			VirusIdent:                           true,
			DataIdent:                            true,
			AlgDisableCapability:                 "alg disable",
			ParentApp:                            "parent app",
			NoAppIdCaching:                       true,
			raw: map[string]string{
				"sigs": "a bunch of sig data",
			},
		}},
		{"v2 default ports", version.Number{8, 1, 0, ""}, Entry{
			Name:                 "v2",
			DefaultType:          DefaultTypePort,
			DefaultPorts:         []string{"tcp/dynamic", "udp/32"},
			Category:             "category",
			Subcategory:          "subcat",
			Technology:           "tech",
			Description:          "my desc",
			Timeout:              5,
			TcpTimeout:           6,
			UdpTimeout:           7,
			TcpHalfClosedTimeout: 8,
			TcpTimeWaitTimeout:   9,
			Risk:                 2,
			FileTypeIdent:        true,
			VirusIdent:           true,
			DataIdent:            true,
			AlgDisableCapability: "alg disable",
			ParentApp:            "parent app",
			NoAppIdCaching:       true,
		}},
		{"v2 default ip protocol", version.Number{8, 1, 0, ""}, Entry{
			Name:                                 "v2",
			DefaultType:                          DefaultTypeIpProtocol,
			DefaultIpProtocol:                    "101",
			Category:                             "category",
			Subcategory:                          "subcat",
			Technology:                           "tech",
			Description:                          "my desc",
			Timeout:                              5,
			TcpTimeout:                           6,
			UdpTimeout:                           7,
			TcpHalfClosedTimeout:                 8,
			TcpTimeWaitTimeout:                   9,
			Risk:                                 2,
			AbleToFileTransfer:                   true,
			ExcessiveBandwidth:                   true,
			TunnelsOtherApplications:             true,
			HasKnownVulnerability:                true,
			UsedByMalware:                        true,
			EvasiveBehavior:                      true,
			PervasiveUse:                         true,
			ProneToMisuse:                        true,
			ContinueScanningForOtherApplications: true,
			AlgDisableCapability:                 "alg disable",
			ParentApp:                            "parent app",
		}},
		{"v2 default icmp", version.Number{8, 1, 0, ""}, Entry{
			Name:                  "v2",
			DefaultType:           DefaultTypeIcmp,
			DefaultIcmpType:       55,
			DefaultIcmpCode:       44,
			Category:              "category",
			Subcategory:           "subcat",
			Technology:            "tech",
			Description:           "my desc",
			Timeout:               5,
			TcpTimeout:            6,
			UdpTimeout:            7,
			TcpHalfClosedTimeout:  8,
			TcpTimeWaitTimeout:    9,
			Risk:                  2,
			ExcessiveBandwidth:    true,
			HasKnownVulnerability: true,
			EvasiveBehavior:       true,
			ProneToMisuse:         true,
			FileTypeIdent:         true,
			DataIdent:             true,
			AlgDisableCapability:  "alg disable",
			ParentApp:             "parent app",
			NoAppIdCaching:        true,
		}},
		{"v2 no default", version.Number{8, 1, 0, ""}, Entry{
			Name:                                 "v2",
			DefaultType:                          DefaultTypeIcmp6,
			DefaultIcmpType:                      33,
			DefaultIcmpCode:                      22,
			Category:                             "category",
			Subcategory:                          "subcat",
			Technology:                           "tech",
			Description:                          "my desc",
			Timeout:                              5,
			TcpTimeout:                           6,
			UdpTimeout:                           7,
			TcpHalfClosedTimeout:                 8,
			TcpTimeWaitTimeout:                   9,
			Risk:                                 2,
			AbleToFileTransfer:                   true,
			TunnelsOtherApplications:             true,
			UsedByMalware:                        true,
			PervasiveUse:                         true,
			ContinueScanningForOtherApplications: true,
			VirusIdent:                           true,
			AlgDisableCapability:                 "alg disable",
			ParentApp:                            "parent app",
		}},
	}
}
