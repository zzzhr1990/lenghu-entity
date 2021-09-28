package offline

type OfflineType int

const (
	_                     = iota
	NormalURL OfflineType = 10 * iota
	Magnet
	TorrentFile
	IPFS
	ThunderLink
)

type DownloadStatus int

const (
	Added                      = iota
	Downloading DownloadStatus = 10 * iota
	Complete
	Error
)
