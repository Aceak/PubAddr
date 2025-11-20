# PubAddr

<sub>[English Documentation](README.md)</sub>

PubAddr 是一个轻量级的公网 IP 查询服务，旨在提供简单、高效、安全的 IP 地址查询功能。它可以部署在自己的服务器上，完全控制数据隐私，并且提供了丰富的安全特性，如 IP 限流、用户代理过滤和访问令牌认证。

## ✨ 核心特性

- **公网 IP 查询**：在根路径 `/` 提供简单的 IP 查询服务
- **健康检查**：通过 `/health` 端点监控服务状态
- **智能 IP 获取**：支持多种 IP 获取策略，正确处理代理链
- **基于 IP 的速率限制**：防止滥用，保护服务稳定
- **白名单支持**：可配置信任的 IP 地址，白名单中的 IP 地址不受限流机制限制
- **用户代理（UA）黑名单**：过滤不需要的客户端

## 🚀 快速开始

### 前提条件

- Go 1.18 或更高版本（仅开发环境需要）
- 任意 Windows（amd64/arm64）、Linux（amd64/arm64/arm/riscv64）或 macOS（amd64/arm64）

### 安装方法

#### 二进制安装

从 [GitHub Releases](https://github.com/Aceak/PubAddr/releases) 下载适合你操作系统的二进制文件：

```bash
# Linux
wget https://github.com/Aceak/PubAddr/releases/latest/download/pubaddr_linux_amd64
chmod +x pubaddr_linux_amd64
```

#### 从源码编译

```bash
# 克隆仓库
git clone https://github.com/Aceak/PubAddr.git
cd PubAddr

# 编译
./build.sh

# 或者直接使用 Go 命令
go build -o pubaddr ./cmd/main.go
```

### 基本使用

1. **准备配置文件**

   复制示例配置文件并根据需要修改：

   ```bash
   wget https://raw.githubusercontent.com/Aceak/PubAddr/refs/heads/main/config.example.yaml
   cp config.example.yaml config.yaml
   ```

2. **启动服务**

   ```bash
   # 使用默认配置文件
   ./pubaddr
   ```

3. **测试服务**

   ```bash
   # 查询公网 IP
   curl http://localhost:8080/

   # 健康检查
   curl http://localhost:8080/health
   ```

### 配置开机自启

#### Linux 系统

1. **创建 Systemd 服务文件**

   创建 `/etc/systemd/system/pubaddr.service` 文件：

   ```ini
   [Unit]
   Description=PubAddr - Public IP Query Service
   After=network.target

   [Service]
   Type=simple
   User=pubaddr
   Group=pubaddr
   WorkingDirectory=/path/to/pubaddr
   ExecStart=/path/to/pubaddr/pubaddr --config /path/to/config.yaml
   Restart=on-failure
   RestartSec=3

   [Install]
   WantedBy=multi-user.target
   ```

2. **启用并启动服务**

   ```bash
   sudo systemctl enable pubaddr
   sudo systemctl start pubaddr
   ```

3. **检查服务状态**

   ```bash
   sudo systemctl status pubaddr
   ```

## ⚙️ 配置指南

PubAddr 使用 YAML 格式的配置文件。以下是配置选项的详细说明：

```yaml
# 服务器配置
server:
  addr: ":8080"         # 服务监听地址和端口
  log_level: "info"      # 日志级别: debug, info, warn, error, fatal

# IP 头配置
ip_header:
  trusted_real_ip_header: ""  # 可信代理头名称，留空表示不使用

# 安全配置
security:
  access_token: ""       # 访问令牌，留空表示不启用认证
  rate_limit_per_ip: 1   # 每个 IP 的请求频率限制（次/秒）
  rate_duration: "5s"    # 速率限制的时间窗口
  enable_ua_block: false # 是否启用用户代理过滤
  blocked_user_agents:   # 被阻止的用户代理关键词列表
    - "bot"
    - "spider"
```

## 📄 许可证

本项目采用 MIT 许可证 - 详情请查看 [LICENSE](LICENSE) 文件。