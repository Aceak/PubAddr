# v0.1.0 - 2025-11-18
## 🚀 Initial Release | 初始版本

First test release of PubAddr service.
PubAddr 服务的第一个测试版本，提供了公网 IP 查询、限流、黑名单等核心能力。

## ✨ Features | 功能特性

- **Public IP Query | 公网 IP 查询**
  - Simple IP query service at the root path `/`, returning the visitor's public IP address.
  - 在根路径 `/` 提供简单的 IP 查询服务，返回访问者的公网 IP 地址

- **Health Check Endpoint | 健康检查端点**
  - `/health` endpoint for service health status checking.
  - 提供 `/health` 端点用于服务健康状态检查

- **Intelligent IP Retrieval | 智能 IP 获取**
  - Multi-level IP retrieval strategy with priority:
    - Trusted proxy header (configured trusted_real_ip_header)
    - X-Forwarded-For header (first IP in multi-level proxy chain)
    - X-Real-IP header
    - RemoteAddr field (fallback solution)
  - 支持多种 IP 获取策略，优先级为：
    - 可信代理头（配置的 trusted_real_ip_header）
    - X-Forwarded-For 头（多级代理链中的第一个 IP）
    - X-Real-IP 头
    - RemoteAddr 字段（兜底方案）

- **IP-based Rate Limiting | 基于 IP 的速率限制**
  - Rate limiting mechanism implemented with fixed window algorithm, configurable frequency.
  - 通过固定窗口算法实现的限流机制，可配置限制频率

- **Whitelist Support | 白名单支持**
  - Configurable whitelist for trusted IP addresses, whitelisted IPs are not subject to rate limiting.
  - 可配置的白名单，用于信任的 IP 地址，白名单中的 IP 地址不受限流机制限制

- **User Agent (UA) Blacklist | 用户代理黑名单**
  - Support for filtering specific user agent keywords.
  - 支持配置特定的用户代理关键词进行过滤


