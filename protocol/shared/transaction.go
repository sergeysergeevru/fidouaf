package shared

import "github.com/sergeysergeevru/fidouaf/protocol/metadata"

//dictionary Transaction {
//    required DOMString                  contentType;
//    required DOMString                  content;
//    DisplayPNGCharacteristicsDescriptor tcDisplayPNGCharacteristics;
//};

type Transaction struct {
	ContentType string
	Content string
	TcDisplayPNGCharacteristics metadata.DisplayPNGCharacteristicsDescriptor
}