package offline

type OfflineType int

const (
	_                     = iota
	NormalURL OfflineType = 10 * iota
	TorrentFile
	IPFS
	ThunderLink
)
