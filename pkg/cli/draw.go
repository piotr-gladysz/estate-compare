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

func drawOfferList(writer io.Writer, offers []*api.OfferResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)
	t.AppendHeader(table.Row{"#", "Id", "SiteId", "Site", "Created", "Updated", "Name", "Url", "Area", "Rooms",
		"Floor", "BuildingFloors", "Year", "Heating", "Market", "Window", "Elevator", "Balcony", "Media"})

	for i, offer := range offers {
		t.AppendRow(
			table.Row{
				i,
				offer.Id,
				offer.SiteId,
				offer.Site,
				time.Unix(offer.Created, 0).Format(time.RFC3339),
				time.Unix(offer.Updated, 0).Format(time.RFC3339),
				offer.Name,
				offer.Url,
				offer.Area,
				offer.Rooms,
				offer.Floor,
				offer.BuildingFloors,
				offer.Year,
				offer.Heating,
				offer.Market,
				offer.Window,
				offer.Elevator,
				offer.Balcony,
				offer.Media,
			})
	}

	t.Render()
}

func drawOfferHistory(writer io.Writer, histories []*api.OfferHistory) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)
	t.AppendHeader(table.Row{"#", "Updated", "Price"})

	for i, history := range histories {
		t.AppendRow(
			table.Row{
				i,
				time.Unix(history.Updated, 0).Format(time.RFC3339),
				history.Price,
			})
	}

	t.Render()
}

func drawProcessorStatus(writer io.Writer, state *api.ProcessingStatus) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)

	t.AppendHeader(table.Row{"Param", "Value"})

	t.AppendRow(table.Row{"IsRunning", state.IsRunning})
	t.AppendRow(table.Row{"LastRun", state.LastRun.AsTime().Format(time.RFC3339)})
	t.AppendRow(table.Row{"NextRun", state.NextRun.AsTime().Format(time.RFC3339)})

	t.Render()
}

func drawConditionList(writer io.Writer, conditions []*api.ConditionResponse) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)
	t.AppendHeader(table.Row{"#", "Id", "Name", "Created", "Updated"})

	for i, condition := range conditions {
		t.AppendRow(
			table.Row{
				i,
				condition.Id,
				condition.Name,
				time.Unix(condition.Created, 0).Format(time.RFC3339),
				time.Unix(condition.Updated, 0).Format(time.RFC3339),
			})
	}

	t.Render()
}
