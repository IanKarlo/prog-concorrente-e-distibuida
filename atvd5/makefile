run_app:
	echo "Select a type of connection please"

app: run_client

run_client: build_client
	./bin/client 10000 2 &

build_client: run_server
	cd client && go build -o ../bin/client

run_server: build_server
	./bin/server &

build_server:
	cd server && go build -o ../bin/server