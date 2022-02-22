package srv

import (
	apiv1 "github.com/easyCZ/qfy/gen/v1"
)

type RunService struct {
}

func (s *RunService) Create(context.Context, r *apiv1.CreateRequest) (*apiv1.CreateResponse, error) {
	return nil
}
func (s *RunService) List(context.Context, r *apiv1.ListRequest) (*apiv1.ListResponse, error) {
	return nil
}
