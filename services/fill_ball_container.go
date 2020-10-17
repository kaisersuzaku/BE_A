package services

import (
	"context"

	"github.com/kaisersuzaku/BE_A/models"
)

const (
	ContainerAlreadyFull string = "ALREADY_FULL"

	// ContainerVerified : number of current ball in container after addition is same with ball container size
	ContainerVerified string = "VERIFIED"

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
}

func (fbcs FillBallContainerService) IsContainerFull(ctx context.Context, req models.FillBallContainerReq) (resp models.FillBallContainerResp) {
	totalBall := req.BallContainer.CurrentBallInContainer + req.ThrownBall.NumberOfBall
	if totalBall == req.BallContainer.BallContainerSize {
		resp.Status = ContainerVerified
		resp.BallContainer = req.BallContainer
		resp.BallContainer.CurrentBallInContainer = totalBall
		return
	}
	return
}
