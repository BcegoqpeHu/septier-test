package models

type WsMessage struct {
	Event string `json:"event"`
	Data  WsData `json:"data"`
}

type WsData struct {
	GpsExt GpsExt
	Points []Point
}
