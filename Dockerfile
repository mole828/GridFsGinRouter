FROM golang:alpine
# ENV TZ Asia/Shanghai

WORKDIR /app
ADD . .
RUN ["go", "build", "-o", "example", "."]
EXPOSE 7999:7999
CMD ["./example"]
