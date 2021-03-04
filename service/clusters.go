package service

import (
	"github.com/khorevaa/kubodin/errors"
	"github.com/khorevaa/ras-client/serialize"
)

func (s *service) GetClusters(client ClientContext) ([]*serialize.ClusterInfo, error) {

	return s.getClusters(client)
}

func (s *service) GetClusterInfo(client ClientContext) (*serialize.ClusterInfo, error) {

	clusterID, ok := client.GetClusterID()

	if !ok {
		return nil, errors.BadRequest.New("incorrect or not set <cluster-id>")
	}

	info, err := client.GetClusterInfo(clusterID)

	if err != nil {
		return nil, err
	}
	return &info, nil
}
