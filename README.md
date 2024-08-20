# kubectl-decodesecret
A kubectl plugin for printing a secret with all values decoded

## Build
```/bin/bash
go mod tidy
go build
```

## Install
Place the binary in your $PATH

## Usage
```
$ kubectl decodesecret jfrog-oci
apiVersion: v1
kind: Secret
metadata:
    name: jfrog-oci
    namespace: argocd
    labels:
        argocd.argoproj.io/secret-type: repo-creds
    annotations:
        managed-by: argocd.argoproj.io
stringData:
    password: eyJ2ZXIiOiIyIiwcThLOGFNUGl5b0o4SGdZNEJXeW8xWVNHQkdkaVdUMUJaemFXdExiVFI4In0.eyJzdWIiOiJqZmFjQDAxZmU0bXNoMGdiZnZtMGptcHpiOHkwdzcxL3VzZXJzL2VkZXZvcHMtamZyb2ctc3ZjLXVzZXIiLCJzY3AiOiJhcHBsaWVkLXBlcm1pc3Npb25zL3VzZXIiLCJhdWQiOiIqQCoiLCJpc3MiOiJqZmZlQDAxZmU0bXNoMGdiZnZtMGptcHpiOHkwdzcxIiwiaWF0IjoxNzI0MTY1Mjg0LCJqdGkiOiJkNmY5OGE3MC00ODM5LTRlYjgtYjZkYy01NzljMWFkNzFjM2QifQ.Qbhc7FO0lViyJ6uWnKsmWqj6kwFwM7u0oKJxPn6Lb5dJgRJtjSIzgf0FwnU9Tog9g1t8qRUldBVWs0jrqR5U0Lie0x3Bb9O2_ieFQXcdMteoiLHJFNdnwAkgQQjWpuovM26X0ZuwDktAGjhZs7zzeItu9YYeisG3O497_SMbAbThJYHNPdK28OFXwnNaCbEsO--zFmB7VliuqpEzw2X0TCFXCf9h04U8r-dipzBeZ_dSL0nJ3C_jS7x-wpCrxqAhZ31NkKFU8h83fmQc5Pj67Xv92-hkibt27Ki--Pou8lqSyBOOI27qaYJTVoIgXdaFqiwVmIzxasPye6e9nO2wCw
    url: example.jfrog.io
    username: jfrog-svc-user
type: Opaque
```
