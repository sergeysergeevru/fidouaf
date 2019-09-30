package shared

//dictionary DeregisterAuthenticator {
//    required AAID  aaid;
//    required KeyID keyID;
//};

type DeregisterAuthenticator struct {
	AAID  AAID
	KeyID KeyID
}
