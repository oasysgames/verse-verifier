package gen

func (x *MsgCommonTopic) TryGetReqSubmitterTopicSub() (*ReqSubmitterTopicSub, bool) {
	if x, ok := x.GetBody().(*MsgCommonTopic_ReqSubmitterTopicSub); ok {
		return x.ReqSubmitterTopicSub, true
	}
	return nil, false
}

func (x *MsgSubmitterTopic) TryGetReqOptimismSignature() (*ReqOptimismSignature, bool) {
	if x, ok := x.GetBody().(*MsgSubmitterTopic_ReqOptimismSignature); ok {
		return x.ReqOptimismSignature, true
	}
	return nil, false
}

func (x *MsgSubmitterTopic) TryGetPubOptimismSignature() (*PubOptimismSignature, bool) {
	if x, ok := x.GetBody().(*MsgSubmitterTopic_PubOptimismSignature); ok {
		return x.PubOptimismSignature, true
	}
	return nil, false
}
