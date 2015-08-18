# sns-sink

Sinks HTTPS subscription confirmations and messages from SNS.

Subscription confirmations are immediately accepted and received messages are logged to stdout.

## Endpoints

### `/sns`

This endpoint requires no authentication.

### `/sns/with-auth`

This endpoint requires basic authentication, with credentials of `user` and `pass`.
