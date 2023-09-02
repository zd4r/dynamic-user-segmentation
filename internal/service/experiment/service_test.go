package experiment

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/zd4r/dynamic-user-segmentation/internal/errs"
	experimentModel "github.com/zd4r/dynamic-user-segmentation/internal/model/experiment"
	"github.com/zd4r/dynamic-user-segmentation/internal/service/experiment/mocks"
)

type experimentTestSuite struct {
	suite.Suite
	experimentService *Service
	experimentRepo    *mocks.ExperimentRepository
	context           context.Context
}

func (suite *experimentTestSuite) SetupTest() {
	suite.experimentRepo = mocks.NewExperimentRepository(suite.T())
	suite.experimentService = NewService(suite.experimentRepo)
	suite.context = context.TODO()
}

func (suite *experimentTestSuite) TestService_Create() {
	var tests = []struct {
		name     string
		mock     func()
		input    *experimentModel.Experiment
		expected error
	}{
		{
			name: "create success",
			mock: func() {
				suite.experimentRepo.On(
					"Create",
					suite.context,
					mock.Anything,
				).Return(nil).Once()
			},
			input:    &experimentModel.Experiment{},
			expected: nil,
		},
		{
			name: "create fail",
			mock: func() {
				suite.experimentRepo.On(
					"Create",
					suite.context,
					mock.Anything,
				).Return(errors.New("some error")).Once()
			},
			input:    &experimentModel.Experiment{},
			expected: errors.New("some error"),
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			test.mock()

			err := suite.experimentService.Create(suite.context, test.input)
			require.Equal(t, test.expected, err)
		})
	}
}

func (suite *experimentTestSuite) TestService_CreateBatch() {
	var tests = []struct {
		name     string
		mock     func()
		input    []experimentModel.Experiment
		expected error
	}{
		{
			name: "create batch success",
			mock: func() {
				suite.experimentRepo.On(
					"CreateBatch",
					suite.context,
					mock.Anything,
				).Return(nil).Once()
			},
			input:    make([]experimentModel.Experiment, 1),
			expected: nil,
		},
		{
			name:     "create batch empty slice success",
			mock:     func() {},
			input:    make([]experimentModel.Experiment, 0),
			expected: nil,
		},
		{
			name: "create batch fail",
			mock: func() {
				suite.experimentRepo.On(
					"CreateBatch",
					suite.context,
					mock.Anything,
				).Return(errors.New("some error")).Once()
			},
			input:    make([]experimentModel.Experiment, 1),
			expected: errors.New("some error"),
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			test.mock()

			err := suite.experimentService.CreateBatch(suite.context, test.input)
			require.Equal(t, test.expected, err)
		})
	}
}

func (suite *experimentTestSuite) TestService_Delete() {
	tests := []struct {
		name     string
		mock     func()
		input    *experimentModel.Experiment
		expected error
	}{
		{
			name: "delete success",
			mock: func() {
				suite.experimentRepo.On(
					"Delete",
					suite.context,
					mock.Anything,
				).Return(int64(1), nil).Once()
			},
			input:    &experimentModel.Experiment{},
			expected: nil,
		},
		{
			name: "delete fail",
			mock: func() {
				suite.experimentRepo.On(
					"Delete",
					suite.context,
					mock.Anything,
				).Return(int64(1), errors.New("some error")).Once()
			},
			input:    &experimentModel.Experiment{},
			expected: errors.New("some error"),
		},
		{
			name: "delete not found",
			mock: func() {
				suite.experimentRepo.On(
					"Delete",
					suite.context,
					mock.Anything,
				).Return(int64(0), nil).Once()
			},
			input:    &experimentModel.Experiment{},
			expected: errs.ErrExperimentNotFound,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			test.mock()

			err := suite.experimentService.Delete(suite.context, test.input)
			require.Equal(t, test.expected, err)
		})
	}
}

func (suite *experimentTestSuite) TestService_DeleteBatch() {
	tests := []struct {
		name     string
		mock     func()
		input    []experimentModel.Experiment
		expected error
	}{
		{
			name: "delete batch success",
			mock: func() {
				suite.experimentRepo.On(
					"DeleteBatch",
					suite.context,
					mock.Anything,
				).Return(int64(1), nil).Once()
			},
			input:    make([]experimentModel.Experiment, 1),
			expected: nil,
		},
		{
			name:     "delete batch empty slice success",
			mock:     func() {},
			input:    make([]experimentModel.Experiment, 0),
			expected: nil,
		},
		{
			name: "delete batch fail",
			mock: func() {
				suite.experimentRepo.On(
					"DeleteBatch",
					suite.context,
					mock.Anything,
				).Return(int64(1), errors.New("some error")).Once()
			},
			input:    make([]experimentModel.Experiment, 1),
			expected: errors.New("some error"),
		},
		{
			name: "delete batch not found",
			mock: func() {
				suite.experimentRepo.On(
					"DeleteBatch",
					suite.context,
					mock.Anything,
				).Return(int64(0), nil).Once()
			},
			input:    make([]experimentModel.Experiment, 1),
			expected: errs.ErrExperimentNotFound,
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			test.mock()

			err := suite.experimentService.DeleteBatch(suite.context, test.input)
			require.Equal(t, test.expected, err)
		})
	}
}

func TestExperimentTestSuite(t *testing.T) {
	suite.Run(t, new(experimentTestSuite))
}
