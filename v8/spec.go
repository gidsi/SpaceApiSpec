// Code generated by schema-generate. DO NOT EDIT.

package spaceapistruct

import (
	"bytes"
	"encoding/json"
	"errors"
)

// EventsItems
type EventsItems struct {

	// A custom text field to give more information about the event
	Extra string `json:"extra,omitempty"`

	// Name or other identity of the subject (e.g. <samp>J. Random Hacker</samp>, <samp>fridge</samp>, <samp>3D printer</samp>, …)
	Name string `json:"name"`

	// Unix timestamp when the event occurred
	T float64 `json:"t"`

	// Action (e.g. <samp>check-in</samp>, <samp>check-out</samp>, <samp>finish-print</samp>, …). Define your own actions and use them consistently, canonical actions are not (yet) specified
	Type string `json:"type"`
}

// Root SpaceAPI 0.8
type Root struct {

	// The main website
	Address string `json:"address,omitempty"`

	// The version of SpaceAPI your endpoint uses
	Api string `json:"api"`

	// URL(s) of webcams in your space
	Cam []string `json:"cam,omitempty"`

	// Events which happened recently in your space and which could be interesting to the public, like 'User X has entered/triggered/did something at timestamp Z'
	Events []*EventsItems `json:"events,omitempty"`

	// The Unix timestamp when the space status changed most recently
	Lastchange float64 `json:"lastchange,omitempty"`

	// Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.
	Lat float64 `json:"lat,omitempty"`

	// The space logo
	Logo string `json:"logo"`

	// Latitude of your space location, in degree with decimal places. Use positive values for locations north of the equator, negative values for locations south of equator.
	Lon float64 `json:"lon,omitempty"`

	// A boolean which indicates if the space is currently open
	Open bool `json:"open"`

	// Phone number, including country code with a leading plus sign. Example: <samp>+1 800 555 4567</samp>
	Phone string `json:"phone,omitempty"`

	// The name of your space
	Space string `json:"space"`

	// An additional free-form string, could be something like <samp>'open for public'</samp>, <samp>'members only'</samp> or whatever you want it to be
	Status string `json:"status,omitempty"`

	// A mapping of stream types to stream URLs. Example: <samp>{'mp4':'http//example.org/stream.mpg', 'mjpeg':'http://example.org/stream.mjpeg'}</samp>
	Stream *Stream `json:"stream,omitempty"`
	Url    string  `json:"url"`
}

// Stream A mapping of stream types to stream URLs. Example: <samp>{'mp4':'http//example.org/stream.mpg', 'mjpeg':'http://example.org/stream.mjpeg'}</samp>
type Stream struct {
}

func (strct *EventsItems) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "extra" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"extra\": ")
	if tmp, err := json.Marshal(strct.Extra); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Name" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "name" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "T" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "t" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"t\": ")
	if tmp, err := json.Marshal(strct.T); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Type" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "type" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"type\": ")
	if tmp, err := json.Marshal(strct.Type); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *EventsItems) UnmarshalJSON(b []byte) error {
	nameReceived := false
	tReceived := false
	typeReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "extra":
			if err := json.Unmarshal([]byte(v), &strct.Extra); err != nil {
				return err
			}
		case "name":
			if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
				return err
			}
			nameReceived = true
		case "t":
			if err := json.Unmarshal([]byte(v), &strct.T); err != nil {
				return err
			}
			tReceived = true
		case "type":
			if err := json.Unmarshal([]byte(v), &strct.Type); err != nil {
				return err
			}
			typeReceived = true
		}
	}
	// check if name (a required property) was received
	if !nameReceived {
		return errors.New("\"name\" is required but was not present")
	}
	// check if t (a required property) was received
	if !tReceived {
		return errors.New("\"t\" is required but was not present")
	}
	// check if type (a required property) was received
	if !typeReceived {
		return errors.New("\"type\" is required but was not present")
	}
	return nil
}

