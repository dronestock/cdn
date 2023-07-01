FROM storezhang/alpine:3.18.2


LABEL author="storezhang<华寅>" \
email="storezhang@gmail.com" \
qq="160290688" \
wechat="storezhang" \
description="Drone持续集成系统CDN插件，提供常见的CDN对接服务（腾讯云、阿里云、百度云、创世云等），提供地址或目录刷新功能，方便在CI/CD过程中及时刷新相应的文件"


# 复制文件
COPY cdn /bin


RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/cdn \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 执行命令
ENTRYPOINT /bin/cdn
