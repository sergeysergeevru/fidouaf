package shared

//dictionary FinalChallengeParams {
//    required DOMString       appID;
//    required ServerChallenge challenge;
//    required DOMString       facetID;
//    required ChannelBinding  channelBinding;
//};

type FinalChallengeParams struct {
	AppId          string
	Challenge      ServerChallenge
	FacetID        string
	ChannelBinding ChannelBinding
}