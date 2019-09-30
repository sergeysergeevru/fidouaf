package shared

import (
	"github.com/sergeysergeevru/fidouaf/pkg/uaf/metadata"
	)

//dictionary AuthenticatorRegistrationAssertion {
//    required DOMString                    assertionScheme;
//    required DOMString                    assertion;
//    DisplayPNGCharacteristicsDescriptor[] tcDisplayPNGCharacteristics;
//    Extension[]                           exts;
//};

type AuthenticatorRegistrationAssertion struct {
	AssertionScheme string
	Assertion string
	TcDisplayPNGCharacteristics []metadata.DisplayPNGCharacteristicsDescriptor
	Exts []Extension
}