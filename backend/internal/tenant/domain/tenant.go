package domain

type Tenant struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Slug   string         `json:"slug"`
	Domain string         `json:"domain"`
	Status string         `json:"status"`
	Config map[string]any `json:"config"`
}
