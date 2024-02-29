package models

import (
	"testing"
	"time"
)

func TestProduct_CheckId(t *testing.T) {
	type fields struct {
		Id        string
		Name      string
		Price     uint64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid_uuid",
			fields: fields{
				Id: "05c87fd9-0c11-4992-9a58-c8ff1531ef9e",
			},
			wantErr: false,
		},
		{
			name: "invalid_uuid",
			fields: fields{
				Id: "05c87fd9-0c11-4992-9a58-c8ff1531ef9e434",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				Id:        tt.fields.Id,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := p.CheckId(); (err != nil) != tt.wantErr {
				t.Errorf("CheckId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_CheckMaxPrice(t *testing.T) {
	type fields struct {
		Id        string
		Name      string
		Price     uint64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "valid_price",
			fields: fields{
				Price: 10,
			},
			wantErr: false,
		},
		{
			name: "invalid_price",
			fields: fields{
				Price: 100000000,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				Id:        tt.fields.Id,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := p.CheckMaxPrice(); (err != nil) != tt.wantErr {
				t.Errorf("CheckMaxPrice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProduct_CheckValidName(t *testing.T) {
	type fields struct {
		Id        string
		Name      string
		Price     uint64
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "invalid_name",
			fields: fields{
				Name: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets ",
			},
			wantErr: true,
		},
		{
			name: "valid_name",
			fields: fields{
				Name: "lorem lorem lorem lorem",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{
				Id:        tt.fields.Id,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := p.CheckValidName(); (err != nil) != tt.wantErr {
				t.Errorf("CheckValidName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
