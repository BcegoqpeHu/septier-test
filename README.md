# Septier Test

Тестовое задание от Septier

### Описание каталогов
 - frontend - фронтенд, выполнен с использование Angular 7
 - backend-go - бэкенд, реализованный на Golang 1.12 (Echo)
 - backend-java - бэкенд, реализованный на Java 8 (Micronaut)

### Технические требования для запуска
На машине должно быть установлено:
  - Golang 1.11+
  - JDK 8
  - NodeJs 10+
  - NPM 6+
  - Angular CLI 7+

### Если неохота всё ставить, но есть докер
Если имеется докер, то можно запустить уже собранные образы:
```sh
docker run -d -p 9000:8080 docker.io/sotnikov/septier-test-go
docker run -d -p 10000:8080 docker.io/sotnikov/septier-test-java
```
### Сборка приложений
Для сборки можно использовать скрипты build.sh или build.cmd расположенные в каждом из каталогов, описанных выше, а так же в корне проекта. Скрипты расположенные в корне проекта собирают всё (и фронт и go-сервис и java-сервис), помещая фронт в бэкенд в static-ресурсы.

### Запуск приложений
Для запуска Java-сервиса: 
```sh
$ java -jar backend-java/target/septier-test-backend-0.1.jar
```
Для запуска Go-сервиса:
```sh
$ ./backend-go/app
```
или
```sh
$ backend-go\app.exe
```