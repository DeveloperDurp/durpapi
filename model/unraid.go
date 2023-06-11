package model

type PowerSupply struct {
	TwelveVoltLoad  int    `json:"12v_load"`
	TwelveVoltWatts int    `json:"12v_watts"`
	ThreeVoltLoad   int    `json:"3v_load"`
	ThreeVoltWatts  int    `json:"3v_watts"`
	FiveVoltLoad    int    `json:"5v_load"`
	FiveVoltWatts   int    `json:"5v_watts"`
	Capacity        string `json:"capacity"`
	Efficiency      int    `json:"efficiency"`
	FanRPM          int    `json:"fan_rpm"`
	Load            int    `json:"load"`
	PoweredOn       string `json:"poweredon"`
	PoweredOnRaw    string `json:"poweredon_raw"`
	Product         string `json:"product"`
	Temp1           int    `json:"temp1"`
	Temp2           int    `json:"temp2"`
	Uptime          string `json:"uptime"`
	UptimeRaw       string `json:"uptime_raw"`
	Vendor          string `json:"vendor"`
	Watts           int    `json:"watts"`
}
