# qbit exporter

A simple prometheus exporter for qbittorrent written in golang.

## Metrics

### Torrents

| Name | Type | Description |
| ---- | ---- | ----------- |
| `qbit_torrent_uploaded`   | gauge | Torrent total bytes uploaded |
| `qbit_torrent_downloaded` | gauge | Torrent total bytes downloaded |
| `qbit_torrent_ratio`      | gauge | Torrent ratio |
| `qbit_torrent_size`       | gauge | Torrent size |
| `qbit_torrent_seeds`      | gauge | Torrent number of seeders connected to |
| `qbit_torrent_leechs`     | gauge | Torrent number of leechers connected to |
| `qbit_torrent_upspeed`    | gauge | Torrent upload speed in bytes |
| `qbit_torrent_dlspeed`    | gauge | Torrent download speed in bytes |

Labels: `name`, `tracker`, `category`, `save_path`

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

For a `docker-compose` setup, see [example folder](./example)

### Prometheus

Add this to your prometheus configuration

```yml
scrape_configs:
  - job_name: qbit
    static_configs:
      - targets: [ "localhost:9897" ]
```
