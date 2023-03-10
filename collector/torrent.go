package collector

import (
	"net/url"
	qbit "qbit-exporter/qbit"

	"github.com/prometheus/client_golang/prometheus"
)

type torrentCollector struct {
	qbit qbit.Client
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
		ch <- getTorrentMetric(torrentUploaded, float64(torrent.Uploaded), torrent)
		ch <- getTorrentMetric(torrentDownloaded, float64(torrent.Downloaded), torrent)
	}
}

func getTorrentMetric(desc *prometheus.Desc, val float64, torrent qbit.Torrent) prometheus.Metric {

	// TODO: mapping should not stay here
	trackerHost := "error"
	tracker, _ := url.Parse(torrent.Tracker)

	if tracker.Host != "" {
		trackerHost = tracker.Host
	}

	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		val,
		torrent.Name,
		trackerHost,
		torrent.Category,
	)
}
