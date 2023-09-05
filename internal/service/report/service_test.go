package report

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	reportModel "github.com/zd4r/dynamic-user-segmentation/internal/model/report"
	"github.com/zd4r/dynamic-user-segmentation/internal/service/report/mocks"
)

type reportTestSuite struct {
	suite.Suite
	reportService *Service
	reportRepo    *mocks.ReportRepository
	context       context.Context
}

func (suite *reportTestSuite) SetupTest() {
	suite.reportRepo = mocks.NewReportRepository(suite.T())
	suite.reportService = NewService(suite.reportRepo)
	suite.context = context.TODO()
}

func (suite *reportTestSuite) TestService_CreateBatchRecord() {
	var tests = []struct {
		name     string
		mock     func()
		input    []reportModel.Record
		expected error
	}{
		{
			name: "create batch record success",
			mock: func() {
				suite.reportRepo.On(
					"CreateBatchRecord",
					suite.context,
					mock.Anything,
				).Return(nil).Once()
			},
			input:    make([]reportModel.Record, 1),
			expected: nil,
		},
		{
			name:     "create batch record empty slice success",
			mock:     func() {},
			input:    make([]reportModel.Record, 0),
			expected: nil,
		},
		{
			name: "create batch record fail",
			mock: func() {
				suite.reportRepo.On(
					"CreateBatchRecord",
					suite.context,
					mock.Anything,
				).Return(errors.New("some error")).Once()
			},
			input:    make([]reportModel.Record, 1),
			expected: errors.New("some error"),
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			test.mock()

			err := suite.reportService.CreateBatchRecord(suite.context, test.input)
			require.Equal(t, test.expected, err)
		})
	}
}

func (suite *reportTestSuite) TestService_GetRecordsInIntervalByUser() {
	var tests = []struct {
		name  string
		mock  func()
		input struct {
			userId int
			from   time.Time
			to     time.Time
		}
		expected struct {
			result []reportModel.Record
			err    error
		}
	}{
		{
			name: "get records in interval by user success",
			mock: func() {
				suite.reportRepo.On(
					"GetRecordsInIntervalByUser",
					suite.context,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil, nil).Once()
			},
			input: struct {
				userId int
				from   time.Time
				to     time.Time
			}{},
			expected: struct {
				result []reportModel.Record
				err    error
			}{nil, nil},
		},
		{
			name: "get records in interval by user fail",
			mock: func() {
				suite.reportRepo.On(
					"GetRecordsInIntervalByUser",
					suite.context,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil, errors.New("some error")).Once()
			},
			input: struct {
				userId int
				from   time.Time
				to     time.Time
			}{},
			expected: struct {
				result []reportModel.Record
				err    error
			}{nil, errors.New("some error")},
		},
	}

	for _, test := range tests {
		suite.T().Run(test.name, func(t *testing.T) {
			test.mock()

			records, err := suite.reportService.GetRecordsInIntervalByUser(
				suite.context,
				test.input.userId,
				test.input.from,
				test.input.to,
			)
			require.Equal(t, test.expected.result, records)
			require.Equal(t, test.expected.err, err)
		})
	}
}

func TestReportTestSuite(t *testing.T) {
	suite.Run(t, new(reportTestSuite))
}
