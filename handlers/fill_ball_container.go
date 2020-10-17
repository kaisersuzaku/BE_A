package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(bodyByte, &req)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
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
