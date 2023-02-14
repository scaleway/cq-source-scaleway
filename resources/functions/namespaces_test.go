package functions

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
	"github.com/scaleway/cq-source-scaleway/client"
	function "github.com/scaleway/scaleway-sdk-go/api/function/v1beta1"
)

func createNamespaces(router *mux.Router) error {
	var obj function.Namespace
	if err := faker.FakeObject(&obj); err != nil {
		return err
	}
	obj.Region = "fr-par"

	var tobj function.Token
	if err := faker.FakeObject(&tobj); err != nil {
		return err
	}
	tobj.NamespaceID = &obj.ID

	router.HandleFunc("/functions/v1beta1/regions/fr-par/namespaces", func(w http.ResponseWriter, r *http.Request) {
		list := function.ListNamespacesResponse{
			Namespaces: []*function.Namespace{&obj},
			TotalCount: 1,
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	router.HandleFunc("/functions/v1beta1/regions/fr-par/tokens", func(w http.ResponseWriter, r *http.Request) {
		list := function.ListTokensResponse{
			Tokens:     []*function.Token{&tobj},
			TotalCount: 1,
		}

		b, err := json.Marshal(&list)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	return nil
}

func TestNamespaces(t *testing.T) {
	client.TestHelper(t, Namespaces(), createNamespaces)
}
