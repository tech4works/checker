package checker

import (
	"testing"
	"time"
)

type timeCase struct {
	name  string
	a     any
	b     any
	want  bool
	panic bool
}

func TestIsBeforeNow(t *testing.T) {
	now := time.Now()
	inThePast := now.Add(-time.Minute)
	inTheFuture := now.Add(time.Minute)

	tests := []baseCase{
		{
			name: "IsBeforeNow with time in the past",
			arg:  inThePast,
			want: true,
		},
		{
			name: "IsBeforeNow with time in the future",
			arg:  inTheFuture,
			want: false,
		},
		{
			name: "IsBeforeNow with current time",
			arg:  time.Now(),
			want: true,
		},
		{
			name: "IsBeforeNow with uint",
			arg:  uint(time.Now().UnixMilli()),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBeforeNow(tt.arg); got != tt.want {
				t.Errorf("IsBeforeNow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBeforeToday(t *testing.T) {
	tests := []baseCase{
		{
			name: "Test with time today",
			arg:  time.Now(),
			want: false,
		},
		{
			name: "Test with time tomorrow",
			arg:  time.Now().Add(24 * time.Hour),
			want: false,
		},
		{
			name: "Test with time yesterday",
			arg:  time.Now().Add(-24 * time.Hour),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBeforeToday(tt.arg); got != tt.want {
				t.Errorf("IsBeforeToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBeforeDate(t *testing.T) {
	cases := []timeCase{
		{
			name: "Normal Case",
			a:    "2022-01-01",
			b:    "2022-01-02",
			want: true,
		},
		{
			name: "Equal dates",
			a:    "2022-01-02",
			b:    "2022-01-02",
			want: false,
		},
		{
			name: "a is After b",
			a:    "2022-01-03",
			b:    "2022-01-02",
			want: false,
		},
		{
			name:  "Incorrect date format in a",
			a:     "202201-31",
			b:     "2022-12-31",
			panic: true,
		},
		{
			name:  "Incorrect date format in b",
			a:     "2022-12-31",
			b:     "202212-31",
			panic: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !c.panic {
					t.Errorf("Got panic when none was expected")
				} else if r == nil && c.panic {
					t.Errorf("Expected panic but got nothing")
				}
			}()
			if got := IsBeforeDate(c.a, c.b); got != c.want {
				t.Errorf("IsBeforeDate() = %v, want %v", got, c.want)
			}
		})
	}
}

func TestIsBefore(t *testing.T) {
	tests := []timeCase{
		{
			name: "Same time",
			a:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			b:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "A before B",
			a:    time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC),
			b:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: true,
		},
		{
			name: "A after B",
			a:    time.Date(2024, time.January, 1, 0, 0, 0, 1, time.UTC),
			b:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
		{
			name: "Non time type",
			a:    "2024-01-01T00:00:00Z",
			b:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsBefore(tc.a, tc.b); got != tc.want {
				t.Errorf("IsBefore() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsAfterNow(t *testing.T) {
	now := time.Now()

	testCases := []baseCase{
		{
			name: "TimeTypeExactlyNow",
			arg:  now,
			want: false,
		},
		{
			name: "TimeStringFormatRFC3339",
			arg:  now.Format(time.RFC3339),
			want: false,
		},
		{
			name: "TimeTypeAfterNow",
			arg:  now.Add(time.Minute * 1),
			want: true,
		},
		{
			name: "TimeStringFormatUnixDate",
			arg:  now.Add(time.Minute * 1).Format(time.UnixDate),
			want: true,
		},
		{
			name: "TimeTypeBeforeNow",
			arg:  now.Add(time.Minute * -1),
			want: false,
		},
		{
			name: "TimeStringFormatRubyDate",
			arg:  now.Add(time.Minute * -1).Format(time.RubyDate),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsAfterNow(tc.arg); got != tc.want {
				t.Errorf("IsAfterNow() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestIsAfterToday(t *testing.T) {
	testCases := []baseCase{
		{
			name: "future time in time.Time type",
			arg:  time.Now().Add(24 * time.Hour),
			want: true,
		},
		{
			name: "current time in time.Time type",
			arg:  time.Now(),
			want: false,
		},
		{
			name: "past time in time.Time type",
			arg:  time.Now().Add(-24 * time.Hour),
			want: false,
		},
		{
			name: "future time in string type",
			arg:  time.Now().Add(24 * time.Hour).Format(time.DateOnly),
			want: true,
		},
		{
			name: "current time in string type",
			arg:  time.Now().Format(time.DateTime),
			want: false,
		},
		{
			name: "past time in string type",
			arg:  time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
			want: false,
		},
		{
			name:  "invalid time string",
			arg:   "invalid string",
			panic: true,
		},
		{
			name: "future UNIX timestamp",
			arg:  time.Now().Add(24 * time.Hour).UnixMilli(),
			want: true,
		},
		{
			name: "current UNIX timestamp",
			arg:  time.Now().UnixMilli(),
			want: false,
		},
		{
			name: "past UNIX timestamp",
			arg:  time.Now().Add(-24 * time.Hour).UnixMilli(),
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if !tc.panic {
						t.Errorf("Unexpected panic for case: %s. Panic: %v", tc.name, r)
					}
				}
			}()

			if got := IsAfterToday(tc.arg); got != tc.want {
				t.Errorf("Unexpected result for case: %s. Expected %v, but got: %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestIsAfterDate(t *testing.T) {
	now := time.Now()

	cases := []timeCase{
		{
			name: "After",
			a:    now.AddDate(0, 0, 1),
			b:    now,
			want: true,
		},
		{
			name: "Before",
			a:    now.AddDate(0, 0, -1),
			b:    now,
			want: false,
		},
		{
			name: "Equal",
			a:    now,
			b:    now,
			want: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if res := IsAfterDate(tc.a, tc.b); res != tc.want {
				t.Errorf("IsAfterDate() = %v; want %v", res, tc.want)
			}
		})
	}
}

func TestIsAfter(t *testing.T) {
	now := time.Now()
	past := now.Add(-time.Hour)
	future := now.Add(time.Hour)

	testCases := []timeCase{
		{
			name: "AIsAfterB",
			a:    future,
			b:    past,
			want: true,
		},
		{
			name: "AIsBeforeB",
			a:    past,
			b:    future,
			want: false,
		},
		{
			name: "AIsSameAsB",
			a:    now,
			b:    now,
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := IsAfter(tc.a, tc.b); result != tc.want {
				t.Errorf("For %v and %v, expected %v, but got %v", tc.a, tc.b, tc.want, result)
			}
		})
	}
}

func TestIsToday(t *testing.T) {
	tests := []baseCase{
		{
			name: "TestToday",
			arg:  time.Now(),
			want: true,
		},
		{
			name: "TestTomorrow",
			arg:  time.Now().Add(24 * time.Hour),
			want: false,
		},
		{
			name: "TestYesterday",
			arg:  time.Now().Add(-24 * time.Hour),
			want: false,
		},
		{
			name:  "TestOtherType",
			arg:   "stringVal",
			panic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					if tt.panic {
						t.Logf("Caught an expected panic: %v", r)
					} else {
						t.Errorf("Caught an unexpected panic: %v", r)
					}
				} else if tt.panic {
					t.Error("Expected a panic, but did not get one")
				}
			}()

			if got := IsToday(tt.arg); got != tt.want {
				t.Errorf("IsToday() = %v, want %v", got, tt.want)
			}
		})
	}
}
