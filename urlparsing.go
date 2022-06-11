package main

import (
	"fmt"
	"net"
	"net/url"
)

func urlParseTest() {
	s := "https://swyrik:pass123@www.google.com:5432/search?q=rrr&oq=rr&aqs=edge.0.69i59j69i57j69i59j0i131i433i512j0i433i512j0i131i433i512j0i433i512j0i512l2.3967j0j1&sourceid=chrome&ie=UTF-8#frg"
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, pexist := u.User.Password()
	if pexist {
		fmt.Printf("is password available : %v\n", pexist)
		fmt.Println(p)
	}

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["q"])

}
