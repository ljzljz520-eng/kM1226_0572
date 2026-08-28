package report

import (
	"spacetrash/model"
	"spacetrash/service"
	"spacetrash/storage"
	"strings"
	"testing"
)

func TestRenderSummary(t *testing.T) {
	store, err := storage.Open(t.TempDir() + "/report.db")
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()
	player := model.NewPlayer("p", "Nova")
	vessel := model.NewVessel("v", "p", "Wayfarer")
	session, err := service.NewSession(store, player, vessel, []model.Weapon{model.NewWeapon("w", "v", "Pulse", 0)})
	if err != nil {
		t.Fatal(err)
	}
	text := RenderSummary(BuildSummary(session))
	if !strings.Contains(text, "Nova") || !strings.Contains(text, "Slot 0") {
		t.Fatalf("summary missing fields: %s", text)
	}
}
