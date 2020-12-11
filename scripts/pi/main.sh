#!/bin/sh

main() {
    sleep 1
    cd /home/pi/projects/v2ray-core
    sudo ./v2ctl server 1>>./v2ctl.log 2>&1 &
    sleep 3
    curl -X POST -H "Content-Type: application/json" --data-binary "@./startApp.curl.json" http://localhost:3000/startApp
    sleep 30
    curl -X POST -H "Content-Type: application/json" -d '{"index": 9}' http://localhost:3000/updateShadowIndex
}

main 2>&1 | logger -t "v2ray-core"
