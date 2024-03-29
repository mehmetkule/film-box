package netflix

import (
	"os"
	"strconv"

	"github.com/mehmetkule/film-box/core"
	"github.com/mehmetkule/film-box/parser"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func NetFlixCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "netflix",
		Short: "run netflix",
		RunE:  netflix,
	}
	rootCmd.Flags().BoolP("five", "f", false, "Kategori sayısı 3 yerine => 5 kabul edilir.")
	rootCmd.Flags().IntP("category", "c", 1, "Netflix'den çekilecek Kategori Grubu: '-f'(5li grup) var ise 1 ile 3 yok is 1 ile 5 arasında")
	rootCmd.Flags().IntP("rowcount", "r", 5, "Kategori bazlı toplam çekilecek film sayısı. 1 ile 10 arasında")
	
	return rootCmd
}

func netflix(cmd *cobra.Command, args []string) error {

	isFive, err := cmd.Flags().GetBool("five")
	if err != nil {
		return err
	}
	category, err := cmd.Flags().GetInt("category")
	if err != nil {
		return err
	}

	rowcount, err := cmd.Flags().GetInt("rowcount")
	if err != nil {
		return err
	}

	getMovieCategorys(&category, rowcount, isFive)
	return nil
}

func getMovieCategorys(categoryID *int, count int, isFive bool) {
	moveData := parser.ParserWeb(*categoryID, count, isFive)

	categories := make([]string, 0)
	movieList := make([][]string, 0)

	keys := core.SortedKeys(moveData, isFive)

	for i := 0; i < count; i++ {
		movieLine := make([]string, 0)
		for _, category := range keys {
			categories = core.UniqueAppend(categories, category)
			for titleIndex, item := range moveData[category] {
				if i == titleIndex {
					movieLine = append(movieLine, item)
				} else if titleIndex > i {
					break
				}
			}
		}
		movieList = append(movieList, movieLine)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(categories)

	table.SetCaption(true, "Netflix'de Kategori Bazlı Top "+strconv.Itoa(count)+" Film\n ®wermidon")
	table.AppendBulk(movieList)
	if !isWindows() {
		if !isFive {
			table.SetHeaderColor(tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgMagentaColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgGreenColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgYellowColor})
		} else {
			table.SetHeaderColor(tablewriter.Colors{
				tablewriter.Bold, tablewriter.BgMagentaColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgGreenColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgYellowColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgBlueColor},
				tablewriter.Colors{
					tablewriter.Bold, tablewriter.BgRedColor})

		}
	}
	table.Render()
}

func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
