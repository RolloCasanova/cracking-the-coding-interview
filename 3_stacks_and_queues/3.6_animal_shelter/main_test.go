package main

import (
	"reflect"
	"testing"
	"time"
)

var now time.Time = time.Now()

func TestAnimalShelter(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want *shelter
	}{
		{
			name: "empty input should return nil",
			args: args{[]string{}},
			want: nil,
		},
		{
			name: "input with one dog should return a shelter with one dog",
			args: args{[]string{"dog", "fido"}},
			want: &shelter{
				dogs: &linkedlist{
					value: animal{
						kind: "dog",
						name: "fido",
					},
					next: nil,
				},
				cats: nil,
			},
		},
		{
			name: "input with one cat should return a shelter with one cat",
			args: args{[]string{"cat", "fluffy"}},
			want: &shelter{
				dogs: nil,
				cats: &linkedlist{
					value: animal{
						kind: "cat",
						name: "fluffy",
					},
					next: nil,
				},
			},
		},
		{
			name: "input with alternating dogs and cats should return a shelter with alternating dogs and cats",
			args: args{[]string{"dog", "fido", "cat", "fluffy", "dog", "spot", "cat", "mittens"}},
			want: &shelter{
				dogs: &linkedlist{
					value: animal{
						kind: "dog",
						name: "fido",
					},
					next: &linkedlist{
						value: animal{
							kind: "dog",
							name: "spot",
						},
						next: nil,
					},
				},
				cats: &linkedlist{
					value: animal{
						kind: "cat",
						name: "fluffy",
					},
					next: &linkedlist{
						value: animal{
							kind: "cat",
							name: "mittens",
						},
						next: nil,
					},
				},
			},
		},
		{
			name: "invalid animal type should return nil",
			args: args{[]string{"invalid", "fido"}},
			want: nil,
		},
		{
			name: "odd number of elements in input should return nil",
			args: args{[]string{"dog", "fido", "cat"}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// validate the values of the linked lists, not the actual linked lists
			got := AnimalShelter(tt.args.data)
			if got == nil {
				if tt.want != nil {
					t.Errorf("AnimalShelter() = %v, want %v", got, tt.want)
				}
				return
			}

			for got.dogs != nil && tt.want.dogs != nil {
				if got.dogs.value.name != tt.want.dogs.value.name {
					t.Errorf("AnimalShelter() = %v, want %v", got, tt.want)
				}
				got.dogs = got.dogs.next
				tt.want.dogs = tt.want.dogs.next
			}

			for got.cats != nil && tt.want.cats != nil {
				if got.cats.value.name != tt.want.cats.value.name {
					t.Errorf("AnimalShelter() = %v, want %v", got, tt.want)
				}
				got.cats = got.cats.next
				tt.want.cats = tt.want.cats.next
			}

		})
	}
}

func Test_shelter_Enqueue(t *testing.T) {
	type fields struct {
		dogs *linkedlist
		cats *linkedlist
	}
	type args struct {
		animal animal
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "enqueue dog in a shelter with no dogs should add dog to dogs linked list",
			fields: fields{
				dogs: nil,
				cats: nil,
			},
			args: args{
				animal: animal{
					kind: "dog",
					name: "fido",
				},
			},
		},
		{
			name: "enqueue cat in a shelter with no cats should add cat to cats linked list",
			fields: fields{
				dogs: nil,
				cats: nil,
			},
			args: args{
				animal: animal{
					kind: "cat",
					name: "fluffy",
				},
			},
		},
		{
			name: "enqueue dog should add dog to dogs linked list",
			fields: fields{
				dogs: &linkedlist{
					value: animal{
						kind: "dog",
						name: "fido",
					},
					next: &linkedlist{
						value: animal{
							kind: "dog",
							name: "spot",
						},
					},
				},
				cats: nil,
			},
			args: args{
				animal: animal{
					kind: "dog",
					name: "rover",
				},
			},
		},
		{
			name: "enqueue cat should add cat to cats linked list",
			fields: fields{
				dogs: nil,
				cats: &linkedlist{
					value: animal{
						kind: "cat",
						name: "fluffy",
					},
					next: &linkedlist{
						value: animal{
							kind: "cat",
							name: "mittens",
						},
					},
				},
			},
			args: args{
				animal: animal{
					kind: "cat",
					name: "whiskers",
				},
			},
		},
		{
			name: "enqueueing an animal with an invalid type should not add it to the shelter",
			fields: fields{
				dogs: nil,
				cats: nil,
			},
			args: args{
				animal: animal{
					kind: "invalid",
					name: "invalid",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &shelter{
				dogs: tt.fields.dogs,
				cats: tt.fields.cats,
			}
			s.Enqueue(tt.args.animal)
		})
	}
}

