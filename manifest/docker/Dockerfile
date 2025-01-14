FROM --platform=linux/amd64 alpine
###############################################################################
#                                LiveDog-build
###############################################################################

LABEL maintainer="shichen437 <shichen437@126.com>"

ENV WORKDIR=/LiveDog
ENV TZ=Asia/Shanghai \
    PROJECT_UPLOAD=$WORKDIR/upload \
    PROJECT_OUTPUT=$WORKDIR/video

RUN mkdir -p $PROJECT_OUTPUT && \
    mkdir -p $PROJECT_UPLOAD && \
    apk --no-cache add nginx ffmpeg libc6-compat

ADD web/nginx.conf /etc/nginx/nginx.conf
COPY web/dist /usr/share/nginx/html
COPY i18n $WORKDIR/i18n
COPY manifest/config/config_docker.yaml $WORKDIR/config.yaml
COPY manifest/migrate $WORKDIR/manifest/migrate
COPY temp/linux_amd64/main $WORKDIR/main
COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

VOLUME $PROJECT_OUTPUT
VOLUME $PROJECT_UPLOAD

ENV DATABASE_DEFAULT_LINK="" \
    PROJECT_SM4KEY="abcdefghijklmnopqrstuvwxyz123456"

EXPOSE 9876

WORKDIR $WORKDIR
ENTRYPOINT ["sh"]
CMD ["/entrypoint.sh"]