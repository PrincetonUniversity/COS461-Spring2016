#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: wget.sh <port_of_your_server>"
else
    port=$1
    mkdir tmp
    cd tmp
    curl -0 -o index.html  http://localhost:$port/ 2> /dev/null
    curl -0 -o buses.jpg  http://localhost:$port/images/buses.jpg 2> /dev/null
    curl -0 -o spacer.gif http://localhost:$port/images/spacer.gif 2> /dev/null
    curl -0 -o tranhd.gif  http://localhost:$port/images/tranhd.gif 2> /dev/null
    curl -0 -o transithd.gif http://localhost:$port/images/transithd.gif 2> /dev/null
    diff -q index.html ../../www/index.html
    diff -q buses.jpg ../../www/images/buses.jpg
    diff -q spacer.gif ../../www/images/spacer.gif
    diff -q tranhd.gif ../../www/images/tranhd.gif
    diff -q transithd.gif ../../www/images/transithd.gif
    cd ..
    rm -rf tmp
    echo If no line were printed before this one, your implementation probably passed the tests.
fi
