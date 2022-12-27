package convert

import (
	"errors"

	"github.com/TheCrether/tachiyomi-paperback-converter/config"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

func ConvertTachiyomiToPaperback(tBackup *tachiyomi.Backup) (*paperback.Backup, error) {
	backup := config.DefaultPaperbackBackup()
	return backup, errors.New("not implemented")
}
