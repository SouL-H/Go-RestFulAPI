package dataloaders

import(
	"encoding/json"
	. "api/models" // import your models
	util "api/utils"  
)

func LoadUsers() []User{
	bytes,_ := util.ReadFile("../json/users.json")
	var data []User
	json.Unmarshal([]byte(bytes), &data)
	return data 
}

func LoadInterests() []Interest{
	bytes,_ := util.ReadFile("../json/interests.json")
	var data []Interest
	json.Unmarshal([]byte(bytes), &data)
	return data

}
func LoadInterestMappings() []InterestMapping{
	bytes,_ := util.ReadFile("../json/interestMappings.json")
	var data []InterestMapping
	json.Unmarshal([]byte(bytes), &data)
	return data

}