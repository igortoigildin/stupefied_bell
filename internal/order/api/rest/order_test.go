package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	config "github.com/igortoigildin/stupefied_bell/config/order"
	"github.com/igortoigildin/stupefied_bell/internal/order/api/rest/mocks"
	"github.com/igortoigildin/stupefied_bell/internal/order/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_addOrderHandler(t *testing.T) {
	tests := []struct {
		name           string
		order          model.Order
		respStatusCode int
	}{
		{
			name: "Success",
			order: model.Order{
				Number:   "111",
				Quantity: 5,
				Title:    "First order",
			},
			respStatusCode: http.StatusOK,
		},
		{
			name:           "Emty body",
			order:          model.Order{},
			respStatusCode: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var cfg config.Config
			rep := mocks.NewOrderRepository(t)
			if tt.respStatusCode == http.StatusOK {
				rep.On("SaveOrder", mock.Anything, mock.Anything).Return("111", nil).Once()
			}
			handler := addOrderHandler(&cfg, rep)

			js, err := json.Marshal(tt.order)
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/api/order", bytes.NewReader(js))
			require.NoError(t, err)

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			require.Equal(t, tt.respStatusCode, rr.Code)
		})
	}
}
