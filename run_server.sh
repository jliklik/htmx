#!/bin/bash

PORT=8080

while true; do
	clear

	pid=$(sudo lsof -t -i:$PORT)
	kill -9 $pid

	echo "Running app..."
	go run .

	echo
	echo "Waiting for changes..."
	inotifywait -r -e modify,create,delete,move . --include './*.(go|html)$'

	echo "Change detected. Restarting..."
done
