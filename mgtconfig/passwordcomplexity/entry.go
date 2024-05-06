package passwordcomplexity

import (
	"github.com/fpluchorg/pango/util"
	"github.com/fpluchorg/pango/version"
)

// Entry is a normalized, version independent representation of a password complexity.
type Entry struct {
	MinimumLength                  int
	Enabled                        bool
	MinimumUppercaseLetters        int
	MinimumLowercaseLetters        int
	MinimumNumericLetters          int
	MinimumSpecialCharacters       int
	BlockRepeatedCharacters        int
	BlockUsernameInclusion         bool
	NewPasswordDiffersByCharacters int
	PasswordChangeOnFirstLogin     bool
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

	if o.Enabled != util.EmptyString {
		ans.Enabled = util.AsBool(o.Enabled)
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

	if o.BlockUsernameInclusion != util.EmptyString {
		ans.BlockUsernameInclusion = util.AsBool(o.BlockUsernameInclusion)
	}

	if o.NewPasswordDiffersByCharacters != nil {
		ans.NewPasswordDiffersByCharacters = *o.NewPasswordDiffersByCharacters
	}

	if o.PasswordChangeOnFirstLogin != util.EmptyString {
		ans.PasswordChangeOnFirstLogin = util.AsBool(o.PasswordChangeOnFirstLogin)
	}

	if o.PasswordHistoryCount != nil {
		ans.PasswordHistoryCount = *o.PasswordHistoryCount
	}

	if o.PasswordChangePeriodBlock != nil {
		ans.PasswordChangePeriodBlock = *o.PasswordChangePeriodBlock
	}

	if o.PasswordChange.ExpirationPeriod != nil {
		ans.ExpirationPeriod = *o.PasswordChange.ExpirationPeriod
	}

	if o.PasswordChange.ExpirationWarningPeriod != nil {
		ans.ExpirationWarningPeriod = *o.PasswordChange.ExpirationWarningPeriod
	}

	if o.PasswordChange.PostExpirationAdminLoginCount != nil {
		ans.PostExpirationAdminLoginCount = *o.PasswordChange.PostExpirationAdminLoginCount
	}

	if o.PasswordChange.PostExpirationGracePeriod != nil {
		ans.PostExpirationGracePeriod = *o.PasswordChange.PostExpirationGracePeriod
	}

	return ans
}

type entry_v1 struct {
	MinimumLength                  *int            `xml:"minimum-length,omitempty"`
	Enabled                        string          `xml:"enabled,omitempty"`
	MinimumUppercaseLetters        *int            `xml:"minimum-uppercase-letters,omitempty"`
	MinimumLowercaseLetters        *int            `xml:"minimum-lowercase-letters,omitempty"`
	MinimumNumericLetters          *int            `xml:"minimum-numeric-letters,omitempty"`
	MinimumSpecialCharacters       *int            `xml:"minimum-special-characters,omitempty"`
	BlockRepeatedCharacters        *int            `xml:"block-repeated-characters,omitempty"`
	BlockUsernameInclusion         string          `xml:"block-username-inclusion,omitempty"`
	NewPasswordDiffersByCharacters *int            `xml:"new-password-differs-by-characters,omitempty"`
	PasswordChangeOnFirstLogin     string          `xml:"password-change-on-first-login,omitempty"`
	PasswordHistoryCount           *int            `xml:"password-history-count,omitempty"`
	PasswordChangePeriodBlock      *int            `xml:"password-change-period-block,omitempty"`
	PasswordChange                 *passwordChange `xml:"password-change"`
}
type passwordChange struct {
	ExpirationPeriod              *int `xml:"expiration-period,omitempty"`
	ExpirationWarningPeriod       *int `xml:"expiration-warning-period,omitempty"`
	PostExpirationAdminLoginCount *int `xml:"post-expiration-admin-login-count,omitempty"`
	PostExpirationGracePeriod     *int `xml:"post-expiration-grace-period,omitempty"`
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
	}

	enabled := util.YesNo(e.Enabled)
	ans.Enabled = enabled

	blockUsernameInclusion := util.YesNo(e.BlockUsernameInclusion)
	ans.BlockUsernameInclusion = blockUsernameInclusion

	passwordChangeOnFirstLogin := util.YesNo(e.PasswordChangeOnFirstLogin)
	ans.PasswordChangeOnFirstLogin = passwordChangeOnFirstLogin

	if ans.PasswordChange == nil {
		ans.PasswordChange = &passwordChange{
			ExpirationPeriod:              &e.ExpirationPeriod,
			ExpirationWarningPeriod:       &e.ExpirationWarningPeriod,
			PostExpirationAdminLoginCount: &e.PostExpirationAdminLoginCount,
			PostExpirationGracePeriod:     &e.PostExpirationGracePeriod,
		}
	}

	return ans
}
