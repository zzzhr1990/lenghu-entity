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

type OofFileMapping struct {
	//logger *zap.Logger

	ID        string    `bson:"_id"`
	CreatedAt time.Time `bson:"createdAt"`
	AccessAt  time.Time `bson:"accessAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
	SliceHash string    `bson:"sliceHash"`
	FileSize  int64     `bson:"fileSize"`
	Bucket    string    `bson:"bucket"`
	Key       string    `bson:"key"`
	Source    string    `bson:"source"`
	Name      string    `bson:"name"`
	Type      int       `bson:"type"`
	Status    int       `bson:"status"`
	WcsEtag   string    `bson:"wcsEtag"`
}
