package cassandrawrapper

func CreateCassandraQueryHandler(config ParamConfig) func(...interface{}) []map[string]any {
	// params must match config.Query
	return func(params ...interface{}) []map[string]any {
		cluster := getCassandraClusterConfig(config.Nodes, config.DatacenterName)

		// TODO figure out when/how this wille error
		session, _ := cluster.CreateSession()
		defer session.Close()

		// TODO add a log or error if for some reason the table or keyspace does not exist
		iter := session.Query(config.Query, params...).Iter()

		// TODO look at mapScan for large pages of query results
		// TODO handle this error
		result, _ := iter.SliceMap()

		return result
	}
}
