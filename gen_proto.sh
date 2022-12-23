#!/usr/bin/env bash
set -eux
protoc -I=. --go_out=. ./tachiyomi.proto
