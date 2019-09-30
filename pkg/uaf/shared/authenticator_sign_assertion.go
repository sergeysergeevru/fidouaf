package shared

//dictionary AuthenticatorSignAssertion {
//    required DOMString assertionScheme;
//    required DOMString assertion;
//    Extension[]        exts;
//};

type AuthenticatorSignAssertion struct {
	AssertionScheme string
	Assertion string
	Exts []Extension
}
