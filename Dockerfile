FROM alpine:3.1

COPY ./server /opt/codefight/

WORKDIR /opt/codefight

EXPOSE 3000

ENV PORT=3000
CMD ["./server"]
