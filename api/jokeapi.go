package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

type Joke struct {
	Id   int    `json:Id`
	Joke string `json:Joke`
}

type User struct {
	Name    string `json:"name"`
	Anumber string `json:"anumber"`
}

var jokes []Joke

func JokeDAO() {

	jokes = append(jokes, Joke{1, "Chuck Norris uses ribbed condoms inside out, so he gets the pleasure."})
	jokes = append(jokes, Joke{2, "MacGyver can build an airplane out of gum and paper clips. Chuck Norris can kill him and take it."})
	jokes = append(jokes, Joke{3, "Chuck Norris doesn't read books. He stares them down until he gets the information he wants."})
	jokes = append(jokes, Joke{4, "Chuck Norris lost his virginity before his dad did."})
	jokes = append(jokes, Joke{5, "Since 1940, the year Chuck Norris was born, roundhouse kick related deaths have increased 13,000 percent."})
	jokes = append(jokes, Joke{6, "Chuck Norris sheds his skin twice a year."})
	jokes = append(jokes, Joke{7, "Chuck Norris once challenged Lance Armstrong in a &quot;Who has more testicles?&quot; contest. Chuck Norris won by 5."})
	jokes = append(jokes, Joke{8, "There are no steroids in baseball. Just players Chuck Norris has breathed on."})
	jokes = append(jokes, Joke{9, "When Chuck Norris goes to donate blood, he declines the syringe, and instead requests a hand gun and a bucket."})
}

func HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to Muzi's Chuck Norris Page!!...Have Fun")
}

func AllJokes(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(jokes)
}

func RandJoke(w http.ResponseWriter, req *http.Request) {
	ranNum := rand.Intn(9)
	fmt.Println(ranNum)
	for _, item := range jokes {
		if item.Id == ranNum {
			json.NewEncoder(w).Encode(item)
			return
		} // end if
	} // end for loop
} // end RandJOke

func NameWriter(w http.ResponseWriter, req *http.Request) {
	var user User
	if req.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprint(w, "Your name is: "+user.Name)
}
