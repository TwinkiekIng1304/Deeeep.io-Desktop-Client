#!/bin/bash

# Build
if [ -a "dist" ]; then
    rm -rf "dist"
fi
go build -o dist/Deeeep.io-Desktop-Client
cp -r plugins dist/plugins

# Zip
tar -zcf Deeeep.io-Desktop-Client-Linux.tar.gz dist