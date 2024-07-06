#!/usr/bin/env bash

set -euxo pipefail

echo 1 | sudo tee /sys/module/vfio/parameters/enable_unsafe_noiommu_mode
ip a del 10.0.66.1 dev enp144s0f0np0 || true
ip a add 10.0.66.1/24 dev enp144s0f0np0 || true
modprobe vfio_pci
dpdk-devbind.py -b vfio-pci 90:00.1
echo 8192 > /sys/devices/system/node/node0/hugepages/hugepages-2048kB/nr_hugepages
echo 8192 > /sys/devices/system/node/node1/hugepages/hugepages-2048kB/nr_hugepages
pktgen -l 1,16,17,18,19 -a 90:00.1 -n 16 -- -T -m "[16/17/18/19].0" -f "$(dirname "$0")"/cmd.pkt

