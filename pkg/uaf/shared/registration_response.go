package shared

//dictionary RegistrationResponse {
//    required OperationHeader                      header;
//    required DOMString                            fcParams;
//    required AuthenticatorRegistrationAssertion[] assertions;
//};

type RegistrationResponse struct {
	Header     OperationHeader
	FcParams   string
	Assertions []AuthenticatorRegistrationAssertion
}