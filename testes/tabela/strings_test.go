package strings

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	t.Parallel()
	type args struct {
		str     string
		strPart string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Mass 1", args{"Cod3r é show", "Cod3r"}, 0},
		{"Mass 2", args{"", ""}, 0},
		{"Mass 3", args{"Opa", "opa"}, -1},
		{"Mass 4", args{"leonardo", "o"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strings.Index(tt.args.str, tt.args.strPart); got != tt.want {
				t.Errorf("strings.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste da aula do curso
// const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

// func TestIndex(t *testing.T) {
// 	t.Parallel()
// 	testes := []struct {
// 		texto    string
// 		parte    string
// 		esperado int
// 	}{
// 		{"Cod3r é show", "Cod3r", 0},
// 		{"", "", 0},
// 		{"Opa", "opa", -1},
// 		{"leonardo", "o", 2},
// 	}

// 	for _, teste := range testes {
// 		t.Logf("Massa: %v", teste)
// 		atual := strings.Index(teste.texto, teste.parte)

// 		if atual != teste.esperado {
// 			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
// 		}
// 	}
// }
