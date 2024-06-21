package p2p

var (
	DefaultTimeout         = int64(30)
	DefaultGatewayEndpoint = ":8080"
)

type Option interface {
	Apply(*Node2) error
}

type IsHandleCommiterSubReq bool

func (o IsHandleCommiterSubReq) Apply(s *Node2) error {
	s.isHandleSubmitterSubReq = bool(o)
	return nil
}
func WithIsHandleCommiterSubReq(isHandleCommiterSubReq bool) IsHandleCommiterSubReq {
	return IsHandleCommiterSubReq(isHandleCommiterSubReq)
}

type IsHandleOptimismSigReq bool

func (o IsHandleOptimismSigReq) Apply(s *Node2) error {
	s.isHandleSubmitterSubReq = bool(o)
	return nil
}
func WithIsHandleOptimismSigReq(isHandleOptimismSigReq bool) IsHandleOptimismSigReq {
	return IsHandleOptimismSigReq(isHandleOptimismSigReq)
}
