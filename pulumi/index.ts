import * as pulumi from "@pulumi/pulumi";
import * as k8s from "@pulumi/kubernetes";

// Minikube does not implement services of type `LoadBalancer`; require the user to specify if we're
// running on minikube, and if so, create only services of type ClusterIP.
const config = new pulumi.Config();
const isMinikube = config.requireBoolean("isMinikube");

const namespace = new k8s.core.v1.Namespace("athens-pulumi", {
    metadata: {
        name: "athens-pulumi"
    }
});

const appName = "athens";
const image = "gomods/athens:v0.5.0"
const appLabels = { app: appName };
const deployment = new k8s.apps.v1.Deployment(appName, {
    metadata: {
        namespace: namespace.metadata.name
    },
    spec: {
        selector: { matchLabels: appLabels },
        replicas: 50,
        template: {
            metadata: { labels: appLabels },
            spec: {
                containers: [
                    {
                        name: appName,
                        image: image,
                        env: [{
                            name: "ATHENS_GOGET_WORKERS",
                            value: "5"
                        }]
                    }
                ]
            }
        }
    }
});

// Allocate an IP to the Deployment.
const frontend = new k8s.core.v1.Service(appName, {
    metadata: {
        namespace: namespace.metadata.name,
        labels: deployment.spec.template.metadata.labels
    },
    spec: {
        type: isMinikube ? "ClusterIP" : "LoadBalancer",
        ports: [{ port: 80, targetPort: 3000, protocol: "TCP" }],
        selector: appLabels
    }
});

// When "done", this will print the public IP.
export const ip = isMinikube
    ? frontend.spec.clusterIP
    : frontend.status.apply(status => status.loadBalancer.ingress[0].ip);
