package main

import "fmt"

type LV252Trainee struct {
	Name string
}

func PleaseCheckAndApprovekMyPR(teamMates []LV252Trainee){
	teamMates = []LV252Trainee{{"Bohdan"}, {"Volodymyr"}, {"Yuriy"}}
	for range teamMates {
		fmt.Println("I need one more PR approve , please check and approve PR #38")
	}
}
