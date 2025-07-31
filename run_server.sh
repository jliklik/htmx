#!/bin/bash

PORT=8080

while true; do
	clear

	# Kill by port
	PORT=8080
	PID=$(lsof -ti tcp:$PORT)
	if [[ -n "$PID" ]]; then
  	kill $PID
	fi

	echo "Running app..."
	go build -o htmx_app .
	./htmx_app &
	GO_PID=$!


	echo ""
	echo "Waiting for changes..."
	inotifywait -r -e modify,create,delete,move . --include './*.(go|html)$'

	echo "Change detected. Restarting..."
	kill -9 $GO_PID
done
