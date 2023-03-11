package qbit

import (
	"encoding/json"

	"github.com/imroc/req/v3"
)

type Client struct {
	user string
	pass string
	url  string
}

type Torrent struct {
	Name         string  `json:"name"`
	AddedOn      int64   `json:"added_on"`
	LastActivity int64   `json:"last_activity"`
	TimeActive   int64   `json:"time_active"`
	State        string  `json:"state"`
	Tracker      string  `json:"tracker"`
	Category     string  `json:"category"`
	SavePath     string  `json:"save_path"`
	Size         int64   `json:"size"`
	Seeds        int64   `json:"num_seeds"`
	Leechs       int64   `json:"num_leechs"`
	DlSpeed      int64   `json:"dlspeed"`
	UpSpeed      int64   `json:"upspeed"`
	Uploaded     int64   `json:"uploaded"`
	Downloaded   int64   `json:"downloaded"`
	Ratio        float64 `json:"ratio"`
}

func NewClient(user string, pass string, url string) Client {
	return Client{user: user, pass: pass, url: url}
}

func logout(cli *req.Client) {
	cli.R().Get("/api/v2/auth/logout")
}

func authedRequest[T any](q *Client, fun func(cli *req.Client) (T, error)) (t T, err error) {
	c := req.C().SetBaseURL(q.url)

	payload := map[string]string{
		"username": q.user,
		"password": q.pass,
	}

	res, err := c.R().
		SetFormData(payload).
		Post("/api/v2/auth/login")

	if err != nil {
		return t, err
	}

	cookie := res.GetHeader("set-cookie")

	c.SetCommonHeader("Cookie", cookie)

	defer logout(c)

	return fun(c)
}

func (q *Client) GetTorrentsInfo() ([]Torrent, error) {
	return authedRequest(q, func(cli *req.Client) ([]Torrent, error) {
		var torrents []Torrent

		res, err := cli.R().
			Get("api/v2/torrents/info")

		if err != nil {
			return torrents, err
		}

		json.Unmarshal(res.Bytes(), &torrents)

		return MapTorrents(torrents), nil
	})
}
