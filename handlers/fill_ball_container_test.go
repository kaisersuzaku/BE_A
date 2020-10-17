package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kaisersuzaku/BE_A/services/mocks"

	"github.com/kaisersuzaku/BE_A/services"

	"github.com/kaisersuzaku/BE_A/handlers"
	"github.com/kaisersuzaku/BE_A/models"
	"github.com/stretchr/testify/assert"
)

func TestBallContainer(t *testing.T) {
	assert := assert.New(t)

	fbcs := mocks.IFillBallContainerService{}
	fillBallContainerHandler := handlers.BuildFBCHandler(&fbcs)
	var tests = []struct {
		testName string
		expected func() (int, string)
		prepare  func() (*http.Request, *httptest.ResponseRecorder)
	}{
		{
			"TestBallContainerSizeFour : Ball Container VERIFIED",
			func() (int, string) {
				ballContainer := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: services.CommonBallContainerSizeFour,
				}
				resp := models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerVerified,
				}
				b, _ := json.Marshal(resp)
				return http.StatusOK, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {

				ballContainerReq := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: 3,
				}
				thrownBall := models.ThrownBall{
					NumberOfBall: 1,
				}
				req := models.FillBallContainerReq{
					BallContainer: ballContainerReq,
					ThrownBall:    thrownBall,
				}

				ballContainerResp := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeFour,
					CurrentBallInContainer: services.CommonBallContainerSizeFour,
				}
				resp := models.FillBallContainerResp{
					BallContainer: ballContainerResp,
					Status:        services.ContainerVerified,
				}

				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))

				r, _ := http.NewRequest(http.MethodPost, "/ball-container-check", br)

				fbcs.On("ValidateRequest", r.Context(), req).Return(models.RespError{})
				fbcs.On("IsContainerFull", r.Context(), req).Return(resp)
				return r, httptest.NewRecorder()
			},
		},
		{
			"TestBallContainerSizeThree : Ball Container VERIFIED",
			func() (int, string) {
				ballContainer := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: services.CommonBallContainerSizeThree,
				}
				resp := models.FillBallContainerResp{
					BallContainer: ballContainer,
					Status:        services.ContainerVerified,
				}
				b, _ := json.Marshal(resp)
				return http.StatusOK, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {

				ballContainerReq := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: 3,
				}
				thrownBall := models.ThrownBall{
					NumberOfBall: 1,
				}
				req := models.FillBallContainerReq{
					BallContainer: ballContainerReq,
					ThrownBall:    thrownBall,
				}

				ballContainerResp := models.BallContainer{
					BallContainerSize:      services.CommonBallContainerSizeThree,
					CurrentBallInContainer: services.CommonBallContainerSizeThree,
				}
				resp := models.FillBallContainerResp{
					BallContainer: ballContainerResp,
					Status:        services.ContainerVerified,
				}

				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))

				r, _ := http.NewRequest(http.MethodPost, "/ball-container-check", br)

				fbcs.On("ValidateRequest", r.Context(), req).Return(models.RespError{})
				fbcs.On("IsContainerFull", r.Context(), req).Return(resp)
				return r, httptest.NewRecorder()
			},
		},
		{
			"Invalid Payload : Missing Field",
			func() (int, string) {
				resp := models.GetInvalidPayloadResp()
				b, _ := json.Marshal(resp)
				return http.StatusBadRequest, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {

				ballContainerReq := models.BallContainer{
					BallContainerSize: services.CommonBallContainerSizeFour,
				}
				thrownBall := models.ThrownBall{
					NumberOfBall: 1,
				}
				req := models.FillBallContainerReq{
					BallContainer: ballContainerReq,
					ThrownBall:    thrownBall,
				}

				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))

				r, _ := http.NewRequest(http.MethodPost, "/ball-container-check", br)
				return r, httptest.NewRecorder()
			},
		},
		{
			"Invalid Payload : Invalid Content-type",
			func() (int, string) {
				resp := models.GetInvalidPayloadResp()
				b, _ := json.Marshal(resp)
				return http.StatusBadRequest, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {
				b := []byte("This is just a plain text")
				br := ioutil.NopCloser(bytes.NewReader(b))
				r, _ := http.NewRequest(http.MethodPost, "/ball-container-check", br)
				return r, httptest.NewRecorder()
			},
		},
		{
			"Unhandled Request : Ball Container Size more than 4",
			func() (int, string) {
				resp := models.GetUnhandledRequest()
				b, _ := json.Marshal(resp)
				return http.StatusBadRequest, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {

				ballContainerReq := models.BallContainer{
					BallContainerSize:      5,
					CurrentBallInContainer: 3,
				}
				thrownBall := models.ThrownBall{
					NumberOfBall: 1,
				}
				req := models.FillBallContainerReq{
					BallContainer: ballContainerReq,
					ThrownBall:    thrownBall,
				}

				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))

				r, _ := http.NewRequest(http.MethodPost, "/ball-container-check", br)

				fbcs.On("ValidateRequest", r.Context(), req).Return(models.GetUnhandledRequest())
				return r, httptest.NewRecorder()
			},
		},
		{
			"Unhandled Request : Ball Container Size less than 3",
			func() (int, string) {
				resp := models.GetUnhandledRequest()
				b, _ := json.Marshal(resp)
				return http.StatusBadRequest, string(b)
			},
			func() (*http.Request, *httptest.ResponseRecorder) {

				ballContainerReq := models.BallContainer{
					BallContainerSize:      2,
					CurrentBallInContainer: 1,
				}
				thrownBall := models.ThrownBall{
					NumberOfBall: 1,
				}
				req := models.FillBallContainerReq{
					BallContainer: ballContainerReq,
					ThrownBall:    thrownBall,
				}

				b, _ := json.Marshal(req)
				br := ioutil.NopCloser(bytes.NewReader(b))

				r, _ := http.NewRequest(http.MethodPost, "/ball-container-check", br)

				fbcs.On("ValidateRequest", r.Context(), req).Return(models.GetUnhandledRequest())
				return r, httptest.NewRecorder()
			},
		},
	}

	for _, test := range tests {
		req, rr := test.prepare()

		handler := http.HandlerFunc(fillBallContainerHandler.CheckBallContainer)
		handler.ServeHTTP(rr, req)

		httpStatus, bodyString := test.expected()
		assert.Equal(httpStatus, rr.Code, fmt.Sprintf("%s : expected http status %d, got %v", test.testName, httpStatus, rr.Code))
		assert.Equal(bodyString, rr.Body.String(), fmt.Sprintf("%s : expected body %s, got %v", test.testName, bodyString, rr.Body.String()))
	}
}
