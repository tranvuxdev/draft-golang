# Step 1: Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Step 2: Run stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
# Copy thêm file .env hoặc các folder cần thiết nếu có
COPY --from=builder /app/.env . 

EXPOSE 8080
CMD ["./main"]