// Code generated by protoc-gen-go.
// source: google/pubsub/v1/pubsub.proto
// DO NOT EDIT!

package google_pubsub_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/api"
import _ "go.pedge.io/pb/go/google/protobuf"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// A topic resource.
type Topic struct {
	// The name of the topic. It must have the format
	// `"projects/{project}/topics/{topic}"`. `{topic}` must start with a letter,
	// and contain only letters (`[A-Za-z]`), numbers (`[0-9]`), dashes (`-`),
	// underscores (`_`), periods (`.`), tildes (`~`), plus (`+`) or percent
	// signs (`%`). It must be between 3 and 255 characters in length, and it
	// must not start with `"goog"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Topic) Reset()                    { *m = Topic{} }
func (m *Topic) String() string            { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()               {}
func (*Topic) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// A message data and its attributes. The message payload must not be empty;
// it must contain either a non-empty data field, or at least one attribute.
type PubsubMessage struct {
	// The message payload. For JSON requests, the value of this field must be
	// base64-encoded.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Optional attributes for this message.
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ID of this message, assigned by the server when the message is published.
	// Guaranteed to be unique within the topic. This value may be read by a
	// subscriber that receives a `PubsubMessage` via a `Pull` call or a push
	// delivery. It must not be populated by the publisher in a `Publish` call.
	MessageId string `protobuf:"bytes,3,opt,name=message_id" json:"message_id,omitempty"`
	// The time at which the message was published, populated by the server when
	// it receives the `Publish` call. It must not be populated by the
	// publisher in a `Publish` call.
	PublishTime *google_protobuf2.Timestamp `protobuf:"bytes,4,opt,name=publish_time" json:"publish_time,omitempty"`
}

func (m *PubsubMessage) Reset()                    { *m = PubsubMessage{} }
func (m *PubsubMessage) String() string            { return proto.CompactTextString(m) }
func (*PubsubMessage) ProtoMessage()               {}
func (*PubsubMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PubsubMessage) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *PubsubMessage) GetPublishTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.PublishTime
	}
	return nil
}

