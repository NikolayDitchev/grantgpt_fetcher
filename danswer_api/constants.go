package danswer_api

const (
	base_url = `https://ask.grantgpt.eu/api`
)

const (
	METHOD_GET    = "GET"
	METHOD_POST   = "POST"
	METHOD_PATCH  = "PATCH"
	METHOD_HEAD   = "HEAD"
	METHOD_DELETE = "DELETE"
)

// endpoints
const (

	//GET
	get_connectors      = `/manage/connector`
	get_connector_by_id = `/manage/connector/` //{connector_id}

	//POST
	// file_upload = `/manage/admin/connector/file/upload`

	// //DELETE
	// delete_connector = `/manage/admin/connector/` //{connector_id}
)
