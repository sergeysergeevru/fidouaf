package shared

//dictionary DeregistrationRequest {
//    required OperationHeader           header;
//    required DeregisterAuthenticator[] authenticators;
//};

type DeregistrationRequest struct {
	Header OperationHeader
	Authenticators []DeregisterAuthenticator
}