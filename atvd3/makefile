run_app: build_server
	./bin/server & ./bin/client

build_server:
	cd server && go build -o ../bin/server
	cd client && go build -o ../bin/client