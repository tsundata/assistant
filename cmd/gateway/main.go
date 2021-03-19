package main

import "flag"

var configFile = flag.String("f", "gateway.yml", "set config file which will loading")

func main() {
	//rollbarOptions, err := rollbar.NewOptions(viper)
	//if err != nil {
	//	return nil, err
	//}
	//rollbar.Config(rollbarOptions)

	flag.Parse()

	a, err := CreateApp(*configFile)
	if err != nil {
		panic(err)
	}

	if err := a.Start(); err != nil {
		panic(err)
	}

	a.AwaitSignal()
}
