package main

import "testing"

func Test_sortedPersonalDataRecordsBySurname_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    sortedPersonalDataRecordsBySurname
		args args
		want bool
	}{
		{
			name: "surname before",
			s:    sortedPersonalDataRecordsBySurname{{"A", "A"}, {"A", "B"}},
			args: args{0, 1},
			want: true,
		},
		{
			name: "surname after",
			s:    sortedPersonalDataRecordsBySurname{{"A", "B"}, {"A", "A"}},
			args: args{0, 1},
			want: false,
		},
		{
			name: "equal",
			s:    sortedPersonalDataRecordsBySurname{{"A", "A"}, {"A", "A"}},
			args: args{0, 1},
			want: false,
		},
		{
			name: "surname equal, first name before",
			s:    sortedPersonalDataRecordsBySurname{{"A", "A"}, {"B", "A"}},
			args: args{0, 1},
			want: true,
		},
		{
			name: "surname equal, first name after",
			s:    sortedPersonalDataRecordsBySurname{{"B", "A"}, {"A", "A"}},
			args: args{0, 1},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
