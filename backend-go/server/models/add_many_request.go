package models

type AddManyRequest struct {
	DftrxId   string
	Config    Config
	State     State
	GpsExt    GpsExt
	Compass   Compass
	Timestamp int
	Points    []Point
}

type Config struct {
	RxGain    int
	Network   int
	Params_2g CellNetworkParams `json:"2g_params"`
	Params_3g CellNetworkParams `json:"3g_params"`
	Params_4g CellNetworkParams `json:"4g_params"`
	Alpha     float64
	Slot      string
	Celltrack int
	Watcher   int
	Antenna   int
	GpsSrc    int
	Version   int
	App       int
	GsmMode   int
	Url       string
	Sound     float64
}

type CellNetworkParams struct {
	Arfcn int
	Ulsc  int
	Dlsc  int
}

type State struct {
	DlMode       int
	Tune         string
	SignalFound  int
	Hsdpa        int
	Standard     string
	Usrp         int
	Capabilities int
	Expiration   int
	Compression  int
	Error        int
	Remote       int
}

type GpsExt struct {
	Status int
	Lat    float64
	Lon    float64
	Alt    float64
	TS     int
}

type Compass struct {
	Hdg    float64
	FitErr string
}

type Point struct {
	Id              int
	SettingsVersion int
	Channel         int
	SNR             float64
	Angles          []Angle
	RSSI            float64
	Angle2a         float64
	Antenna2a       int
	Compression     int
	Antenna         int
	TSi             int
	TSf             float64
	Clocks          []int64
}

type Angle struct {
	TSi     int
	TSf     float64
	RSSI_m  float64
	RSSI_s  float64
	SNR_m   float64
	SNR_s   float64
	Master  int
	An      float64
	Phase   float64
	Visible int
	Int_m   float64
	Int_s   float64
	Int     float64
	Ant     int
	RxGain  int
	V       int
}
