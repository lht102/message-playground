package job_test

import (
	"testing"

	"github.com/lht102/message-playground/jobworker"
	"github.com/lht102/message-playground/jobworker/ent"
	"github.com/lht102/message-playground/jobworker/ent/enttest"
	"github.com/lht102/message-playground/jobworker/job"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type JobTestSuite struct {
	suite.Suite

	entClient  *ent.Client
	jobService *job.Service
}

func TestJob(t *testing.T) {
	logger, err := zap.NewDevelopment()
	require.NoError(t, err)

	entClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	jobService := job.NewService(entClient, &jobworker.NopMessageBus{}, logger)

	s := &JobTestSuite{
		entClient:  entClient.Debug(),
		jobService: jobService,
	}

	suite.Run(t, s)
}

func (s *JobTestSuite) TearDownSuite() {
	err := s.entClient.Close()
	s.Require().NoError(err)
}
