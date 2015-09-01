# sns-sink

Sinks HTTPS subscription confirmations and messages from SNS.

Subscription confirmations are immediately accepted and received messages are logged to stdout.

## Endpoints

### `/sns`

This endpoint requires no authentication.

### `/sns/with-auth`

This endpoint requires basic authentication, with credentials of `user` and `pass`.

## Forwarding (designed around CW alarms)

> Designed for forwarding message payload as json only.

```
FORWARD=http://some.example.com/json/api/path PORT=3000 go run main.go
```
