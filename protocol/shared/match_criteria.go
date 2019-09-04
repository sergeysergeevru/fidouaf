package shared

//dictionary MatchCriteria {
//    AAID[]           aaid;
//    DOMString[]      vendorID;
//    KeyID[]          keyIDs;
//    unsigned long    userVerification;
//    unsigned short   keyProtection;
//    unsigned short   matcherProtection;
//    unsigned long    attachmentHint;
//    unsigned short   tcDisplay;
//    unsigned short[] authenticationAlgorithms;
//    DOMString[]      assertionSchemes;
//    unsigned short[] attestationTypes;
//    unsigned short   authenticatorVersion;
//    Extension[]      exts;
//};

type MatchCriteria struct {
	AAID AAID
	VendorID []string
	KeyIDs   []KeyID
	UserVerification uint64
	MatcherProtection uint16
	AttachmentHint uint64
	TcDisplay uint16
	AuthenticationAlgorithms []uint16
	AssertionSchemes []string
	AttestationTypes []uint16
	AuthenticatorVersion uint16
	Exts []Extension
}