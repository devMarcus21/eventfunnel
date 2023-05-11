package scheme

import (
	"github.com/devMarcus21/eventfunnel/pkg/utils/dbutils/cassandrawrapper"
)

func GetSchemeTable() func(string, string) Scheme {
	// TODO look at making this dynamic and use env variables
	nodes := [3]string{
		"scheme-table-0.scheme-table-service.default.svc.cluster.local",
		"scheme-table-1.scheme-table-service.default.svc.cluster.local",
		"scheme-table-2.scheme-table-service.default.svc.cluster.local",
	}
	config := cassandrawrapper.ParamConfig{
		Query:          "SELECT * FROM schemes.schemes WHERE name = ? AND stage = ?",
		DatacenterName: "DC1-scheme-table-cluster",
		Nodes:          nodes[:],
	}
	cassandra := cassandrawrapper.CreateCassandraQueryHandler(config)

	return CreateSchemeTable(cassandra)
}
