package shared

//dictionary AuthenticationRequest {
//    required OperationHeader header;
//    required ServerChallenge challenge;
//    Transaction[]            transaction;
//    required Policy          policy;
//};

type AuthenticationRequest struct {
	Header      OperationHeader
	Challenge   ServerChallenge
	Transaction []Transaction
	Policy      Policy
}