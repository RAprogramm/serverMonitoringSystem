#!/bin/sh

# This is the bare minimum to run in development. For full list of flags,
# run ./smsystem -help

go build -v -o smsystem cmd/web/*.go && ./smsystem \
	-dbuser='postgres' \
	-pusherHost='localhost' \
	-pusherKey='abc123' \
	-pusherSecret='123abc' \
	-pusherApp="1"
-pusherPort="4001"
-pusherSecure=false
