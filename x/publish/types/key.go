package types

const (
	//Modulename is the name of the module
	ModuleName = "publish"

	// StoreKey to be used when creating the KV Store
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)
