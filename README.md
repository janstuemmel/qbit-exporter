# qbit exporter

A simple prometheus exporter written in golang.

## Usage

Make sure webui is enabled in qbittorrent settings.

### Environment

* `QB_USER` Your qbittorrent webui username
* `QB_PASS` Your qbittorrent webui password
* `QB_URL` Your qbittorrent webui url

### Docker

```sh
docker run \
    -e QB_USER=admin \
    -e QB_PASS=adminadmin \
    -e QB_URL=http://localhost:8080 \
    -p 9897:9897 \
    ghcr.io/janstuemmel/qbit-exporter
```

For a docker-compose setup file, see [exampel folder](./example)

### Prometheus

Add this to your prometheus configuration

```yml
scrape_configs:
  - job_name: qbit
    static_configs:
      - targets: [ "localhost:9897" ]
```