func (strct *Root) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
	comma := false
	// Marshal the "address" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"address\": ")
	if tmp, err := json.Marshal(strct.Address); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Api" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "api" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"api\": ")
	if tmp, err := json.Marshal(strct.Api); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "cam" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"cam\": ")
	if tmp, err := json.Marshal(strct.Cam); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "events" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"events\": ")
	if tmp, err := json.Marshal(strct.Events); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "lastchange" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"lastchange\": ")
	if tmp, err := json.Marshal(strct.Lastchange); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "lat" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"lat\": ")
	if tmp, err := json.Marshal(strct.Lat); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Logo" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "logo" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"logo\": ")
	if tmp, err := json.Marshal(strct.Logo); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "lon" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"lon\": ")
	if tmp, err := json.Marshal(strct.Lon); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Open" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "open" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"open\": ")
	if tmp, err := json.Marshal(strct.Open); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "phone" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"phone\": ")
	if tmp, err := json.Marshal(strct.Phone); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Space" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "space" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"space\": ")
	if tmp, err := json.Marshal(strct.Space); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "status" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"status\": ")
	if tmp, err := json.Marshal(strct.Status); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// Marshal the "stream" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"stream\": ")
	if tmp, err := json.Marshal(strct.Stream); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true
	// "Url" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "url" field
	if comma {
		buf.WriteString(",")
	}
	buf.WriteString("\"url\": ")
	if tmp, err := json.Marshal(strct.Url); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Root) UnmarshalJSON(b []byte) error {
	apiReceived := false
	logoReceived := false
	openReceived := false
	spaceReceived := false
	urlReceived := false
	var jsonMap map[string]json.RawMessage
	if err := json.Unmarshal(b, &jsonMap); err != nil {
		return err
	}
	// parse all the defined properties
	for k, v := range jsonMap {
		switch k {
		case "address":
			if err := json.Unmarshal([]byte(v), &strct.Address); err != nil {
				return err
			}
		case "api":
			if err := json.Unmarshal([]byte(v), &strct.Api); err != nil {
				return err
			}
			apiReceived = true
		case "cam":
			if err := json.Unmarshal([]byte(v), &strct.Cam); err != nil {
				return err
			}
		case "events":
			if err := json.Unmarshal([]byte(v), &strct.Events); err != nil {
				return err
			}
		case "lastchange":
			if err := json.Unmarshal([]byte(v), &strct.Lastchange); err != nil {
				return err
			}
		case "lat":
			if err := json.Unmarshal([]byte(v), &strct.Lat); err != nil {
				return err
			}
		case "logo":
			if err := json.Unmarshal([]byte(v), &strct.Logo); err != nil {
				return err
			}
			logoReceived = true
		case "lon":
			if err := json.Unmarshal([]byte(v), &strct.Lon); err != nil {
				return err
			}
		case "open":
			if err := json.Unmarshal([]byte(v), &strct.Open); err != nil {
				return err
			}
			openReceived = true
		case "phone":
			if err := json.Unmarshal([]byte(v), &strct.Phone); err != nil {
				return err
			}
		case "space":
			if err := json.Unmarshal([]byte(v), &strct.Space); err != nil {
				return err
			}
			spaceReceived = true
		case "status":
			if err := json.Unmarshal([]byte(v), &strct.Status); err != nil {
				return err
			}
		case "stream":
			if err := json.Unmarshal([]byte(v), &strct.Stream); err != nil {
				return err
			}
		case "url":
			if err := json.Unmarshal([]byte(v), &strct.Url); err != nil {
				return err
			}
			urlReceived = true
		}
	}
	// check if api (a required property) was received
	if !apiReceived {
		return errors.New("\"api\" is required but was not present")
	}
	// check if logo (a required property) was received
	if !logoReceived {
		return errors.New("\"logo\" is required but was not present")
	}
	// check if open (a required property) was received
	if !openReceived {
		return errors.New("\"open\" is required but was not present")
	}
	// check if space (a required property) was received
	if !spaceReceived {
		return errors.New("\"space\" is required but was not present")
	}
	// check if url (a required property) was received
	if !urlReceived {
		return errors.New("\"url\" is required but was not present")
	}
	return nil
}
