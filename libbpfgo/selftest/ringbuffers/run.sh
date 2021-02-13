#!/bin/bash

VERSION_LIMIT=5.8
CURRENT_VERSION=$(uname -r | cut -c1-3)
if (( $(echo "$CURRENT_VERSION < $VERSION_LIMIT" |bc -l) )); then
    echo [*] OUTDATED KERNEL VERSION
    exit 1
else
    echo [*] SUFFICIENT KERNEL VERSION
fi

make -f $PWD/Makefile
if [ $? -ne 0 ]; then
    echo MAKE FAILED
    exit 2
else
    echo [*] MAKE RAN SUCCESFULLY
fi

timeout 5 $PWD/self
RETURN_CODE=$?
if [ $RETURN_CODE -eq 124 ]; then
    echo [*] SELFTEST TIMEDOUT
    exit 3
fi

if [ $RETURN_CODE -ne 0 ]; then
    echo [*] ERROR IN SELFTEST
    exit 4
fi

echo [*] SUCCESS
exit 0
