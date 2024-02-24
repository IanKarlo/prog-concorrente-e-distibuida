run_app:
	echo "Select a type of connection please"

http: run_client_http

run_client_http: build_client_http
	./bin/client HTTP 10000 2 false &

build_client_http: run_server_http
	cd client && go build -o ../bin/client

run_server_http: build_server_http
	./bin/server HTTP &

build_server_http:
	cd server && go build -o ../bin/server

tcp: run_client_tcp

run_client_tcp: build_client_tcp
	./bin/client TCP 10000 2 false &

build_client_tcp: run_server_tcp
	cd client && go build -o ../bin/client

run_server_tcp: build_server_tcp
	./bin/server TCP &

build_server_tcp:
	cd server && go build -o ../bin/server
