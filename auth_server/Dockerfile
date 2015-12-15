FROM ubuntu:14.04 
EXPOSE 5001
EXPOSE 8081
# RUN sed -i 's/http:\/\/archive.ubuntu.com\/ubuntu\//http:\/\/mirrors.yun-idc.com\/ubuntu\//g' /etc/apt/sources.list
RUN apt-get update && apt-get install -y apache2-utils
ENTRYPOINT ["/auth_server"]
CMD ["/config/auth_config.yml"]
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./server/manager/conf /conf
COPY ./server/manager/static /static
COPY ./server/manager/views /views
COPY auth_server .
