package convert

import (
	"errors"

	"github.com/TheCrether/tachiyomi-paperback-converter/config"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

func ConvertPaperbackToTachiyomi(pBackup *paperback.Backup) (*tachiyomi.Backup, error) {
	backup := config.DefaultTachiyomiBackup()
	return backup, errors.New("not implemented")
}
