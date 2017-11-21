// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package main

import (
	"sync"
)

var (
	lockNatsConnectionMockPublish sync.RWMutex
)

// NatsConnectionMock is a mock implementation of NatsConnection.
//
//     func TestSomethingThatUsesNatsConnection(t *testing.T) {
//
//         // make and configure a mocked NatsConnection
//         mockedNatsConnection := &NatsConnectionMock{
//             PublishFunc: func(subj string, data []byte) error {
// 	               panic("TODO: mock out the Publish method")
//             },
//         }
//
//         // TODO: use mockedNatsConnection in code that requires NatsConnection
//         //       and then make assertions.
//
//     }
type NatsConnectionMock struct {
	// PublishFunc mocks the Publish method.
	PublishFunc func(subj string, data []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// Publish holds details about calls to the Publish method.
		Publish []struct {
			// Subj is the subj argument value.
			Subj string
			// Data is the data argument value.
			Data []byte
		}
	}
}

// Publish calls PublishFunc.
func (mock *NatsConnectionMock) Publish(subj string, data []byte) error {
	if mock.PublishFunc == nil {
		panic("moq: NatsConnectionMock.PublishFunc is nil but NatsConnection.Publish was just called")
	}
	callInfo := struct {
		Subj string
		Data []byte
	}{
		Subj: subj,
		Data: data,
	}
	lockNatsConnectionMockPublish.Lock()
	mock.calls.Publish = append(mock.calls.Publish, callInfo)
	lockNatsConnectionMockPublish.Unlock()
	return mock.PublishFunc(subj, data)
}

// PublishCalls gets all the calls that were made to Publish.
// Check the length with:
//     len(mockedNatsConnection.PublishCalls())
func (mock *NatsConnectionMock) PublishCalls() []struct {
	Subj string
	Data []byte
} {
	var calls []struct {
		Subj string
		Data []byte
	}
	lockNatsConnectionMockPublish.RLock()
	calls = mock.calls.Publish
	lockNatsConnectionMockPublish.RUnlock()
	return calls
}