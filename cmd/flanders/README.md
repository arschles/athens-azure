# `Deployment`

## User Configurable

```hcl
// optional, defaults to a moniker-generated name
app = "myapp"

web {
    // optional, defaults to 3
    Replicas = 123
    Name = "Athens"
    Image = "quay.io/gomods/athens:v0.3.1"
    // optional
    Env = {
        "SOMETHING": "SOMETHING_ELSE"
    }
    // optional
    HealthyHTTPPath = "/healthy"
    // optional
    ReadyHTTPPath = "/ready"
    // optional, defaults to no ports exposed
    Port = 8080
}

job {
    Name = "Crathens"
    Image = "quay.io/arschles/crathens:canary"
}
```

## Internal

- `spec.replicas` (matches `Replicas` above)
- `spec.selector.matchLabels`
- `metadata.labels`
- `metadata.name`
- `spec.template.metadata.labels`

# `Service`

## User Configurable

Nothing

## Internal

- `spec.type = ClusterIP` if `Port` was set in the `Deployment`
- `spec.ports` for each container that has a `Port` in it in the `Deployment`
- `spec.selector` matches `spec.template.metadata.labels` in the `Deployment`

# `Ingress`

## User Configurable

Nothing

## Internal

- `spec.rules.host` - we need to get this from somewhere
  - `http`
    - `paths.path[0].path` - `/`
    - `paths.path[0].backend.serviceName` - the name of the service created
    - `paths.path[0].backend.servicePort` - the external port on the service created
