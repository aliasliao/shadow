#!/bin/sh

if [ "$1" = "set" ]; then
  echo "setting routes..."
  iptables -t mangle -N V2RAY
  # hijack dns queries
  # Ignore LANs and any other addresses you'd like to bypass the proxy
  # See Wikipedia and RFC5735 for full list of reserved networks.
  iptables -t mangle -A V2RAY -d 0.0.0.0/8 -j RETURN
  iptables -t mangle -A V2RAY -d 10.0.0.0/8 -j RETURN
  iptables -t mangle -A V2RAY -d 127.0.0.0/8 -j RETURN
  iptables -t mangle -A V2RAY -d 169.254.0.0/16 -j RETURN
  iptables -t mangle -A V2RAY -d 172.16.0.0/12 -j RETURN
  iptables -t mangle -A V2RAY -d 192.168.0.0/16 -p udp ! --dport 53 -j RETURN
  iptables -t mangle -A V2RAY -d 192.168.0.0/16 -j RETURN
  iptables -t mangle -A V2RAY -d 224.0.0.0/4 -j RETURN
  iptables -t mangle -A V2RAY -d 240.0.0.0/4 -j RETURN
  # redirect all other packets to v2ray
  ## iptables -t mangle -A V2RAY -p tcp -j REDIRECT --to-ports 1080
  iptables -t mangle -A V2RAY -p udp -j TPROXY --on-port 1080 --tproxy-mark 0x1/0x1
  iptables -t mangle -A V2RAY -p tcp -j TPROXY --on-port 1080 --tproxy-mark 0x1/0x1

  # only handle packets from lan
  iptables -t mangle -A PREROUTING -j V2RAY
  echo "done"

  exit 0
fi

if [ "$1" = "restore" ]; then
  echo "restoring routes..."
  ## iptables -t mangle -D PREROUTING -i br0 -j V2RAY
  iptables -t mangle -D PREROUTING -j V2RAY
  iptables -t mangle -F V2RAY
  iptables -t mangle -X V2RAY
  echo "done"

  exit 0
fi

echo "usage: iptables set/restore"
exit 1
