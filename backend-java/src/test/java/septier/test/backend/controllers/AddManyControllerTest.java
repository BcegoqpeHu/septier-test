package septier.test.backend.controllers;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import io.micronaut.context.ApplicationContext;
import io.micronaut.core.io.ResourceResolver;
import io.micronaut.core.io.scan.ClassPathResourceLoader;
import io.micronaut.http.HttpRequest;
import io.micronaut.http.client.HttpClient;
import io.micronaut.runtime.server.EmbeddedServer;
import org.junit.AfterClass;
import org.junit.BeforeClass;
import org.junit.Test;

import java.io.BufferedInputStream;
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.stream.Collectors;

import static junit.framework.TestCase.assertEquals;
import static org.junit.Assert.assertNotNull;

public class AddManyControllerTest {
    private static EmbeddedServer server;
    private static HttpClient client;
    private static ObjectMapper mapper;

    @BeforeClass
    public static void setupServer() {
        server = ApplicationContext.run(EmbeddedServer.class);
        client = server
                .getApplicationContext()
                .createBean(HttpClient.class, server.getURL());
        mapper = server
                .getApplicationContext()
                .createBean(ObjectMapper.class);
    }

    @AfterClass
    public static void stopServer() {
        if (server != null) {
            server.stop();
        }
        if (client != null) {
            client.stop();
        }
    }

    @Test
    public void addMany() throws IOException {
        String body = client.toBlocking().retrieve(HttpRequest.POST("/addmany", loadJsonReq()));
        assertNotNull(body);

        JsonNode expected = mapper.readTree(loadJsonRes());
        JsonNode actual = mapper.readTree(body);

        assertEquals(expected, actual);
    }

    private String loadJsonReq() {
        return loadJson("classpath:add_many_request.json");
    }

    private String loadJsonRes() {
        return loadJson("classpath:add_many_response.json");
    }

    private String loadJson(String path) {
        return new ResourceResolver()
                .getLoader(ClassPathResourceLoader.class)
                .map(loader -> loader.getResource(path)
                        .map(url -> {
                            try {
                                BufferedInputStream content = (BufferedInputStream) url.getContent();
                                BufferedReader reader = new BufferedReader(new InputStreamReader(content));
                                return reader.lines().collect(Collectors.joining("\n"));
                            } catch (IOException e) {
                                return null;
                            }
                        })
                        .orElse(null))
                .orElse(null);
    }
}
