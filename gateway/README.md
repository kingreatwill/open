traefik https://www.qikqiak.com/traefik-book/
kong
tyk
NGINX
APISIX
Zuul
Gravitee 

The Cloud-Native API Gateway：[Apache APISIX](https://github.com/apache/apisix)

[开源API网关架构分析](https://zhuanlan.zhihu.com/p/358862217)


### Unkey/ API 管理平台
- 密钥管理：定期轮换 API 密钥，确保密钥的安全性。
- 速率限制：根据业务需求设置合理的速率限制，防止 API 被滥用。
- 访问控制：使用 Unkey 提供的访问控制功能，确保只有授权的用户可以访问敏感 API。

[Unkey](https://github.com/unkeyed/unkey) 可以与以下开源项目结合使用，以增强其功能：

- Kong：一个开源的 API 网关，可以与 Unkey 结合使用，提供更强大的 API 管理功能。
- Prometheus：用于监控和报警，可以监控 Unkey 服务的性能和健康状态。
- Grafana：用于数据可视化，可以展示 Unkey 服务的监控数据。