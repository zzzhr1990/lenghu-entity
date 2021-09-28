package offline

import "time"

type OfflineInfo struct {
	//logger *zap.Logger

	Hash      string    `json:"hash" gorm:"primarykey, type:varchar(80)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	URL       string    `json:"url"`
	Torrent   string    `json:"torrent"` // Torrent File Info (Base64)
	FileSize  int64     `json:"fileSize"`
	Status    int       `json:"status"`
	Type      int       `json:"type"`
	Source    string    `json:"source"`
	Name      string    `json:"name"`
}
