//go:build go1.7
// +build go1.7

package redis

import (
	"errors"
	"testing"
	"time"
)

func TestParseURL(t *testing.T) {
	cases := []struct {
		u    string
		addr string
		db   int
		tls  bool
		err  error
	}{
		{
			"redis-set://localhost:123/1",
			"localhost:123",
			1, false, nil,
		},
		{
			"redis-set://localhost:123",
			"localhost:123",
			0, false, nil,
		},
		{
			"redis-set://localhost/1",
			"localhost:6379",
			1, false, nil,
		},
		{
			"redis-set://12345",
			"12345:6379",
			0, false, nil,
		},
		{
			"rediss://localhost:123",
			"localhost:123",
			0, true, nil,
		},
		{
			"redis-set://localhost/?abc=123",
			"",
			0, false, errors.New("no options supported"),
		},
		{
			"http://google.com",
			"",
			0, false, errors.New("invalid redis-set URL scheme: http"),
		},
		{
			"redis-set://localhost/1/2/3/4",
			"",
			0, false, errors.New("invalid redis-set URL path: /1/2/3/4"),
		},
		{
			"12345",
			"",
			0, false, errors.New("invalid redis-set URL scheme: "),
		},
		{
			"redis-set://localhost/iamadatabase",
			"",
			0, false, errors.New(`invalid redis-set database number: "iamadatabase"`),
		},
	}

	for _, c := range cases {
		t.Run(c.u, func(t *testing.T) {
			o, err := ParseURL(c.u)
			if c.err == nil && err != nil {
				t.Fatalf("unexpected error: %q", err)
				return
			}
			if c.err != nil && err != nil {
				if c.err.Error() != err.Error() {
					t.Fatalf("got %q, expected %q", err, c.err)
				}
				return
			}
			if o.Addr != c.addr {
				t.Errorf("got %q, want %q", o.Addr, c.addr)
			}
			if o.DB != c.db {
				t.Errorf("got %q, expected %q", o.DB, c.db)
			}
			if c.tls && o.TLSConfig == nil {
				t.Errorf("got nil TLSConfig, expected a TLSConfig")
			}
		})
	}
}

// Test ReadTimeout option initialization, including special values -1 and 0.
// And also test behaviour of WriteTimeout option, when it is not explicitly set and use
// ReadTimeout value.
func TestReadTimeoutOptions(t *testing.T) {
	testDataInputOutputMap := map[time.Duration]time.Duration{
		-1: 0 * time.Second,
		0:  3 * time.Second,
		1:  1 * time.Nanosecond,
		3:  3 * time.Nanosecond,
	}

	for in, out := range testDataInputOutputMap {
		o := &Options{ReadTimeout: in}
		o.init()
		if o.ReadTimeout != out {
			t.Errorf("got %d instead of %d as ReadTimeout option", o.ReadTimeout, out)
		}

		if o.WriteTimeout != o.ReadTimeout {
			t.Errorf("got %d instead of %d as WriteTimeout option", o.WriteTimeout, o.ReadTimeout)
		}
	}
}
