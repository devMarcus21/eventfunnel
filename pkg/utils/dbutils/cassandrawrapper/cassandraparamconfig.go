package cassandrawrapper

type ParamConfig struct {
	Query          string
	DatacenterName string
	Nodes          []string // names of the nodes in the cassandra cluster
}
