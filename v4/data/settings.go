package data

type Settings struct {
	AwsAccessKey            interface{}   `json:"aws_access_key"`
	AwsBucket               interface{}   `json:"aws_bucket"`
	AwsPrefix               interface{}   `json:"aws_prefix"`
	AwsSecretKeyEncrypted   interface{}   `json:"aws_secret_key_encrypted"`
	AzureAccount            interface{}   `json:"azure_account"`
	AzureBucket             interface{}   `json:"azure_bucket"`
	AzurePrefix             interface{}   `json:"azure_prefix"`
	AzureSasExpiry          interface{}   `json:"azure_sas_expiry"`
	AzureSasStringEncrypted interface{}   `json:"azure_sas_string_encrypted"`
	AzureServiceType        interface{}   `json:"azure_service_type"`
	CacheTTLBehavior        string        `json:"cache_ttl_behavior"`
	CacheTTLError           int           `json:"cache_ttl_error"`
	CacheTTLValue           int           `json:"cache_ttl_value"`
	CrossdomainCorsEnabled  bool          `json:"crossdomain_cors_enabled"`
	CrossdomainXMLEnabled   bool          `json:"crossdomain_xml_enabled"`
	CustomDomains           []interface{} `json:"custom_domains"`
	DefaultParams           struct{}      `json:"default_params"`
	GcsAccessKey            interface{}   `json:"gcs_access_key"`
	GcsBucket               interface{}   `json:"gcs_bucket"`
	GcsPrefix               interface{}   `json:"gcs_prefix"`
	GcsSecretKeyEncrypted   interface{}   `json:"gcs_secret_key_encrypted"`
	ImageError              interface{}   `json:"image_error"`
	ImageErrorAppendQs      bool          `json:"image_error_append_qs"`
	ImageMissing            interface{}   `json:"image_missing"`
	ImageMissingAppendQs    bool          `json:"image_missing_append_qs"`
	ImgixSubdomains         []string      `json:"imgix_subdomains"`
	SecureURLEnabled        bool          `json:"secure_url_enabled"`
	Type                    string        `json:"type"`
	WebfolderPrefix         string        `json:"webfolder_prefix"`
}
