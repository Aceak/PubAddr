# Change Log | 变更日志

## v0.2.0 (2025-11-24)

### 🆕 New Features | 新特性

- **TCP Service Support | TCP 服务支持**
  - Added TCP protocol support for lightweight IP querying
  - 新增TCP协议支持，提供更轻量的IP查询接口

- **TCP SDK Client | TCP SDK 客户端**
  - Provided client SDK for connecting to TCP service
  - 提供了用于连接TCP服务的客户端SDK

- **Enhanced Access Control | 增强的访问控制**
  - Support for token-based authentication in TCP service
  - 支持基于令牌(Token)的TCP服务认证

### 🔧 Improvements | 改进

- **Optimized Rate Limiting | 优化的限流机制**
  - Implemented specialized rate limiting for TCP service
  - 为TCP服务实现了专门的限流策略

- **Enhanced Logging | 日志系统增强**
  - Added support for runtime log level adjustment
  - 支持运行时动态调整日志级别

- **Extended Configuration | 配置项扩展**
  - Added TCP-related configurations: `enable_tcp` and `tcp_addr`
  - 添加TCP相关配置：`enable_tcp`和`tcp_addr`

- **Build System Optimization | 构建系统优化**
  - Improved cross-platform build script with support for more architectures
  - 改进跨平台构建脚本，支持更多架构

## v0.1.0 (2025-11-18)

### ✨ Features | 功能特性

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


