#!/bin/bash
set -e

# Docker uses tmpfs for /dev, which doesn't auto-create device nodes.
# Mount devtmpfs to get automatic /dev/input/eventX creation from uinput.
mount -t devtmpfs devtmpfs /dev 2>/dev/null || true

exec "$@"
