package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

// Executable mongo shell
var mongo = "mongo"

// Options
var nonVerbose = "--quiet"
var evaluate = "--eval"

// Mongo commands
var showDbs = "db.adminCommand('listDatabases')"
var showCollections = "db.getCollectionNames()"

type CmdManager struct {
}

func (c *CmdManager) GetDbs() interface{} {
	var dbs interface{}

	cmd := exec.Command(mongo, nonVerbose, evaluate, showDbs)
	output, _ := cmd.Output()
	err := json.Unmarshal(output, &dbs)
	if err != nil {
		fmt.Println(err)
	}

	return dbs
}

func (c *CmdManager) GetCollections(db string) (result []string, err error) {
	var collections []string

	cmd := exec.Command(mongo, db, nonVerbose, evaluate, showCollections)
	output, _ := cmd.Output()
	decodingErr := json.Unmarshal(output, &collections)
	if decodingErr != nil {
		fmt.Println(err)
	}
	fmt.Println(collections)
	if len(output) == 0 {
		return nil, errors.New("There are no collections under db: " + db)
	}

	return collections, nil
}
