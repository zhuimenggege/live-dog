#!/bin/sh

exec sh -c /LiveDog/main & nginx -g 'daemon off;'
