package cell

import "testing"

func TestInitCell (t *testing.T) {
	wanted := false;
	state := InitCell(); // Bruker semantikken i Golangs konseptuelle modell
	if state != wanted {
		t.Errorf("Feil, fikk %q, ikke ønsket %q.", state, wanted)
	}
}