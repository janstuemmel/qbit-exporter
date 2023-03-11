package qbit

import (
	"net/url"
)

func getOrDefault(str string, def string) string {
	if str != "" {
		return str
	}
	return def
}

func MapTorrent(torrent Torrent) Torrent {
	tracker, _ := url.Parse(torrent.Tracker)
	return Torrent{
		Name:         torrent.Name,
		Tracker:      getOrDefault(tracker.Host, "error"),
		Category:     getOrDefault(torrent.Category, "Uncategorized"),
		SavePath:     torrent.SavePath,
		Size:         torrent.Size,
		Seeds:        torrent.Seeds,
		Leechs:       torrent.Leechs,
		DlSpeed:      torrent.DlSpeed,
		UpSpeed:      torrent.UpSpeed,
		Downloaded:   torrent.Downloaded,
		Uploaded:     torrent.Uploaded,
		Ratio:        torrent.Ratio,
		LastActivity: torrent.LastActivity,
	}
}

func MapTorrents(torrents []Torrent) []Torrent {
	var mappedTorrents []Torrent
	for _, torrent := range torrents {
		mappedTorrents = append(mappedTorrents, MapTorrent(torrent))
	}
	return mappedTorrents
}
