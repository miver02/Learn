# 基于 Ubuntu 22.04 镜像
FROM ubuntu:22.04

# Ubuntu 使用 apt 而非 apk，更新包索引并安装必要证书
RUN apt update && \
    apt install -y ca-certificates && \
    rm -rf /var/lib/apt/lists/*  # 清理缓存减小镜像体积

# 创建应用目录
WORKDIR /app

# 复制 webook 可执行文件
COPY webook .

# 赋予执行权限
RUN chmod +x ./webook

# 启动应用
CMD ["./webook"]
    