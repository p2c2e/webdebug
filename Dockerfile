FROM alpine:3.7
# RUN apk add --no-cache ca-certificates
CMD ["mkdir /bin"]
COPY ./go_http_debug /bin/go_http_debug
CMD ["chmod","+x","/bin/go_http_debug"]
EXPOSE 8888
ENTRYPOINT ["/bin/go_http_debug"]
