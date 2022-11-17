package testing_theme

import (
	"testing"
)

// TestSum - протестировали Sum()
func TestSum(t *testing.T) {
	tests := []struct {
		name       string
		values     []int
		wantResult int
	}{
		{
			name:       "simple test",
			values:     []int{1, 2, 3},
			wantResult: 6,
		},
		{
			name:       "a lot of number test",
			values:     []int{1, 2, 3, 78, 0},
			wantResult: 84,
		},
		{
			name:       "negative values",
			values:     []int{-3, 2, 4},
			wantResult: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if sum := Sum(tt.values...); sum != tt.wantResult {
				t.Errorf("Add: %v, want: %v", sum, tt.wantResult)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name    string
		value   float64
		trueRes float64
	}{
		{
			name:    "test small negative number",
			value:   -0.00000000006,
			trueRes: 0.00000000006,
		},
		{
			name:    "test small positive number",
			value:   0.00000000006,
			trueRes: 0.00000000006,
		},
		{
			name:    "simple test №1",
			value:   3,
			trueRes: 3,
		},
		{
			name:    "simple test №2",
			value:   -3,
			trueRes: 3,
		},
		{
			name:    "simple test №3",
			value:   2.00000001,
			trueRes: 2.00000001,
		},
		{
			name:    "simple test №1",
			value:   -2.00000001,
			trueRes: 2.00000001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := Abs(tt.value); res != tt.trueRes {
				t.Errorf("Abs: %f, want: %f", res, tt.trueRes)
			}
		})
	}
}

func TestFullName(t *testing.T) {
	tests := []struct {
		testNema  string
		firstName string
		lastName  string
		fullName  string
	}{
		{
			testNema:  "test simple name",
			firstName: "Egor",
			lastName:  "Kozlov",
			fullName:  "Egor Kozlov",
		},
		{
			testNema:  "test short name",
			firstName: "o",
			lastName:  "k",
			fullName:  "o k",
		},
		{
			testNema:  "test long name",
			firstName: "Austerlic Augustin II Joseph",
			lastName:  "Genrivolts",
			fullName:  "Austerlic Augustin II Joseph Genrivolts",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testNema, func(t *testing.T) {
			u := User{FirstName: tt.firstName, LastName: tt.lastName}
			if res := u.FullName(); res != tt.fullName {
				t.Errorf("wnat: %s, had: %s", tt.fullName, res)
			}
		})

	}
}

func TestFamily_AddNew(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name           string
		existedMembers map[Relationship]Person
		newPerson      newPerson
		wantErr        bool
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Maria",
					LastName:  "Popova",
					Age:       36,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			wantErr: false,
		},
		{
			name: "catch error",
			existedMembers: map[Relationship]Person{
				Father: {
					FirstName: "Misha",
					LastName:  "Popov",
					Age:       42,
				},
			},
			newPerson: newPerson{
				r: Father,
				p: Person{
					FirstName: "Ken",
					LastName:  "Gymsohn",
					Age:       32,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Family{
				Members: tt.existedMembers,
			}
			if err := f.AddNew(tt.newPerson.r, tt.newPerson.p); (err != nil) != tt.wantErr {
				t.Errorf("AddNew() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
