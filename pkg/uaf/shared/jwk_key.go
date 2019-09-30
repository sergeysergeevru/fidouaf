package shared

//dictionary JwkKey {
//    required DOMString kty = "EC";
//    required DOMString crv = "P-256";
//    required DOMString x;
//    required DOMString y;
//};

//TODO: to investigate best practice for the default value of structures
type JwkKey struct {
	Kty string //="EC";
	Crv string //="P-256";
	X string
	Y string
}