package ipfs

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/gorilla/mux"
	"github.com/scaleway/cq-source-scaleway/client"
	ipfs "github.com/scaleway/scaleway-sdk-go/api/ipfs/v1alpha1"
)

func createVolumes(router *mux.Router) error {
	var obj ipfs.Volume
	if err := faker.FakeObject(&obj); err != nil {
		return err
	}
	obj.Region = "fr-par"

	router.HandleFunc("/ipfs/v1alpha1/regions/fr-par/volumes", func(w http.ResponseWriter, r *http.Request) {
		list := ipfs.ListVolumesResponse{
			Volumes:    []*ipfs.Volume{&obj},
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

	var pin ipfs.Pin
	if err := faker.FakeObject(&pin); err != nil {
		return err
	}

	router.HandleFunc("/ipfs/v1alpha1/regions/fr-par/pins", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "unable to parse form: "+err.Error(), http.StatusBadRequest)
			return
		}
		if v := r.Form.Get("volume_id"); v != obj.ID {
			http.Error(w, "invalid volume_id: "+v, http.StatusBadRequest)
			return
		}

		list := ipfs.ListPinsResponse{
			Pins:       []*ipfs.Pin{&pin},
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

func TestVolumes(t *testing.T) {
	client.TestHelper(t, Volumes(), createVolumes)
}
