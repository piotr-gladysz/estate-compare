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

func createExampleWrapper(ctx context.Context) (*Wrapper, error) {
	wasmFile, err := os.Open("../../../bin/plugins/condition/example.wasm")

	if err != nil {
		return nil, err
	}

	wrapper, err := NewWrapper(ctx, wasmFile)

	if err != nil {
		return nil, err
	}

	return wrapper, nil
}

func testFullProcess(t *testing.T) []error {
	ctx := context.Background()
	wrapper, err := createExampleWrapper(ctx)

	if err != nil {
		return []error{errors.Join(errors.New("failed to create wrapper"), err)}

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

	_, err = wrapper.CheckOffer(ctx, testOffer, config)

	if err != nil {
		return []error{errors.Join(errors.New("failed to check offer"), err)}
	}

	//TODO: check ret

	return nil
}

func TestNewWrapper(t *testing.T) {
	ctx := context.Background()
	wrapper, err := createExampleWrapper(ctx)

	if err != nil {
		t.Error("failed to create wrapper", err.Error())
		return
	}

	defer wrapper.Close(ctx)
}

func TestWrapper_FullProcess(t *testing.T) {
	errorsArr := testFullProcess(t)

	if len(errorsArr) > 0 {
		for _, err := range errorsArr {
			t.Error(err)
		}
	}

}

func BenchmarkWrapper_FullProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testFullProcess(nil)
	}
}

func BenchmarkWrapper_FullProcess_Parallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testFullProcess(nil)
		}
	})
}

func BenchmarkWrapper_CheckOffer(b *testing.B) {
	ctx := context.Background()
	wrapper, err := createExampleWrapper(ctx)

	if err != nil {
		b.Error("failed to create wrapper", err)
		return

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
		_, err = wrapper.CheckOffer(ctx, testOffer, config)

		if err != nil {
			b.Error("failed to check offer", err)
		}
	}
}
