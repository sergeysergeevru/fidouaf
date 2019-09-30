package shared

//dictionary Policy {
//    required MatchCriteria[][] accepted;
//    MatchCriteria[]            disallowed;
//};

type Policy struct {
	Accepted [][]MatchCriteria
	Disallowed []MatchCriteria
}
