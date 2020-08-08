/*
#######
##                 __
##        ___ ___ / /
##       (_-<(_-</ _ \
##      /___/___/_//_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

// Package ssh AFAIRE.
//
package ssh

import (
	"github.com/arnumina/failure"
)

type (
	// Clients AFAIRE.
	Clients map[string]map[string]*Client
)

// NewClients AFAIRE.
func NewClients(cos []*ClientOptions) Clients {
	clients := Clients{}

	for _, co := range cos {
		_, ok := clients[co.Host]
		if !ok {
			clients[co.Host] = make(map[string]*Client)
		}

		clients[co.Host][co.Username] = co.NewClient()
	}

	return clients
}

// Connect AFAIRE.
func (c Clients) Connect(host, username string) (*Connection, error) {
	client, ok := c[host][username]
	if !ok {
		return nil,
			failure.New(nil).
				Set("server", host).
				Set("user", username).
				Msg("this SSH server or user does not exist") //////////////////////////////////////////////////////////
	}

	return client.Connect()
}

/*
######################################################################################################## @(°_°)@ #######
*/
