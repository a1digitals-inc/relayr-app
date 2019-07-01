# Relayr Challenge

The challenge goals is to run databases and microservices applications in kubernetes. Create an application that connects to a database, reads some data and returns this data upon HTTP request

To achieve all the **requirements** of the challenge, ***minikube*** was used as an alternative to run kubernetes localy. For application, **Golang** was chosen due to a massive impact in kubernetes ecosystem. ***Mysql*** was the database for this project.

## Persistence

- To persistent the data, HostPath was the choice as volume type on minukube. It mounts a file or directory from the host node’s filesystem into your Pod allow it to restart without losing data. In a production environment this isn't a better solution due to an ephemerality of nodes.

## Security Credentials

- To handle with database credentials, a random password is created at runtime within a k8s secret, this feature is provided by oficial helm [stable/mysql](https://github.com/helm/charts/tree/master/stable/mysql)

```text
  mysql-root-password: {{ randAlphaNum 10 | b64enc | quote }}
  mysql-password:  {{ randAlphaNum 10 | b64enc | quote }}
```

## Resilience

- Liveness and Readiness probes are configured to ensure that pods always are in a consistent state, when a readiness fails it will remove endpoints of workload, if a liveness check fails it will restart the pod:

    ```yaml
        (...)
        livenessProbe:
          httpGet:
            path: /healthz
            port: http
          initialDelaySeconds: 15
          timeoutSeconds: 5
        readinessProbe:
          httpGet:
            path: /healthz
            port: http
          initialDelaySeconds: 5
          timeoutSeconds: 1
        (...)
    ```

## Security [Bonus]

- [**Kubernetes Administration**] Allow network access to database for application pod by label:

    ```yaml
    apiVersion: networking.k8s.io/v1
    kind: NetworkPolicy
    metadata:
      name: access-mysql
    spec:
      podSelector:
        matchLabels:
          app: {{ template "relayr-app.name" . }}-mysql
      ingress:
      - from:
        - podSelector:
            matchLabels:
              app: "{{ template "relayr-app.name" . }}"
        ports:
        - protocol: TCP
          port: 3306
    ```

- [**Kubernetes Administration**] Don't let workloads run as root user, use security contexts to improve security:

    ```yaml
      (...)
      securityContext:
        runAsUser: 1000
        fsGroup: 2000
      initContainers:
      - name: check-db-ready
        image: mysql:latest
        securityContext:
          allowPrivilegeEscalation: false
      (...)
      containers:
      - name: {{ template "relayr-app.name" . }}
        image: "{{ .Values.image }}:{{ .Values.imageTag }}"
        imagePullPolicy: {{ .Values.imagePullPolicy | quote }}
        securityContext:
          allowPrivilegeEscalation: false
      (...)
    ```

## Continuous Integration

It will automatically build and push images to [quay.io](https://quay.io) using [travis](https://travis-ci.org) as CI.

## Metrics
Instrumenting Application with [prometheus](https://prometheus.io). Prometheus has an official Go client library that you can use to instrument Go applications.

To expose Prometheus metrics appplication use /metrics HTTP endpoint
here's a result example.

```text
# HELP go_gc_duration_seconds A summary of the GC invocation durations.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0.000202427
go_gc_duration_seconds{quantile="0.25"} 0.000202427
go_gc_duration_seconds{quantile="0.5"} 0.100661494
go_gc_duration_seconds{quantile="0.75"} 0.100661494
go_gc_duration_seconds{quantile="1"} 0.100661494
go_gc_duration_seconds_sum 0.100863921
go_gc_duration_seconds_count 2
# HELP go_goroutines Number of g
(...)
```