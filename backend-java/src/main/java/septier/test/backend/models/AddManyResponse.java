package septier.test.backend.models;

import com.fasterxml.jackson.annotation.JsonProperty;
import septier.test.backend.models.AddManyRequest.Config;
import septier.test.backend.models.AddManyRequest.Config.CellNetworkParams;

public class AddManyResponse {
    public UsrpCfg UsrpCfg;
    public Integer Command;

    public static class UsrpCfg {
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
        public Integer Mode;
        public String ScanUARFCN;
        public String ScanPSC;
        public Integer UlScStart;
        public Integer UlScEnd;
        public String ScanULSC;
        public String ScanTimeouts;

        public UsrpCfg(Config config) {
            this.RxGain = config.RxGain;
            this.Network = config.Network;
            this.Params2g = config.Params2g;
            this.Params3g = config.Params3g;
            this.Params4g = config.Params4g;
            this.Alpha = config.Alpha;
            this.Slot = config.Slot;
            this.Celltrack = config.Celltrack;
            this.Watcher = config.Watcher;
            this.Antenna = config.Antenna;
            this.GpsSrc = config.GpsSrc;
            this.Version = config.Version;
            this.App = config.App;
        }
    }
}
