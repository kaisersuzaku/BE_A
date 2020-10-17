package services

import (
	"context"

	"github.com/kaisersuzaku/BE_A/models"
)

const (
	// ContainerAlreadyFull : number of current ball in container is already same with ball container size
	// Will return also verified
	ContainerAlreadyFull string = "FULL"

	// ContainerFull : number of current ball in container after addition is same with ball container size
	ContainerFull string = "FULL"

	ContainerNotFull string = "NOT_FULL"
)

const (
	CommonBallContainerSizeFour  int = 4
	CommonBallContainerSizeThree int = 3
)

type FillBallContainerService struct {
}

type IFillBallContainerService interface {
	IsContainerFull(ctx context.Context, req models.FillBallContainerReq) models.FillBallContainerResp
	ValidateRequest(ctx context.Context, req models.FillBallContainerReq) models.RespError
}

func (fbcs FillBallContainerService) IsContainerFull(ctx context.Context, req models.FillBallContainerReq) (resp models.FillBallContainerResp) {
	if req.BallContainer.CurrentBallInContainer == req.BallContainer.BallContainerSize {
		resp.Status = ContainerAlreadyFull
		resp.BallContainer = req.BallContainer
		return
	}
	totalBall := req.BallContainer.CurrentBallInContainer + 1
	if totalBall == req.BallContainer.BallContainerSize {
		resp.Status = ContainerFull
		resp.BallContainer = req.BallContainer
		resp.BallContainer.CurrentBallInContainer = totalBall
		return
	}
	if totalBall < req.BallContainer.BallContainerSize {
		resp.Status = ContainerNotFull
		resp.BallContainer = req.BallContainer
		resp.BallContainer.CurrentBallInContainer = totalBall
	}
	return
}

func (fbcs FillBallContainerService) ValidateRequest(ctx context.Context, req models.FillBallContainerReq) (respError models.RespError) {
	if req.BallContainer.BallContainerSize > CommonBallContainerSizeFour || req.BallContainer.BallContainerSize < CommonBallContainerSizeThree {
		respError = models.GetUnhandledRequest()
		return
	}
	return
}
