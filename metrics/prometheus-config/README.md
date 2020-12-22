## Running Prometheus in Docker:

### Just so fast and so simple command as below:
```
docker run --name prometheus -d -p 9090:9090 -v ./prometheus.yml:/etc/prometheus/prometheus.yml  prom/prometheus
```