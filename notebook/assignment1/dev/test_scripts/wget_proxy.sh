#!/bin/bash

if [ $# -ne 2 ]; then
    echo "Usage: wget_proxy.sh <port_of_your_http_server> <port_of_your_proxy>"
else
    port=$1
    pproxy=$2
    mkdir tmp
    cd tmp
    wget -e use_proxy=yes http_proxy=localhost:$pproxy http://localhost:$port/ 2> /dev/null
    wget -e use_proxy=yes http_proxy=localhost:$pproxy http://localhost:$port/images/buses.jpg 2> /dev/null
    wget -e use_proxy=yes http_proxy=localhost:$pproxy http://localhost:$port/images/spacer.gif 2> /dev/null
    wget -e use_proxy=yes http_proxy=localhost:$pproxy http://localhost:$port/images/tranhd.gif 2> /dev/null
    wget -e use_proxy=yes http_proxy=localhost:$pproxy http://localhost:$port/images/transithd.gif 2> /dev/null
    diff -q index.html ../../www/index.html
    diff -q buses.jpg ../../www/images/buses.jpg
    diff -q spacer.gif ../../www/images/spacer.gif
    diff -q tranhd.gif ../../www/images/tranhd.gif
    diff -q transithd.gif ../../www/images/transithd.gif
    cd ..
    rm -rf tmp
    echo If no line were printed before this one, your implementation probably passed the tests.
fi
