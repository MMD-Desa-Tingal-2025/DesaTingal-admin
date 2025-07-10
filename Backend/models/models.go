package models

import (
	"time"

	"github.com/google/uuid"
)

// Area represents a geographical area on the map
type Area struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Code        string    `json:"code" db:"code"`
	Coordinates string    `json:"coordinates" db:"coordinates"` // GeoJSON format
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// DashboardData represents data that can be sent to Looker Studio
type DashboardData struct {
	ID          uuid.UUID   `json:"id" db:"id"`
	AreaID      uuid.UUID   `json:"area_id" db:"area_id"`
	Title       string      `json:"title" db:"title"`
	Description string      `json:"description" db:"description"`
	Value       float64     `json:"value" db:"value"`
	Unit        string      `json:"unit" db:"unit"`
	Category    string      `json:"category" db:"category"`
	Tags        []string    `json:"tags" db:"tags"`
	Metadata    interface{} `json:"metadata" db:"metadata"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
	SyncedAt    *time.Time  `json:"synced_at" db:"synced_at"`
}

// LookerReport represents a report from Looker Studio
type LookerReport struct {
	ID          string                 `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	URL         string                 `json:"url"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
	Data        map[string]interface{} `json:"data"`
	Filters     map[string]interface{} `json:"filters"`
}

// MapClickEvent represents a map click event
type MapClickEvent struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	AreaID    string  `json:"area_id,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

// AreaDataRequest represents a request for area data
type AreaDataRequest struct {
	AreaID    string    `json:"area_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Category  string    `json:"category,omitempty"`
	Filters   map[string]interface{} `json:"filters,omitempty"`
}

// AreaDataResponse represents the response with area data
type AreaDataResponse struct {
	Area     Area            `json:"area"`
	Data     []DashboardData `json:"data"`
	Summary  DataSummary     `json:"summary"`
	Metadata map[string]interface{} `json:"metadata"`
}

// DataSummary represents summary statistics for area data
type DataSummary struct {
	TotalRecords int     `json:"total_records"`
	TotalValue   float64 `json:"total_value"`
	AverageValue float64 `json:"average_value"`
	MaxValue     float64 `json:"max_value"`
	MinValue     float64 `json:"min_value"`
}

// LookerSyncRequest represents a request to sync data to Looker
type LookerSyncRequest struct {
	DataIDs   []uuid.UUID `json:"data_ids"`
	ReportID  string      `json:"report_id,omitempty"`
	BatchSize int         `json:"batch_size,omitempty"`
}

// LookerSyncResponse represents the response from Looker sync
type LookerSyncResponse struct {
	Success     bool      `json:"success"`
	SyncedCount int       `json:"synced_count"`
	FailedCount int       `json:"failed_count"`
	Errors      []string  `json:"errors,omitempty"`
	SyncedAt    time.Time `json:"synced_at"`
}

// CreateDashboardDataRequest represents a request to create dashboard data
type CreateDashboardDataRequest struct {
	AreaID      uuid.UUID   `json:"area_id" binding:"required"`
	Title       string      `json:"title" binding:"required"`
	Description string      `json:"description"`
	Value       float64     `json:"value" binding:"required"`
	Unit        string      `json:"unit"`
	Category    string      `json:"category" binding:"required"`
	Tags        []string    `json:"tags"`
	Metadata    interface{} `json:"metadata"`
}

// UpdateDashboardDataRequest represents a request to update dashboard data
type UpdateDashboardDataRequest struct {
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Value       *float64    `json:"value,omitempty"`
	Unit        string      `json:"unit,omitempty"`
	Category    string      `json:"category,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
	Metadata    interface{} `json:"metadata,omitempty"`
}

// GetAreasRequest represents a request to get areas
type GetAreasRequest struct {
	Page     int    `form:"page" binding:"min=1"`
	PageSize int    `form:"page_size" binding:"min=1,max=100"`
	Search   string `form:"search"`
	SortBy   string `form:"sort_by"`
	SortDir  string `form:"sort_dir"`
}

// GetAreasResponse represents the response with areas
type GetAreasResponse struct {
	Areas      []Area `json:"areas"`
	TotalCount int    `json:"total_count"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
	TotalPages int    `json:"total_pages"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    int    `json:"code,omitempty"`
}