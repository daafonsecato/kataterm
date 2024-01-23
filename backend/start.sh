#!/bin/bash
useradd git-katas-user -m
ttyd -p 7681 -i 0.0.0.0 --writable -w /home/git-katas-user bash &
code-server --auth none --bind-addr 0.0.0.0:8080 /home/git-katas-user &
cd /app
air 