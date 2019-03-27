package models

type AddManyResponse struct {
	UsrpCfg UsrpCfg
	Command int
}

type UsrpCfg struct {
	RxGain       int
	Network      int
	Params_2g    CellNetworkParams `json:"2g_params"`
	Params_3g    CellNetworkParams `json:"3g_params"`
	Params_4g    CellNetworkParams `json:"4g_params"`
	Alpha        float64
	Slot         string
	Celltrack    int
	Watcher      int
	Antenna      int
	GpsSrc       int
	Version      int
	App          int
	Mode         int
	ScanUARFCN   string
	ScanPSC      string
	UlScStart    int
	UlScEnd      int
	ScanULSC     string
	ScanTimeouts string
}
