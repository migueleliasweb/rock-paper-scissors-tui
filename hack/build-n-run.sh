#!/usr/bin/env bash

reset

# We need extra flags to ensure all debugging information is kept on the final binary.
# Very useful for debugging with remote attach.
go build -gcflags "all=-N -l" -o rock-paper-scissors && ./rock-paper-scissors