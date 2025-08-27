# ใช้ Go official image
FROM golang:1.25-trixie

# ตั้ง working directory
WORKDIR /app

# คัดลอกไฟล์ go.mod และ go.sum
COPY go.mod go.sum ./

# ติดตั้ง dependencies
RUN go mod download

# คัดลอก source code ทั้งหมด
COPY . .

# build binary
RUN go build -o main .

# expose port
EXPOSE 8080

# run binary
CMD ["./main"]
