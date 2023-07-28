package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	jsonFile, err := os.Open("region.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cont *Content
	err = json.Unmarshal(byteValue, &cont)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cont.Count)
	regions := cont.Regions
	no_leader_id := []int{}
	for _, region := range regions {
		if region.Leader.Id != 0 {
			continue
		} else {
			no_leader_id = append(no_leader_id, region.Id)
		}
	}
	fmt.Println(len(no_leader_id))
	shell, err := os.OpenFile("recreate.sh", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer shell.Close()
	for _, id := range no_leader_id {
		idstring := strconv.Itoa(id)
		s := "/tikv-ctl --data-dir /data --config=/tikv.toml recreate-region -p ${PD_ADDR} -r " + idstring + "\n"
		shell.WriteString(s)
	}
}

type Content struct {
	Count   int      `name:"count"`
	Regions []Region `name:"regions"`
}

type Region struct {
	Id     int    `name:"id"`
	Leader Leader `name:"leader"`
}

type Leader struct {
	Id      int    `name:"id"`
	StoreID int    `name:"store_id"`
	Role    string `name:"role_name"`
}
