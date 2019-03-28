package septier.test.backend.models;

import com.fasterxml.jackson.annotation.JsonProperty;

@SuppressWarnings({"WeakerAccess", "unused"})
public class AddManyRequest {
    public String DftrxId;
    public Config Config;
    public State State;
    public GpsExt GpsExt;
    public Compass Compass;
    public Integer Timestamp;
    public Point[] Points;

    public static class Config {
        public Integer RxGain;
        public Integer Network;
        @JsonProperty("2g_params")
        public CellNetworkParams Params2g;
        @JsonProperty("3g_params")
        public CellNetworkParams Params3g;
        @JsonProperty("4g_params")
        public CellNetworkParams Params4g;
        public Double Alpha;
        public String Slot;
        public Integer Celltrack;
        public Integer Watcher;
        public Integer Antenna;
        public Integer GpsSrc;
        public Integer Version;
        public Integer App;
        public Integer GsmMode;
        public String Url;
        public Double Sound;

        public static class CellNetworkParams {
            public Integer Arfcn;
            public Integer Ulsc;
            public Integer Dlsc;
        }
    }

    public static class State {
        public Integer DlMode;
        public String Tune;
        public Integer SignalFound;
        public Integer Hsdpa;
        public String Standard;
        public Integer Usrp;
        public Integer Capabilities;
        public Integer Expiration;
        public Integer Compression;
        public Integer Error;
        public Integer Remote;
    }

    public static class GpsExt {
        public Integer Status;
        public Double Lat;
        public Double Lon;
        public Double Alt;
        public Integer TS;
    }

    public static class Compass {
        public Double Hdg;
        public String FitErr;
    }

    public static class Point {
        public Integer Id;
        public Integer SettingsVersion;
        public Integer Channel;
        public Double SNR;
        public Angle[] Angles;
        public Double RSSI;
        public Double Angle2a;
        public Integer Antenna2a;
        public Integer Compression;
        public Integer Antenna;
        public Integer TSi;
        public Double TSf;
        public Long[] Clocks;

        public static class Angle {
            public Integer TSi;
            public Double TSf;
            public Double RSSI_m;
            public Double RSSI_s;
            public Double SNR_m;
            public Double SNR_s;
            public Integer Master;
            public Double An;
            public Double Phase;
            public Integer Visible;
            public Double Int_m;
            public Double Int_s;
            public Double Int;
            public Integer Ant;
            public Integer RxGain;
            public Integer V;
        }
    }
}