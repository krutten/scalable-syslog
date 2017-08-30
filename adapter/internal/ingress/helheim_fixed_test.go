// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package ingress_test

import (
	"time"

	v2 "code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
	"code.cloudfoundry.org/scalable-syslog/adapter/internal/egress"
	"code.cloudfoundry.org/scalable-syslog/adapter/internal/ingress"
	v1 "code.cloudfoundry.org/scalable-syslog/internal/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type mockReceiverClient struct {
	RecvCalled chan bool
	RecvOutput struct {
		Ret0 chan *v2.Envelope
		Ret1 chan error
	}
	HeaderCalled chan bool
	HeaderOutput struct {
		Ret0 chan metadata.MD
		Ret1 chan error
	}
	TrailerCalled chan bool
	TrailerOutput struct {
		Ret0 chan metadata.MD
	}
	CloseSendCalled chan bool
	CloseSendOutput struct {
		Ret0 chan error
	}
	ContextCalled chan bool
	ContextOutput struct {
		Ret0 chan context.Context
	}
	SendMsgCalled chan bool
	SendMsgInput  struct {
		M chan interface{}
	}
	SendMsgOutput struct {
		Ret0 chan error
	}
	RecvMsgCalled chan bool
	RecvMsgInput  struct {
		M chan interface{}
	}
	RecvMsgOutput struct {
		Ret0 chan error
	}
}

func newMockReceiverClient() *mockReceiverClient {
	m := &mockReceiverClient{}
	m.RecvCalled = make(chan bool, 100)
	m.RecvOutput.Ret0 = make(chan *v2.Envelope, 100)
	m.RecvOutput.Ret1 = make(chan error, 100)
	m.HeaderCalled = make(chan bool, 100)
	m.HeaderOutput.Ret0 = make(chan metadata.MD, 100)
	m.HeaderOutput.Ret1 = make(chan error, 100)
	m.TrailerCalled = make(chan bool, 100)
	m.TrailerOutput.Ret0 = make(chan metadata.MD, 100)
	m.CloseSendCalled = make(chan bool, 100)
	m.CloseSendOutput.Ret0 = make(chan error, 100)
	m.ContextCalled = make(chan bool, 100)
	m.ContextOutput.Ret0 = make(chan context.Context, 100)
	m.SendMsgCalled = make(chan bool, 100)
	m.SendMsgInput.M = make(chan interface{}, 100)
	m.SendMsgOutput.Ret0 = make(chan error, 100)
	m.RecvMsgCalled = make(chan bool, 100)
	m.RecvMsgInput.M = make(chan interface{}, 100)
	m.RecvMsgOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockReceiverClient) Recv() (*v2.Envelope, error) {
	m.RecvCalled <- true
	return <-m.RecvOutput.Ret0, <-m.RecvOutput.Ret1
}
func (m *mockReceiverClient) Header() (metadata.MD, error) {
	m.HeaderCalled <- true
	return <-m.HeaderOutput.Ret0, <-m.HeaderOutput.Ret1
}
func (m *mockReceiverClient) Trailer() metadata.MD {
	m.TrailerCalled <- true
	return <-m.TrailerOutput.Ret0
}
func (m *mockReceiverClient) CloseSend() error {
	m.CloseSendCalled <- true
	return <-m.CloseSendOutput.Ret0
}
func (m *mockReceiverClient) Context() context.Context {
	m.ContextCalled <- true
	return <-m.ContextOutput.Ret0
}
func (m *mockReceiverClient) SendMsg(m_ interface{}) error {
	m.SendMsgCalled <- true
	m.SendMsgInput.M <- m_
	return <-m.SendMsgOutput.Ret0
}
func (m *mockReceiverClient) RecvMsg(m_ interface{}) error {
	m.RecvMsgCalled <- true
	m.RecvMsgInput.M <- m_
	return <-m.RecvMsgOutput.Ret0
}

type mockClientPool struct {
	NextCalled chan bool
	NextOutput struct {
		Client chan ingress.LogsProviderClient
	}
}

func newMockClientPool() *mockClientPool {
	m := &mockClientPool{}
	m.NextCalled = make(chan bool, 100)
	m.NextOutput.Client = make(chan ingress.LogsProviderClient, 100)
	return m
}
func (m *mockClientPool) Next() (client ingress.LogsProviderClient) {
	m.NextCalled <- true
	return <-m.NextOutput.Client
}

type mockSyslogConnector struct {
	ConnectCalled chan bool
	ConnectInput  struct {
		Ctx     chan context.Context
		Binding chan *v1.Binding
	}
	ConnectOutput struct {
		W   chan egress.Writer
		Err chan error
	}
}

