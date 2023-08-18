# wait

[![Build](https://github.com/ViBiOh/wait/workflows/Build/badge.svg)](https://github.com/ViBiOh/wait/actions)
[![codecov](https://codecov.io/gh/ViBiOh/wait/branch/main/graph/badge.svg)](https://codecov.io/gh/ViBiOh/wait)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ViBiOh_wait&metric=alert_status)](https://sonarcloud.io/dashboard?id=ViBiOh_wait)

## Getting started

Golang binary is built with static link. You can download it directly from the [GitHub Release page](https://github.com/ViBiOh/wait/releases) or build it by yourself by cloning this repo and running `make`.

You can configure app by passing CLI args or environment variables (cf. [Usage](#usage) section). CLI override environment variables.

## Usage

The application can be configured by passing CLI args described below or their equivalent as environment variable. CLI values take precedence over environments variables.

Be careful when using the CLI values, if someone list the processes on the system, they will appear in plain-text. Pass secrets by environment variables: it's less easily visible.

```bash
Usage of wait:
  --address           string slice  [wait] Dial address in the form network:host:port, e.g. tcp:localhost:5432 ${WAIT_ADDRESS}, as a string slice, environment variable separated by ","
  --loggerJson                      [logger] Log format as JSON ${WAIT_LOGGER_JSON} (default false)
  --loggerLevel       string        [logger] Logger level ${WAIT_LOGGER_LEVEL} (default "INFO")
  --loggerLevelKey    string        [logger] Key for level in JSON ${WAIT_LOGGER_LEVEL_KEY} (default "level")
  --loggerMessageKey  string        [logger] Key for message in JSON ${WAIT_LOGGER_MESSAGE_KEY} (default "msg")
  --loggerTimeKey     string        [logger] Key for timestamp in JSON ${WAIT_LOGGER_TIME_KEY} (default "time")
  --next              string        [wait] Action to execute after ${WAIT_NEXT}
  --nextArg           string slice  [wait] Args for the action to execute ${WAIT_NEXT_ARG}, as a string slice, environment variable separated by ","
  --timeout           duration      [wait] Timeout of retries ${WAIT_TIMEOUT} (default 10s)
```
