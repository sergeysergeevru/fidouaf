package shared

//dictionary RegistrationRequest {
//    required OperationHeader header;
//    required ServerChallenge challenge;
//    required DOMString       username;
//    required Policy          policy;
//};

type RegistrationRequest struct {
	Header OperationHeader
	Challenge ServerChallenge
	Username string
	Policy Policy
}
