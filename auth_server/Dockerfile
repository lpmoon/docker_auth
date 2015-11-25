FROM busybox
EXPOSE 5001
EXPOSE 8081
ENTRYPOINT ["/auth_server"]
CMD ["/config/auth_config.yml"]
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY ./server/manager/conf /conf
COPY ./server/manager/static /static
COPY ./server/manager/views /views
COPY auth_server .
