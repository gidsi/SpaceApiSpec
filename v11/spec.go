package spaceapi

type Contact struct {
	Email     string   `json:"email,omitempty"`
	Irc       string   `json:"irc,omitempty"`
	Jabber    string   `json:"jabber,omitempty"`
	Keymaster []string `json:"keymaster,omitempty"`
	Ml        string   `json:"ml,omitempty"`
	Phone     string   `json:"phone,omitempty"`
	Sip       string   `json:"sip,omitempty"`
	Twitter   string   `json:"twitter,omitempty"`
}

type Events struct {
	Extra string  `json:"extra,omitempty"`
	Name  string  `json:"name"`
	T     float64 `json:"t"`
	Type  string  `json:"type"`
}

type Icon struct {
	Closed string `json:"closed"`
	Open   string `json:"open"`
}

type SpaceAPI011 struct {
	Address    string   `json:"address,omitempty"`
	Api        string   `json:"api"`
	Cam        []string `json:"cam,omitempty"`
	Contact    *Contact `json:"contact,omitempty"`
	Events     []Events `json:"events,omitempty"`
	Icon       *Icon    `json:"icon"`
	Lastchange float64  `json:"lastchange,omitempty"`
	Lat        float64  `json:"lat,omitempty"`
	Logo       string   `json:"logo"`
	Lon        float64  `json:"lon,omitempty"`
	Open       bool     `json:"open"`
	Space      string   `json:"space"`
	Status     string   `json:"status,omitempty"`
	Stream     *Stream  `json:"stream,omitempty"`
	Url        string   `json:"url"`
}

type Stream struct {
}
