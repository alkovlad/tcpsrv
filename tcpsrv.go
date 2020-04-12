package main

import(
	 "github.com/firstrow/tcp_server"
	 "log"
	// "github.com/davecgh/go-spew/spew"
	 "strings"
	 "fmt"
	 //"strconv"
	 "net/http"
	// "net/url"
	 "io/ioutil"
)

func main() {
	server := tcp_server.New(":4096")

	server.OnNewClient(func(c *tcp_server.Client) {
		// new client connected
		// lets send some message
		//c.Send("Hello")
		log.Printf("new connect",c.Conn().RemoteAddr())
		//spew.Dump(c.Conn().raddr)
		//fmt.Printf(c)
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new message received
		fmt.Printf("%s", message)
		sdata := ""
		if(len(message)>5){
		inarray := strings.Split(strings.TrimSpace(message), ":")
		for i := range inarray {
			elm := strings.Split(strings.TrimSpace(inarray[i]), "=")
			//for j := range elm {
				sdata +="&"+strings.TrimSpace(elm[0])+"="+strings.TrimSpace(elm[1])
				//f/mt.Println(result[i])
			//}
		}
		fmt.Println("http://app.pitaya.io/iot/save.php?"+sdata)
		 u, err := http.Get("http://app.pitaya.io/iot/save.php?"+strings.TrimSpace(sdata))
		if err != nil {
			log.Fatal(err)
		}
		 defer u.Body.Close()
			body, err := ioutil.ReadAll(u.Body)
			if err != nil {
					panic(err)
			}
			fmt.Printf("%s", body)
		c.Close()
		}else{
			c.Send("Too small body\r\n")
			c.Close()
		}
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection with client lost
		log.Printf("close connect")
		//spew.Dump(c)
	})

	server.Listen()
}