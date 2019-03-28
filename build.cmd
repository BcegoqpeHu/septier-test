set FRONT_DIST=frontend\dist\septier-test-front
set GO_STATIC_PATH=backend-go\public
set JAVA_STATIC_PATH=backend-java\src\main\resources\public

cd frontend
rmdir /S /Q dist 2>nul
call build.cmd
cd ..

rmdir /S /Q %JAVA_STATIC_PATH% 2>nul
xcopy %FRONT_DIST% %JAVA_STATIC_PATH% /e/i/h
cd backend-java
call build.cmd
cd ..

rmdir /q/s %GO_STATIC_PATH% 2>nul
xcopy %FRONT_DIST% %GO_STATIC_PATH% /e/i/h
cd backend-go
call build.cmd
cd ..