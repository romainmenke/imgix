package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/romainmenke/imgix/v4/data"
)

func (c *Client) Sources(ctx context.Context, page int, size int) (*data.SourcesData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.imgix.com/v4/sources", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("filter[enabled]", "true")
	q.Add("page[number]", fmt.Sprint(page))
	q.Add("page[size]", fmt.Sprint(size))
	req.URL.RawQuery = q.Encode()

	req = req.Clone(ctx)

	out := &data.SourcesData{}

	err = c.do(req, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) SourcesFind(ctx context.Context, search string, page int, size int) (*data.SourcesData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.imgix.com/v4/sources", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("filter[enabled]", "true")
	q.Add("filter[or:name]", search)
	q.Add("filter[or:production.aws_bucket]", search)
	q.Add("filter[or:production.custom_domains]", search)
	q.Add("filter[or:production.gcs_bucket]", search)
	q.Add("filter[or:production.imgix_subdomains]", search)
	q.Add("filter[or:production.webfolder_prefix]", search)
	q.Add("page[number]", fmt.Sprint(page))
	q.Add("page[size]", fmt.Sprint(size))
	req.URL.RawQuery = q.Encode()

	req = req.Clone(ctx)

	out := &data.SourcesData{}

	err = c.do(req, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) Source(ctx context.Context, id string) (*data.SourceData, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.imgix.com/v4/sources/"+id, nil)
	if err != nil {
		return nil, err
	}

	req = req.Clone(ctx)

	out := &data.SourceData{}

	err = c.do(req, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
