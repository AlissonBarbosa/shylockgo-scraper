package models

type ProjectData struct {
  ID string
  Sponsor string
  Name string
}

type ProjectDesc struct {
  ID uint `gorm:"primaryKey"`
  Timestamp int64 `json:"timestamp"`
  ProjectID string `json:"project_id"`
  ProjectName string `json:"project_name"`
  ProjectSponsor string `json:"project_sponsor"`
}

type ProjectQuota struct {
  ID uint `gorm:"primaryKey"`
  Timestamp int64 `json:"timestamp"`
  ProjectID string `json:"project_id"`
  QuotaRam int64 `json:"quota_ram"`
  QuotaVcpu int64 `json:"quota_vcpu"` 
}

type ProjectQuotaUsage struct {
  ID uint `gorm:"primaryKey"`
  Timestamp int64 `json:"timestamp"`
  ProjectID string `json:"project_id"`
  RamUsage int64 `json:"ram_usage"`
  VcpuUsage int64 `json:"vcpu_usage"` 
}

type ProjectUsers struct {
  ID uint `gorm:"primaryKey"`
  Timestamp int64 `json:"timestamp"`
  ProjectID string `json:"project_id"`
  UserID string `json:"user_id"` 
}
