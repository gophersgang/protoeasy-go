// Code generated by protoc-gen-gogo.
// source: google/pubsub/v1beta2/pubsub.proto
// DO NOT EDIT!

package google_pubsub_v1beta2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

// A topic resource.
type Topic struct {
	// Name of the topic.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{0} }

// A message data and its attributes.
type PubsubMessage struct {
	// The message payload. For JSON requests, the value of this field must be
	// base64-encoded.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Optional attributes for this message.
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of this message assigned by the server at publication time. Guaranteed
	// to be unique within the topic. This value may be read by a subscriber
	// that receives a PubsubMessage via a Pull call or a push delivery. It must
	// not be populated by a publisher in a Publish call.
	MessageId string `protobuf:"bytes,3,opt,name=message_id,proto3" json:"message_id,omitempty"`
}

func (m *PubsubMessage) Reset()                    { *m = PubsubMessage{} }
func (m *PubsubMessage) String() string            { return proto.CompactTextString(m) }
func (*PubsubMessage) ProtoMessage()               {}
func (*PubsubMessage) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{1} }

func (m *PubsubMessage) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Request for the GetTopic method.
type GetTopicRequest struct {
	// The name of the topic to get.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (m *GetTopicRequest) Reset()                    { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()               {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{2} }

// Request for the Publish method.
type PublishRequest struct {
	// The messages in the request will be published on this topic.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// The messages to publish.
	Messages []*PubsubMessage `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{3} }

func (m *PublishRequest) GetMessages() []*PubsubMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Response for the Publish method.
type PublishResponse struct {
	// The server-assigned ID of each published message, in the same order as
	// the messages in the request. IDs are guaranteed to be unique within
	// the topic.
	MessageIds []string `protobuf:"bytes,1,rep,name=message_ids" json:"message_ids,omitempty"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{4} }

// Request for the ListTopics method.
type ListTopicsRequest struct {
	// The name of the cloud project that topics belong to.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Maximum number of topics to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,proto3" json:"page_size,omitempty"`
	// The value returned by the last ListTopicsResponse; indicates that this is
	// a continuation of a prior ListTopics call, and that the system should
	// return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,proto3" json:"page_token,omitempty"`
}

func (m *ListTopicsRequest) Reset()                    { *m = ListTopicsRequest{} }
func (m *ListTopicsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTopicsRequest) ProtoMessage()               {}
func (*ListTopicsRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{5} }

// Response for the ListTopics method.
type ListTopicsResponse struct {
	// The resulting topics.
	Topics []*Topic `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
	// If not empty, indicates that there may be more topics that match the
	// request; this value should be passed in a new ListTopicsRequest.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,proto3" json:"next_page_token,omitempty"`
}

func (m *ListTopicsResponse) Reset()                    { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()               {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{6} }

func (m *ListTopicsResponse) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

// Request for the ListTopicSubscriptions method.
type ListTopicSubscriptionsRequest struct {
	// The name of the topic that subscriptions are attached to.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Maximum number of subscription names to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,proto3" json:"page_size,omitempty"`
	// The value returned by the last ListTopicSubscriptionsResponse; indicates
	// that this is a continuation of a prior ListTopicSubscriptions call, and
	// that the system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,proto3" json:"page_token,omitempty"`
}

func (m *ListTopicSubscriptionsRequest) Reset()         { *m = ListTopicSubscriptionsRequest{} }
func (m *ListTopicSubscriptionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTopicSubscriptionsRequest) ProtoMessage()    {}
func (*ListTopicSubscriptionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPubsub, []int{7}
}

// Response for the ListTopicSubscriptions method.
type ListTopicSubscriptionsResponse struct {
	// The names of the subscriptions that match the request.
	Subscriptions []string `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// If not empty, indicates that there may be more subscriptions that match
	// the request; this value should be passed in a new
	// ListTopicSubscriptionsRequest to get more subscriptions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,proto3" json:"next_page_token,omitempty"`
}

func (m *ListTopicSubscriptionsResponse) Reset()         { *m = ListTopicSubscriptionsResponse{} }
func (m *ListTopicSubscriptionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTopicSubscriptionsResponse) ProtoMessage()    {}
func (*ListTopicSubscriptionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPubsub, []int{8}
}

// Request for the DeleteTopic method.
type DeleteTopicRequest struct {
	// Name of the topic to delete.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (m *DeleteTopicRequest) Reset()                    { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()               {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{9} }

// A subscription resource.
type Subscription struct {
	// Name of the subscription.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The name of the topic from which this subscription is receiving messages.
	// This will be present if and only if the subscription has not been detached
	// from its topic.
	Topic string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	// If push delivery is used with this subscription, this field is
	// used to configure it. An empty pushConfig signifies that the subscriber
	// will pull and ack messages using API methods.
	PushConfig *PushConfig `protobuf:"bytes,4,opt,name=push_config" json:"push_config,omitempty"`
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	//
	// For pull delivery this value
	// is used as the initial value for the ack deadline. It may be overridden
	// for a specific message by calling ModifyAckDeadline.
	//
	// For push delivery, this value is also used to set the request timeout for
	// the call to the push endpoint.
	//
	// If the subscriber never acknowledges the message, the Pub/Sub
	// system will eventually redeliver the message.
	AckDeadlineSeconds int32 `protobuf:"varint,5,opt,name=ack_deadline_seconds,proto3" json:"ack_deadline_seconds,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{10} }

func (m *Subscription) GetPushConfig() *PushConfig {
	if m != nil {
		return m.PushConfig
	}
	return nil
}

// Configuration for a push delivery endpoint.
type PushConfig struct {
	// A URL locating the endpoint to which messages should be pushed.
	// For example, a Webhook endpoint might use "https://example.com/push".
	PushEndpoint string `protobuf:"bytes,1,opt,name=push_endpoint,proto3" json:"push_endpoint,omitempty"`
	// Endpoint configuration attributes.
	//
	// Every endpoint has a set of API supported attributes that can be used to
	// control different aspects of the message delivery.
	//
	// The currently supported attribute is `x-goog-version`, which you can
	// use to change the format of the push message. This attribute
	// indicates the version of the data expected by the endpoint. This
	// controls the shape of the envelope (i.e. its fields and metadata).
	// The endpoint version is based on the version of the Pub/Sub
	// API.
	//
	// If not present during the CreateSubscription call, it will default to
	// the version of the API used to make such call. If not present during a
	// ModifyPushConfig call, its value will not be changed. GetSubscription
	// calls will always return a valid version, even if the subscription was
	// created without this attribute.
	//
	// The possible values for this attribute are:
	//
	// * `v1beta1`: uses the push format defined in the v1beta1 Pub/Sub API.
	// * `v1beta2`: uses the push format defined in the v1beta2 Pub/Sub API.
	//
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *PushConfig) Reset()                    { *m = PushConfig{} }
func (m *PushConfig) String() string            { return proto.CompactTextString(m) }
func (*PushConfig) ProtoMessage()               {}
func (*PushConfig) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{11} }

func (m *PushConfig) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A message and its corresponding acknowledgment ID.
type ReceivedMessage struct {
	// This ID can be used to acknowledge the received message.
	AckId string `protobuf:"bytes,1,opt,name=ack_id,proto3" json:"ack_id,omitempty"`
	// The message.
	Message *PubsubMessage `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ReceivedMessage) Reset()                    { *m = ReceivedMessage{} }
func (m *ReceivedMessage) String() string            { return proto.CompactTextString(m) }
func (*ReceivedMessage) ProtoMessage()               {}
func (*ReceivedMessage) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{12} }

func (m *ReceivedMessage) GetMessage() *PubsubMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

// Request for the GetSubscription method.
type GetSubscriptionRequest struct {
	// The name of the subscription to get.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (m *GetSubscriptionRequest) Reset()                    { *m = GetSubscriptionRequest{} }
func (m *GetSubscriptionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSubscriptionRequest) ProtoMessage()               {}
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{13} }

// Request for the ListSubscriptions method.
type ListSubscriptionsRequest struct {
	// The name of the cloud project that subscriptions belong to.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Maximum number of subscriptions to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,proto3" json:"page_size,omitempty"`
	// The value returned by the last ListSubscriptionsResponse; indicates that
	// this is a continuation of a prior ListSubscriptions call, and that the
	// system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,proto3" json:"page_token,omitempty"`
}

func (m *ListSubscriptionsRequest) Reset()                    { *m = ListSubscriptionsRequest{} }
func (m *ListSubscriptionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSubscriptionsRequest) ProtoMessage()               {}
func (*ListSubscriptionsRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{14} }

// Response for the ListSubscriptions method.
type ListSubscriptionsResponse struct {
	// The subscriptions that match the request.
	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// If not empty, indicates that there may be more subscriptions that match
	// the request; this value should be passed in a new ListSubscriptionsRequest
	// to get more subscriptions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,proto3" json:"next_page_token,omitempty"`
}

func (m *ListSubscriptionsResponse) Reset()                    { *m = ListSubscriptionsResponse{} }
func (m *ListSubscriptionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSubscriptionsResponse) ProtoMessage()               {}
func (*ListSubscriptionsResponse) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{15} }

func (m *ListSubscriptionsResponse) GetSubscriptions() []*Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

// Request for the DeleteSubscription method.
type DeleteSubscriptionRequest struct {
	// The subscription to delete.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (m *DeleteSubscriptionRequest) Reset()                    { *m = DeleteSubscriptionRequest{} }
func (m *DeleteSubscriptionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubscriptionRequest) ProtoMessage()               {}
func (*DeleteSubscriptionRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{16} }

// Request for the ModifyPushConfig method.
type ModifyPushConfigRequest struct {
	// The name of the subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// The push configuration for future deliveries.
	//
	// An empty pushConfig indicates that the Pub/Sub system should
	// stop pushing messages from the given subscription and allow
	// messages to be pulled and acknowledged - effectively pausing
	// the subscription if Pull is not called.
	PushConfig *PushConfig `protobuf:"bytes,2,opt,name=push_config" json:"push_config,omitempty"`
}

func (m *ModifyPushConfigRequest) Reset()                    { *m = ModifyPushConfigRequest{} }
func (m *ModifyPushConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyPushConfigRequest) ProtoMessage()               {}
func (*ModifyPushConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{17} }

func (m *ModifyPushConfigRequest) GetPushConfig() *PushConfig {
	if m != nil {
		return m.PushConfig
	}
	return nil
}

// Request for the Pull method.
type PullRequest struct {
	// The subscription from which messages should be pulled.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// If this is specified as true the system will respond immediately even if
	// it is not able to return a message in the Pull response. Otherwise the
	// system is allowed to wait until at least one message is available rather
	// than returning no messages. The client may cancel the request if it does
	// not wish to wait any longer for the response.
	ReturnImmediately bool `protobuf:"varint,2,opt,name=return_immediately,proto3" json:"return_immediately,omitempty"`
	// The maximum number of messages returned for this request. The Pub/Sub
	// system may return fewer than the number specified.
	MaxMessages int32 `protobuf:"varint,3,opt,name=max_messages,proto3" json:"max_messages,omitempty"`
}

func (m *PullRequest) Reset()                    { *m = PullRequest{} }
func (m *PullRequest) String() string            { return proto.CompactTextString(m) }
func (*PullRequest) ProtoMessage()               {}
func (*PullRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{18} }

// Response for the Pull method.
type PullResponse struct {
	// Received Pub/Sub messages. The Pub/Sub system will return zero messages if
	// there are no more available in the backlog. The Pub/Sub system may return
	// fewer than the maxMessages requested even if there are more messages
	// available in the backlog.
	ReceivedMessages []*ReceivedMessage `protobuf:"bytes,1,rep,name=received_messages" json:"received_messages,omitempty"`
}

func (m *PullResponse) Reset()                    { *m = PullResponse{} }
func (m *PullResponse) String() string            { return proto.CompactTextString(m) }
func (*PullResponse) ProtoMessage()               {}
func (*PullResponse) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{19} }

func (m *PullResponse) GetReceivedMessages() []*ReceivedMessage {
	if m != nil {
		return m.ReceivedMessages
	}
	return nil
}

// Request for the ModifyAckDeadline method.
type ModifyAckDeadlineRequest struct {
	// The name of the subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// The acknowledgment ID.
	AckId string `protobuf:"bytes,2,opt,name=ack_id,proto3" json:"ack_id,omitempty"`
	// The new ack deadline with respect to the time this request was sent to the
	// Pub/Sub system. Must be >= 0. For example, if the value is 10, the new ack
	// deadline will expire 10 seconds after the ModifyAckDeadline call was made.
	// Specifying zero may immediately make the message available for another pull
	// request.
	AckDeadlineSeconds int32 `protobuf:"varint,3,opt,name=ack_deadline_seconds,proto3" json:"ack_deadline_seconds,omitempty"`
}

func (m *ModifyAckDeadlineRequest) Reset()                    { *m = ModifyAckDeadlineRequest{} }
func (m *ModifyAckDeadlineRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyAckDeadlineRequest) ProtoMessage()               {}
func (*ModifyAckDeadlineRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{20} }

// Request for the Acknowledge method.
type AcknowledgeRequest struct {
	// The subscription whose message is being acknowledged.
	Subscription string `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
	// The acknowledgment ID for the messages being acknowledged that was returned
	// by the Pub/Sub system in the Pull response. Must not be empty.
	AckIds []string `protobuf:"bytes,2,rep,name=ack_ids" json:"ack_ids,omitempty"`
}

func (m *AcknowledgeRequest) Reset()                    { *m = AcknowledgeRequest{} }
func (m *AcknowledgeRequest) String() string            { return proto.CompactTextString(m) }
func (*AcknowledgeRequest) ProtoMessage()               {}
func (*AcknowledgeRequest) Descriptor() ([]byte, []int) { return fileDescriptorPubsub, []int{21} }

func init() {
	proto.RegisterType((*Topic)(nil), "google.pubsub.v1beta2.Topic")
	proto.RegisterType((*PubsubMessage)(nil), "google.pubsub.v1beta2.PubsubMessage")
	proto.RegisterType((*GetTopicRequest)(nil), "google.pubsub.v1beta2.GetTopicRequest")
	proto.RegisterType((*PublishRequest)(nil), "google.pubsub.v1beta2.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "google.pubsub.v1beta2.PublishResponse")
	proto.RegisterType((*ListTopicsRequest)(nil), "google.pubsub.v1beta2.ListTopicsRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "google.pubsub.v1beta2.ListTopicsResponse")
	proto.RegisterType((*ListTopicSubscriptionsRequest)(nil), "google.pubsub.v1beta2.ListTopicSubscriptionsRequest")
	proto.RegisterType((*ListTopicSubscriptionsResponse)(nil), "google.pubsub.v1beta2.ListTopicSubscriptionsResponse")
	proto.RegisterType((*DeleteTopicRequest)(nil), "google.pubsub.v1beta2.DeleteTopicRequest")
	proto.RegisterType((*Subscription)(nil), "google.pubsub.v1beta2.Subscription")
	proto.RegisterType((*PushConfig)(nil), "google.pubsub.v1beta2.PushConfig")
	proto.RegisterType((*ReceivedMessage)(nil), "google.pubsub.v1beta2.ReceivedMessage")
	proto.RegisterType((*GetSubscriptionRequest)(nil), "google.pubsub.v1beta2.GetSubscriptionRequest")
	proto.RegisterType((*ListSubscriptionsRequest)(nil), "google.pubsub.v1beta2.ListSubscriptionsRequest")
	proto.RegisterType((*ListSubscriptionsResponse)(nil), "google.pubsub.v1beta2.ListSubscriptionsResponse")
	proto.RegisterType((*DeleteSubscriptionRequest)(nil), "google.pubsub.v1beta2.DeleteSubscriptionRequest")
	proto.RegisterType((*ModifyPushConfigRequest)(nil), "google.pubsub.v1beta2.ModifyPushConfigRequest")
	proto.RegisterType((*PullRequest)(nil), "google.pubsub.v1beta2.PullRequest")
	proto.RegisterType((*PullResponse)(nil), "google.pubsub.v1beta2.PullResponse")
	proto.RegisterType((*ModifyAckDeadlineRequest)(nil), "google.pubsub.v1beta2.ModifyAckDeadlineRequest")
	proto.RegisterType((*AcknowledgeRequest)(nil), "google.pubsub.v1beta2.AcknowledgeRequest")
}

var fileDescriptorPubsub = []byte{
	// 950 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x72, 0xdb, 0x44,
	0x14, 0x1e, 0xc5, 0xf9, 0x3d, 0x72, 0x6a, 0xb2, 0x34, 0xa9, 0x23, 0x0a, 0x13, 0x54, 0xc8, 0xa4,
	0x0c, 0xc8, 0xd8, 0xb4, 0x0c, 0x53, 0xae, 0xd2, 0x36, 0x03, 0x0c, 0x74, 0x30, 0x01, 0x86, 0x0e,
	0x30, 0xf1, 0xc8, 0xd2, 0x89, 0x23, 0x2c, 0x4b, 0x42, 0x5a, 0x99, 0x9a, 0x3b, 0x2e, 0x98, 0xe1,
	0x25, 0x78, 0x05, 0x5e, 0x81, 0x57, 0x63, 0xa5, 0x5d, 0xd9, 0x92, 0xad, 0x95, 0x45, 0x2e, 0xb5,
	0x7b, 0xce, 0x77, 0xfe, 0xbe, 0x6f, 0x75, 0x40, 0x1f, 0xf9, 0xfe, 0xc8, 0xc5, 0x4e, 0x10, 0x0f,
	0xa3, 0x78, 0xd8, 0x99, 0x76, 0x87, 0x48, 0xcd, 0x9e, 0xf8, 0x34, 0x82, 0xd0, 0xa7, 0x3e, 0x39,
	0xe4, 0x36, 0x86, 0x38, 0x14, 0x36, 0xda, 0x1b, 0x99, 0x6b, 0x62, 0x34, 0x8c, 0xaf, 0x3b, 0x38,
	0x09, 0xe8, 0x8c, 0xfb, 0xe8, 0x87, 0xb0, 0xf5, 0x9d, 0x1f, 0x38, 0x16, 0x69, 0xc2, 0xa6, 0x67,
	0x4e, 0xb0, 0xad, 0x9c, 0x28, 0x67, 0x7b, 0xfa, 0x3f, 0x0a, 0xec, 0xf7, 0x53, 0x98, 0x17, 0x18,
	0x45, 0xe6, 0x08, 0x93, 0x7b, 0xdb, 0xa4, 0x66, 0x7a, 0xdf, 0x24, 0x9f, 0x03, 0x98, 0x94, 0x86,
	0xce, 0x30, 0xa6, 0x18, 0xb5, 0x37, 0x4e, 0x1a, 0x67, 0x6a, 0xef, 0x91, 0x51, 0x1a, 0xdf, 0x28,
	0xe0, 0x18, 0xe7, 0x73, 0xb7, 0x0b, 0x8f, 0x86, 0x33, 0x42, 0x00, 0x26, 0xfc, 0x6a, 0xe0, 0xd8,
	0xed, 0x46, 0x12, 0x5d, 0xeb, 0x42, 0x6b, 0xd9, 0x4c, 0x85, 0xc6, 0x18, 0x67, 0x3c, 0x3b, 0xb2,
	0x0f, 0x5b, 0x53, 0xd3, 0x8d, 0x91, 0x05, 0x66, 0x9f, 0x4f, 0x36, 0x3e, 0x51, 0xf4, 0x13, 0x68,
	0x7d, 0x86, 0x34, 0x2d, 0xe5, 0x12, 0x7f, 0x8d, 0x31, 0xa2, 0x89, 0x15, 0x4d, 0xbe, 0x45, 0x49,
	0x3f, 0xc0, 0x1d, 0x96, 0x89, 0xeb, 0x44, 0x37, 0xe5, 0x06, 0xe4, 0x63, 0xd8, 0x15, 0x99, 0x64,
	0x15, 0xbd, 0x53, 0xa7, 0x22, 0xfd, 0x14, 0x5a, 0x73, 0xe0, 0x28, 0xf0, 0xbd, 0x08, 0xc9, 0xeb,
	0xa0, 0x2e, 0x8a, 0x8a, 0x18, 0x7e, 0x83, 0x25, 0xf0, 0x25, 0x1c, 0x7c, 0xe5, 0x44, 0x3c, 0xc7,
	0x28, 0xcb, 0xa1, 0x05, 0x3b, 0x6c, 0x10, 0xbf, 0xa0, 0x45, 0x45, 0x16, 0x07, 0xb0, 0x17, 0x24,
	0x7e, 0x91, 0xf3, 0x3b, 0xaf, 0x6f, 0x2b, 0x69, 0x51, 0x7a, 0x44, 0xfd, 0x31, 0x7a, 0xbc, 0x45,
	0xfa, 0x4f, 0x40, 0xf2, 0x60, 0x22, 0xee, 0xfb, 0xb0, 0x9d, 0x56, 0xc4, 0x43, 0xaa, 0xbd, 0xfb,
	0x92, 0x02, 0xf8, 0xc8, 0xef, 0x41, 0xcb, 0xc3, 0x57, 0x74, 0x90, 0x03, 0x4f, 0x1b, 0xaa, 0x7f,
	0x0f, 0x6f, 0xce, 0xc1, 0xbf, 0x65, 0x9e, 0x56, 0xe8, 0x04, 0xd4, 0x61, 0x11, 0x24, 0x9d, 0xab,
	0x99, 0x73, 0x1f, 0xde, 0x92, 0xc1, 0x8a, 0xfc, 0x0f, 0x61, 0x3f, 0xca, 0x5f, 0xf0, 0xce, 0xc9,
	0x13, 0x7d, 0x00, 0xe4, 0x39, 0xba, 0x48, 0xb1, 0x6a, 0xf0, 0x7f, 0x28, 0xd0, 0xcc, 0x87, 0x2b,
	0x52, 0x7d, 0x61, 0xbd, 0x21, 0x58, 0xa0, 0x06, 0x71, 0x74, 0x33, 0xb0, 0x7c, 0xef, 0xda, 0x19,
	0xb5, 0x37, 0xd9, 0xa1, 0xda, 0x7b, 0x5b, 0x4a, 0x84, 0xe8, 0xe6, 0x59, 0x6a, 0x48, 0xee, 0xc3,
	0x5d, 0xd3, 0x1a, 0x0f, 0x6c, 0x34, 0x6d, 0xd7, 0xf1, 0x58, 0x2f, 0x90, 0x21, 0xb0, 0xd9, 0x6f,
	0x25, 0xed, 0xd0, 0xff, 0x56, 0x00, 0x72, 0xc6, 0xac, 0xce, 0x34, 0x08, 0x7a, 0x76, 0xe0, 0x3b,
	0x5e, 0x36, 0xfb, 0x8b, 0x12, 0x55, 0x75, 0xd7, 0x86, 0x5e, 0x96, 0xd4, 0x6d, 0xe4, 0xf3, 0x12,
	0x5a, 0x97, 0x68, 0xa1, 0x33, 0x45, 0x3b, 0x13, 0xfc, 0x1d, 0xd8, 0x4e, 0x0a, 0x62, 0xa2, 0xe4,
	0x5e, 0x8f, 0x61, 0x47, 0x70, 0x3a, 0xf5, 0xab, 0xab, 0x0e, 0x03, 0x8e, 0x98, 0x30, 0xf3, 0xfd,
	0xcf, 0xc6, 0x74, 0x17, 0x9a, 0xf9, 0x61, 0x8b, 0x69, 0x5d, 0x42, 0x3b, 0x21, 0x49, 0x29, 0xed,
	0x6e, 0x2b, 0x96, 0x00, 0x8e, 0x4b, 0x30, 0x05, 0xe7, 0x9e, 0x94, 0x71, 0x4e, 0xed, 0x3d, 0x90,
	0x54, 0x57, 0x60, 0x92, 0x94, 0x98, 0x5d, 0x38, 0xe6, 0xc4, 0xac, 0x5f, 0xf8, 0x08, 0xee, 0xbd,
	0xf0, 0x6d, 0xe7, 0x7a, 0xb6, 0x98, 0x6c, 0xa5, 0xc3, 0x32, 0x53, 0x37, 0x6a, 0x32, 0x95, 0xa9,
	0x5b, 0xed, 0xc7, 0xae, 0x5b, 0x0d, 0xae, 0x01, 0x09, 0x91, 0xc6, 0xa1, 0x37, 0x70, 0x26, 0x13,
	0xb4, 0x1d, 0x93, 0xa2, 0x3b, 0x4b, 0x63, 0xec, 0x26, 0x1e, 0x13, 0xf3, 0xd5, 0x60, 0xfe, 0x58,
	0x36, 0x52, 0x8a, 0x7f, 0x03, 0x4d, 0x0e, 0x2b, 0xfa, 0x7a, 0x0e, 0x07, 0xa1, 0xa0, 0xd4, 0xc2,
	0x94, 0xf7, 0xf6, 0x54, 0x92, 0xe4, 0x12, 0x05, 0xf5, 0x2b, 0x68, 0xf3, 0x96, 0x9c, 0x5b, 0xe3,
	0xe7, 0x42, 0x58, 0xd5, 0x69, 0x2f, 0x48, 0xcb, 0xd5, 0x2c, 0x53, 0x25, 0x4f, 0xf9, 0x53, 0x20,
	0x0c, 0xd9, 0xf3, 0x7f, 0x73, 0xd1, 0x1e, 0xad, 0x41, 0x66, 0xdc, 0xe3, 0xc8, 0x5c, 0x98, 0x7b,
	0xbd, 0xbf, 0xb6, 0x01, 0xc4, 0x74, 0x87, 0x18, 0x92, 0x2b, 0x20, 0xcf, 0x42, 0x34, 0x8b, 0x13,
	0x27, 0x75, 0x58, 0xa4, 0xd5, 0xa2, 0x1a, 0xa6, 0x3f, 0xb8, 0xc2, 0xd1, 0x07, 0x12, 0xbf, 0x72,
	0xbd, 0xd5, 0x0b, 0x33, 0xe5, 0x3f, 0xa9, 0x82, 0x54, 0x48, 0x47, 0xe2, 0x29, 0x13, 0xaa, 0xf6,
	0x61, 0x7d, 0x07, 0xc1, 0x96, 0xab, 0xec, 0x25, 0x2f, 0x64, 0x23, 0xc3, 0x91, 0x6a, 0x4b, 0x3b,
	0x9a, 0x7b, 0x88, 0x6d, 0xc7, 0xb8, 0x48, 0xb6, 0x1d, 0xf2, 0x33, 0x1c, 0xac, 0x50, 0x49, 0x5a,
	0x97, 0x8c, 0x74, 0x52, 0xf4, 0x3e, 0xa8, 0x39, 0x22, 0x91, 0x87, 0x12, 0xdc, 0x55, 0xb2, 0x49,
	0x11, 0xbf, 0x86, 0xcd, 0x44, 0x4d, 0x44, 0x97, 0xea, 0x79, 0xae, 0x60, 0xe9, 0x60, 0x0b, 0x72,
	0xfc, 0x11, 0x5e, 0x5b, 0x7e, 0x5e, 0x88, 0x51, 0x59, 0xff, 0xca, 0x3b, 0x24, 0x4b, 0xb6, 0xf7,
	0xef, 0x26, 0xec, 0x89, 0x15, 0x88, 0x29, 0xe1, 0x0b, 0x50, 0xb9, 0x12, 0xf8, 0x96, 0x51, 0xb9,
	0x83, 0x68, 0xd5, 0x1b, 0xca, 0x4b, 0xd8, 0x11, 0xb8, 0xe4, 0x5d, 0xf9, 0xdf, 0x26, 0xb7, 0xd3,
	0x69, 0xa7, 0xeb, 0xcc, 0x44, 0x3b, 0xfa, 0xb0, 0x9b, 0xed, 0x8b, 0xe4, 0x54, 0xae, 0xa3, 0xfc,
	0x5e, 0xb1, 0x26, 0x57, 0x13, 0x60, 0xb1, 0x91, 0x91, 0xb3, 0x0a, 0x05, 0x14, 0x36, 0x40, 0xed,
	0x61, 0x0d, 0x4b, 0x91, 0xf4, 0x9f, 0x0a, 0x1c, 0x95, 0x6f, 0x50, 0xe4, 0xd1, 0x3a, 0x94, 0x52,
	0x9d, 0x3e, 0xfe, 0x9f, 0x5e, 0xf3, 0xe6, 0xa9, 0xb9, 0xb5, 0x4b, 0x4a, 0xf7, 0xd5, 0xd5, 0x4c,
	0xc6, 0xa0, 0xa7, 0xef, 0xc1, 0xb1, 0xe5, 0x4f, 0xca, 0x71, 0x9e, 0xaa, 0x7c, 0xa3, 0xe8, 0x27,
	0x2e, 0x7d, 0x65, 0xb8, 0x9d, 0xfa, 0x7e, 0xf4, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x12,
	0xa9, 0x2d, 0x13, 0x0d, 0x00, 0x00,
}
