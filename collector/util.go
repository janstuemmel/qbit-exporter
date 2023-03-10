package collector

import (
	"net/url"
	qbit "qbit-exporter/qbit"
)

func getOrDefault(str string, def string) string {
	if str != "" {
		return str
	}
	return def
}

func MapTorrent(torrent qbit.Torrent) Torrent {
	tracker, _ := url.Parse(torrent.Tracker)
	return Torrent{
		Name:       torrent.Name,
		Tracker:    getOrDefault(tracker.Host, "error"),
		Category:   getOrDefault(torrent.Category, "Uncategorized"),
		SavePath:   torrent.SavePath,
		Downloaded: float64(torrent.Downloaded),
		Uploaded:   float64(torrent.Uploaded),
	}
}
