package main

import (
	"fmt"
"github.com/to4to/Go-Discord-Bot1/bot"
"github.com/to4to/Go-Discord-Bot1/config"
	"honnef.co/go/tools/config"
)




func main(){
	err:=config.ReadConfig();

	if err!=nil{

		fmt.Println(err.Error())
		return
	}


	bot.Start()

<-make(chan stuct{})
return


}