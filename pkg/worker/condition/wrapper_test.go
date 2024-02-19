package condition

import (
	"context"
	"errors"
	"github.com/piotr-gladysz/estate-compare/pkg/worker/db/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"testing"
	"time"
)

var config = map[string]any{
	"test": "test",
}

func TestWrapper_CheckOffer(t *testing.T) {
	ctx := context.Background()
	wrapper, err := createExampleWrapper(ctx)
	if err != nil {
		t.Fatal("failed to create wrapper", err.Error())
	}

	defer wrapper.Close(ctx)
}

func TestNewWrapper(t *testing.T) {
	tests := []struct {
		path  string
		error error
	}{
		{
			path:  "../../../bin/plugins/condition/test/valid.wasm",
			error: nil,
		},
		{
			path:  "../../../bin/plugins/condition/test/invalid-export.wasm",
			error: FunctionNotFoundError,
		},
		{
			path:  "../../../bin/plugins/condition/test/invalid-input.wasm",
			error: InvalidFunctionDefinitionError,
		},
		{
			path:  "../../../bin/plugins/condition/test/invalid-output.wasm",
			error: InvalidFunctionDefinitionError,
		},
	}

	for _, test := range tests {
		file, err := os.Open(test.path)
		if err != nil {
			t.Error("failed to open file", err)
			continue
		}

		_, err = NewWrapper(context.Background(), file)

		if err != test.error {
			t.Error("unexpected error", err)
		}
	}

}

func TestWrapper_FullProcess(t *testing.T) {
	errorsArr := testFullProcess()

	if len(errorsArr) > 0 {
		for _, err := range errorsArr {
			t.Error(err)
		}
	}

}

func BenchmarkWrapper_FullProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		errs := testFullProcess()
		if len(errs) > 0 {
			for _, err := range errs {
				b.Error(err)
			}
		}
	}
}

func BenchmarkWrapper_FullProcess_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			errs := testFullProcess()
			if len(errs) > 0 {
				for _, err := range errs {
					b.Error(err)
				}
			}
		}
	})
}

func BenchmarkWrapper_CheckOffer(b *testing.B) {
	ctx := context.Background()
	wrapper, err := createExampleWrapper(ctx)

	if err != nil {
		b.Fatal("failed to create wrapper", err)

	}

	defer wrapper.Close(ctx)

	t1, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "2021-01-02T00:00:00Z")

	testOffer := &model.Offer{
		Name: "Test",
		Url:  "http://test.com",
		History: []*model.OfferHistory{
			{
				Price:   100,
				Updated: primitive.NewDateTimeFromTime(t1),
			},
			{
				Price:   200,
				Updated: primitive.NewDateTimeFromTime(t2),
			},
		},
	}

	config := map[string]any{
		"test": "test",
	}

	for i := 0; i < b.N; i++ {
		_, err = wrapper.CheckOffer(ctx, testOffer, model.OfferActionAdd, config)

		if err != nil {
			b.Error("failed to check offer", err)
		}
	}
}

func createTestOffer(t time.Time) *model.Offer {
	return &model.Offer{
		Name: "Test",
		Url:  "http://test.com",
		History: []*model.OfferHistory{
			{
				Price:   100,
				Updated: primitive.NewDateTimeFromTime(t),
			},
			{
				Price:   200,
				Updated: primitive.NewDateTimeFromTime(t.Add(24 * time.Hour)),
			},
		},
	}
}
func createExampleWrapper(ctx context.Context) (*Wrapper, error) {
	wasmFile, err := os.Open("../../../bin/plugins/condition/test/valid.wasm")

	if err != nil {
		return nil, err
	}

	wrapper, err := NewWrapper(ctx, wasmFile)

	if err != nil {
		return nil, err
	}

	return wrapper, nil
}

func testFullProcess() []error {
	ctx := context.Background()
	wrapper, err := createExampleWrapper(ctx)

	if err != nil {
		return []error{errors.Join(errors.New("failed to create wrapper"), err)}

	}

	defer wrapper.Close(ctx)

	ti, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")

	testOffer := createTestOffer(ti)

	ret, err := wrapper.CheckOffer(ctx, testOffer, model.OfferActionAdd, config)

	if err != nil {
		return []error{errors.Join(errors.New("failed to check offer"), err)}
	}

	if ret == nil || ret.Message == "" {
		return []error{errors.New("empty message")}
	}

	return nil
}
