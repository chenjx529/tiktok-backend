// Code generated by Kitex v0.7.0. DO NOT EDIT.

package messageservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	message "tiktok-backend/kitex_gen/message"
)

func serviceInfo() *kitex.ServiceInfo {
	return messageServiceServiceInfo
}

var messageServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MessageService"
	handlerType := (*message.MessageService)(nil)
	methods := map[string]kitex.MethodInfo{
		"MessageChat":    kitex.NewMethodInfo(messageChatHandler, newMessageChatArgs, newMessageChatResult, false),
		"MessageActioin": kitex.NewMethodInfo(messageActioinHandler, newMessageActioinArgs, newMessageActioinResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "message",
		"ServiceFilePath": "",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.7.0",
		Extra:           extra,
	}
	return svcInfo
}

func messageChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(message.DouyinMessageChatRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(message.MessageService).MessageChat(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MessageChatArgs:
		success, err := handler.(message.MessageService).MessageChat(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MessageChatResult)
		realResult.Success = success
	}
	return nil
}
func newMessageChatArgs() interface{} {
	return &MessageChatArgs{}
}

func newMessageChatResult() interface{} {
	return &MessageChatResult{}
}

type MessageChatArgs struct {
	Req *message.DouyinMessageChatRequest
}

func (p *MessageChatArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(message.DouyinMessageChatRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MessageChatArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MessageChatArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MessageChatArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MessageChatArgs) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(message.DouyinMessageChatRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessageChatArgs_Req_DEFAULT *message.DouyinMessageChatRequest

func (p *MessageChatArgs) GetReq() *message.DouyinMessageChatRequest {
	if !p.IsSetReq() {
		return MessageChatArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessageChatArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MessageChatArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MessageChatResult struct {
	Success *message.DouyinMessageChatResponse
}

var MessageChatResult_Success_DEFAULT *message.DouyinMessageChatResponse

func (p *MessageChatResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(message.DouyinMessageChatResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MessageChatResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MessageChatResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MessageChatResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MessageChatResult) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(message.DouyinMessageChatResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessageChatResult) GetSuccess() *message.DouyinMessageChatResponse {
	if !p.IsSetSuccess() {
		return MessageChatResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessageChatResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.DouyinMessageChatResponse)
}

func (p *MessageChatResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessageChatResult) GetResult() interface{} {
	return p.Success
}

func messageActioinHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(message.DouyinRelationActionRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(message.MessageService).MessageActioin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MessageActioinArgs:
		success, err := handler.(message.MessageService).MessageActioin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MessageActioinResult)
		realResult.Success = success
	}
	return nil
}
func newMessageActioinArgs() interface{} {
	return &MessageActioinArgs{}
}

func newMessageActioinResult() interface{} {
	return &MessageActioinResult{}
}

type MessageActioinArgs struct {
	Req *message.DouyinRelationActionRequest
}

func (p *MessageActioinArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(message.DouyinRelationActionRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MessageActioinArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MessageActioinArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MessageActioinArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MessageActioinArgs) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(message.DouyinRelationActionRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MessageActioinArgs_Req_DEFAULT *message.DouyinRelationActionRequest

func (p *MessageActioinArgs) GetReq() *message.DouyinRelationActionRequest {
	if !p.IsSetReq() {
		return MessageActioinArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MessageActioinArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MessageActioinArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MessageActioinResult struct {
	Success *message.DouyinRelationActionResponse
}

var MessageActioinResult_Success_DEFAULT *message.DouyinRelationActionResponse

func (p *MessageActioinResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(message.DouyinRelationActionResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MessageActioinResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MessageActioinResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MessageActioinResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MessageActioinResult) Unmarshal(in []byte) error {
	if len(in) == 0 {
		return nil
	}
	msg := new(message.DouyinRelationActionResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MessageActioinResult) GetSuccess() *message.DouyinRelationActionResponse {
	if !p.IsSetSuccess() {
		return MessageActioinResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MessageActioinResult) SetSuccess(x interface{}) {
	p.Success = x.(*message.DouyinRelationActionResponse)
}

func (p *MessageActioinResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MessageActioinResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) MessageChat(ctx context.Context, Req *message.DouyinMessageChatRequest) (r *message.DouyinMessageChatResponse, err error) {
	var _args MessageChatArgs
	_args.Req = Req
	var _result MessageChatResult
	if err = p.c.Call(ctx, "MessageChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessageActioin(ctx context.Context, Req *message.DouyinRelationActionRequest) (r *message.DouyinRelationActionResponse, err error) {
	var _args MessageActioinArgs
	_args.Req = Req
	var _result MessageActioinResult
	if err = p.c.Call(ctx, "MessageActioin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
