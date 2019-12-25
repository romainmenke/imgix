package api

import (
	"context"
	"fmt"
	"net/http"
	"path"

	"github.com/romainmenke/imgix/v4/data"
)

func (c *Client) Assets(ctx context.Context, sourceID string, offset int, size int) (*data.AssetsData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.imgix.com/v4/assets/"+sourceID, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("sort", "-date_created")
	q.Add("page[cursor]", fmt.Sprint(offset))
	q.Add("page[size]", fmt.Sprint(size))
	req.URL.RawQuery = q.Encode()

	req = req.Clone(ctx)

	out := &data.AssetsData{}

	err = c.do(req, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) Asset(ctx context.Context, sourceID string, assetID string) (*data.AssetData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.imgix.com/v4/assets/"+path.Join(sourceID, assetID), nil)
	if err != nil {
		return nil, err
	}

	req = req.Clone(ctx)

	out := &data.AssetData{}

	err = c.do(req, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
