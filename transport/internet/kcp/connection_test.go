package kcp_test

import (
	"net"
	"testing"
	"time"
	"v2ray.com/core/testing/assert"
	"v2ray.com/core/transport/internet"
	"v2ray.com/core/transport/internet/internal"
	. "v2ray.com/core/transport/internet/kcp"
)

type NoOpConn struct{}

func (o *NoOpConn) Write(b []byte) (int, error) {
	return len(b), nil
}

func (o *NoOpConn) Close() error {
	return nil
}

func (o *NoOpConn) Read([]byte) (int, error) {
	panic("Should not be called.")
}

func (o *NoOpConn) LocalAddr() net.Addr {
	return nil
}

func (o *NoOpConn) RemoteAddr() net.Addr {
	return nil
}

func (o *NoOpConn) SetDeadline(time.Time) error {
	return nil
}

func (o *NoOpConn) SetReadDeadline(time.Time) error {
	return nil
}

func (o *NoOpConn) SetWriteDeadline(time.Time) error {
	return nil
}

func (o *NoOpConn) Id() internal.ConnectionId {
	return internal.ConnectionId{}
}

func (o *NoOpConn) Reset(auth internet.Authenticator, input func([]byte)) {}

type NoOpRecycler struct{}

func (o *NoOpRecycler) Put(internal.ConnectionId, net.Conn) {}

func TestConnectionReadTimeout(t *testing.T) {
	assert := assert.On(t)

	conn := NewConnection(1, &NoOpConn{}, &NoOpRecycler{}, NewSimpleAuthenticator(), &Config{})
	conn.SetReadDeadline(time.Now().Add(time.Second))

	b := make([]byte, 1024)
	nBytes, err := conn.Read(b)
	assert.Int(nBytes).Equals(0)
	assert.Error(err).IsNotNil()

	conn.Terminate()
}
