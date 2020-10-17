package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/kaisersuzaku/BE_A/models"
	"github.com/kaisersuzaku/BE_A/services"
	"github.com/stretchr/testify/assert"
)

func TestBallContainerSizeFour(t *testing.T) {
	assert := assert.New(t)

	fillBallContainerService := services.FillBallContainerService{}
	var tests = []struct {
		testName string
		input1   func() context.Context
		input2   func() models.FillBallContainerReq
		expected func() models.FillBallContainerResp
	}{
		{
			"TestBallContainerSizeFour : Ball Container FULL",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: 3,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: services.CommonBallContainerSizeFour,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerFull,
				}
			},
		},
		{
			"TestBallContainerSizeFour : Ball Container already FULL",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: 4,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: services.CommonBallContainerSizeFour,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerFull,
				}
			},
		},
		{
			"TestBallContainerSizeFour : Ball Container NOT FULL",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: 2,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: 3,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerNotFull,
				}
			},
		},
	}
	for _, test := range tests {
		resp := fillBallContainerService.IsContainerFull(test.input1(), test.input2())
		assert.Equal(test.expected(), resp, fmt.Sprintf("%s : Object not same, expected %v, got %v", test.testName, test.expected(), resp))
	}
}

func TestBallContainerSizeThree(t *testing.T) {
	assert := assert.New(t)

	fillBallContainerService := services.FillBallContainerService{}
	var tests = []struct {
		testName string
		input1   func() context.Context
		input2   func() models.FillBallContainerReq
		expected func() models.FillBallContainerResp
	}{
		{
			"TestBallContainerSizeThree : Ball Container VERIFIED",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: 2,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: services.CommonBallContainerSizeThree,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerFull,
				}
			},
		},
		{
			"TestBallContainerSizeThree : Ball Container already FULL",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: 3,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: services.CommonBallContainerSizeThree,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerFull,
				}
			},
		},
		{
			"TestBallContainerSizeThree : Ball Container NOT FULL",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: 1,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: 2,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerNotFull,
				}
			},
		},
	}
	for _, test := range tests {
		resp := fillBallContainerService.IsContainerFull(test.input1(), test.input2())
		assert.Equal(test.expected(), resp, fmt.Sprintf("%s : Object not same, expected %v, got %v", test.testName, test.expected(), resp))
	}
}

func TestValidateRequest(t *testing.T) {
	assert := assert.New(t)

	fillBallContainerService := services.FillBallContainerService{}
	var tests = []struct {
		testName string
		input1   func() context.Context
		input2   func() models.FillBallContainerReq
		expected func() models.RespError
	}{
		{
			"TestValidateRequest : Valid Size 4",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: 3,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.RespError {
				return models.RespError{}
			},
		},
		{
			"TestValidateRequest : Valid Size 3",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: 3,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.RespError {
				return models.RespError{}
			},
		},
		{
			"TestValidateRequest : Invalid Size more than 4",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      5,
					CurrentBallInContainer: 3,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.RespError {
				return models.GetUnhandledRequest()
			},
		},
		{
			"TestValidateRequest : Invalid Size less than 3",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					ID:                     1,
					BallContainerSize:      2,
					CurrentBallInContainer: 1,
				}

				return models.FillBallContainerReq{
					BallContainer: ballContainer,
				}
			},
			func() models.RespError {
				return models.GetUnhandledRequest()
			},
		},
	}
	for _, test := range tests {
		resp := fillBallContainerService.ValidateRequest(test.input1(), test.input2())
		assert.Equal(test.expected(), resp, fmt.Sprintf("%s : Object not same, expected %v, got %v", test.testName, test.expected(), resp))
	}
}
