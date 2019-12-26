package main

import (
	"encoding/json"
	"fmt"
	"log"
	"picollo/types"

	r "gopkg.in/rethinkdb/rethinkdb-go.v5"
)

func rdbConnect(url string) (session *r.Session, err error) {
	session, err = r.Connect(r.ConnectOpts{
		Address: url,
	})
	if err != nil {
		return nil, err
	}
	return session, nil
}

func main() {
	sess, err := rdbConnect("localhost:28015")
	if err != nil {
		log.Fatal(err)
	}
	insertRecord(sess)
	fetchOne(sess)
	fetchAllRecords(sess)

}

func fetchOne(session *r.Session) {
	cur, err := r.DB("picollo").Table("rup").Run(session)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close()

	row := new(types.Packet)
	err = cur.One(&row)
	if err != nil {
		log.Fatalln(err)
	}

	printObj(row)
}
func insertRecord(session *r.Session) string {
	data := types.Packet{}
	data.Kode = "fetch4"
	data.Nama = "ABCD Packet"
	data.HPS = "40000000"
	data.Instansi = "ABCDE"
	data.Tahap = "PQER"
	data.TahapSekarang = "NOW"

	result, err := r.DB("picollo").Table("rup").Insert(data).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println("*** Insert result: ***")
	printObj(result)

	return result.GeneratedKeys[0]
}

func fetchAllRecords(session *r.Session) {
	cur, err := r.DB("picollo").Table("rup").Run(session)
	if err != nil {
		log.Fatal(err)
	}
	// Read records into  slice
	var packets []types.Packet
	err2 := cur.All(&packets)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	printStr("*** Fetch all rows: ***")
	for _, p := range packets {
		printObj(p)
	}
	printStr("\n")
}

func printObj(v interface{}) {
	vBytes, _ := json.Marshal(v)
	fmt.Println(string(vBytes))
}
func printStr(v string) {
	fmt.Println(v)
}
