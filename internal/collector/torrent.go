package collector

import (
	"fmt"
	"qbit-exporter/internal/qbit"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type torrentCollector struct {
	qbit qbit.Client
}

var torrentLabels = []string{"name", "tracker", "category"}

var torrentInfo = prometheus.NewDesc(
	"qbit_torrent_info",
	"Torrent info",
	[]string{
		"name",
		"category",
		"state",
		"tracker",
		"save_path",
		"downloaded",
		"uploaded",
		"ratio",
		"seeds",
		"leechs",
		"added_on",
		"last_activity",
	},
	nil,
)

var torrentUploaded = prometheus.NewDesc(
	"qbit_torrent_uploaded_bytes_total",
	"Torrent total bytes uploaded",
	torrentLabels,
	nil,
)

var torrentDownloaded = prometheus.NewDesc(
	"qbit_torrent_downloaded_bytes_total",
	"Torrent total bytes downloaded",
	torrentLabels,
	nil,
)

var torrentRatio = prometheus.NewDesc(
	"qbit_torrent_ratio_total",
	"Torrent ratio",
	torrentLabels,
	nil,
)

var torrentSize = prometheus.NewDesc(
	"qbit_torrent_size_bytes_total",
	"Torrent size",
	torrentLabels,
	nil,
)

var torrentSeeds = prometheus.NewDesc(
	"qbit_torrent_seeds_total",
	"Torrent number of seeds connected to",
	torrentLabels,
	nil,
)

var torrentLeechs = prometheus.NewDesc(
	"qbit_torrent_leechs_total",
	"Torrent number of leechers connected to",
	torrentLabels,
	nil,
)

var torrentUpSpeed = prometheus.NewDesc(
	"qbit_torrent_upspeed_bytes",
	"Torrent upload speed in bytes",
	torrentLabels,
	nil,
)

var torrentDlSpeed = prometheus.NewDesc(
	"qbit_torrent_dlspeed_bytes",
	"Torrent download speed in bytes",
	torrentLabels,
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
		ch <- getTorrentInfoMetric(torrent)
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
	)
}

func getTorrentInfoMetric(torrent qbit.Torrent) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		torrentInfo,
		prometheus.GaugeValue,
		1,
		torrent.Name,
		torrent.Category,
		torrent.State,
		torrent.Tracker,
		torrent.SavePath,
		strconv.FormatInt(torrent.Downloaded, 10),
		strconv.FormatInt(torrent.Uploaded, 10),
		fmt.Sprintf("%.2f", torrent.Ratio),
		strconv.FormatInt(torrent.Seeds, 10),
		strconv.FormatInt(torrent.Leechs, 10),
		strconv.FormatInt(torrent.AddedOn, 10),
		strconv.FormatInt(torrent.LastActivity, 10),
	)
}
