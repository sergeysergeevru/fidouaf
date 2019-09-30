package shared

// typedef DOMString AAID;

type AAID string

// The AAID is a string with format "V#M", where
// "#" is a separator
// "V" indicates the authenticator Vendor Code. This code consists of 4 hexadecimal digits.
// "M" indicates the authenticator Model Code. This code consists of 4 hexadecimal digits.