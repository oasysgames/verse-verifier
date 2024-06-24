package p2p

var (
	DefaultTimeout         = int64(30)
	DefaultGatewayEndpoint = ":8080"
)

type Option interface {
	Apply(*Node2) error
}

type IsHandlingSignatureRequest bool

func (o IsHandlingSignatureRequest) Apply(s *Node2) error {
	s.isHandlingSignatureRequest = bool(o)
	return nil
}
func WithIsHandlingSignatureRequest(isHandlingSignatureRequest bool) IsHandlingSignatureRequest {
	return IsHandlingSignatureRequest(isHandlingSignatureRequest)
}

type IsHandlingPublishedSignatures bool

func (o IsHandlingPublishedSignatures) Apply(s *Node2) error {
	s.isHandlingPublishedSignatures = bool(o)
	return nil
}
func WithIsHandlingPublishedSignatures(isHandlingPublishedSignatures bool) IsHandlingPublishedSignatures {
	return IsHandlingPublishedSignatures(isHandlingPublishedSignatures)
}
