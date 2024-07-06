#!/usr/bin/env bash

set -euo pipefail

modprobe vfio_pci
dpdk-devbind.py -b ice 90:00.1
