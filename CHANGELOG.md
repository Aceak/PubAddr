## v0.1.0 - 2025-11-18
### 🚀 Initial Release | 初始版本

PubAddr 服务的第一个正式版本，提供了公网 IP 查询、限流、黑名单等核心能力。

### ✨ Features | 功能特性
- Provide `/` 接口用于返回客户端公网 IP  
  提供根路径 `/` 作为公网 IP 查询接口
- Provide `/health` 健康检查接口  
  新增 `/health` 端点，用于服务状态检测
- Add trusted real IP header support  
  支持从可信代理 Header（如 `X-Internal-Real-IP`）获取真实 IP
- Add rate limit middleware  
  内置基于令牌桶的限流中间件
- Add User-Agent blacklist  
  支持 UA 黑名单过滤（如 curl、nmap 等）

