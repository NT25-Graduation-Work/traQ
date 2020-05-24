package repository

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"github.com/leandro-lugaresi/hub"
	"github.com/traPtitech/traQ/migration"
	"github.com/traPtitech/traQ/model"
	"github.com/traPtitech/traQ/service/rbac/role"
	"github.com/traPtitech/traQ/utils/gormutil"
	"github.com/traPtitech/traQ/utils/storage"
	"go.uber.org/zap"
)

// GormRepository リポジトリ実装
type GormRepository struct {
	db     *gorm.DB
	hub    *hub.Hub
	logger *zap.Logger
	chTree *channelTreeImpl
	stamps *stampRepository
	fs     storage.FileStorage
}

// Channel implements ReplaceMapper interface.
func (repo *GormRepository) Channel(path string) (uuid.UUID, bool) {
	id := repo.chTree.GetChannelIDFromPath(path)
	return id, id != uuid.Nil
}

// Group implements ReplaceMapper interface.
func (repo *GormRepository) Group(name string) (uuid.UUID, bool) {
	g, err := repo.GetUserGroupByName(name)
	if err != nil {
		return uuid.Nil, false
	}
	return g.ID, true
}

// User implements ReplaceMapper interface.
func (repo *GormRepository) User(name string) (uuid.UUID, bool) {
	u, err := repo.GetUserByName(name, false)
	if err != nil {
		return uuid.Nil, false
	}
	return u.GetID(), true
}

// Sync implements Repository interface.
func (repo *GormRepository) Sync() (init bool, err error) {
	if err := migration.Migrate(repo.db); err != nil {
		return false, err
	}

	// チャンネルツリー構築
	var channels []*model.Channel
	if err := repo.db.Where(&model.Channel{IsPublic: true}).Find(&channels).Error; err != nil {
		return false, err
	}
	repo.chTree, err = makeChannelTree(channels)
	if err != nil {
		return false, err
	}

	// スタンプをキャッシュ
	var stamps []*model.Stamp
	if err := repo.db.Find(&stamps).Error; err != nil {
		return false, err
	}
	repo.stamps = makeStampRepository(stamps)

	// 管理者ユーザーの確認
	if exists, err := gormutil.RecordExists(repo.db, &model.User{Role: role.Admin}); err != nil {
		return false, err
	} else if !exists {
		_, err := repo.CreateUser(CreateUserArgs{
			Name:     "traq",
			Password: "traq",
			Role:     role.Admin,
		})
		if err != nil {
			return false, err
		}
		return true, err
	}
	return false, nil
}

// NewGormRepository リポジトリ実装を初期化して生成します
func NewGormRepository(db *gorm.DB, fs storage.FileStorage, hub *hub.Hub, logger *zap.Logger) (Repository, error) {
	repo := &GormRepository{
		db:     db,
		hub:    hub,
		logger: logger.Named("repository"),
		fs:     fs,
	}
	return repo, nil
}
