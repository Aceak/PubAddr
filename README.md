# PubAddr

<sub>[中文文档](README_ZH.md)</sub>

PubAddr is a lightweight public IP lookup service designed to provide simple, efficient, and secure IP address retrieval. It can be deployed on your own server, giving you full control over data privacy, and includes useful security features such as rate limiting, User-Agent filtering, and access token authentication.

## ✨ Key Features

- **Public IP Lookup**: Provides a simple IP query endpoint at `/`
- **Health Check**: Monitor service status via the `/health` endpoint
- **Intelligent IP Extraction**: Supports multiple strategies and properly handles proxy chains
- **IP-based Rate Limiting**: Protects your service from abuse
- **Whitelist Support**: Trusted IPs can be excluded from rate limits
- **User-Agent Blacklist**: Block unnecessary or unwanted clients

## 🚀 Getting Started

### Requirements

- Go 1.18 or later (only required for development or building from source)
- Supports Windows (amd64/arm64), Linux (amd64/arm64/arm/riscv64), and macOS (amd64/arm64)

### Installation

#### Prebuilt Binaries

Download the binary for your platform from  
[GitHub Releases](https://github.com/Aceak/PubAddr/releases):

```bash
# Linux
wget https://github.com/Aceak/PubAddr/releases/latest/download/pubaddr_linux_amd64
chmod +x pubaddr_linux_amd64
```

#### Build from Source

```bash
# Clone the repository
git clone https://github.com/Aceak/PubAddr.git
cd PubAddr

# Build
./build.sh

# Or build directly using Go
go build -o pubaddr ./cmd/main.go
```

### Basic Usage

1. **Prepare the configuration file**

   ```bash
   wget https://raw.githubusercontent.com/Aceak/PubAddr/refs/heads/main/config.example.yaml
   cp config.example.yaml config.yaml
   ```

2. **Start the service**

   ```bash
   ./pubaddr
   ```

3. **Test the service**

   ```bash
   curl http://localhost:8080/        # Query public IP
   curl http://localhost:8080/health  # Health check
   ```

## Systemd Autostart (Linux)

1. **Create service file**

   Create `/etc/systemd/system/pubaddr.service`:

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

2. **Enable and start**

   ```bash
   sudo systemctl enable pubaddr
   sudo systemctl start pubaddr
   ```

3. **Check status**

   ```bash
   sudo systemctl status pubaddr
   ```

## ⚙️ Configuration Guide

PubAddr uses a YAML configuration file. Below is the configuration reference:

```yaml
server:
  addr: ":8080"         # Listen address and port
  log_level: "info"     # Log level: debug, info, warn, error, fatal

ip_header:
  trusted_real_ip_header: ""  # Trusted proxy header (optional)

security:
  access_token: ""       # Access token (optional)
  rate_limit_per_ip: 1   # Requests per second per IP
  rate_duration: "5s"    # Rate limit duration window
  enable_ua_block: false # Enable User-Agent blocking
  blocked_user_agents:
    - "bot"
    - "spider"
```

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
