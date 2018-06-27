FROM docker.io/golang:alpine
MAINTAINER devfeel

# 创建基础目录
RUN mkdir -p /home

# 拷贝运行文件及配置文件
COPY performance /home/performance

# 拷贝启动脚本
COPY docker-entrypoint.sh /home/docker-entrypoint.sh

# 暴露端口
EXPOSE 80 81 8080

# 赋权
RUN chmod +x /home/performance
RUN chmod +x /home/docker-entrypoint.sh

ENTRYPOINT ["/home/docker-entrypoint.sh"]