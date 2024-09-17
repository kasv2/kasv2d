package server

import (
	"context"
	"github.com/kasv2/kasv2d/cmd/kaspawallet/daemon/pb"
	"github.com/kasv2/kasv2d/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
