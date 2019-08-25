/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	prettytable "github.com/tatsushid/go-prettytable"
	"github.com/spf13/cobra"
)

type TableRow = map[string]string;
type DataTable = map[int]TableRow;

func formatOutputJson(rows DataTable) string {
	return "json";
}

func formatOutputTable(rows DataTable) string {
	tbl, err := prettytable.NewTable([]prettytable.Column{
		{Header: "COL1"},
		{Header: "COL2", MinWidth: 6},
		{Header: "COL3", AlignRight: true},
	}...)
	if err != nil {
		panic(err)
	}
	tbl.Separator = " | "

	for row := range rows {
		for cel := range row {
			tbl.AddRow(cel)
		}
	}

	tbl.Print()

	return "table";
}

func formatOutput(rows DataTable) string {
	output := "format: ";

	switch rootCmd.Flag("format").Value.String() {
	case "table": return formatOutputTable(rows);
	case "json": return formatOutputJson(rows);
	}

	return output;
}

func nodeList(cmd *cobra.Command, args []string) {
	rows := make(DataTable, 0);

	row := make(TableRow);
	row["name"] = "one";
	rows[0] = row;


	fmt.Println(formatOutput(rows))
}

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Node details",
}

var nodeListCmd = &cobra.Command {
	Use: "list",
	Run: nodeList,
}

func init() {
	nodeCmd.AddCommand(nodeListCmd)
	rootCmd.AddCommand(nodeCmd)
	rootCmd.PersistentFlags().StringP("format", "f", "table", "output format");

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
