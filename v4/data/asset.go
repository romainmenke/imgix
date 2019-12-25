package data

type AssetsData struct {
	Data     []Asset       `json:"data"`
	Included []interface{} `json:"included"`
	Jsonapi  struct {
		Version string `json:"version"`
	} `json:"jsonapi"`
	Meta struct {
		Authentication Authentication `json:"authentication"`
		Cursor         Cursor         `json:"cursor"`
		Pagination     Pagination     `json:"pagination"`
		Server         Server         `json:"server"`
	} `json:"meta"`
}

type AssetData struct {
	Data     Asset         `json:"data"`
	Included []interface{} `json:"included"`
	Jsonapi  struct {
		Version string `json:"version"`
	} `json:"jsonapi"`
	Meta struct {
		Authentication Authentication `json:"authentication"`
		Server         Server         `json:"server"`
	} `json:"meta"`
}

type Asset struct {
	Attributes struct {
		AccountID        string        `json:"account_id"`
		Categories       []interface{} `json:"categories"`
		ColorProfileName interface{}   `json:"color_profile_name"`
		Colors           struct {
			DominantColors map[string]float64 `json:"dominant_colors"`
		} `json:"colors"`
		ContentType     string             `json:"content_type"`
		DateCreated     int                `json:"date_created"`
		DateModified    int                `json:"date_modified"`
		Description     interface{}        `json:"description"`
		DpiHeight       int                `json:"dpi_height"`
		DpiWidth        int                `json:"dpi_width"`
		FaceCount       int                `json:"face_count"`
		FileSize        int                `json:"file_size"`
		HasColors       bool               `json:"has_colors"`
		HasFaces        bool               `json:"has_faces"`
		HasFrames       bool               `json:"has_frames"`
		HasLabels       bool               `json:"has_labels"`
		HasUserDefined  bool               `json:"has_user_defined"`
		HasWarnings     bool               `json:"has_warnings"`
		ImageMode       string             `json:"image_mode"`
		IsIdentified    bool               `json:"is_identified"`
		IsRenderable    bool               `json:"is_renderable"`
		Labels          map[string]float64 `json:"labels"`
		MediaHeight     int                `json:"media_height"`
		MediaKind       string             `json:"media_kind"`
		MediaWidth      int                `json:"media_width"`
		Metadata        interface{}        `json:"metadata"`
		Name            interface{}        `json:"name"`
		OriginPath      string             `json:"origin_path"`
		SourceID        string             `json:"source_id"`
		WarningAdult    int                `json:"warning_adult"`
		WarningMedical  int                `json:"warning_medical"`
		WarningRacy     int                `json:"warning_racy"`
		WarningSpoof    int                `json:"warning_spoof"`
		WarningViolence int                `json:"warning_violence"`
	} `json:"attributes"`
	ID            string `json:"id"`
	Relationships struct {
		Account Relationship `json:"account"`
		Source  Relationship `json:"source"`
	} `json:"relationships"`
	Type string `json:"type"`
}
