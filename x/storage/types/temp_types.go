package types

import "time"

type StoragePaymentInfo struct {
	Start          time.Time                                `protobuf:"bytes,1,opt,name=start,proto3,stdtime" json:"start"`
	End            time.Time                                `protobuf:"bytes,2,opt,name=end,proto3,stdtime" json:"end"`
	SpaceAvailable int64                                    `protobuf:"varint,3,opt,name=spaceAvailable,proto3" json:"spaceAvailable,omitempty"`
	SpaceUsed      int64                                    `protobuf:"varint,4,opt,name=spaceUsed,proto3" json:"spaceUsed,omitempty"`
	Address        string                                   `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Coins          github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

type MsgBuyStorage struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Receiver     string `protobuf:"bytes,2,opt,name=receiver,json=receiver,proto3" json:"receiver,omitempty"`
	Duration     int64  `protobuf:"varint,3,opt,name=duration_days,json=durationDays,proto3" json:"duration_days,omitempty"`
	Bytes        int64  `protobuf:"varint,4,opt,name=bytes,proto3" json:"bytes,omitempty"`
	PaymentDenom string `protobuf:"bytes,5,opt,name=payment_denom,json=paymentDenom,proto3" json:"payment_denom,omitempty"`
	Referral     string `protobuf:"bytes,6,opt,name=referral,proto3" json:"referral,omitempty"`
}
type MsgBuyStorageResponse struct {
}
