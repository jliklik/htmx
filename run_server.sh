#!/bin/bash

PORT=8080

while true; do
	clear

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