func newMockSyslogConnector() *mockSyslogConnector {
	m := &mockSyslogConnector{}
	m.ConnectCalled = make(chan bool, 100)
	m.ConnectInput.Ctx = make(chan context.Context, 100)
	m.ConnectInput.Binding = make(chan *v1.Binding, 100)
	m.ConnectOutput.W = make(chan egress.Writer, 100)
	m.ConnectOutput.Err = make(chan error, 100)
	return m
}
func (m *mockSyslogConnector) Connect(ctx context.Context, binding *v1.Binding) (w egress.Writer, err error) {
	m.ConnectCalled <- true
	m.ConnectInput.Ctx <- ctx
	m.ConnectInput.Binding <- binding
	return <-m.ConnectOutput.W, <-m.ConnectOutput.Err
}

type mockLogsProviderClient struct {
	ReceiverCalled chan bool
	ReceiverInput  struct {
		Ctx chan context.Context
		In  chan *v2.EgressRequest
	}
	ReceiverOutput struct {
		Ret0 chan v2.Egress_ReceiverClient
		Ret1 chan error
	}
}

func newMockLogsProviderClient() *mockLogsProviderClient {
	m := &mockLogsProviderClient{}
	m.ReceiverCalled = make(chan bool, 100)
	m.ReceiverInput.Ctx = make(chan context.Context, 100)
	m.ReceiverInput.In = make(chan *v2.EgressRequest, 100)
	m.ReceiverOutput.Ret0 = make(chan v2.Egress_ReceiverClient, 100)
	m.ReceiverOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockLogsProviderClient) Receiver(ctx context.Context, in *v2.EgressRequest) (v2.Egress_ReceiverClient, error) {
	m.ReceiverCalled <- true
	m.ReceiverInput.Ctx <- ctx
	m.ReceiverInput.In <- in
	return <-m.ReceiverOutput.Ret0, <-m.ReceiverOutput.Ret1
}

type mockContext struct {
	DeadlineCalled chan bool
	DeadlineOutput struct {
		Deadline chan time.Time
		Ok       chan bool
	}
	DoneCalled chan bool
	DoneOutput struct {
		Ret0 chan (<-chan struct{})
	}
	ErrCalled chan bool
	ErrOutput struct {
		Ret0 chan error
	}
	ValueCalled chan bool
	ValueInput  struct {
		Key chan interface{}
	}
	ValueOutput struct {
		Ret0 chan interface{}
	}
}

func newMockContext() *mockContext {
	m := &mockContext{}
	m.DeadlineCalled = make(chan bool, 100)
	m.DeadlineOutput.Deadline = make(chan time.Time, 100)
	m.DeadlineOutput.Ok = make(chan bool, 100)
	m.DoneCalled = make(chan bool, 100)
	m.DoneOutput.Ret0 = make(chan (<-chan struct{}), 100)
	m.ErrCalled = make(chan bool, 100)
	m.ErrOutput.Ret0 = make(chan error, 100)
	m.ValueCalled = make(chan bool, 100)
	m.ValueInput.Key = make(chan interface{}, 100)
	m.ValueOutput.Ret0 = make(chan interface{}, 100)
	return m
}
func (m *mockContext) Deadline() (deadline time.Time, ok bool) {
	m.DeadlineCalled <- true
	return <-m.DeadlineOutput.Deadline, <-m.DeadlineOutput.Ok
}
func (m *mockContext) Done() <-chan struct{} {
	m.DoneCalled <- true
	return <-m.DoneOutput.Ret0
}
func (m *mockContext) Err() error {
	m.ErrCalled <- true
	return <-m.ErrOutput.Ret0
}
func (m *mockContext) Value(key interface{}) interface{} {
	m.ValueCalled <- true
	m.ValueInput.Key <- key
	return <-m.ValueOutput.Ret0
}

type mockWriter struct {
	WriteCalled chan bool
	WriteInput  struct {
		Arg0 chan *v2.Envelope
	}
	WriteOutput struct {
		Ret0 chan error
	}
}

func newMockWriter() *mockWriter {
	m := &mockWriter{}
	m.WriteCalled = make(chan bool, 100)
	m.WriteInput.Arg0 = make(chan *v2.Envelope, 100)
	m.WriteOutput.Ret0 = make(chan error, 100)
	return m
}
func (m *mockWriter) Write(arg0 *v2.Envelope) error {
	m.WriteCalled <- true
	m.WriteInput.Arg0 <- arg0
	return <-m.WriteOutput.Ret0
}
