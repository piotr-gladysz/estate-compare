package cli

import (
	"github.com/piotr-gladysz/estate-compare/pkg/api"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func newAddConditionCmd() *cobra.Command {
	var name string
	var path string

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a condition",
		Long:  "Add a condition",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			defer file.Close()

			wasm, err := io.ReadAll(file)
			if err != nil {
				return err
			}

			client := api.NewConditionServiceClient(conn)

			in := &api.AddConditionRequest{
				Name: name,
				Wasm: wasm,
			}

			cond, err := client.AddCondition(cmd.Context(), in)

			if err != nil {
				return err
			}

			drawConditionList(cmd.OutOrStdout(), []*api.ConditionResponse{cond})

			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "name of the condition")
	cmd.Flags().StringVarP(&path, "path", "p", "", "path to the wasm file")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("path")

	return cmd
}

func newGetConditionCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get condition",
		Long:  "Get condition",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewConditionServiceClient(conn)

			in := &api.GetConditionRequest{
				Id: id,
			}

			ret, err := client.GetCondition(cmd.Context(), in)

			if err != nil {
				return err
			}

			drawConditionList(cmd.OutOrStdout(), []*api.ConditionResponse{ret})

			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "id of the condition")
	cmd.MarkFlagRequired("name")

	return cmd
}

func newListConditionCmd() *cobra.Command {
	var page int32
	var pageSize int32

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List conditions",
		Long:  "List conditions",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewConditionServiceClient(conn)
			in := &api.GetConditionsRequest{
				Page:     page,
				PageSize: pageSize,
			}
			ret, err := client.GetConditions(cmd.Context(), in)
			if err != nil {
				return err
			}

			drawConditionList(cmd.OutOrStdout(), ret.Conditions)

			return nil
		},
	}

	cmd.Flags().Int32VarP(&page, "page", "p", 1, "page number")
	cmd.Flags().Int32VarP(&pageSize, "pageSize", "s", 10, "page size")

	return cmd
}

func newDeleteConditionCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete condition",
		Long:  "Delete condition",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := getHostConn()
			if err != nil {
				return err
			}
			client := api.NewConditionServiceClient(conn)
			in := &api.GetConditionRequest{
				Id: id,
			}
			_, err = client.DeleteCondition(cmd.Context(), in)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "id of the condition")

	cmd.MarkFlagRequired("id")

	return cmd
}
