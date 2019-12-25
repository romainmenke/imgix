package data

import "context"

type SourceData struct {
	Data     Source        `json:"data"`
	Included []interface{} `json:"included"`
	Jsonapi  struct {
		Version string `json:"version"`
	} `json:"jsonapi"`
	Meta struct {
		Authentication Authentication `json:"authentication"`
		Server         Server         `json:"server"`
	} `json:"meta"`
}

type Source struct {
	Attributes struct {
		Capabilities struct {
			Multiframe bool `json:"multiframe"`
		} `json:"capabilities"`
		DateCreated      int      `json:"date_created"`
		DateDeployed     int      `json:"date_deployed"`
		DateModified     int      `json:"date_modified"`
		DeploymentStatus string   `json:"deployment_status"`
		EnableMultiframe bool     `json:"enable_multiframe"`
		Enabled          bool     `json:"enabled"`
		GlobalVersion    int      `json:"global_version"`
		Name             string   `json:"name"`
		Production       Settings `json:"production"`
		SecureURLToken   string   `json:"secure_url_token"`
	} `json:"attributes"`
	ID            string                  `json:"id"`
	Relationships map[string]Relationship `json:"relationships"`
	Type          string                  `json:"type"`
}

type SourcesData struct {
	Data     []Source      `json:"data"`
	Included []interface{} `json:"included"`
	Jsonapi  struct {
		Version string `json:"version"`
	} `json:"jsonapi"`
	Meta struct {
		Authentication Authentication `json:"authentication"`
		Pagination     Pagination     `json:"pagination"`
		Server         Server         `json:"server"`
	} `json:"meta"`
}

func (x Source) Get(ctx context.Context, client SourceGetter) (*SourceData, error) {
	return client.Source(ctx, x.ID)
}

type SourceGetter interface {
	Source(ctx context.Context, id string) (*SourceData, error)
}

func (x Source) Assets(ctx context.Context, client AssetsGetter, offset int, size int) (*AssetsData, error) {
	return client.Assets(ctx, x.ID, offset, size)
}

type AssetsGetter interface {
	Assets(ctx context.Context, sourceID string, offset int, size int) (*AssetsData, error)
}
