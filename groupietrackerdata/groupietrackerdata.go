package groupietrackerdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Artists struct {
	ID            int      //`json:"id"`
    Image         string   //`json:"image"`
    Name          string   //`json:"name"`
    Members       []string //`json:"members"`
    CreationDate  int      //`json:"creationDate"`
    FirstAlbum    string   //`json:"firstAlbum"`
    Locations     string   `json:"locations"`
    ConcertDates  string   `json:"concertDates"`
    Relations     string   `json:"relations"`
}

type Relation struct{
	ID            		int      					`json:"id"`
    DatesLocations      map[string][]string   `json:"datesLocations"`
}

func GetArtistRelation(id int) map[string][]string{
	data,state := getData("https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id))
	if !state {
		return nil
	} else {
		var Data Relation
		json.Unmarshal(data,&Data)
		return Data.DatesLocations
	}
}
func GetAllArtists() []Artists{
	data,state := getData("https://groupietrackers.herokuapp.com/api/artists")
	if !state {
		return nil
	} else {
		var Data []Artists
		json.Unmarshal(data,&Data)
		return Data
	}
}

func GetArtist(id int) *Artists {
	data,state := getData("https://groupietrackers.herokuapp.com/api/artists/"+strconv.Itoa(id)+"")
	if !state {
		return nil
	} else {
		var Data Artists
		json.Unmarshal(data,&Data)
		return &Data
	}
}

/*func GetArtistLocations(id int) []string {
	data,state := getData("https://groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(id)+"")
	if !state {
		return nil
	} else {
		var Data ArtistsLocations
		json.Unmarshal(data,&Data)
		return Data.Locations
	}
}
func GetArtistLocationsDates(id int ) *Dates{
	data,state := getData("https://groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(id)+"")
	if !state {
		return nil
	} else {
		var Data Dates
		json.Unmarshal(data,&Data)
		return &Data
	}
}
*/
func GetArtistImage(name string ) []byte {
	data,state := getData("https://groupietrackers.herokuapp.com/api/images/"+name+".jpeg")
	if !state {
		fmt.Println("no data")
		return nil
	} else {
		return data
	}
}


func getData(url string) ([]byte,bool){
	data, err := http.Get(url)
	if err != nil {
		return nil,false
	}
	body , err := io.ReadAll(data.Body)
	if err != nil {
		return nil,false
	}
	return body, true
}