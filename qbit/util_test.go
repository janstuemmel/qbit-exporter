package qbit_test

import (
	"qbit-exporter/qbit"
	"testing"

	easy_assert "github.com/janstuemmel/easy-assert"
)

func TestMapTorrent(t *testing.T) {
	assert := easy_assert.New(t)

	t.Run("should map all fields", func(t *testing.T) {
		assert.Equal(qbit.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "some.linux-tracker.org:1337",
			Category:   "ubuntu",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337,
			Uploaded:   1337,
		}, qbit.MapTorrent(qbit.Torrent{
			Name:       "ubuntu22.04",
			Tracker:    "https://some.linux-tracker.org:1337",
			Category:   "ubuntu",
			SavePath:   "/downloads/ubuntu",
			Downloaded: 1337,
			Uploaded:   1337,
		}))
	})

	t.Run("should map tracker error when tracker missing", func(t *testing.T) {
		assert.Equal(qbit.Torrent{
			Tracker:  "error",
			Category: "ubuntu",
		}, qbit.MapTorrent(qbit.Torrent{
			Tracker:  "",
			Category: "ubuntu",
		}))
	})

	t.Run("should map category to 'Uncategorized' when empty", func(t *testing.T) {
		assert.Equal(qbit.Torrent{
			Tracker:  "some.linux-tracker.org:1337",
			Category: "Uncategorized",
		}, qbit.MapTorrent(qbit.Torrent{
			Tracker:  "https://some.linux-tracker.org:1337",
			Category: "",
		}))
	})
}
