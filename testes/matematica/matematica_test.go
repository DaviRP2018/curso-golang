package matematica

import "testing"

// Teste gerado pelo Go
func TestMedia(t *testing.T) {
	t.Parallel()
	type args struct {
		numeros []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"teste 1", args{[]float64{7.2, 9.9, 6.1, 5.9}}, 7.28},
		{"teste 2", args{[]float64{7.2, 9.9, 6.1, 5.9}}, 7.29},
		{"teste 3", args{[]float64{7.2, 9.9, 6.1, 5.9}}, 7.30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Media(tt.args.numeros...); got != tt.want {
				t.Errorf("Media() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Teste da aula do curso
// const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// func TestMedia(t *testing.T) {
// 	valorEsperado := 7.28
// 	valor := Media(7.2, 9.9, 6.1, 5.9)

// 	if valor != valorEsperado {
// 		t.Errorf(erroPadrao, valorEsperado, valor)
// 	}
// }
