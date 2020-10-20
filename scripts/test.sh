#!/bin/sh

cd /tmp/mnt/liao/addons/shadow
echo "pwd: $(pwd)"

mkdir -p /tmp/shadow_app
cp v2ray v2ctl /tmp/shadow_app
cd /tmp/shadow_app
echo "pwd: $(pwd)"

# Starting v2ctl deamon...
echo "Starting v2ctl deamon..."
./v2ctl server 1>>./v2ctl.log 2>&1 &

