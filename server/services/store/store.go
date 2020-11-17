//go:generate mockgen -destination=mockstore/mockstore.go -package mockstore . Store
package store

import "github.com/mattermost/mattermost-octo-tasks/server/model"

// Store represents the abstraction of the data storage.
type Store interface {
	GetBlocksWithParentAndType(parentID string, blockType string) ([]model.Block, error)
	GetBlocksWithParent(parentID string) ([]model.Block, error)
	GetBlocksWithType(blockType string) ([]model.Block, error)
	GetSubTree2(blockID string) ([]model.Block, error)
	GetSubTree3(blockID string) ([]model.Block, error)
	GetAllBlocks() ([]model.Block, error)
	GetParentID(blockID string) (string, error)
	InsertBlock(block model.Block) error
	DeleteBlock(blockID string) error
	Shutdown() error
	GetSystemSettings() (map[string]string, error)
	SetSystemSetting(key string, value string) error
	GetUserById(userID string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
}