func Test_shelter_DequeueAny(t *testing.T) {
	type fields struct {
		dogs *linkedlist
		cats *linkedlist
	}
	tests := []struct {
		name       string
		fields     fields
		wantAnimal animal
	}{
		{
			name: "dequeueing from a shelter with no animals should not fail",
			fields: fields{
				dogs: nil,
				cats: nil,
			},
			wantAnimal: animal{},
		},
		{
			name: "dequeueing from a shelter with one dog should return the dog",
			fields: fields{
				dogs: &linkedlist{
					value: animal{
						kind: "dog",
						name: "fido",
					},
					next: nil,
				},
				cats: nil,
			},
			wantAnimal: animal{
				kind: "dog",
				name: "fido",
			},
		},
		{
			name: "dequeueing from a shelter with one cat should return the cat",
			fields: fields{
				dogs: nil,
				cats: &linkedlist{
					value: animal{
						kind: "cat",
						name: "fluffy",
					},
					next: nil,
				},
			},
			wantAnimal: animal{
				kind: "cat",
				name: "fluffy",
			},
		},
		{
			name: "dequeueing from a shelter with alternating dogs and cats should return the animal with the oldest timestamp - dog",
			fields: fields{
				dogs: &linkedlist{
					value: animal{
						kind:      "dog",
						name:      "fido",
						timestamp: now.Add(-2 * time.Second),
					},
					next: nil,
				},
				cats: &linkedlist{
					value: animal{
						kind:      "cat",
						name:      "fluffy",
						timestamp: now.Add(-1 * time.Second),
					},
					next: nil,
				},
			},
			wantAnimal: animal{
				kind:      "dog",
				name:      "fido",
				timestamp: now.Add(-2 * time.Second),
			},
		},
		{
			name: "dequeueing from a shelter with alternating dogs and cats should return the animal with the oldest timestamp - cat",
			fields: fields{
				dogs: &linkedlist{
					value: animal{
						kind:      "dog",
						name:      "fido",
						timestamp: now.Add(-1 * time.Second),
					},
					next: nil,
				},
				cats: &linkedlist{
					value: animal{
						kind:      "cat",
						name:      "fluffy",
						timestamp: now.Add(-2 * time.Second),
					},
					next: nil,
				},
			},
			wantAnimal: animal{
				kind:      "cat",
				name:      "fluffy",
				timestamp: now.Add(-2 * time.Second),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shelter{
				dogs: tt.fields.dogs,
				cats: tt.fields.cats,
			}
			if gotAnimal := s.DequeueAny(); !reflect.DeepEqual(gotAnimal, tt.wantAnimal) {
				t.Errorf("shelter.DequeueAny() = %v, want %v", gotAnimal, tt.wantAnimal)
			}
		})
	}
}

func Test_shelter_DequeueDog(t *testing.T) {
	type fields struct {
		dogs *linkedlist
		cats *linkedlist
	}
	tests := []struct {
		name       string
		fields     fields
		wantAnimal animal
	}{
		{
			name: "dequeueing from a shelter with no dogs should not fail",
			fields: fields{
				dogs: nil,
				cats: nil,
			},
			wantAnimal: animal{},
		},
		{
			name: "dequeueing from a shelter with one dog should return the dog",
			fields: fields{
				dogs: &linkedlist{
					value: animal{
						kind: "dog",
						name: "fido",
					},
					next: nil,
				},
			},
			wantAnimal: animal{
				kind: "dog",
				name: "fido",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shelter{
				dogs: tt.fields.dogs,
				cats: tt.fields.cats,
			}
			if gotAnimal := s.DequeueDog(); !reflect.DeepEqual(gotAnimal, tt.wantAnimal) {
				t.Errorf("shelter.DequeueDog() = %v, want %v", gotAnimal, tt.wantAnimal)
			}
		})
	}
}

func Test_shelter_DequeueCat(t *testing.T) {
	type fields struct {
		dogs *linkedlist
		cats *linkedlist
	}
	tests := []struct {
		name       string
		fields     fields
		wantAnimal animal
	}{
		{
			name: "dequeueing from a shelter with no cats should not fail",
			fields: fields{
				dogs: nil,
				cats: nil,
			},
			wantAnimal: animal{},
		},
		{
			name: "dequeueing from a shelter with one cat should return the cat",
			fields: fields{
				cats: &linkedlist{
					value: animal{
						kind: "cat",
						name: "fluffy",
					},
					next: nil,
				},
			},
			wantAnimal: animal{
				kind: "cat",
				name: "fluffy",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shelter{
				dogs: tt.fields.dogs,
				cats: tt.fields.cats,
			}
			if gotAnimal := s.DequeueCat(); !reflect.DeepEqual(gotAnimal, tt.wantAnimal) {
				t.Errorf("shelter.DequeueCat() = %v, want %v", gotAnimal, tt.wantAnimal)
			}
		})
	}
}

func Test_shelter_Print(t *testing.T) {
	type fields struct {
		dogs *linkedlist
		cats *linkedlist
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "print empty shelter should not fail",
			fields: fields{
				dogs: nil,
				cats: nil,
			},
		},
		{
			name: "print shelter with one dog should print one dog",
			fields: fields{
				dogs: &linkedlist{
					value: animal{
						kind: "dog",
						name: "fido",
					},
				},
			},
		},
		{
			name: "print shelter with one cat should print one cat",
			fields: fields{
				cats: &linkedlist{
					value: animal{
						kind: "cat",
						name: "fluffy",
					},
				},
			},
		},
		{
			name: "print shelter with alternating dogs and cats should print them by timestamp",
			fields: fields{
				dogs: &linkedlist{
					value: animal{
						kind:      "dog",
						name:      "fido",
						timestamp: now.Add(-4 * time.Second),
					},
					next: &linkedlist{
						value: animal{
							kind:      "dog",
							name:      "spot",
							timestamp: now.Add(-2 * time.Second),
						},
					},
				},
				cats: &linkedlist{
					value: animal{
						kind:      "cat",
						name:      "fluffy",
						timestamp: now.Add(-3 * time.Second),
					},
					next: &linkedlist{
						value: animal{
							kind:      "cat",
							name:      "mittens",
							timestamp: now.Add(-1 * time.Second),
						},
						next: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			s := &shelter{
				dogs: tt.fields.dogs,
				cats: tt.fields.cats,
			}
			s.Print()
		})
	}
}
