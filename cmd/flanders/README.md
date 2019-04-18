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


