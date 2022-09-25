package calc

import "testing"

func TestCalc(t *testing.T) {
	type args struct {
		a        int
		b        int
		operator string
	}
	type Tests struct {
		name   string
		args   args
		want   int
		hasErr bool
	}
	tests := []Tests{}
	test_1 := Tests{
		name: "足し算",
		args: args{
			a:        10,
			b:        2,
			operator: "+",
		},
		want:   12,
		hasErr: false,
	}
	test_2 := Tests{
		name: "不正な演算子を指定",
		args: args{
			a:        10,
			b:        2,
			operator: "?",
		},
		hasErr: true,
	}
	tests = append(tests, test_1)
	tests = append(tests, test_2)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			if (err != nil) != tt.hasErr {
				t.Errorf("Calc() error = %v, hasErr %v", err, tt.hasErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
