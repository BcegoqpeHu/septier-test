call build.cmd

call docker build -t sotnikov/septier-test-go:latest ./backend-go/
call docker build -t sotnikov/septier-test-java:latest ./backend-java/

call docker push sotnikov/septier-test-go:latest
call docker push sotnikov/septier-test-java:latest
