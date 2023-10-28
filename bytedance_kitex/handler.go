package main

import (
	"context"
	echo "micah/echo/kitex_gen/echo"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *echo.Request) (resp *echo.Response, err error) {
	cli := echo.EchoClient{}
	_, _ = cli.Echo(ctx, nil)
	return
}
