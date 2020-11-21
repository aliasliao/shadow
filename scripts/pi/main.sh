#!/bin/sh

cd /home/pi/projects/v2ray-core
./routes.sh start
./v2ctl server 1>>./v2ctl.log 2>&1 &
sleep 3
curl -X POST -H "Content-Type: application/json" --data-binary "@./startApp.curl.json" http://localhost:3000/startApp
