package dlog

import (
	"testing"

	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"gorm.io/gorm"
)

func TestNewZapGorm2(t *testing.T) {
	logger := NewZapGorm2(kitexzap.NewLogger())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	db, _ := gorm.Open(nil, &gorm.Config{Logger: logger})

	// do stuff normally
	var _ = db // avoid "unused variable" warn
}
