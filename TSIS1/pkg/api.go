package api

type Response struct {
	Tennis_players []Tennis_player `json:"players"`
}

type Tennis_player struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender string `json:"gender"`
	Country        string    `json:"country"`
	Wins int `json:"wins"`

}

func prepareResponse() []Tennis_player {
	var tennis_players []Tennis_player
   
	var player Tennis_player
	player.FirstName = "Novak"
	player.LastName = "Djokovic"
	player.Gender = "Male"
	player.Country = "Serbia"
	player.Wins = 40
	tennis_players = append(tennis_players, player)
   
	player.FirstName = "Alexander"
	player.LastName = "Zverev"
	player.Gender = "Male"
	player.Country = "Germany"
	player.Wins = 21
	tennis_players = append(tennis_players, player)
   
	player.FirstName = "Elena"
	player.LastName = "Rybakina"
	player.Gender = "Female"
	player.Country = "Kazakhstan"
	player.Wins = 6
	tennis_players = append(tennis_players, player)
   
	player.FirstName = "Serena"
	player.LastName = "Williams"
	player.Gender = "Female"
	player.Country = "USA"
	player.Wins = 39
	tennis_players = append(tennis_players, player)
   
	player.FirstName = "Rafael"
	player.LastName = "Nadal"
	player.Gender = "Male"
	player.Country = "Spain"
	player.Wins = 92
	tennis_players = append(tennis_players, player)
   
	return tennis_players
   }