package cli

import (
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var host string

func rootFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&host, "host", "127.0.0.1:11080", "host to connect to")
}

func CreateCLICommand() *cobra.Command {
	root := &cobra.Command{}
	rootFlags(root)

	watchUrl := &cobra.Command{
		Use:   "watch-url",
		Long:  "Watch url commands",
		Short: "Watch url commands",
	}

	watchUrl.AddCommand(newAddWatchUrlCmd())
	watchUrl.AddCommand(newListWatchUrlCmd())
	watchUrl.AddCommand(newSetWatchUrlStateCmd())

	root.AddCommand(watchUrl)
	return root
}

func getHost() string {
	if host != "" {
		return host
	} else {
		return "127.0.0.1:11080"
	}
}

func getHostConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(getHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	return conn, nil
}