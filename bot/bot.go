package bot 
import(
 "fmt "
"github.com/aadarsh/discord-ping/config"
"github.com/bwmarrin/discordgo"
) 
var BotID string
var goBot "discordgo.Session"
func Start(){
	goBot discordgo.New("Bot " + config.Token)
	if err! = nil {
		fmt.Println(err.error())
		error
	}
	u,err:=goBot.User("@me")
	if err! =nil{
	fmt.Println(err,Error())
	}
	BotID =u.ID 
	goBot.AddHandler(messageHandler)
	err = goBot.Open()
	if err! = nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running ...")
}
func messageHandler(s "discordgo.session, m"discordgo.MessageCreate){
	if m.Author.ID == BotID
	return
}
if m.Content == "ping"
 _,_=s.ChannelMessageSend(m.channelID,"pong")