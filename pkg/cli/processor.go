package cli

import (
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newProcessorStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start processor",
		Long:  "Start processor",
		RunE: func(cmd *cobra.Command, args []string) error {

			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewProcessorServiceClient(conn)

			_, err = client.StartProcessing(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func newProcessorStopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stop",
		Short: "Stop processor",
		Long:  "Stop processor",
		RunE: func(cmd *cobra.Command, args []string) error {

			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewProcessorServiceClient(conn)

			_, err = client.StopProcessing(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func newProcessorStatusCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Processor status",
		Long:  "Processor status",
		RunE: func(cmd *cobra.Command, args []string) error {

			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewProcessorServiceClient(conn)

			ret, err := client.GetProcessingStatus(cmd.Context(), &emptypb.Empty{})
			if err != nil {
				return err
			}

			drawProcessorStatus(cmd.OutOrStdout(), ret)

			return nil
		},
	}

	return cmd
}
