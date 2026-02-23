# Step 1: Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
# Trỏ đúng vào file main.go trong thư mục cmd/router
RUN go build -o main ./cmd/main.go 

# Step 2: Run stage
FROM alpine:latest
WORKDIR /root/
# Chỉ copy file thực thi, KHÔNG copy .env
COPY --from=builder /app/main . 

# Render sẽ tự cấp PORT, nhưng EXPOSE giúp tài liệu hóa image
EXPOSE 8080
CMD ["./main"]