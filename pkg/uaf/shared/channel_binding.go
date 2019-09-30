package shared

//dictionary ChannelBinding {
//    DOMString serverEndPoint;
//    DOMString tlsServerCertificate;
//    DOMString tlsUnique;
//    DOMString cid_pubkey;
//};

type ChannelBinding struct {
	ServerEndPoint string
	TlsServerCertificate string
	TlsUnique string
	CidPubKey string
}