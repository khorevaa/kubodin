package service

import (
	db "github.com/khorevaa/kubodin/database"
	"github.com/khorevaa/kubodin/models"
	"github.com/khorevaa/kubodin/service/cache"
	"github.com/khorevaa/ras-client/serialize"
)

const (
	clustersTpl  = "%s.clusters"
	infobasesTpl = "%s.infobases"
)

var _ Service = (*service)(nil)

//Service interface allows us to access the CRUD Operations
type Service interface {
	Repository() db.Repository

	CreateInfobase(ctt ClientContext, info *serialize.InfobaseInfo, createDB bool) (*serialize.InfobaseInfo, error)
	DropInfobase(ctt ClientContext, deleteDB bool) error

	GetInfobases(ctt ClientContext) (serialize.InfobaseSummaryList, error)

	GetClusters(ctt ClientContext) ([]*serialize.ClusterInfo, error)
	GetClusterInfo(ctt ClientContext) (*serialize.ClusterInfo, error)

	GetAppServers() (apps []*models.AppServer, err error)
	GetAppServer(name string) (*models.AppServer, error)

	GetCache() cache.Cache
}

func NewService(cache cache.Cache, repository db.Repository) (Service, error) {

	return &service{
		cache:      cache,
		repository: repository,
	}, nil
}

type service struct {
	repository db.Repository
	cache      cache.Cache
}

func (s service) GetCache() cache.Cache {
	return s.cache
}

func (s *service) Repository() db.Repository {
	return s.repository
}

func (s *service) GetAppServers() (apps []*models.AppServer, err error) {
	return s.repository.GetAppServers()
}

func (s *service) GetAppServer(name string) (*models.AppServer, error) {
	return s.repository.GetAppServer(name)
}
