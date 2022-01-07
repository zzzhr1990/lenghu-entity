package mapping

type MappingFile struct {
	ID             string `bson:"_id" json:"id"`
	FileName       string `bson:"fileName" json:"fileName"`
	Sha1           string `bson:"sha1" json:"sha1"`
	BlockSha1      string `bson:"blockSha1" json:"blockSha1"`
	FileSize       int64  `bson:"fileSize" json:"fileSize"`
	Bucket         string `bson:"bucket" json:"bucket"`
	ObjectKey      string `bson:"objectKey" json:"objectKey"`
	Store          int32  `bson:"store" json:"store"`
	Status         int32  `bson:"status" json:"status"`
	LastAccessTime int64  `bson:"lastAccessTime" json:"lastAccessTime"`
}
