
DISPLAY_CLIENT_METRICS_ENDPOINT="http://192.168.100.195/dashboard/summary"

displaylocal:
	DISPLAY_CLIENT_METRICS_ENDPOINT="$(DISPLAY_CLIENT_METRICS_ENDPOINT)" \
		go run main.go display

buildraspios64:
	GOOS=linux \
		GOARCH=arm64 \
		go build \
		-v \
		-a \
		-ldflags \
		'-extldflags "-static"' \
		.
