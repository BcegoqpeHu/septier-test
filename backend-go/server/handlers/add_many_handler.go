package handlers

import (
	"encoding/json"
	"github.com/BcegoqpeHu/septier-test/backend-go/server/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func (h *Handler) AddManyHandler(c echo.Context) (err error) {
	request := new(models.AddManyRequest)
	if err = c.Bind(request); err != nil {
		return
	}

	wsMessage := models.WsMessage{
		Event: "notification",
		Data: models.WsData{
			GpsExt: request.GpsExt,
			Points: request.Points,
		},
	}

	message, err := json.Marshal(wsMessage)
	if err != nil {
		log.Printf("Marshaling error: %v", err)
		return
	}
	h.Hub.Broadcast <- message

	return c.JSON(http.StatusOK, makeResponse(request))
}

func makeResponse(amReq *models.AddManyRequest) models.AddManyResponse {
	result := models.AddManyResponse{
		UsrpCfg: models.UsrpCfg{
			RxGain:       amReq.Config.RxGain,
			Network:      amReq.Config.Network,
			Params_2g:    amReq.Config.Params_2g,
			Params_3g:    amReq.Config.Params_3g,
			Params_4g:    amReq.Config.Params_4g,
			Alpha:        amReq.Config.Alpha,
			Slot:         amReq.Config.Slot,
			Celltrack:    amReq.Config.Celltrack,
			Watcher:      amReq.Config.Watcher,
			Antenna:      amReq.Config.Antenna,
			GpsSrc:       amReq.Config.GpsSrc,
			Version:      amReq.Config.Version + 1,
			App:          amReq.Config.App + 1,
			Mode:         0,
			ScanUARFCN:   `10788`,
			ScanPSC:      ` `,
			UlScStart:    8000000,
			UlScEnd:      9000000,
			ScanULSC:     " ",
			ScanTimeouts: " ",
		},
		Command: 0,
	}
	return result
}
