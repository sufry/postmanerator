package postman

type Collection struct {
	Name        string
	Description string
	Requests    []Request
	Folders     []Folder
	Structures  []StructureDefinition
}

type Request struct {
	ID            string
	Name          string
	Description   string
	Method        string
	URL           string
	PayloadType   string
	PayloadRaw    string
	PayloadParams []KeyValuePair
	PathVariables []KeyValuePair
	Headers       []KeyValuePair
	Responses     []Response
	Tests         string
	Auth          []collectionV210Auth
	AuthParams    []KeyValuePair
}

type Response struct {
	ID         string
	Name       string
	Status     string
	StatusCode int
	Body       string
	Headers    []KeyValuePair
}

type Folder struct {
	ID          string
	Name        string
	Description string
	Folders     []Folder
	Requests    []Request
}

type StructureDefinition struct {
	Name        string
	Description string
	Fields      []StructureFieldDefinition
}

type StructureFieldDefinition struct {
	Name        string
	Description string
	Type        string
}

type KeyValuePair struct {
	Name        string
	Key         string
	Value       interface{}
	Description string
	Source      string
	Type        string
	Disabled    bool
}

type collectionV210Auth struct {
	Type   string
	Basic  []collectionV210KeyValuePair
	Bearer []collectionV210KeyValuePair
	Oauth2 []collectionV210KeyValuePair
}
