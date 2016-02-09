FROM alpine:3.1

COPY ./server-linux-amd64 /opt/codefight/

WORKDIR /opt/codefight

EXPOSE 3000

ENV PORT=3000
CMD ["./server-linux-amd64"]
