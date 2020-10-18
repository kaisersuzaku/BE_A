package models

type FillBallContainerResp struct {
	Status        string        `json:"status"`
	BallContainer BallContainer `json:"ball_container"`
}

type OrderProductResp struct {
	Status string `json:"status"`
}
