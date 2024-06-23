package eu_client

const (
	TypeTopics  = "1"
	TypeGrant   = "2"
	TypeCascade = "8"

	StatusOpen        = "31094501"
	StatusForthcoming = "31094502"
)

type Query struct {
	Bool struct {
		Must []struct {
			Terms map[string][]string `json:"terms,omitempty"`
		} `json:"must"`
	} `json:"bool"`
}

type Option func(*Query)

func NewQuery(opts ...Option) *Query {
	query := &Query{}

	for _, opt := range opts {
		opt(query)
	}

	return query
}

func WithStatus(status []string) Option {
	return func(q *Query) {
		q.Bool.Must = append(q.Bool.Must, struct {
			Terms map[string][]string "json:\"terms,omitempty\""
		}{
			Terms: map[string][]string{
				"status": status,
			},
		})
	}
}

func WithTypes(types []string) Option {
	return func(q *Query) {
		q.Bool.Must = append(q.Bool.Must, struct {
			Terms map[string][]string "json:\"terms,omitempty\""
		}{
			Terms: map[string][]string{
				"type": types,
			},
		})
	}
}
