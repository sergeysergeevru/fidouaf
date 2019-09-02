package shared

//dictionary OperationHeader {
//    required Version   upv;
//    required Operation op;
//    DOMString          appID;
//    DOMString          serverData;
//    Extension[]        exts;
//};

type OperationHeader struct {
	Upv Version
	Op Operation
	AppID string
	ServerData string
	Exts []Extension
}
