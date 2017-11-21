package main

// NatsConnection is an interface defining nats connection functionality
type NatsConnection interface {
	Publish(subj string, data []byte) error
}
