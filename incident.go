package cachet

// Incident Cachet data model
type Incident struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Visible int    `json"visible"`
	Notify  bool   `json:"notify"`

	ComponentID     int `json:"component_id"`
	ComponentStatus int `json:"component_status"`
}

// Send - Create or Update incident
func (incident *Incident) Send(cfg *CachetMonitor) error {
	switch incident.Status {
		case 1, 2, 3:
			// performance issue
			incident.ComponentStatus = 2

			compInfo := cfg.API.GetComponentData(incident.ComponentID)
			if compInfo.Status == 2 {
				// major outage
				incident.ComponentStatus = 4
			}
		case 4:
			// fixed
			incident.ComponentStatus = 1
	}
	return nil
}

// SetInvestigating sets status to Investigating
func (incident *Incident) SetInvestigating() {
	incident.Status = 1
}

// SetIdentified sets status to Identified
func (incident *Incident) SetIdentified() {
	incident.Status = 2
}

// SetWatching sets status to Watching
func (incident *Incident) SetWatching() {
	incident.Status = 3
}

// SetFixed sets status to Fixed
func (incident *Incident) SetFixed() {
	incident.Status = 4
}
