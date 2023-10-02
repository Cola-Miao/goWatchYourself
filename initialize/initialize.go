package initialize

func Initialize() (err error) {
	//if err = initLog(); err != nil {
	//	return
	//}
	initEngine()
	initClient()

	return
}
