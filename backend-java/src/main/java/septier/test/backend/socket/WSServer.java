package septier.test.backend.socket;

import io.micronaut.websocket.annotation.OnClose;
import io.micronaut.websocket.annotation.OnMessage;
import io.micronaut.websocket.annotation.OnOpen;
import io.micronaut.websocket.annotation.ServerWebSocket;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@ServerWebSocket
public class WSServer {
    private final Logger log = LoggerFactory.getLogger(WSServer.class);

    @OnOpen
    public void onOpen() {
        log.info("New client");
    }

    @OnClose
    public void onClose() {
        log.info("Client gone");
    }

    @OnMessage
    public void onMessage(String message) {
        log.info("Message from peer '{}' ignored", message);
    }
}