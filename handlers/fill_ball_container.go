package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/kaisersuzaku/BE_A/models"
	"github.com/kaisersuzaku/BE_A/services"
)

type FillBallContainerHandler struct {
	fbcs services.IFillBallContainerService
}

func BuildFBCHandler(fbcs services.IFillBallContainerService) FillBallContainerHandler {
	return FillBallContainerHandler{
		fbcs: fbcs,
	}
}

type IFillBallContainerHandler interface {
	CheckBallContainer(w http.ResponseWriter, r *http.Request)
}

func (fbch FillBallContainerHandler) CheckBallContainer(w http.ResponseWriter, r *http.Request) {
	var req models.FillBallContainerReq
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := models.GetInvalidPayloadResp()
		log.Println(respErr.ErrorCode, respErr.ErrorMessage, err)
		fbch.respError(w, respErr)
		return
	}
	err = json.Unmarshal(bodyByte, &req)
	if err != nil {
		respErr := models.GetInvalidPayloadResp()
		log.Println(respErr.ErrorCode, respErr.ErrorMessage, err)
		fbch.respError(w, respErr)
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		respErr := models.GetInvalidPayloadResp()
		log.Println(respErr.ErrorCode, respErr.ErrorMessage, err)
		fbch.respError(w, respErr)
		return
	}
	respErr := fbch.fbcs.ValidateRequest(r.Context(), req)
	if respErr.ErrorCode != "" {
		log.Println(respErr.ErrorCode, respErr.ErrorMessage, err)
		fbch.respError(w, respErr)
		return
	}
	resp := fbch.fbcs.IsContainerFull(r.Context(), req)
	fbch.respSuccess(w, resp)
}

func (fbch FillBallContainerHandler) respSuccess(w http.ResponseWriter, data interface{}) {
	body, _ := json.Marshal(data)
	m := make(map[string]string)
	w.Header().Add("Content-type", "application/json")
	m["Content-type"] = "application/json"
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (fbch FillBallContainerHandler) respError(w http.ResponseWriter, e models.RespError) {
	body, _ := json.Marshal(e)
	m := make(map[string]string)
	w.Header().Add("Content-type", "application/json")
	m["Content-type"] = "application/json"
	w.WriteHeader(http.StatusBadRequest)
	w.Write(body)
}
