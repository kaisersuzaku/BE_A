package models

type BallContainer struct {
	BallContainerSize      int `json:"ball_container_size" valid:"required"`
	CurrentBallInContainer int `json:"current_ball_in_container" valid:"required"`
}

type ThrownBall struct {
	NumberOfBall int `json:"number_of_ball" valid:"required"`
}

type FillBallContainerReq struct {
	BallContainer BallContainer `json:"ball_container" valid:"required"`
	ThrownBall    ThrownBall    `json:"thrown_ball" valid:"required"`
}
