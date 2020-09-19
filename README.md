# DTU-Admin

顾名思义，DTU-Admin就是DTU的通用管理后台，主要用于简化物联网或工业互联网的后端开发。

【传送门】另有衍生项目IoT-Admin，集成了Modbus和主流PLC协议，支持数据自动采集，
历史存储，定时和自动控制策略，可以直接用于大部分物联网项目后端，真正实现零代码开发。
请移步 [<真格智能实验室>](https://labs.zgwit.com)

## 起源
作者曾经接触多个物联网实际项目的后端，需求大同小异，
因为团队不同，实现就千奇百怪了，大家都在重复地造轮子。
痛定思痛，于是提取共同的功能，做成了通用的DTU-Admin，
并且开源免费分享给小伙伴儿们使用。

## 主要功能
1. 自建通道，支持TCP/UDP，Server/Client；
2. 接收注册包，过滤心跳包；
3. 数据透传，通过MQTT转发到其他应用；
4. 连接状态监控，方便协议开发，和数据检查；
5. 虚拟串口透传，方便远程调试。

## 支持的数据通道
1. DTU：数据传输单元；
2. 2G/3G/4G/5G 移动通信模组；
3. WIFI芯片或模组，比如：ESP8266、ESP32；
4. RJ45芯片或模组，比如：W5500、ENC28J60。

## 其他
1. 暂不支持MQTT5





