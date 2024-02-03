package cli

import (
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/spf13/cobra"
)

func newAddWatchUrlCmd() *cobra.Command {

	var url string
	var isList bool

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a watchurl",
		Long:  `Add a watchurl`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewWatchUrlServiceClient(conn)

			in := &api.AddUrlRequest{
				Url:    url,
				IsList: isList,
			}

			ret, err := client.AddUrl(cmd.Context(), in)
			if err != nil {
				return err
			}

			drawWatchUrlList(cmd.OutOrStdout(), []*api.UrlResponse{ret})

			return nil
		},
	}

	cmd.Flags().StringVarP(&url, "url", "u", "", "url to watch")
	cmd.Flags().BoolVarP(&isList, "list", "l", false, "is list")
	cmd.MarkFlagRequired("url")

	return cmd
}

func newListWatchUrlCmd() *cobra.Command {

	var page int32
	var pageSize int32

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List watchurls",
		Long:  `List watchurls`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewWatchUrlServiceClient(conn)
			in := &api.GetUrlsRequest{
				Page:     page,
				PageSize: pageSize,
			}
			ret, err := client.GetUrls(cmd.Context(), in)
			if err != nil {
				return err
			}

			drawWatchUrlList(cmd.OutOrStdout(), ret.Urls)

			return nil
		},
	}

	cmd.Flags().Int32VarP(&page, "page", "p", 1, "page number")
	cmd.Flags().Int32VarP(&pageSize, "page-size", "s", 30, "page size")

	return cmd
}

func newSetWatchUrlStateCmd() *cobra.Command {

	var id string
	var isDisabled bool

	cmd := &cobra.Command{
		Use:   "set-state",
		Short: "Set watchurl state",
		Long:  `Set watchurl state`,
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewWatchUrlServiceClient(conn)
			in := &api.SetStateRequest{
				Id:         id,
				IsDisabled: isDisabled,
			}

			ret, err := client.SetState(cmd.Context(), in)
			if err != nil {
				return err
			}

			drawWatchUrlList(cmd.OutOrStdout(), []*api.UrlResponse{ret})

			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "id")
	cmd.Flags().BoolVarP(&isDisabled, "disabled", "d", false, "is disabled")
	cmd.MarkFlagRequired("id")

	return cmd
}
