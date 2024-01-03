package cli

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"io"
	"time"
)

func drawWatchUrlList(writer io.Writer, urls []*api.UrlResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)
	t.AppendHeader(table.Row{"#", "Id", "Url", "Is List", "Is Disabled", "Created", "Updated"})

	for i, url := range urls {
		t.AppendRow(
			table.Row{
				i,
				url.Id,
				url.Url,
				url.IsList,
				url.IsDisabled,
				time.Unix(url.Created, 0).Format(time.RFC3339),
				time.Unix(url.Updated, 0).Format(time.RFC3339),
			})
	}

	t.Render()
}
