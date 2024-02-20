FROM golang:1.17-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .


FROM alpine:latest
ENV PORT=8081 \
    DB_NAME=AUTH-TODO \
    MONGO_URI=mongodb+srv://ihimrao:sciJz7anvmj2vDWB@lms.e5ahgmo.mongodb.net/?retryWrites=true&w=majority \
    ENVIRONMENT=production
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8081
CMD ["./server"]