#!/bin/bash
nohup ./build/socket-server > ./log/socket-output.log 2>&1&
nohup ./build/web-server > ./log/web-output.log 2>&1&