package septier.test.backend.models;

import septier.test.backend.models.AddManyRequest.GpsExt;
import septier.test.backend.models.AddManyRequest.Point;

public class WsMessage<T> {
    public final String event;
    public final T data;

    public WsMessage(String event, T data) {
        this.event = event;
        this.data = data;
    }

    public static class WsData {
        public GpsExt GpsExt;
        public Point[] Points;
    }
}
