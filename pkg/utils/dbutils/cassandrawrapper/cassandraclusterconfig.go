package cassandrawrapper

import (
	"github.com/gocql/gocql"
)

func getCassandraClusterConfig(nodesInCluster []string, datacenterName string) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(nodesInCluster[0], nodesInCluster[1], nodesInCluster[2])
	// Routes request to a coordinator based on datacenter
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.DCAwareRoundRobinPolicy(datacenterName))

	return cluster
}
