#!/usr/bin/env bash

go build -gcflags "all=-N -l" -o rock-paper-scissors && ./rock-paper-scissors