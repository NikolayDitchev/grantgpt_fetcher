package eu_client

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

func WithStatus(statuses ...string) Option {
	return func(q *Query) {
		q.Bool.Must = append(q.Bool.Must, struct {
			Terms map[string][]string "json:\"terms,omitempty\""
		}{
			Terms: map[string][]string{
				"status": statuses,
			},
		})
	}
}

func WithTypes(types ...string) Option {
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