// Request for the GetTopic method.
type GetTopicRequest struct {
	// The name of the topic to get.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *GetTopicRequest) Reset()                    { *m = GetTopicRequest{} }
func (m *GetTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTopicRequest) ProtoMessage()               {}
func (*GetTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Request for the Publish method.
type PublishRequest struct {
	// The messages in the request will be published on this topic.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	// The messages to publish.
	Messages []*PubsubMessage `protobuf:"bytes,2,rep,name=messages" json:"messages,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (m *PublishRequest) String() string            { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PublishRequest) GetMessages() []*PubsubMessage {
	if m != nil {
		return m.Messages
	}
	return nil
}

// Response for the `Publish` method.
type PublishResponse struct {
	// The server-assigned ID of each published message, in the same order as
	// the messages in the request. IDs are guaranteed to be unique within
	// the topic.
	MessageIds []string `protobuf:"bytes,1,rep,name=message_ids" json:"message_ids,omitempty"`
}

func (m *PublishResponse) Reset()                    { *m = PublishResponse{} }
func (m *PublishResponse) String() string            { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()               {}
func (*PublishResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Request for the `ListTopics` method.
type ListTopicsRequest struct {
	// The name of the cloud project that topics belong to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// Maximum number of topics to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size" json:"page_size,omitempty"`
	// The value returned by the last `ListTopicsResponse`; indicates that this is
	// a continuation of a prior `ListTopics` call, and that the system should
	// return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token" json:"page_token,omitempty"`
}

func (m *ListTopicsRequest) Reset()                    { *m = ListTopicsRequest{} }
func (m *ListTopicsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTopicsRequest) ProtoMessage()               {}
func (*ListTopicsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Response for the `ListTopics` method.
type ListTopicsResponse struct {
	// The resulting topics.
	Topics []*Topic `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
	// If not empty, indicates that there may be more topics that match the
	// request; this value should be passed in a new `ListTopicsRequest`.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token" json:"next_page_token,omitempty"`
}

func (m *ListTopicsResponse) Reset()                    { *m = ListTopicsResponse{} }
func (m *ListTopicsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTopicsResponse) ProtoMessage()               {}
func (*ListTopicsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListTopicsResponse) GetTopics() []*Topic {
	if m != nil {
		return m.Topics
	}
	return nil
}

// Request for the `ListTopicSubscriptions` method.
type ListTopicSubscriptionsRequest struct {
	// The name of the topic that subscriptions are attached to.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	// Maximum number of subscription names to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size" json:"page_size,omitempty"`
	// The value returned by the last `ListTopicSubscriptionsResponse`; indicates
	// that this is a continuation of a prior `ListTopicSubscriptions` call, and
	// that the system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token" json:"page_token,omitempty"`
}

func (m *ListTopicSubscriptionsRequest) Reset()                    { *m = ListTopicSubscriptionsRequest{} }
func (m *ListTopicSubscriptionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListTopicSubscriptionsRequest) ProtoMessage()               {}
func (*ListTopicSubscriptionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

// Response for the `ListTopicSubscriptions` method.
type ListTopicSubscriptionsResponse struct {
	// The names of the subscriptions that match the request.
	Subscriptions []string `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// If not empty, indicates that there may be more subscriptions that match
	// the request; this value should be passed in a new
	// `ListTopicSubscriptionsRequest` to get more subscriptions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token" json:"next_page_token,omitempty"`
}

func (m *ListTopicSubscriptionsResponse) Reset()                    { *m = ListTopicSubscriptionsResponse{} }
func (m *ListTopicSubscriptionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListTopicSubscriptionsResponse) ProtoMessage()               {}
func (*ListTopicSubscriptionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

// Request for the `DeleteTopic` method.
type DeleteTopicRequest struct {
	// Name of the topic to delete.
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
}

func (m *DeleteTopicRequest) Reset()                    { *m = DeleteTopicRequest{} }
func (m *DeleteTopicRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTopicRequest) ProtoMessage()               {}
func (*DeleteTopicRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// A subscription resource.
type Subscription struct {
	// The name of the subscription. It must have the format
	// `"projects/{project}/subscriptions/{subscription}"`. `{subscription}` must
	// start with a letter, and contain only letters (`[A-Za-z]`), numbers
	// (`[0-9]`), dashes (`-`), underscores (`_`), periods (`.`), tildes (`~`),
	// plus (`+`) or percent signs (`%`). It must be between 3 and 255 characters
	// in length, and it must not start with `"goog"`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The name of the topic from which this subscription is receiving messages.
	// The value of this field will be `_deleted-topic_` if the topic has been
	// deleted.
	Topic string `protobuf:"bytes,2,opt,name=topic" json:"topic,omitempty"`
	// If push delivery is used with this subscription, this field is
	// used to configure it. An empty `pushConfig` signifies that the subscriber
	// will pull and ack messages using API methods.
	PushConfig *PushConfig `protobuf:"bytes,4,opt,name=push_config" json:"push_config,omitempty"`
	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time (on a best-effort basis).
	//
	// For pull subscriptions, this value is used as the initial value for the ack
	// deadline. To override this value for a given message, call
	// `ModifyAckDeadline` with the corresponding `ack_id` if using
	// pull.
	//
	// For push delivery, this value is also used to set the request timeout for
	// the call to the push endpoint.
	//
	// If the subscriber never acknowledges the message, the Pub/Sub
	// system will eventually redeliver the message.
	//
	// If this parameter is not set, the default value of 10 seconds is used.
	AckDeadlineSeconds int32 `protobuf:"varint,5,opt,name=ack_deadline_seconds" json:"ack_deadline_seconds,omitempty"`
}

func (m *Subscription) Reset()                    { *m = Subscription{} }
func (m *Subscription) String() string            { return proto.CompactTextString(m) }
func (*Subscription) ProtoMessage()               {}
func (*Subscription) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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
	PushEndpoint string `protobuf:"bytes,1,opt,name=push_endpoint" json:"push_endpoint,omitempty"`
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
	// If not present during the `CreateSubscription` call, it will default to
	// the version of the API used to make such call. If not present during a
	// `ModifyPushConfig` call, its value will not be changed. `GetSubscription`
	// calls will always return a valid version, even if the subscription was
	// created without this attribute.
	//
	// The possible values for this attribute are:
	//
	// * `v1beta1`: uses the push format defined in the v1beta1 Pub/Sub API.
	// * `v1` or `v1beta2`: uses the push format defined in the v1 Pub/Sub API.
	Attributes map[string]string `protobuf:"bytes,2,rep,name=attributes" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *PushConfig) Reset()                    { *m = PushConfig{} }
func (m *PushConfig) String() string            { return proto.CompactTextString(m) }
func (*PushConfig) ProtoMessage()               {}
func (*PushConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *PushConfig) GetAttributes() map[string]string {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// A message and its corresponding acknowledgment ID.
type ReceivedMessage struct {
	// This ID can be used to acknowledge the received message.
	AckId string `protobuf:"bytes,1,opt,name=ack_id" json:"ack_id,omitempty"`
	// The message.
	Message *PubsubMessage `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *ReceivedMessage) Reset()                    { *m = ReceivedMessage{} }
func (m *ReceivedMessage) String() string            { return proto.CompactTextString(m) }
func (*ReceivedMessage) ProtoMessage()               {}
func (*ReceivedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReceivedMessage) GetMessage() *PubsubMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

// Request for the GetSubscription method.
type GetSubscriptionRequest struct {
	// The name of the subscription to get.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *GetSubscriptionRequest) Reset()                    { *m = GetSubscriptionRequest{} }
func (m *GetSubscriptionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSubscriptionRequest) ProtoMessage()               {}
func (*GetSubscriptionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

// Request for the `ListSubscriptions` method.
type ListSubscriptionsRequest struct {
	// The name of the cloud project that subscriptions belong to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// Maximum number of subscriptions to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size" json:"page_size,omitempty"`
	// The value returned by the last `ListSubscriptionsResponse`; indicates that
	// this is a continuation of a prior `ListSubscriptions` call, and that the
	// system should return the next page of data.
	PageToken string `protobuf:"bytes,3,opt,name=page_token" json:"page_token,omitempty"`
}

func (m *ListSubscriptionsRequest) Reset()                    { *m = ListSubscriptionsRequest{} }
func (m *ListSubscriptionsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListSubscriptionsRequest) ProtoMessage()               {}
func (*ListSubscriptionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

// Response for the `ListSubscriptions` method.
type ListSubscriptionsResponse struct {
	// The subscriptions that match the request.
	Subscriptions []*Subscription `protobuf:"bytes,1,rep,name=subscriptions" json:"subscriptions,omitempty"`
	// If not empty, indicates that there may be more subscriptions that match
	// the request; this value should be passed in a new
	// `ListSubscriptionsRequest` to get more subscriptions.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token" json:"next_page_token,omitempty"`
}

func (m *ListSubscriptionsResponse) Reset()                    { *m = ListSubscriptionsResponse{} }
func (m *ListSubscriptionsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListSubscriptionsResponse) ProtoMessage()               {}
func (*ListSubscriptionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *ListSubscriptionsResponse) GetSubscriptions() []*Subscription {
	if m != nil {
		return m.Subscriptions
	}
	return nil
}

// Request for the DeleteSubscription method.
type DeleteSubscriptionRequest struct {
	// The subscription to delete.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
}

func (m *DeleteSubscriptionRequest) Reset()                    { *m = DeleteSubscriptionRequest{} }
func (m *DeleteSubscriptionRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSubscriptionRequest) ProtoMessage()               {}
func (*DeleteSubscriptionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

// Request for the ModifyPushConfig method.
type ModifyPushConfigRequest struct {
	// The name of the subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// The push configuration for future deliveries.
	//
	// An empty `pushConfig` indicates that the Pub/Sub system should
	// stop pushing messages from the given subscription and allow
	// messages to be pulled and acknowledged - effectively pausing
	// the subscription if `Pull` is not called.
	PushConfig *PushConfig `protobuf:"bytes,2,opt,name=push_config" json:"push_config,omitempty"`
}

func (m *ModifyPushConfigRequest) Reset()                    { *m = ModifyPushConfigRequest{} }
func (m *ModifyPushConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyPushConfigRequest) ProtoMessage()               {}
func (*ModifyPushConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *ModifyPushConfigRequest) GetPushConfig() *PushConfig {
	if m != nil {
		return m.PushConfig
	}
	return nil
}

// Request for the `Pull` method.
type PullRequest struct {
	// The subscription from which messages should be pulled.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// If this is specified as true the system will respond immediately even if
	// it is not able to return a message in the `Pull` response. Otherwise the
	// system is allowed to wait until at least one message is available rather
	// than returning no messages. The client may cancel the request if it does
	// not wish to wait any longer for the response.
	ReturnImmediately bool `protobuf:"varint,2,opt,name=return_immediately" json:"return_immediately,omitempty"`
	// The maximum number of messages returned for this request. The Pub/Sub
	// system may return fewer than the number specified.
	MaxMessages int32 `protobuf:"varint,3,opt,name=max_messages" json:"max_messages,omitempty"`
}

func (m *PullRequest) Reset()                    { *m = PullRequest{} }
func (m *PullRequest) String() string            { return proto.CompactTextString(m) }
func (*PullRequest) ProtoMessage()               {}
func (*PullRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

// Response for the `Pull` method.
type PullResponse struct {
	// Received Pub/Sub messages. The Pub/Sub system will return zero messages if
	// there are no more available in the backlog. The Pub/Sub system may return
	// fewer than the `maxMessages` requested even if there are more messages
	// available in the backlog.
	ReceivedMessages []*ReceivedMessage `protobuf:"bytes,1,rep,name=received_messages" json:"received_messages,omitempty"`
}

func (m *PullResponse) Reset()                    { *m = PullResponse{} }
func (m *PullResponse) String() string            { return proto.CompactTextString(m) }
func (*PullResponse) ProtoMessage()               {}
func (*PullResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *PullResponse) GetReceivedMessages() []*ReceivedMessage {
	if m != nil {
		return m.ReceivedMessages
	}
	return nil
}

// Request for the ModifyAckDeadline method.
type ModifyAckDeadlineRequest struct {
	// The name of the subscription.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// List of acknowledgment IDs.
	AckIds []string `protobuf:"bytes,4,rep,name=ack_ids" json:"ack_ids,omitempty"`
	// The new ack deadline with respect to the time this request was sent to
	// the Pub/Sub system. Must be >= 0. For example, if the value is 10, the new
	// ack deadline will expire 10 seconds after the `ModifyAckDeadline` call
	// was made. Specifying zero may immediately make the message available for
	// another pull request.
	AckDeadlineSeconds int32 `protobuf:"varint,3,opt,name=ack_deadline_seconds" json:"ack_deadline_seconds,omitempty"`
}

func (m *ModifyAckDeadlineRequest) Reset()                    { *m = ModifyAckDeadlineRequest{} }
func (m *ModifyAckDeadlineRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyAckDeadlineRequest) ProtoMessage()               {}
func (*ModifyAckDeadlineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

// Request for the Acknowledge method.
type AcknowledgeRequest struct {
	// The subscription whose message is being acknowledged.
	Subscription string `protobuf:"bytes,1,opt,name=subscription" json:"subscription,omitempty"`
	// The acknowledgment ID for the messages being acknowledged that was returned
	// by the Pub/Sub system in the `Pull` response. Must not be empty.
	AckIds []string `protobuf:"bytes,2,rep,name=ack_ids" json:"ack_ids,omitempty"`
}

func (m *AcknowledgeRequest) Reset()                    { *m = AcknowledgeRequest{} }
func (m *AcknowledgeRequest) String() string            { return proto.CompactTextString(m) }
func (*AcknowledgeRequest) ProtoMessage()               {}
func (*AcknowledgeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func init() {
	proto.RegisterType((*Topic)(nil), "google.pubsub.v1.Topic")
	proto.RegisterType((*PubsubMessage)(nil), "google.pubsub.v1.PubsubMessage")
	proto.RegisterType((*GetTopicRequest)(nil), "google.pubsub.v1.GetTopicRequest")
	proto.RegisterType((*PublishRequest)(nil), "google.pubsub.v1.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "google.pubsub.v1.PublishResponse")
	proto.RegisterType((*ListTopicsRequest)(nil), "google.pubsub.v1.ListTopicsRequest")
	proto.RegisterType((*ListTopicsResponse)(nil), "google.pubsub.v1.ListTopicsResponse")
	proto.RegisterType((*ListTopicSubscriptionsRequest)(nil), "google.pubsub.v1.ListTopicSubscriptionsRequest")
	proto.RegisterType((*ListTopicSubscriptionsResponse)(nil), "google.pubsub.v1.ListTopicSubscriptionsResponse")
	proto.RegisterType((*DeleteTopicRequest)(nil), "google.pubsub.v1.DeleteTopicRequest")
	proto.RegisterType((*Subscription)(nil), "google.pubsub.v1.Subscription")
	proto.RegisterType((*PushConfig)(nil), "google.pubsub.v1.PushConfig")
	proto.RegisterType((*ReceivedMessage)(nil), "google.pubsub.v1.ReceivedMessage")
	proto.RegisterType((*GetSubscriptionRequest)(nil), "google.pubsub.v1.GetSubscriptionRequest")
	proto.RegisterType((*ListSubscriptionsRequest)(nil), "google.pubsub.v1.ListSubscriptionsRequest")
	proto.RegisterType((*ListSubscriptionsResponse)(nil), "google.pubsub.v1.ListSubscriptionsResponse")
	proto.RegisterType((*DeleteSubscriptionRequest)(nil), "google.pubsub.v1.DeleteSubscriptionRequest")
	proto.RegisterType((*ModifyPushConfigRequest)(nil), "google.pubsub.v1.ModifyPushConfigRequest")
	proto.RegisterType((*PullRequest)(nil), "google.pubsub.v1.PullRequest")
	proto.RegisterType((*PullResponse)(nil), "google.pubsub.v1.PullResponse")
	proto.RegisterType((*ModifyAckDeadlineRequest)(nil), "google.pubsub.v1.ModifyAckDeadlineRequest")
	proto.RegisterType((*AcknowledgeRequest)(nil), "google.pubsub.v1.AcknowledgeRequest")
}

var fileDescriptor0 = []byte{
	// 1192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x57, 0x5b, 0x6f, 0x1b, 0xd5,
	0x13, 0xd7, 0xe6, 0x9e, 0x59, 0xa7, 0x6e, 0xce, 0x3f, 0x4d, 0xdc, 0xfd, 0xe7, 0xba, 0x69, 0x9b,
	0xd4, 0x14, 0x6f, 0x62, 0x54, 0x01, 0x81, 0x16, 0xda, 0xa4, 0x42, 0x40, 0x2b, 0x59, 0x69, 0xe1,
	0xd5, 0x5a, 0xdb, 0x27, 0xee, 0xd6, 0xde, 0x0b, 0xde, 0xe3, 0x50, 0x03, 0x95, 0x50, 0x85, 0x78,
	0xe6, 0xf2, 0x82, 0x04, 0x0f, 0x48, 0x7c, 0x03, 0x3e, 0x01, 0xdf, 0x81, 0x47, 0x5e, 0xf9, 0x20,
	0x9c, 0xdb, 0xda, 0xbb, 0xde, 0xb3, 0xb6, 0x93, 0x37, 0xef, 0x9e, 0x39, 0x33, 0xbf, 0x99, 0xf9,
	0xfd, 0x66, 0xc7, 0xb0, 0xd1, 0xf4, 0xfd, 0x66, 0x1b, 0x5b, 0x41, 0xb7, 0x16, 0x76, 0x6b, 0xd6,
	0xf9, 0xa1, 0xfc, 0x55, 0x0a, 0x3a, 0x3e, 0xf1, 0xd1, 0x55, 0x71, 0x5c, 0x92, 0x2f, 0xcf, 0x0f,
	0x8d, 0x75, 0x79, 0xc1, 0x0e, 0x1c, 0xcb, 0xf6, 0x3c, 0x9f, 0xd8, 0xc4, 0xf1, 0xbd, 0x50, 0xd8,
	0x1b, 0xff, 0x8f, 0xdc, 0xb1, 0xa7, 0x5a, 0xf7, 0xcc, 0xc2, 0x6e, 0x40, 0x7a, 0xf2, 0x70, 0x6b,
	0xf8, 0x90, 0x38, 0x2e, 0x0e, 0x89, 0xed, 0x06, 0xc2, 0xc0, 0xbc, 0x06, 0xb3, 0xcf, 0xfc, 0xc0,
	0xa9, 0xa3, 0x1c, 0xcc, 0x78, 0xb6, 0x8b, 0x0b, 0xda, 0xb6, 0xb6, 0xbf, 0x68, 0xfe, 0xa3, 0xc1,
	0x52, 0x85, 0x03, 0x78, 0x82, 0xc3, 0xd0, 0x6e, 0x62, 0x76, 0xde, 0xb0, 0x89, 0xcd, 0xcf, 0x73,
	0xe8, 0x18, 0xc0, 0x26, 0xa4, 0xe3, 0xd4, 0xba, 0x04, 0x87, 0x85, 0xa9, 0xed, 0xe9, 0x7d, 0xbd,
	0x6c, 0x95, 0x86, 0x91, 0x97, 0x12, 0x2e, 0x4a, 0x0f, 0xfa, 0x37, 0x1e, 0x79, 0xa4, 0xd3, 0x43,
	0x08, 0xc0, 0x15, 0x47, 0x55, 0xa7, 0x51, 0x98, 0x66, 0x81, 0xd1, 0x01, 0xe4, 0xe8, 0xf5, 0xb6,
	0x13, 0x3e, 0xaf, 0x32, 0xa8, 0x85, 0x19, 0xfa, 0x56, 0x2f, 0x1b, 0x7d, 0xd7, 0x32, 0x8f, 0xd2,
	0xb3, 0x28, 0x0f, 0xe3, 0x10, 0xf2, 0xc3, 0x8e, 0x75, 0x98, 0x6e, 0xe1, 0x9e, 0x48, 0x05, 0x2d,
	0xc1, 0xec, 0xb9, 0xdd, 0xee, 0x62, 0x8a, 0x92, 0x3e, 0x1e, 0x4d, 0xbd, 0xa3, 0x99, 0xdb, 0x90,
	0xff, 0x08, 0x13, 0x9e, 0xf7, 0x29, 0xfe, 0xa2, 0x4b, 0x1d, 0x31, 0x2b, 0xc2, 0x9e, 0x65, 0xfe,
	0xa7, 0x70, 0xa5, 0x22, 0x60, 0xa8, 0x0d, 0xd0, 0x21, 0x2c, 0x48, 0xec, 0x51, 0xfa, 0x5b, 0x63,
	0xd2, 0x37, 0x6f, 0x41, 0xbe, 0xef, 0x33, 0x0c, 0x68, 0x03, 0x31, 0xfa, 0x1f, 0xe8, 0x83, 0x0a,
	0x84, 0xd4, 0xf5, 0x34, 0x8d, 0xfd, 0x29, 0x2c, 0x3f, 0x76, 0x42, 0x01, 0x2f, 0x8c, 0xc2, 0xe7,
	0x61, 0x9e, 0xe6, 0xfe, 0x02, 0xd7, 0x89, 0x04, 0xb0, 0x0c, 0x8b, 0x01, 0xbb, 0x17, 0x3a, 0x5f,
	0x89, 0xd4, 0x66, 0x59, 0x3d, 0xf9, 0x2b, 0xe2, 0xb7, 0xb0, 0x27, 0xea, 0x69, 0x7e, 0x0e, 0x28,
	0xee, 0x4c, 0xc6, 0xdd, 0x83, 0x39, 0x9e, 0x8c, 0x08, 0xa9, 0x97, 0xd7, 0xd2, 0xd8, 0x05, 0x2b,
	0xd6, 0x20, 0xef, 0xe1, 0x97, 0xa4, 0x1a, 0xf3, 0xcb, 0xcb, 0x68, 0x7e, 0x06, 0x1b, 0x7d, 0xbf,
	0x4f, 0xe9, 0xa5, 0x7a, 0xc7, 0x09, 0x38, 0x2b, 0x33, 0xea, 0x35, 0x21, 0xdc, 0x0a, 0x6c, 0x66,
	0xb9, 0x95, 0xd0, 0xaf, 0xc1, 0x52, 0x18, 0x3f, 0x10, 0x45, 0xcb, 0x06, 0xba, 0x0b, 0xe8, 0x04,
	0xb7, 0x31, 0xc1, 0xa3, 0xda, 0xfd, 0x0d, 0xe4, 0xe2, 0xd1, 0x92, 0x62, 0x18, 0x18, 0x4f, 0xc9,
	0xd6, 0xeb, 0x41, 0x97, 0xf2, 0xb3, 0xee, 0x7b, 0x67, 0x4e, 0x53, 0x32, 0x74, 0x5d, 0xd5, 0xfd,
	0xf0, 0xf9, 0x31, 0xb7, 0x41, 0xeb, 0xb0, 0x62, 0xd7, 0x5b, 0xd5, 0x06, 0xb6, 0x1b, 0x6d, 0xc7,
	0xa3, 0x55, 0xc0, 0xf4, 0x32, 0x6d, 0xf8, 0x2c, 0x2b, 0x84, 0xf9, 0x8b, 0x06, 0x10, 0x33, 0xa6,
	0x19, 0x72, 0xff, 0xd8, 0x6b, 0x04, 0xbe, 0xe3, 0x45, 0x0d, 0xff, 0x50, 0x21, 0xb9, 0x3b, 0xa3,
	0xa2, 0x0e, 0xeb, 0xed, 0x32, 0x4a, 0x79, 0x0a, 0xf9, 0x53, 0x5c, 0xc7, 0xce, 0x39, 0x6e, 0x44,
	0x83, 0xe0, 0x0a, 0xcc, 0xb1, 0x5c, 0xa8, 0x62, 0x35, 0xa9, 0xd8, 0x79, 0xc9, 0x61, 0x7e, 0x6f,
	0x02, 0x21, 0x94, 0x60, 0x95, 0xca, 0x2f, 0x5e, 0xf0, 0xa8, 0x2d, 0x2b, 0x90, 0x8b, 0x37, 0xb7,
	0x2f, 0xc6, 0x02, 0x23, 0x85, 0x92, 0x66, 0x97, 0xd5, 0x45, 0x0b, 0xae, 0x2b, 0x7c, 0x4a, 0x8e,
	0xdd, 0x55, 0x71, 0x4c, 0x2f, 0x6f, 0xa6, 0x13, 0x4b, 0xb0, 0x26, 0x93, 0x83, 0x87, 0x70, 0x5d,
	0x70, 0x70, 0xf2, 0x9c, 0x6b, 0xb0, 0xf6, 0xc4, 0x6f, 0x38, 0x67, 0xbd, 0x41, 0x3f, 0x47, 0x5e,
	0x18, 0x66, 0xe5, 0xd4, 0x78, 0x56, 0x52, 0x0d, 0xeb, 0x95, 0x6e, 0xbb, 0x3d, 0xda, 0xaf, 0x01,
	0xa8, 0x83, 0x49, 0xb7, 0xe3, 0x55, 0x1d, 0xd7, 0xc5, 0x0d, 0xc7, 0x26, 0xb8, 0xdd, 0xe3, 0xee,
	0x17, 0xd8, 0x0d, 0xd7, 0x7e, 0x59, 0xed, 0x0f, 0xc2, 0x69, 0x4e, 0xe7, 0xc7, 0x90, 0x13, 0x6e,
	0x65, 0x35, 0xdf, 0x87, 0xe5, 0x8e, 0xe4, 0xd0, 0xc0, 0x54, 0x54, 0x74, 0x27, 0x8d, 0x6f, 0x88,
	0x6e, 0x66, 0x15, 0x0a, 0xa2, 0x10, 0x0f, 0xea, 0xad, 0x13, 0xa9, 0x9f, 0xd1, 0x88, 0x29, 0x25,
	0x04, 0x41, 0x43, 0xaa, 0x4d, 0x36, 0x1b, 0xb2, 0xd4, 0x27, 0xe0, 0xbe, 0x07, 0x88, 0xba, 0xf6,
	0xfc, 0x2f, 0xdb, 0xb8, 0xd1, 0x9c, 0xdc, 0x35, 0x13, 0xe0, 0x62, 0xf9, 0xaf, 0x45, 0x00, 0xd9,
	0xd4, 0x1a, 0xee, 0xa0, 0xef, 0x35, 0x40, 0xc7, 0x1d, 0x6c, 0x27, 0x3b, 0x8d, 0xc6, 0x10, 0xc7,
	0x18, 0x73, 0x6e, 0x1e, 0xbc, 0xfe, 0xfb, 0xdf, 0x9f, 0xa7, 0x8a, 0xc6, 0x4d, 0xb6, 0x2c, 0x7c,
	0xcd, 0x26, 0xd3, 0x3d, 0x49, 0xfa, 0xd0, 0x2a, 0x5a, 0x09, 0xaa, 0x5a, 0xc5, 0x57, 0x47, 0x5a,
	0x11, 0xfd, 0xa4, 0xf1, 0x4f, 0x5c, 0x02, 0xc5, 0x7e, 0x3a, 0x8a, 0x5a, 0x86, 0x63, 0xf1, 0xdc,
	0xe5, 0x78, 0x2c, 0xf4, 0x26, 0xc7, 0x13, 0x8f, 0x3f, 0x0a, 0x17, 0xfa, 0x4d, 0x13, 0x5f, 0xb6,
	0x84, 0xe8, 0x50, 0x31, 0x1d, 0x2c, 0x4b, 0xed, 0xc6, 0x1b, 0x13, 0xd9, 0x0a, 0xde, 0x99, 0x25,
	0x8e, 0x72, 0x1f, 0xdd, 0xe2, 0x28, 0x25, 0xb0, 0x18, 0xc0, 0x57, 0x49, 0x84, 0xe8, 0x07, 0x2d,
	0xfa, 0x54, 0x24, 0xca, 0xa6, 0x88, 0x99, 0x29, 0x66, 0x63, 0x35, 0xb5, 0xa8, 0x3c, 0x62, 0xdb,
	0x58, 0x54, 0xb1, 0xe2, 0x05, 0x2b, 0xf6, 0x07, 0xad, 0x58, 0x8a, 0xfd, 0xaa, 0x8a, 0x65, 0x49,
	0x24, 0x13, 0xd0, 0x27, 0x1c, 0xd0, 0x89, 0xf9, 0xc1, 0x85, 0x00, 0x1d, 0xb9, 0xc3, 0x71, 0x18,
	0xd9, 0x7e, 0xd4, 0x40, 0x8f, 0x49, 0x08, 0xdd, 0x48, 0xe3, 0x4b, 0x2b, 0x2c, 0x13, 0xd9, 0x09,
	0x47, 0x76, 0xdf, 0x7c, 0xf7, 0x62, 0xc8, 0xec, 0x41, 0x04, 0x86, 0xe9, 0x3b, 0x0d, 0x66, 0xd8,
	0x14, 0x42, 0x1b, 0xaa, 0x11, 0xd8, 0x1f, 0x7a, 0x2a, 0xaa, 0xc7, 0x87, 0x97, 0x79, 0x8f, 0xa3,
	0x79, 0xdb, 0x2c, 0x5f, 0x0c, 0x4d, 0x40, 0x7d, 0x30, 0x18, 0xbf, 0x6b, 0x70, 0x75, 0x78, 0x8e,
	0xa3, 0xdb, 0x59, 0xfd, 0x4b, 0xcd, 0xfa, 0xcc, 0x22, 0x7d, 0xcc, 0x61, 0x1d, 0x9b, 0xf7, 0x2f,
	0xd3, 0xbe, 0x41, 0x18, 0x0a, 0xb1, 0xfc, 0xeb, 0x1c, 0x2c, 0xca, 0xbd, 0x94, 0x4e, 0xb0, 0x17,
	0xa0, 0x8b, 0x01, 0x26, 0xf7, 0xbf, 0x8c, 0xc5, 0xd0, 0xc8, 0x3a, 0x30, 0x6f, 0x73, 0x64, 0xbb,
	0xc6, 0xa6, 0x72, 0x56, 0x89, 0xad, 0x53, 0x0e, 0xa9, 0xd7, 0x1a, 0xcc, 0xcb, 0xc8, 0x68, 0x5b,
	0xb9, 0x34, 0xc4, 0x16, 0x70, 0x63, 0x67, 0x84, 0x85, 0x6c, 0x56, 0x99, 0xc7, 0xbe, 0x63, 0xee,
	0xf1, 0xd8, 0x3c, 0x96, 0x3a, 0xb8, 0xfc, 0x8b, 0xc1, 0x40, 0xf8, 0xb0, 0x10, 0xfd, 0x17, 0x40,
	0x3b, 0xca, 0x09, 0x19, 0x5f, 0x1c, 0xb3, 0xf3, 0xde, 0xe3, 0xb1, 0x77, 0xd0, 0xd6, 0x98, 0xd8,
	0xe8, 0x5b, 0xba, 0xed, 0x0d, 0x56, 0x72, 0xb4, 0xab, 0x1e, 0x69, 0x89, 0xed, 0xdf, 0xb8, 0x31,
	0xda, 0x48, 0xa6, 0x9f, 0x84, 0xa0, 0x1a, 0x78, 0x02, 0x05, 0xfa, 0x53, 0x83, 0x55, 0xf5, 0x9a,
	0x8d, 0xac, 0x11, 0x91, 0x94, 0x23, 0xf9, 0x60, 0xf2, 0x0b, 0x12, 0x66, 0xf2, 0xeb, 0x91, 0x5d,
	0xa9, 0xa1, 0xf1, 0x4c, 0x40, 0x8f, 0x2d, 0xf2, 0xaa, 0x21, 0x93, 0xde, 0xf3, 0x33, 0xf5, 0x23,
	0x4b, 0x55, 0x1c, 0xd7, 0xad, 0x87, 0x37, 0x61, 0xa5, 0xee, 0xbb, 0xa9, 0x58, 0x0f, 0x75, 0xb1,
	0xd2, 0x56, 0x98, 0xdb, 0x8a, 0x56, 0x9b, 0xe3, 0xfe, 0xdf, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff,
	0xa6, 0xfd, 0x1e, 0x19, 0xdc, 0x0f, 0x00, 0x00,
}