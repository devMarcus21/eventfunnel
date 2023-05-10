package model

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/cassandrawrapper"
)

func GetModelTable() func(string, string) Model {
	// TODO look at making this dynamic and use env variables
	nodes := [3]string{
		"models-table-0.models-table-service.default.svc.cluster.local",
		"models-table-1.models-table-service.default.svc.cluster.local",
		"models-table-2.models-table-service.default.svc.cluster.local",
	}
	config := cassandrawrapper.ParamConfig{
		Query:          "SELECT * FROM models.models WHERE name = ? AND stage = ?",
		DatacenterName: "DC1-models-table-cluster",
		Nodes:          nodes[:],
	}
	cassandra := cassandrawrapper.CreateCassandraQueryHandler(config)

	return CreateModelTable(cassandra)
}
