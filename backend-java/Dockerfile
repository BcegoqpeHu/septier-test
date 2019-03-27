FROM openjdk:8u171-alpine3.7
RUN apk --no-cache add curl
COPY target/septier-test-backend*.jar septier-test-backend.jar
CMD java ${JAVA_OPTS} -jar septier-test-backend.jar