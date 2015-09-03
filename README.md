# sns-fwd

Forwards HTTPS subscription confirmations and messages from SNS.

Subscription confirmations are immediately accepted and received messages are logged to stdout.

## Configration

All configs are set via the OS Envrionment:

- `FORWARD_URL`
    - required forward destination
- `FORWARD_ALARM`
    - optional config for forwarding cloudwatch alarms in the sns notification message field only, set to `true` to enable, defaults to `false`
- `USERNAME`
    - optional http basic auth username
- `PASSWORD`
    - optional http basic auth password
- `PORT`
    - optional http listener port, default is `3000`
- `ADDR`
    - optional http listener addr

> Will only require username and password via http basic auth, when `USERNAME` and `PASSWORD` are set.

## Example

```
FORWARD_URL=http://some.example.com/json/api/path \
    FORWARD_ALARM=true \
    USERNAME=username \
    PASSWORD=password \
    PORT=3000 \
    go run main.go
```

