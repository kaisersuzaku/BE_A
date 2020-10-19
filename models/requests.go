package models

type BallContainer struct {
	BallContainerSize      int  `json:"ball_container_size" valid:"range(1|4)"`
	CurrentBallInContainer int  `json:"current_ball_in_container" valid:"range(0|4)"`
	ID                     uint `json:"id" valid:"required"`
}

type FillBallContainerReq struct {
	BallContainer BallContainer `json:"ball_container" valid:"required"`
}

type OrderProductReq struct {
	ID  uint `json:"id" valid:"required"`
	Qty uint `json:"qty" valid:"required"`
}
