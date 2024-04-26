package passwordcomplexity

import (
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
	"strconv"
)

// Entry is a normalized, version independent representation of a password complexity.
type Entry struct {
	MinimumLength                  int
	Enabled                        string
	MinimumUppercaseLetters        int
	MinimumLowercaseLetters        int
	MinimumNumericLetters          int
	MinimumSpecialCharacters       int
	BlockRepeatedCharacters        int
	BlockUsernameInclusion         string
	NewPasswordDiffersByCharacters int
	PasswordChangeOnFirstLogin     string
	PasswordHistoryCount           int
	PasswordChangePeriodBlock      int
	ExpirationPeriod               int
	ExpirationWarningPeriod        int
	PostExpirationAdminLoginCount  int
	PostExpirationGracePeriod      int
}

// Copy copies the information from source's Entry `s` to this object.  As the
// Name field relates to the XPATH of this object, this field is not copied.
func (o *Entry) Copy(s Entry) {
	o.MinimumLength = s.MinimumLength
	o.Enabled = s.Enabled
	o.MinimumUppercaseLetters = s.MinimumUppercaseLetters
	o.MinimumLowercaseLetters = s.MinimumLowercaseLetters
	o.MinimumNumericLetters = s.MinimumNumericLetters
	o.MinimumSpecialCharacters = s.MinimumSpecialCharacters
	o.BlockRepeatedCharacters = s.BlockRepeatedCharacters
	o.BlockUsernameInclusion = s.BlockUsernameInclusion
	o.NewPasswordDiffersByCharacters = s.NewPasswordDiffersByCharacters
	o.PasswordChangeOnFirstLogin = s.PasswordChangeOnFirstLogin
	o.PasswordHistoryCount = s.PasswordHistoryCount
	o.PasswordChangePeriodBlock = s.PasswordChangePeriodBlock
	o.ExpirationPeriod = s.ExpirationPeriod
	o.ExpirationWarningPeriod = s.ExpirationWarningPeriod
	o.PostExpirationAdminLoginCount = s.PostExpirationAdminLoginCount
	o.PostExpirationGracePeriod = s.PostExpirationGracePeriod
}

/** Structs / functions for normalization. **/

func (o Entry) Specify(v version.Number) (string, interface{}) {
	_, fn := versioning(v)
	return util.EmptyString, fn(o)
}

type normalizer interface {
	Normalize() Entry
}

func (o *entry_v1) Normalize() Entry {
	return o.normalize()
}

func (o *entry_v1) normalize() Entry {
	ans := Entry{}

	if o.MinimumLength != nil {
		ans.MinimumLength = *o.MinimumLength
	}

	if o.Enabled != nil {
		ans.Enabled = strconv.FormatBool(util.AsBool(*o.Enabled))
	}

	if o.MinimumUppercaseLetters != nil {
		ans.MinimumUppercaseLetters = *o.MinimumUppercaseLetters
	}

	if o.MinimumLowercaseLetters != nil {
		ans.MinimumLowercaseLetters = *o.MinimumLowercaseLetters
	}

	if o.MinimumNumericLetters != nil {
		ans.MinimumNumericLetters = *o.MinimumNumericLetters
	}

	if o.MinimumSpecialCharacters != nil {
		ans.MinimumSpecialCharacters = *o.MinimumSpecialCharacters
	}

	if o.BlockRepeatedCharacters != nil {
		ans.BlockRepeatedCharacters = *o.BlockRepeatedCharacters
	}

	if o.BlockUsernameInclusion != nil {
		ans.BlockUsernameInclusion = strconv.FormatBool(util.AsBool(*o.BlockUsernameInclusion))
	}

	if o.NewPasswordDiffersByCharacters != nil {
		ans.NewPasswordDiffersByCharacters = *o.NewPasswordDiffersByCharacters
	}

	if o.PasswordChangeOnFirstLogin != nil {
		ans.PasswordChangeOnFirstLogin = strconv.FormatBool(util.AsBool(*o.PasswordChangeOnFirstLogin))
	}

	if o.PasswordHistoryCount != nil {
		ans.PasswordHistoryCount = *o.PasswordHistoryCount
	}

	if o.PasswordChangePeriodBlock != nil {
		ans.PasswordChangePeriodBlock = *o.PasswordChangePeriodBlock
	}

	if o.ExpirationPeriod != nil {
		ans.ExpirationPeriod = *o.ExpirationPeriod
	}

	if o.ExpirationWarningPeriod != nil {
		ans.ExpirationWarningPeriod = *o.ExpirationWarningPeriod
	}

	if o.PostExpirationAdminLoginCount != nil {
		ans.PostExpirationAdminLoginCount = *o.PostExpirationAdminLoginCount
	}

	if o.PostExpirationGracePeriod != nil {
		ans.PostExpirationGracePeriod = *o.PostExpirationGracePeriod
	}

	return ans
}

