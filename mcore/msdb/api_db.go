package msdb

// simple db should return all schemas
type SimpleDB interface{
	GetSchemas()[]string
}

