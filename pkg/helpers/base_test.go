package helpers

import "testing"

func TestCheckValidUuid(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "valid_uuid",
			args:    args{id: "05c87fd9-0c11-4992-9a58-c8ff1531ef9e"},
			wantErr: false,
		},
		{
			name:    "invalid_uuid",
			args:    args{id: "05c87fd9-0c11-4992-9a58-c8ff1531ef9e11"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckValidUuid(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CheckValidUuid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
