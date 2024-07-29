package main

import (
	//"fmt"
	"fmt"
	"groupie-tracker/groupietrackerdata"
	"html/template"
	"net/http"
	"strconv"
)

type Data struct {
	Artists []groupietrackerdata.Artists 
}

type Relation struct {
	Id 		int

}
type ProfileInfo struct{
	Image         string   
    Name          string   
    Members       []string 
    CreationDate  int      
    FirstAlbum    string
}

func HandleProfile(w http.ResponseWriter, r *http.Request){
	temp,err := template.ParseFiles("./Pages/Profile.html")
	if err != nil {
		http.Error(w,"this page not found",http.StatusNotFound)
		return 
	}
	if r.ParseForm() != nil {
		return
	}
	artistID,convertErr := strconv.Atoi(r.FormValue("id"))
	if convertErr != nil {
		http.Error(w,"Bad Request 400 ",http.StatusBadRequest)
		return 
	}
	var profileinfo ProfileInfo
	artist := groupietrackerdata.GetArtist(artistID)
	profileinfo.Name = artist.Name
	profileinfo.Image = artist.Image
	profileinfo.CreationDate = artist.CreationDate
	profileinfo.FirstAlbum=artist.FirstAlbum
	profileinfo.Members=artist.Members
	temp.Execute(w,profileinfo)
}
func main(){
	fmt.Println(groupietrackerdata.GetArtistRelation(1))
	http.Handle("/styles/style.css", http.FileServer(http.Dir("./")))
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		temp,err := template.ParseFiles("./Pages/Artists.html")
		if err != nil {
			http.Error(w,"this page not found",http.StatusNotFound)
			return 
		}
		data := Data{
			Artists: groupietrackerdata.GetAllArtists(),
		}
		temp.Execute(w,data)
	})
	http.HandleFunc("/Profile",HandleProfile)
	http.ListenAndServe(":8080",nil)
}