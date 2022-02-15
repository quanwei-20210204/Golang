# Golang
Jianqiang inspect and pair programming with me.

currentSecond := time.Now().Second()

/*  for every second, it will print*/
	ticker := time.NewTicker(5 * time.Second)
	for x := range ticker.C {
		fmt.Println("current second is: ", x.Second())
}