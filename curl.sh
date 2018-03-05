#!/bin/bash
while true
do
    echo "curl stress"
    curl localhost:30001/stress
    curl localhost:30001/stress
    curl localhost:30001/stress
    sleep "$1"
done
