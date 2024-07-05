```
Addresses:
        eno1           fe80:0000:0000:0000:0210:e0ff:fe5c:9e46  00:10:e0:5c:9e:46  0000:40:00.0  ixgbe
        enp32s0f0np0   fe80:0000:0000:0000:3efd:feff:fe2c:ef20  3c:fd:fe:2c:ef:20  0000:20:00.0  i40e
        enp32s0f1np1   fe80:0000:0000:0000:3efd:feff:fe2c:ef22  3c:fd:fe:2c:ef:22  0000:20:00.1  i40e
        enp48s0f0np0   fe80:0000:0000:0000:527c:6fff:fe24:2740  50:7c:6f:24:27:40  0000:30:00.0  i40e
        enp48s0f1np1   fe80:0000:0000:0000:527c:6fff:fe24:2741  50:7c:6f:24:27:41  0000:30:00.1  i40e
        enp144s0f0np0  fe80:0000:0000:0000:42a6:b7ff:fe98:7208  40:a6:b7:98:72:08  0000:90:00.0  ice
        enp144s0f1np1  fe80:0000:0000:0000:42a6:b7ff:fe98:7209  40:a6:b7:98:72:09  0000:90:00.1  ice
        enp160s0f0np0  fe80:0000:0000:0000:0ec4:7aff:feea:be2a  0c:c4:7a:ea:be:2a  0000:a0:00.0  mlx5_core
        enp160s0f1np1  fe80:0000:0000:0000:0ec4:7aff:feea:be2b  0c:c4:7a:ea:be:2b  0000:a0:00.1  mlx5_core
        enp176s0f0np0  fe80:0000:0000:0000:5e6f:69ff:feef:d250  5c:6f:69:ef:d2:50  0000:b0:00.0  bnxt_en
        enp176s0f1np1  fe80:0000:0000:0000:5e6f:69ff:feef:d251  5c:6f:69:ef:d2:51  0000:b0:00.1  bnxt_en

Connectivity:
        enp32s0f0np0   i40e       | enp160s0f0np0  mlx5_core  |  10 Gbps
        enp32s0f1np1   i40e       | enp160s0f1np1  mlx5_core  |  10 Gbps
        enp48s0f0np0   i40e       | enp176s0f0np0  bnxt_en    |  25 Gbps
        enp48s0f1np1   i40e       | enp176s0f1np1  bnxt_en    |  25 Gbps
        enp144s0f0np0  ice        | enp144s0f1np1  ice        | 100 Gbps
```
