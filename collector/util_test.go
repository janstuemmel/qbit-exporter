package collector_test

import (
	"qbit-exporter/collector"
	"qbit-exporter/qbit"
	"testing"

	easy_assert "github.com/janstuemmel/easy-assert"
)

func TestMapTorrent(t *testing.T) {
	assert := easy_assert.New(t)

	t.Run("should map all fields", func(t *testing.T) {
		assert.Equal(collector.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "some.linux-tracker.org:1337",
			Category:   "ubuntu",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337.0,
			Uploaded:   1337.0,
		}, collector.MapTorrent(qbit.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "https://some.linux-tracker.org:1337",
			Category:   "ubuntu",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337,
			Uploaded:   1337,
		}))
	})

	t.Run("should map tracker error when tracker missing", func(t *testing.T) {
		assert.Equal(collector.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "error",
			Category:   "ubuntu",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337.0,
			Uploaded:   1337.0,
		}, collector.MapTorrent(qbit.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "",
			Category:   "ubuntu",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337,
			Uploaded:   1337,
		}))
	})

	t.Run("should map category to 'Uncategorized' when empty", func(t *testing.T) {
		assert.Equal(collector.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "some.linux-tracker.org:1337",
			Category:   "Uncategorized",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337.0,
			Uploaded:   1337.0,
		}, collector.MapTorrent(qbit.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "https://some.linux-tracker.org:1337",
			Category:   "",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337,
			Uploaded:   1337,
		}))
	})
}
