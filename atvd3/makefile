run_app:
	echo "Select a type of connection please"

udp: run_client_udp

run_client_udp: build_client_udp
	./bin/client UDP 10000 1 &

build_client_udp: run_server_udp
	cd client && go build -o ../bin/client

run_server_udp: build_server_udp
	./bin/server UDP 10000 &

build_server_udp:
	cd server && go build -o ../bin/server

tcp: run_client_tcp

run_client_tcp: build_client_tcp
	./bin/client TCP 10000 1 &

build_client_tcp: run_server_tcp
	cd client && go build -o ../bin/client

run_server_tcp: build_server_tcp
	./bin/server TCP 10000 &

build_server_tcp:
	cd server && go build -o ../bin/server