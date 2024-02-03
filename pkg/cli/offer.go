package cli

import (
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/spf13/cobra"
)

func newGetOfferCmd() *cobra.Command {

	var id string

	cmd := &cobra.Command{
		Use:   "offer",
		Short: "Get offer",
		Long:  "Get offer",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewOfferServiceClient(conn)

			in := &api.GetOfferRequest{
				Id: id,
			}

			ret, err := client.GetOffer(cmd.Context(), in)

			if err != nil {
				return err
			}

			drawOfferList(cmd.OutOrStdout(), []*api.OfferResponse{ret})
			drawOfferHistory(cmd.OutOrStdout(), ret.History)

			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "id of the offer")
	cmd.MarkFlagRequired("id")

	return cmd
}

func newListOfferCmd() *cobra.Command {

	var page int32
	var pageSize int32

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List offers",
		Long:  "List offers",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewOfferServiceClient(conn)

			in := &api.GetOffersRequest{
				Page:     page,
				PageSize: pageSize,
			}

			ret, err := client.GetOffers(cmd.Context(), in)

			if err != nil {
				return err
			}

			drawOfferList(cmd.OutOrStdout(), ret.Offers)

			return nil
		},
	}

	cmd.Flags().Int32VarP(&page, "page", "p", 1, "page number")
	cmd.Flags().Int32VarP(&pageSize, "page-size", "s", 30, "page size")

	return cmd

}
