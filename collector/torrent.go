package collector

import (
	"fmt"
	qbit "qbit-exporter/qbit"

	"github.com/prometheus/client_golang/prometheus"
)

type torrentCollector struct {
	qbit qbit.Client
}

type Torrent struct {
	Name         string
	Tracker      string
	Category     string
	SavePath     string
	Size         int64
	Progress     float64
	Seeds        int64
	Leechs       int64
	DlSpeed      int64
	UpSpeed      int64
	AmountLeft   int64
	LastActivity int64
	Eta          int64
	Uploaded     float64
	Downloaded   float64
	Ratio        float64
	AddedOn      int64
}

var torrentInfo = prometheus.NewDesc(
	"qbit_torrent_info",
	"Torrent information",
	[]string{
		"name",
		"tracker",
		"category",
		"save_path",
		"size",
		"progress",
		"seeds",
		"leechs",
		"dl_speed",
		"up_speed",
		"amount_left",
		"last_activity",
		"eta",
		"uploaded",
		"downloaded",
		"ratio",
		"added_on",
	},
	nil,
)

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

func NewTorrentCollector(qbit qbit.Client) prometheus.Collector {
	return &torrentCollector{qbit}
}

func (t *torrentCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- torrentInfo
	ch <- torrentRatio
	ch <- torrentUploaded
	ch <- torrentDownloaded
}

func (t *torrentCollector) Collect(ch chan<- prometheus.Metric) {
	torrents, _ := t.qbit.GetTorrentsInfo()
	for _, torrent := range torrents {
		t := MapTorrent(torrent)
		ch <- getTorrentInfoMetric(t)
		ch <- getTorrentDataMetric(torrentRatio, t.Ratio, t)
		ch <- getTorrentDataMetric(torrentUploaded, t.Uploaded, t)
		ch <- getTorrentDataMetric(torrentDownloaded, t.Downloaded, t)
	}
}

func getTorrentInfoMetric(torrent Torrent) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		torrentInfo,
		prometheus.GaugeValue,
		1,
		torrent.Name,
		torrent.Tracker,
		torrent.Category,
		torrent.SavePath,
		fmt.Sprintf("%d", torrent.Size),
		fmt.Sprintf("%.2f", torrent.Progress),
		fmt.Sprintf("%d", torrent.Seeds),
		fmt.Sprintf("%d", torrent.Leechs),
		fmt.Sprintf("%d", torrent.DlSpeed),
		fmt.Sprintf("%d", torrent.UpSpeed),
		fmt.Sprintf("%d", torrent.AmountLeft),
		fmt.Sprintf("%d", torrent.LastActivity),
		fmt.Sprintf("%d", torrent.Eta),
		fmt.Sprintf("%d", int64(torrent.Uploaded)),
		fmt.Sprintf("%d", int64(torrent.Downloaded)),
		fmt.Sprintf("%.2f", torrent.Ratio),
		fmt.Sprintf("%d", torrent.AddedOn),
	)
}

func getTorrentDataMetric(desc *prometheus.Desc, val float64, torrent Torrent) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		val,
		torrent.Name,
		torrent.Tracker,
		torrent.Category,
		torrent.SavePath,
	)
}
