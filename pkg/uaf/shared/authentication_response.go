package shared

//dictionary AuthenticationResponse {
//    required OperationHeader              header;
//    required DOMString                    fcParams;
//    required AuthenticatorSignAssertion[] assertions;
//};

type AuthenticationResponse struct {
	Header     OperationHeader
	FcParams   string
	Assertions []AuthenticatorSignAssertion
}