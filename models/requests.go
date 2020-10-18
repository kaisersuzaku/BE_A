package models

type BallContainer struct {
	BallContainerSize      int `json:"ball_container_size" valid:"required"`
	CurrentBallInContainer int `json:"current_ball_in_container" valid:"required"`
	ID                     int `json:"id" valid:"required"`
}

type FillBallContainerReq struct {
	BallContainer BallContainer `json:"ball_container" valid:"required"`
}

type OrderReq struct {
	ID  uint `json:"id" valid:"required"`
	Qty int  `json:"qty" valid:"required"`
}
