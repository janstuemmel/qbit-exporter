package collector

import (
	qbit "qbit-exporter/qbit"

	"github.com/prometheus/client_golang/prometheus"
)

type torrentCollector struct {
	qbit qbit.Client
}

var torrentUploaded = prometheus.NewDesc(
	"qbit_torrent_uploaded",
	"Torrent total bytes uploaded",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

var torrentDownloaded = prometheus.NewDesc(
	"qbit_torrent_downloaded",
	"Torrent total bytes downloaded",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

var torrentRatio = prometheus.NewDesc(
	"qbit_torrent_ratio",
	"Torrent ratio",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

var torrentSize = prometheus.NewDesc(
	"qbit_torrent_size",
	"Torrent size",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

var torrentSeeds = prometheus.NewDesc(
	"qbit_torrent_seeds",
	"Torrent number of seeds connected to",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

var torrentLeechs = prometheus.NewDesc(
	"qbit_torrent_leechs",
	"Torrent number of leechers connected to",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

var torrentUpSpeed = prometheus.NewDesc(
	"qbit_torrent_upspeed",
	"Torrent upload speed in bytes",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

var torrentDlSpeed = prometheus.NewDesc(
	"qbit_torrent_dlspeed",
	"Torrent download speed in bytes",
	[]string{"name", "tracker", "category", "save_path"},
	nil,
)

func NewTorrentCollector(qbit qbit.Client) prometheus.Collector {
	return &torrentCollector{qbit}
}

func (t *torrentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- torrentUploaded
	ch <- torrentDownloaded
	ch <- torrentRatio
	ch <- torrentSize
	ch <- torrentSeeds
	ch <- torrentLeechs
	ch <- torrentUpSpeed
	ch <- torrentDlSpeed
}

func (t *torrentCollector) Collect(ch chan<- prometheus.Metric) {
	torrents, _ := t.qbit.GetTorrentsInfo()
	for _, torrent := range torrents {
		ch <- getTorrentMetric(torrentUploaded, torrent.Uploaded, torrent)
		ch <- getTorrentMetric(torrentDownloaded, torrent.Downloaded, torrent)
		ch <- getTorrentMetric(torrentRatio, torrent.Ratio, torrent)
		ch <- getTorrentMetric(torrentSize, torrent.Size, torrent)
		ch <- getTorrentMetric(torrentSeeds, torrent.Seeds, torrent)
		ch <- getTorrentMetric(torrentLeechs, torrent.Leechs, torrent)
		ch <- getTorrentMetric(torrentUpSpeed, torrent.UpSpeed, torrent)
		ch <- getTorrentMetric(torrentDlSpeed, torrent.DlSpeed, torrent)
	}
}

func getTorrentMetric[T MetricValue](desc *prometheus.Desc, val T, torrent qbit.Torrent) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		float64(val),
		torrent.Name,
		torrent.Tracker,
		torrent.Category,
		torrent.SavePath,
	)
}
