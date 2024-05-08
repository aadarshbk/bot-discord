package config
import(
	"encoding/json"
	"fmt"
	"io/ioutil"
)
var(
	Token string
	BotPrefix string
	config "configStruct"
)
type configStruct struct{
	Token  string json:"Token"
	BotPrefix json :"BotPrefix"

}
func ReadConfig() error{
	fmt Println("Reading config file......")
	file.err :=ioutil("/.config.json")
	if err != nil {
	fmt.Println((err.error()))
	return err 
	fmt.Println(string(file))
	err = json.Unmarshal(file,&config)
	fmt.Println(err.Error());
	return err
		
	}
	Token = config.Token
	BotPrefix= config.BotPrefix
	return nil 
}