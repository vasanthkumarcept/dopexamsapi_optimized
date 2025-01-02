FROM golang:1.20-alpine AS builder
WORKDIR /home/admin1/Documents/recruitment
COPY . .
RUN go build -o /recruit

# Run stage
FROM alpine
#RUN addgroup -S appgroup && adduser -S admin1 -G appgroup


WORKDIR /app
COPY --from=builder /recruit /recruit
COPY .env .
#COPY  /files ./files
#COPY  /INTERFACE ./INTERFACE
#COPY  /keys  ./keys

#VOLUME [ "app/files","app/INTERFACE" ]

EXPOSE 8080

ENTRYPOINT [ "/recruit" ]

#USER "1000"