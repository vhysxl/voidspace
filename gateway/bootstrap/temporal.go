package bootstrap

import (
	"crypto/tls"
	"fmt"
	"os"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"
	"logur.dev/logur"
)

type TemporalService struct {
	Service string
	Client  client.Client
	Worker  worker.Worker
	Logger  *zap.Logger
}

func TemporalServiceInit(logger *zap.Logger, port string) (*TemporalService, error) {
	temporalService := "Voidspace Gateway"
	temporalLogger := logur.LoggerToKV(zapadapter.New(logger))
	temporalOptions := client.Options{
		Logger:      temporalLogger,
		HostPort:    os.Getenv("TEMPORAL_HOST"),
		Namespace:   os.Getenv("TEMPORAL_NAMESPACE"),
		Credentials: client.NewAPIKeyStaticCredentials(os.Getenv("TEMPORAL_API_KEY")),
		ConnectionOptions: client.ConnectionOptions{
			TLS: &tls.Config{},
		},
	}

	temporalClient, err := client.Dial(temporalOptions)
	if err != nil {
		logger.Error("Temporal Client failed to connect", zap.Error(err))
		return nil, err
	}

	logger.Info(fmt.Sprintf("Temporal Client Connected on port %s", temporalOptions.HostPort))

	voidspaceWorker := worker.New(temporalClient, temporalService, worker.Options{})

	return &TemporalService{
		Client:  temporalClient,
		Worker:  voidspaceWorker,
		Service: temporalService,
		Logger:  logger,
	}, nil
}

func (t *TemporalService) RegisterWorkflow(w any, name string) {
	opt := workflow.RegisterOptions{
		Name: name,
	}

	t.Worker.RegisterWorkflowWithOptions(w, opt)
}

func (t *TemporalService) RegisterActivity(a any, name string) {
	opt := activity.RegisterOptions{
		Name: name,
	}

	t.Worker.RegisterActivityWithOptions(a, opt)
}

func (t *TemporalService) TemporalStart() error {
	t.Logger.Info("Starting Temporal Worker...")
	if err := t.Worker.Run(worker.InterruptCh()); err != nil {
		t.Logger.Error("Temporal Worker failed", zap.Error(err))
		return err
	}
	return nil
}

func (t *TemporalService) Stop() {
	t.Logger.Info("Stopping Temporal Worker")
	t.Worker.Stop()
	t.Client.Close()
	t.Logger.Info("Temporal Worker stopped")
}
