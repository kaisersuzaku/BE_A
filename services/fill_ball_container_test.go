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
			"TestBallContainerSizeFour : Ball Container VERIFIED",
			func() context.Context {
				return context.TODO()
			},
			func() models.FillBallContainerReq {
				ballContainer := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: 3,
				}
				thrownBall := models.ThrownBall{
					NumberOfBall: 1,
				}
				return models.FillBallContainerReq{
					BallContainer: ballContainer,
					ThrownBall:    thrownBall,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: services.CommonBallContainerSizeFour,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerVerified,
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
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: 2,
				}
				thrownBall := models.ThrownBall{
					NumberOfBall: 1,
				}
				return models.FillBallContainerReq{
					BallContainer: ballContainer,
					ThrownBall:    thrownBall,
				}
			},
			func() models.FillBallContainerResp {
				ballContainer := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: services.CommonBallContainerSizeThree,
				}
				return models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerVerified,
				}
			},
		},
	}
	for _, test := range tests {
		resp := fillBallContainerService.IsContainerFull(test.input1(), test.input2())
		assert.Equal(test.expected(), resp, fmt.Sprintf("%s : Object not same, expected %v, got %v", test.testName, test.expected(), resp))
	}
}
