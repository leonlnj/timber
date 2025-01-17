package controller

import (
	"context"
	"fmt"
	"testing"

	mlp "github.com/gojek/mlp/api/client"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/caraml-dev/timber/common/errors"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/appcontext"
	"github.com/caraml-dev/timber/dataset-service/services"
	"github.com/caraml-dev/timber/dataset-service/services/mocks"
)

type ObservationServiceControllerTestSuite struct {
	suite.Suite
	ctrl *ObservationServiceController
}

func (s *ObservationServiceControllerTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up ObservationServiceControllerTestSuite")

	// Create mock MLP service and set up with test responses
	mlpSvc := &mocks.MLPService{}
	projectID := int64(0)
	projectName := "test-project"
	expectedProject := &mlp.Project{Id: 0, Name: projectName}
	failedProjectID := int64(4)
	failedProjectName := "failed-test-project"
	expectedFailedProject := &mlp.Project{Id: 4, Name: failedProjectName}
	observationServiceResponse := ""
	mlpSvc.On("GetProject", projectID).Return(expectedProject, nil)
	mlpSvc.On("GetProject", failedProjectID).Return(expectedFailedProject, nil)
	mlpSvc.On(
		"GetProject", int64(3),
	).Return(nil, errors.Newf(errors.NotFound, "MLP Project info for id %d not found in the cache", int64(3)))

	// Create mock Observation service and set up with test responses
	observationSvc := &mocks.ObservationService{}
	observationSvc.On("CreateService", projectName, mock.Anything).Return(&observationServiceResponse, nil)
	observationSvc.On("UpdateService", projectName, int(projectID), mock.Anything).Return(&observationServiceResponse, nil)
	observationSvc.On("CreateService", failedProjectName, mock.Anything).Return(nil, fmt.Errorf("failed create"))
	observationSvc.On("UpdateService", failedProjectName, int(failedProjectID), mock.Anything).Return(nil, fmt.Errorf("failed update"))

	s.ctrl = &ObservationServiceController{
		appCtx: &appcontext.AppContext{
			Services: services.Services{
				MLPService:         mlpSvc,
				ObservationService: observationSvc,
			},
		},
	}
}

func TestObservationServiceControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ObservationServiceControllerTestSuite))
}

func (s *ObservationServiceControllerTestSuite) TestListObservationServices() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.ListObservationServicesRequest
		resp      *timberv1.ListObservationServicesResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.ListObservationServicesRequest{},
			resp:      &timberv1.ListObservationServicesResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.ListObservationServicesRequest{ProjectId: int64(3)},
			err:       "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.ListObservationServices(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *ObservationServiceControllerTestSuite) TestGetObservationService() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.GetObservationServiceRequest
		resp      *timberv1.GetObservationServiceResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.GetObservationServiceRequest{},
			resp:      &timberv1.GetObservationServiceResponse{},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.GetObservationServiceRequest{ProjectId: int64(3)},
			err:       "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.GetObservationService(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *ObservationServiceControllerTestSuite) TestCreateObservationService() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.CreateObservationServiceRequest
		resp      *timberv1.CreateObservationServiceResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.CreateObservationServiceRequest{},
			resp:      &timberv1.CreateObservationServiceResponse{ObservationService: &timberv1.ObservationServiceResponse{}},
		},
		{
			name:      "failure | observation service creation",
			projectID: 4,
			req:       &timberv1.CreateObservationServiceRequest{ProjectId: int64(4)},
			err:       "failed create",
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.CreateObservationServiceRequest{ProjectId: int64(3)},
			err:       "MLP Project info for id 3 not found in the cache",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.CreateObservationService(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}

func (s *ObservationServiceControllerTestSuite) TestUpdateObservationService() {
	ctx := context.Background()
	tests := []struct {
		name      string
		projectID int64
		req       *timberv1.UpdateObservationServiceRequest
		resp      *timberv1.UpdateObservationServiceResponse
		err       string
	}{
		{
			name:      "success",
			projectID: 0,
			req:       &timberv1.UpdateObservationServiceRequest{},
			resp:      &timberv1.UpdateObservationServiceResponse{ObservationService: &timberv1.ObservationServiceResponse{}},
		},
		{
			name:      "failure | project not found",
			projectID: 3,
			req:       &timberv1.UpdateObservationServiceRequest{Id: int64(3), ProjectId: int64(3)},
			err:       "MLP Project info for id 3 not found in the cache",
		},
		{
			name:      "failure | observation service update",
			projectID: 4,
			req:       &timberv1.UpdateObservationServiceRequest{Id: int64(4), ProjectId: int64(4)},
			err:       "failed update",
		},
	}

	for _, data := range tests {
		resp, err := s.ctrl.UpdateObservationService(ctx, data.req)
		if data.err == "" {
			s.Suite.Assert().NoError(err)
			s.Suite.Assert().Equal(data.resp, resp)
		} else {
			s.Suite.Assert().EqualError(err, data.err)
		}
	}
}
