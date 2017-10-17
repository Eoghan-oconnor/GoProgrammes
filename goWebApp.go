//Eoghan O'Connor 
//adapted from https://github.com/SimonWaldherr/golang-examples/blob/master/expert/cookies.go


package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"io"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	
	c, err := r.Cookie("count")
	cookie := &http.Cookie{
		Name:  "count",
		Value: "1",
	}
	http.SetCookie(w, cookie)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, "Hello, you're here for the first time")
	} else {
		timeint, _ := strconv.ParseInt(c.Value, 10, 0)
		fmt.Fprintf(w, "Hi, your last visit was at "+time.Unix(timeint, 0).Format("15:04:05"))
		//fmt.Fprintf(w, "Hi, you've been here"+time.Unix(timeint,0))




	}

	io.WriteString(w,"\nHello World!")
}


func main(){
	http.HandleFunc("/",rootHandler)
	http.ListenAndServe(":8080", nil)
}