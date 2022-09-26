package dto

type OverallStatsResponse struct {
	AccountsCreated uint64 `json:"accounts_created"`
	TotalAllocated  uint64 `json:"total_allocated"`
	StorageUsed     uint64 `json:"storage_used"`
	FilesUploaded   uint64 `json:"files_uploaded"`
	SitesUploaded   uint64 `json:"sites_uploaded"`
}
