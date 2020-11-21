#!/bin/bash

set_iptables() {
  iptables -t mangle -N YUUKO

  iptables -t mangle -A YUUKO -d 0.0.0.0/8 -j RETURN
  iptables -t mangle -A YUUKO -d 10.0.0.0/8 -j RETURN
  iptables -t mangle -A YUUKO -d 127.0.0.0/8 -j RETURN
  iptables -t mangle -A YUUKO -d 169.254.0.0/16 -j RETURN
  iptables -t mangle -A YUUKO -d 172.16.0.0/12 -j RETURN
  iptables -t mangle -A YUUKO -d 192.168.0.0/16 -p tcp -j RETURN
  iptables -t mangle -A YUUKO -d 192.168.0.0/16 -p udp ! --dport 53 -j RETURN
  iptables -t mangle -A YUUKO -d 224.0.0.0/4 -j RETURN
  iptables -t mangle -A YUUKO -d 240.0.0.0/4 -j RETURN

  iptables -t mangle -A YUUKO -p tcp -j TPROXY --on-port 8081 --tproxy-mark 1
  iptables -t mangle -A YUUKO -p udp -j TPROXY --on-port 8081 --tproxy-mark 1

  iptables -t mangle -A PREROUTING -j YUUKO
}

unset_iptables() {
  iptables -t mangle -D PREROUTING -j YUUKO
  iptables -t mangle -F YUUKO
  iptables -t mangle -X YUUKO
}

set_route() {
    ip route add local default dev lo table 100
    ip rule  add fwmark 1             table 100
}

unset_route() {
    ip rule  del   table 100 &>/dev/null
    ip route flush table 100 &>/dev/null
}

status() {
    echo ""
    echo "#####iptables#####"
    iptables -vnL -t mangle
    echo ""
    echo "#####ip route#####"
    ip route list table 100
    echo ""
    echo "#####ip rule#####"
    ip rule list
}

start() {
    set_iptables
    set_route
}

stop() {
    unset_iptables
    unset_route
}

main() {
    if [ $# -eq 0 ]; then
        echo "usage: $0 start|stop|restart ..."
        return 1
    fi

    for funcname in "$@"; do
        if [ "$(type -t $funcname)" != 'function' ]; then
            echo "'$funcname' not a shell function"
            return 1
        fi
    done

    for funcname in "$@"; do
        echo "running '$funcname'..."
        $funcname
        echo "'$funcname'... done"
    done
    return 0
}
main "$@"
