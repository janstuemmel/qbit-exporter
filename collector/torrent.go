package collector

import (
	qbit "qbit-exporter/qbit"

	"github.com/prometheus/client_golang/prometheus"
)

type torrentCollector struct {
	qbit qbit.Client
}

type Torrent struct {
	Name       string
	Tracker    string
	Category   string
	SavePath   string
	Uploaded   float64
	Downloaded float64
}

var torrentUploaded = prometheus.NewDesc(
	"qbit_torrent_uploaded",
	"Torrent total bytes uploaded",
	[]string{"name", "tracker", "category"},
	nil,
)

var torrentDownloaded = prometheus.NewDesc(
	"qbit_torrent_downloaded",
	"Torrent total bytes downloaded",
	[]string{"name", "tracker", "category"},
	nil,
)

func NewTorrentCollector(qbit qbit.Client) prometheus.Collector {
	return &torrentCollector{qbit}
}

func (t *torrentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- torrentUploaded
	ch <- torrentDownloaded
}

func (t *torrentCollector) Collect(ch chan<- prometheus.Metric) {
	torrents, _ := t.qbit.GetTorrentsInfo()
	for _, torrent := range torrents {
		t := MapTorrent(torrent)
		ch <- getTorrentMetric(torrentUploaded, t.Uploaded, t)
		ch <- getTorrentMetric(torrentDownloaded, t.Downloaded, t)
	}
}

func getTorrentMetric(desc *prometheus.Desc, val float64, torrent Torrent) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		val,
		torrent.Name,
		torrent.Tracker,
		torrent.Category,
	)
}
