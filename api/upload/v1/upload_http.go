package v1

type UploadParams struct{}

type UploadResult struct {
	URL string `json:"url"`
}
