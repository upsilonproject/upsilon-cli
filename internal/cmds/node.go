package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/upsilonproject/upsilon-cli/internal/output"
)

func nodeList(cmd *cobra.Command, args []string) {
	tbl := &output.DataTable{}
	tbl.Headers = []string {
		"identifier",
		"type",
		"version",
	}

	tbl.Rows = make(map[int]output.TableRow)

	row := make(output.TableRow)
	row["name"] = "one"
	row["version"] = "blat"
	tbl.Append(row)

	fmt.Println(output.Format(*tbl))
}

var NodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Node details",
}

var NodeListCmd = &cobra.Command{
	Use: "list",
	Run: nodeList,
}
