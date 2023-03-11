# qbit exporter

A simple prometheus exporter for qbittorrent written in golang.

## Metrics

### Torrents

| Name | Type | Description |
| ---- | ---- | ----------- |
| `qbit_torrent_info`                 | gauge | Torrent general info |
| `qbit_torrent_uploaded_bytes_total` | gauge | Torrent total bytes uploaded |
| `qbit_torrent_uploaded_bytes_total` | gauge | Torrent total bytes downloaded |
| `qbit_torrent_ratio_total`          | gauge | Torrent ratio |
| `qbit_torrent_size_bytes_total`     | gauge | Torrent size |
| `qbit_torrent_seeds_total`          | gauge | Torrent number of seeders connected to |
| `qbit_torrent_leechs_total`         | gauge | Torrent number of leechers connected to |
| `qbit_torrent_upspeed_bytes`        | gauge | Torrent upload speed in bytes |
| `qbit_torrent_dlspeed_bytes`        | gauge | Torrent download speed in bytes |

Labels: `name`, `tracker`, `category`

## Usage

Make sure webui is enabled in qbittorrent settings.

### Environment

* `QB_USER` Your qbittorrent webui username
* `QB_PASS` Your qbittorrent webui password
* `QB_URL` Your qbittorrent webui url

### Build and run

```sh
CGO_ENABLED=0 go build
./qbit-exporter
```

### Docker

```sh
docker run \
    -e QB_USER=admin \
    -e QB_PASS=adminadmin \
    -e QB_URL=http://localhost:8080 \
    -p 9897:9897 \
    ghcr.io/janstuemmel/qbit-exporter
```

For a `docker-compose` setup, see [example folder](./example)

### Prometheus

Add this to your prometheus configuration

```yml
scrape_configs:
  - job_name: qbit
    static_configs:
      - targets: [ "localhost:9897" ]
```
