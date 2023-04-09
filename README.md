# wait

[![Build](https://github.com/ViBiOh/wait/workflows/Build/badge.svg)](https://github.com/ViBiOh/wait/actions)
[![codecov](https://codecov.io/gh/ViBiOh/wait/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/wait)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_wait&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_wait)

## Getting started

Golang binary is built with static link. You can download it directly from the [GitHub Release page](https://github.com/ViBiOh/wait/releases) or build it by yourself by cloning this repo and running `make`.

A Docker image is available for `amd64`, `arm` and `arm64` platforms on Docker Hub: [vibioh/wait](https://hub.docker.com/r/vibioh/wait/tags).

You can configure app by passing CLI args or environment variables (cf. [Usage](#usage) section). CLI override environment variables.

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of api:
  -address string
        [server] Listen address {API_ADDRESS}
  -cert string
        [server] Certificate file {API_CERT}
  -corsCredentials
        [cors] Access-Control-Allow-Credentials {API_CORS_CREDENTIALS}
  -corsExpose string
        [cors] Access-Control-Expose-Headers {API_CORS_EXPOSE}
  -corsHeaders string
        [cors] Access-Control-Allow-Headers {API_CORS_HEADERS} (default "Content-Type")
  -corsMethods string
        [cors] Access-Control-Allow-Methods {API_CORS_METHODS} (default "GET")
  -corsOrigin string
        [cors] Access-Control-Allow-Origin {API_CORS_ORIGIN} (default "*")
  -csp string
        [owasp] Content-Security-Policy {API_CSP} (default "default-src 'self'; base-uri 'self'")
  -frameOptions string
        [owasp] X-Frame-Options {API_FRAME_OPTIONS} (default "deny")
  -graceDuration duration
        [http] Grace duration when SIGTERM received {API_GRACE_DURATION} (default 30s)
  -hsts
        [owasp] Indicate Strict Transport Security {API_HSTS} (default true)
  -idleTimeout duration
        [server] Idle Timeout {API_IDLE_TIMEOUT} (default 2m0s)
  -key string
        [server] Key file {API_KEY}
  -location string
        [hello] TimeZone for displaying current time {API_LOCATION} (default "Europe/Paris")
  -loggerJson
        [logger] Log format as JSON {API_LOGGER_JSON}
  -loggerLevel string
        [logger] Logger level {API_LOGGER_LEVEL} (default "INFO")
  -loggerLevelKey string
        [logger] Key for level in JSON {API_LOGGER_LEVEL_KEY} (default "level")
  -loggerMessageKey string
        [logger] Key for message in JSON {API_LOGGER_MESSAGE_KEY} (default "message")
  -loggerTimeKey string
        [logger] Key for timestamp in JSON {API_LOGGER_TIME_KEY} (default "time")
  -okStatus int
        [http] Healthy HTTP Status code {API_OK_STATUS} (default 204)
  -port uint
        [server] Listen port (0 to disable) {API_PORT} (default 1080)
  -prometheusAddress string
        [prometheus] Listen address {API_PROMETHEUS_ADDRESS}
  -prometheusCert string
        [prometheus] Certificate file {API_PROMETHEUS_CERT}
  -prometheusGzip
        [prometheus] Enable gzip compression of metrics output {API_PROMETHEUS_GZIP}
  -prometheusIdleTimeout duration
        [prometheus] Idle Timeout {API_PROMETHEUS_IDLE_TIMEOUT} (default 10s)
  -prometheusIgnore string
        [prometheus] Ignored path prefixes for metrics, comma separated {API_PROMETHEUS_IGNORE}
  -prometheusKey string
        [prometheus] Key file {API_PROMETHEUS_KEY}
  -prometheusPort uint
        [prometheus] Listen port (0 to disable) {API_PROMETHEUS_PORT} (default 9090)
  -prometheusReadTimeout duration
        [prometheus] Read Timeout {API_PROMETHEUS_READ_TIMEOUT} (default 5s)
  -prometheusShutdownTimeout duration
        [prometheus] Shutdown Timeout {API_PROMETHEUS_SHUTDOWN_TIMEOUT} (default 5s)
  -prometheusWriteTimeout duration
        [prometheus] Write Timeout {API_PROMETHEUS_WRITE_TIMEOUT} (default 10s)
  -readTimeout duration
        [server] Read Timeout {API_READ_TIMEOUT} (default 5s)
  -shutdownTimeout duration
        [server] Shutdown Timeout {API_SHUTDOWN_TIMEOUT} (default 10s)
  -tracerRate string
        [tracer] OpenTracing sample rate, 'always', 'never' or a float value {API_TRACER_RATE} (default "always")
  -tracerURL string
        [tracer] OpenTracing gRPC endpoint (e.g. http://otel-exporter:4317) {API_TRACER_URL}
  -url string
        [alcotest] URL to check {API_URL}
  -userAgent string
        [alcotest] User-Agent for check {API_USER_AGENT} (default "Alcotest")
  -writeTimeout duration
        [server] Write Timeout {API_WRITE_TIMEOUT} (default 10s)
```
