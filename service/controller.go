package service

import (
	"context"

	pb "github.com/meateam/permit-service/proto"
)

// Controller is an interface for the business logic of the permit.Service which uses a Store.
type Controller interface {
	CreatePermit(ctx context.Context, reqID string, fileID string, userID string, status pb.Status) (Permit, error)
	GetPermitByFileID(ctx context.Context, fileID string) ([]pb.UserStatus, error)
	UpdatePermitStatus(ctx context.Context, req *pb.UpdatePermitStatusRequest) (*pb.UpdatePermitStatusResponse, error)
	HealthCheck(ctx context.Context) (bool, error)
}
