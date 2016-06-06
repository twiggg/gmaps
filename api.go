package forgoogle

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"strings"
)

//google api server key 1 (created for google maps geocoding)
var serverApiKey1 string = "PUT_YOUR_GOOGLE_API_KEY_VALUE"
var baseHttp string = "http://maps.googleapis.com/maps/api/geocode/json?"
var baseHttps string = "https://maps.googleapis.com/maps/api/geocode/json?"

//endpoint
//http://maps.googleapis.com/maps/api/geocode/output?parameters or https
//for ex http://maps.googleapis.com/maps/api/geocode/json?address=35+rue+Victor+hugo,+95870+Bezons,+France&region=fr&components:=country:FR|administrative_area:ile-de-france or https
//output: json or xml
//parameters: address,key (+bounds, language, region, components)

type Address struct {
	Place string  `json:"place" datastore:"place"`
	Query string  `queryjson:"place" datastore:"query"`
	Addr  string  `json:"addr" datastore:"addr"`
	Lat   float64 `json:"lat" datastore:"lat"`
	Lng   float64 `json:"lng" datastore:"lng"`
}

func (a *Address) Construct(meth string, route string, locality string, administrative_area string, postal_code string, country string) {
	s0 := ""
	s1 := ""
	//base
	switch strings.TrimSpace(strings.ToLower(html.EscapeString(meth))) {
	case "http":
		s0 = fmt.Sprintf("%s%s", s, baseHttp)
	case "https":
		s0 = fmt.Sprintf("%s%s", s, baseHttps)
	default:
		s0 = fmt.Sprintf("%s%s", s, baseHttp)
	}
	firt := true
	//road
	r := strings.TrimSpace(strings.ToLower(html.EscapeString(route)))
	if len(r) > 0 {
		s1 = fmt.Sprintf("%saddress=%s", s1, r)
		firt = false
	}
	compo := ""
	//city
	c := strings.TrimSpace(strings.ToLower(html.EscapeString(locality)))
	if len(r) > 0 {
		if first {
			s1 = fmt.Sprintf("%s%s", s1, r)
			firt = false
		} else {
			s1 = fmt.Sprintf("%s&%s", s1, r)
		}
	}
	//departement
	c := strings.TrimSpace(strings.ToLower(html.EscapeString(administrative_area)))
	if len(r) > 0 {
		if first {
			s1 = fmt.Sprintf("%s%s", s1, r)
			firt = false
		} else {
			s1 = fmt.Sprintf("%s&%s", s1, r)
		}
	}
	//postcode
	p := strings.TrimSpace(strings.ToLower(html.EscapeString(postal_code)))
	if len(r) > 0 {
		if first {
			s1 = fmt.Sprintf("%s%s", s1, r)
			firt = false
		} else {
			s1 = fmt.Sprintf("%s&%s", s1, r)
		}
	}
	//country
	co := strings.TrimSpace(strings.ToLower(html.EscapeString(country)))
	if len(r) > 0 {
		if first {
			s1 = fmt.Sprintf("%s%s", s1, r)
			firt = false
		} else {
			s1 = fmt.Sprintf("%s&%s", s1, r)
		}
	}

	if len(s1) > 0 && len(co) > 0 {
		s1 = fmt.Sprintf("%s&components=region:%s", s1, co)

		a.Query = fmt.Sprintf("%s%s", s0, s1)
	} else {
		a.Query = ""
	}

}

func (a *Address) Geocode() (int32, error) {
	resp, err := http.Get(a.Query)
	//struc:=struct{}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK || err != nil {
		return 1, err
	} else {
		d := json.NewDecoder(resp.Body)
		if err := d.Decode(&n); err != nil {
			//return handler.StatusError{http.StatusInternalServerError, err}
			return 2, err
		}
		/*if err:=d.Decode(&struc);err!=nil{

																}*/
	}
}