type entry_v1 struct {
	MinimumLength                  *int    `xml:"minimum-length"`
	Enabled                        *string `xml:"enabled"`
	MinimumUppercaseLetters        *int    `xml:"minimum-uppercase-letters"`
	MinimumLowercaseLetters        *int    `xml:"minimum-lowercase-letters"`
	MinimumNumericLetters          *int    `xml:"minimum-numeric-letters"`
	MinimumSpecialCharacters       *int    `xml:"minimum-special-characters"`
	BlockRepeatedCharacters        *int    `xml:"block-repeated-characters"`
	BlockUsernameInclusion         *string `xml:"block-username-inclusion"`
	NewPasswordDiffersByCharacters *int    `xml:"new-password-differs-by-characters"`
	PasswordChangeOnFirstLogin     *string `xml:"password-change-on-first-login"`
	PasswordHistoryCount           *int    `xml:"password-history-count"`
	PasswordChangePeriodBlock      *int    `xml:"password-change-period-block"`
	ExpirationPeriod               *int    `xml:"password-change>expiration-period,omitempty"`
	ExpirationWarningPeriod        *int    `xml:"password-change>expiration-warning-period,omitempty"`
	PostExpirationAdminLoginCount  *int    `xml:"password-change>post-expiration-admin-login-count,omitempty"`
	PostExpirationGracePeriod      *int    `xml:"password-change>post-expiration-grace-period,omitempty"`
}

func specify_v1(e Entry) interface{} {
	ans := entry_v1{
		MinimumLength:                  &e.MinimumLength,
		MinimumUppercaseLetters:        &e.MinimumUppercaseLetters,
		MinimumLowercaseLetters:        &e.MinimumLowercaseLetters,
		MinimumNumericLetters:          &e.MinimumNumericLetters,
		MinimumSpecialCharacters:       &e.MinimumSpecialCharacters,
		BlockRepeatedCharacters:        &e.BlockRepeatedCharacters,
		NewPasswordDiffersByCharacters: &e.NewPasswordDiffersByCharacters,
		PasswordHistoryCount:           &e.PasswordHistoryCount,
		PasswordChangePeriodBlock:      &e.PasswordChangePeriodBlock,
		ExpirationPeriod:               &e.ExpirationPeriod,
		ExpirationWarningPeriod:        &e.ExpirationWarningPeriod,
		PostExpirationAdminLoginCount:  &e.PostExpirationAdminLoginCount,
		PostExpirationGracePeriod:      &e.PostExpirationGracePeriod,
		Enabled:                        boolStringToYesNo(e.Enabled),
		BlockUsernameInclusion:         boolStringToYesNo(e.BlockUsernameInclusion),
		PasswordChangeOnFirstLogin:     boolStringToYesNo(e.PasswordChangeOnFirstLogin),
	}

	return ans
}

func boolStringToYesNo(str string) *string {
	if str == util.EmptyString {
		str = "false"
	}
	boolValue, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	boolString := util.YesNo(boolValue)
	return &boolString
}
