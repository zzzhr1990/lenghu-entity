package offline

import (
	"net/http"

	"github.com/zzzhr1990/go-http-utilities/responses"
)

var ErrorMissingURLOrTorrentInfo = &responses.ResponseEntity{
	Status:    http.StatusBadRequest,
	Reference: "MISSING_TORRENT_OR_URL",
	Message:   "You must specfic url or torrent base64 info for download",
}

var ErrorBadURLOrTorrentFile = &responses.ResponseEntity{
	Status:    http.StatusBadRequest,
	Reference: "BAD_URL_OR_TORRENT",
	Message:   "Url or torrent cannot be parsed",
}

var ErrorMissingType = &responses.ResponseEntity{
	Status:    http.StatusBadRequest,
	Reference: "MISSING_TYPE",
	Message:   "Missing download type",
}
