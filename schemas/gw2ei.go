package schemas

type Gw2eiSchema struct {
	EliteInsightsVersion string   `json:"eliteInsightsVersion"`
	TriggerID            int      `json:"triggerID"`
	Players              []Player `json:"players"`
	RecordedBy           string   `json:"recordedBy"`
	RecordedAccountBy    string   `json:"recordedAccountBy"`
	TimeStart            string   `json:"timeStart"`
	TimeEnd              string   `json:"timeEnd"`
	Duration             string   `json:"duration"`
}

type Player struct {
	Account         string           `json:"account"`
	Name            string           `json:"name"`
	Profession      string           `json:"profession"`
	Support         []Support        `json:"support"`
	StatsTargets    [][]StatsTargets `json:"statsTargets"`
	ExtHealingStats ExtHealingStats  `json:"extHealingStats"`
}

type ExtHealingStats struct {
	OutGoingHealing []OutGoingHealing `json:"outgoingHealing"`
}

type OutGoingHealing struct {
	Healing int `json:"healing"`
}

type StatsTargets struct {
	Killed int `json:"killed"`
	Downed int `json:"downed"`
}

type Support struct {
	CondiCleanse int `json:"condiCleanse"`
	BoonStrips   int `json:"boonStrips"`
}
