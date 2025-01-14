package music

import (
    "log"

    "github.com/miekg/dns"
)

//
// TODO: See if there is a better ways to give the insert/remove RRset
//
// Current implementation mimics dns.Insert()/.Remove() in the way that each
// entry in the first array is a call to these functions with the second
// array.
//
type Updater interface {
    Update(signer *Signer, fqdn string, inserts, removes *[][]dns.RR) error
    RemoveRRset(signer *Signer, fqdn string, rrsets [][]dns.RR) error
}

var Updaters map[string]Updater = make(map[string]Updater)

func GetUpdater(type_ string) Updater {
    updater, ok := Updaters[type_]
    if !ok {
        log.Fatal("No updater type", type_)
    }
    return updater
}
