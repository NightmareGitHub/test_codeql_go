package main


import(
    "fmt"
    "os/exec"
)

func getUserShell()(ret string){
	var shell string
	_,_ = fmt.Scanln(&shell)
	return shell
}


func main(){   
	User := getUserShell()
    c := exec.Command("cmd", "/C", User)
    if err := c.Run(); err != nil {
        fmt.Println("Error: ", err)
    }  
}
