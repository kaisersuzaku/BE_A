package models

type InvalidPayload struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type BallContainer struct {
	BallContainerSize      int `json:"ball_container_size"`
	CurrentBallInContainer int `json:"current_ball_in_container"`
}

type ThrownBall struct {
	NumberOfBall int `json:"number_of_ball"`
}

type FillBallContainerReq struct {
	BallContainer BallContainer `json:"ball_container"`
	ThrownBall    ThrownBall    `json:"thrown_ball"`
}
