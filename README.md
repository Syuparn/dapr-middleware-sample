# dapr-middleware-sample
About sample Dapr middleware just for experiments

# Usage

## Prepare

```bash
$ kind create cluster
$ skaffold run

# ensure app is running
$ kubectl port-forward svc/hello 3500:3500
$ curl localhost:3500/v1.0/invoke/hello/method/hello
Hello, World!
```

## Use Middleware

```bash
```
