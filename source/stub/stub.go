package stub

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/IktaS/migrate/source"
)

func init() {
	source.Register("stub", &Stub{})
}

type Config struct{}

// d, _ := source.Open("stub://")
// d.(*stub.Stub).Migrations =

type Stub struct {
	Url        string
	Instance   interface{}
	Migrations *source.Migrations
	Config     *Config
}

func (s *Stub) Open(url string) (source.Driver, error) {
	return &Stub{
		Url:        url,
		Migrations: source.NewMigrations(),
		Config:     &Config{},
	}, nil
}

func WithInstance(instance interface{}, config *Config) (source.Driver, error) {
	return &Stub{
		Instance:   instance,
		Migrations: source.NewMigrations(),
		Config:     config,
	}, nil
}

func (s *Stub) Close() error {
	return nil
}

func (s *Stub) First() (version uint64, err error) {
	if v, ok := s.Migrations.First(); !ok {
		return 0, &os.PathError{Op: "first", Path: s.Url, Err: os.ErrNotExist} // TODO: s.Url can be empty when called with WithInstance
	} else {
		return v, nil
	}
}

func (s *Stub) Prev(version uint64) (prevVersion uint64, err error) {
	if v, ok := s.Migrations.Prev(version); !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("prev for version %v", version), Path: s.Url, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (s *Stub) Next(version uint64) (nextVersion uint64, err error) {
	if v, ok := s.Migrations.Next(version); !ok {
		return 0, &os.PathError{Op: fmt.Sprintf("next for version %v", version), Path: s.Url, Err: os.ErrNotExist}
	} else {
		return v, nil
	}
}

func (s *Stub) ReadUp(version uint64) (r io.ReadCloser, identifier string, err error) {
	if m, ok := s.Migrations.Up(version); ok {
		return io.NopCloser(bytes.NewBufferString(m.Identifier)), fmt.Sprintf("%v.up.stub", version), nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read up version %v", version), Path: s.Url, Err: os.ErrNotExist}
}

func (s *Stub) ReadDown(version uint64) (r io.ReadCloser, identifier string, err error) {
	if m, ok := s.Migrations.Down(version); ok {
		return io.NopCloser(bytes.NewBufferString(m.Identifier)), fmt.Sprintf("%v.down.stub", version), nil
	}
	return nil, "", &os.PathError{Op: fmt.Sprintf("read down version %v", version), Path: s.Url, Err: os.ErrNotExist}
}
