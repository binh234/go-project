package ping

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPing(t *testing.T) {
	api := New()
	rsp := api.Default()
	require.Equal(t, "pong", rsp)
}

func TestHello(t *testing.T) {
	api := New()
	rsp, err := api.Hello(nil, "Foo")
	require.NoError(t, err)
	require.Equal(t, "Hello, Foo", rsp)
}

func TestReqRsp(t *testing.T) {
	api := New()
	req := Request{Name: "Foo"}
	rsp, err := api.ReqRsp(nil, req)
	require.NoError(t, err)
	require.Equal(t, "Hello, Foo", rsp.Response)
}

func TestReqRsp2(t *testing.T) {
	api := New()
	// happy path
	req := Request{Name: "Foo"}
	rsp, err := api.ReqRsp(nil, req)
	require.NoError(t, err)
	require.Equal(t, "Hello, Foo", rsp.Response)

	// error path
	rsp2, err := api.ReqRsp2(nil, nil)
	require.Nil(t, rsp2)
	require.Error(t, err)
}
