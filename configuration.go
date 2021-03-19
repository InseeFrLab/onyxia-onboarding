package main

type configuration struct {
	Authentication authentication
}

type authentication struct {
	BaseUrl string 
	Realm string
}