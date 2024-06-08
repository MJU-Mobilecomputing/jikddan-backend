FROM golang:latest AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o jjikdan ./cmd
COPY config config

# 5. 최종 이미지 생성
FROM build

WORKDIR /app

# 필요한 런타임 의존성 복사
COPY --from=build /app/jjikdan .
COPY --from=build /app/config config

RUN chmod +x ./jjikdan

EXPOSE 8080 
# 어플리케이션 실행
CMD ["./jjikdan","run","cmd/main.go"]
