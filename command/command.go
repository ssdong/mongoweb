package command

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

var showDbs = "db.adminCommand('listDatabases')"
var showCollections = "db.getCollectionNames()"

type CmdManager struct {
}

func (c *CmdManager) GetDbs() {
	var dbs interface{}
	cmd := exec.Command("mongo", "--eval", "db.adminCommand('listDatabases')")
	output, _ := cmd.Output()
	err := json.Unmarshal(output, dbs)
	fmt.Println(dbs)
	fmt.Println(err)
}
