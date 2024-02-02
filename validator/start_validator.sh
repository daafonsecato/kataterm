#!/bin/bash

# Copy or link files from the mounted volume to /app
cp -r /mnt/validator/* /app/ || ln -s /mnt/validator/* /app/

# Ensure the tmp directory exists and has the right permissions
mkdir -p /app/tmp
chmod -R 777 /app/tmp

cd /app
go mod tidy
/usr/bin/air