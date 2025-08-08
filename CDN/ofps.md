
## 负载均衡/软中断
### OSPF（Open Shortest Path First）协议
vtysh -c "show ip ospf neighbor"

https://github.com/FRRouting/frr

### XDP(eBPF)/拦截流量降低软中断

### DPDK /拦截流量降低软中断

## 缓存加速/SSD为HDD加速
目前较为成熟的方案有：
flashcache
enhanceIO
[lvmcache](https://man7.org/linux/man-pages/man7/lvmcache.7.html)
dm-cache
[opencas](https://github.com/Open-CAS/open-cas-linux)
[bcache](https://github.com/g2p/bcache-tools)
ZFS


1. dm-cache
简介：Linux Device Mapper 的一个模块，可以将一个块设备（如 SSD）作为另一个块设备（如 HDD）的缓存。
优点：内核自带，集成度高，配置灵活，支持多种缓存策略（writeback、writethrough、passthrough）。
适用场景：企业级和个人用户均可，LVM 用户集成更方便。
2. bcachefs
简介：是 bcache 作者开发的新一代文件系统，集成了缓存、RAID、快照等功能。
优点：原生支持 SSD 缓存，功能更强大，未来发展潜力大。
适用场景：愿意尝鲜新文件系统的用户。
3. EnhanceIO
简介：最初由 STEC 公司开发，后开源。与 bcache 类似，作为块设备缓存层。
优点：无需格式化原有数据盘，支持在线添加缓存。
适用场景：需要在线无损添加缓存的场景。
4. Flashcache
简介：由 Facebook 开发的块设备缓存方案，基于 dm-cache，适合大规模部署。
优点：成熟稳定，适合高负载环境。
适用场景：大规模服务器、数据中心。
5. ZFS/ZFS L2ARC
简介：ZFS 文件系统自带的二级缓存（L2ARC），可用 SSD 作为缓存。
优点：集成度高，功能强大，支持快照、校验等。
适用场景：ZFS 用户，数据完整性要求高的场景。
6. 其它相关方案
Intel CAS (Cache Acceleration Software)：Intel 的商业级缓存加速软件，支持多种缓存策略。
FSCACHE/CacheFiles：主要用于网络文件系统缓存（如 NFS），不是块设备缓存，但有类似加速效果。

| 方案 | 是否内核集成 | 是否活跃维护 | 适用场景 | 备注 |
|--------------|--------------|--------------|--------------------|----------------|
| bcache | 是 | 是 | 通用 | |
| dm-cache | 是 | 是 | LVM 用户优先 | 推荐 |
| EnhanceIO | 否 | 较少 | 在线添加缓存 | 需第三方支持 |
| Flashcache | 否 | 较少 | 大规模部署 | |
| bcachefs | 否（开发中） | 活跃 | 新文件系统 | 未来可期 |
| ZFS L2ARC | 否 | 是 | ZFS 用户 | |