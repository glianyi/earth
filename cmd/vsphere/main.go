package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/vmware/govmomi"
	vsphere "github.com/vmware/govmomi"
	"github.com/vmware/govmomi/session"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)

func main_P() {
	username := "administrator@vsphere.local"
	// username := "administrator@192.168.90.236"
	password := "Dream001@wh"
	server := "192.168.90.236"

	rawURL := "https://" + url.QueryEscape(username) + ":" + url.QueryEscape(password) + "@" + server + "/sdk"
	fmt.Println(rawURL)
	u, err := url.Parse(rawURL)

	vClient, err := vsphere.NewClient(context.Background(), u, true)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(u.User)
	err = vClient.Login(context.Background(), u.User)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func main() {
	conn()
}

func conn() {
	username := "administrator@vsphere.local"
	// username := "administrator@192.168.90.236"
	password := "Dream001@wh"
	server := "192.168.90.236"

	rawURL := "https://" + url.QueryEscape(username) + ":" + url.QueryEscape(password) + "@" + server + "/sdk"
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println(err.Error())
	}

	var govmcli *govmomi.Client
	{
		insecure := true
		soapCli := soap.NewClient(u, insecure)

		vimCli, err := vim25.NewClient(context.Background(), soapCli)
		if err != nil {
			fmt.Println(err.Error())
		}
		govmcli = &govmomi.Client{
			Client:         vimCli,
			SessionManager: session.NewManager(vimCli),
		}
	}

	userinfo := url.UserPassword(username, password)

	err = govmcli.Login(context.Background(), userinfo)
	if err != nil {
		fmt.Println()
	}

	fmt.Println(govmcli.Client.IsVC())

	fmt.Println("login")
}
