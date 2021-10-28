package offline

import "time"

type OfflineInfo struct {
	//logger *zap.Logger

	Hash      string    `json:"hash" gorm:"primary_key;type:varchar(80)"`
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

type OfflineQueue struct {
	//logger *zap.Logger
	ID        int64     `gorm:"primary_key"`
	Hash      string    `json:"hash" gorm:"type:varchar(80)"`
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

type FastDownloadQueue struct {
	//logger *zap.Logger
	ID        int64     `gorm:"primary_key"`
	WcsHash   string    `json:"wcsHash" gorm:"type:varchar(80)"`
	CloudHash string    `json:"cloudHash" gorm:"type:varchar(80)"`
	CreatedAt time.Time `json:"createdAt"`
	FileSize  int64     `json:"fileSize"`
	Status    int       `json:"status"`
}

type OofFileMapping struct {
	//logger *zap.Logger

	ID        string    `bson:"_id" json:"id"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
	AccessAt  time.Time `bson:"accessAt" json:"accessAt"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt"`
	SliceHash string    `bson:"sliceHash" json:"sliceHash"`
	FileSize  int64     `bson:"fileSize" json:"fileSize"`
	Bucket    string    `bson:"bucket" json:"bucket"`
	Key       string    `bson:"key" json:"key"`
	Source    string    `bson:"source" json:"source"`
	Name      string    `bson:"name" json:"name"`
	Type      int       `bson:"type" json:"type"`
	Status    int       `bson:"status" json:"status"`
	WcsEtag   string    `bson:"wcsEtag" json:"wcsEtag"`
}
