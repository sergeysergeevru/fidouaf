package shared

//dictionary Extension {
//    required DOMString id;
//    required DOMString data;
//    required boolean   fail_if_unknown;
//};

type Extension struct {
	Id string
	Data string
	FailIfUnknown bool
}