#!/bin/bash

rm -f /data/gopen.db && go run cmd/main.go  migration init --dbfile /data/gopen.db