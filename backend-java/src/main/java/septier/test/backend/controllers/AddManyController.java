package septier.test.backend.controllers;

import io.micronaut.http.HttpResponse;
import io.micronaut.http.annotation.Body;
import io.micronaut.http.annotation.Controller;
import io.micronaut.http.annotation.Post;
import io.micronaut.websocket.WebSocketBroadcaster;
import io.reactivex.Single;
import lombok.AllArgsConstructor;
import septier.test.backend.models.AddManyRequest;
import septier.test.backend.models.AddManyResponse;
import septier.test.backend.models.AddManyResponse.UsrpCfg;
import septier.test.backend.models.WsMessage;
import septier.test.backend.models.WsMessage.WsData;

import javax.inject.Inject;

@Controller
@AllArgsConstructor(onConstructor = @__(@Inject))
public class AddManyController {
    private final WebSocketBroadcaster broadcaster;

    @Post("/addmany")
    public Single<HttpResponse<AddManyResponse>> addMany(@Body AddManyRequest request) {
        return Single.just(request)
                .map(this::sendMessageToAllClients)
                .map(this::makeResponse)
                .map(HttpResponse::ok);
    }

    private AddManyRequest sendMessageToAllClients(AddManyRequest request) {
        WsMessage<WsData> notification = new WsMessage<>("notification", makeWsData(request));
        broadcaster.broadcastSync(notification);
        return request;
    }

    private WsData makeWsData(AddManyRequest request) {
        WsData wsData = new WsData();
        wsData.GpsExt = request.GpsExt;
        wsData.Points = request.Points;
        return wsData;
    }

    private AddManyResponse makeResponse(AddManyRequest request) {
        UsrpCfg usrpCfg = new UsrpCfg(request.Config);
        usrpCfg.Mode = 0;
        usrpCfg.ScanUARFCN = "10788";
        usrpCfg.ScanPSC = " ";
        usrpCfg.UlScStart = 8000000;
        usrpCfg.UlScEnd = 9000000;
        usrpCfg.ScanULSC = " ";
        usrpCfg.ScanTimeouts = " ";

        AddManyResponse response = new AddManyResponse();
        response.Command = 0;
        response.UsrpCfg = usrpCfg;

        return response;
    }
}
