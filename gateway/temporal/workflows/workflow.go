package workflow

import (
	"voidspaceGateway/bootstrap"
	temporal_constants "voidspaceGateway/temporal/constants"
)

func RegisterWorkflows(t *bootstrap.TemporalService) {
	t.RegisterWorkflow(DeleteUserWorkflow, temporal_constants.DeleteUserWorkflowName)
}
