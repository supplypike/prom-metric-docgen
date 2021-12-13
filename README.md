## Prometheus Metric Docgen

Scrapes prometheus metrics and generates documentation

## Usage

```bash
kubectl port-forward service/<name> <port>:8081
curl localhost:8081/metrics | go run main.go
```
