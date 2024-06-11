// Code generated by mockery v2.42.2. DO NOT EDIT.

package registry

import (
	http "net/http"

	distribution "github.com/docker/distribution"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BlobExist provides a mock function with given fields: repository, digest
func (_m *Client) BlobExist(repository string, digest string) (bool, error) {
	ret := _m.Called(repository, digest)

	if len(ret) == 0 {
		panic("no return value specified for BlobExist")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(repository, digest)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(repository, digest)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(repository, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Catalog provides a mock function with given fields:
func (_m *Client) Catalog() ([]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Catalog")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Copy provides a mock function with given fields: srcRepository, srcReference, dstRepository, dstReference, override
func (_m *Client) Copy(srcRepository string, srcReference string, dstRepository string, dstReference string, override bool) error {
	ret := _m.Called(srcRepository, srcReference, dstRepository, dstReference, override)

	if len(ret) == 0 {
		panic("no return value specified for Copy")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, bool) error); ok {
		r0 = rf(srcRepository, srcReference, dstRepository, dstReference, override)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlob provides a mock function with given fields: repository, digest
func (_m *Client) DeleteBlob(repository string, digest string) error {
	ret := _m.Called(repository, digest)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(repository, digest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteManifest provides a mock function with given fields: repository, reference
func (_m *Client) DeleteManifest(repository string, reference string) error {
	ret := _m.Called(repository, reference)

	if len(ret) == 0 {
		panic("no return value specified for DeleteManifest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(repository, reference)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Do provides a mock function with given fields: req
func (_m *Client) Do(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request) (*http.Response, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTags provides a mock function with given fields: repository
func (_m *Client) ListTags(repository string) ([]string, error) {
	ret := _m.Called(repository)

	if len(ret) == 0 {
		panic("no return value specified for ListTags")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]string, error)); ok {
		return rf(repository)
	}
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(repository)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManifestExist provides a mock function with given fields: repository, reference
func (_m *Client) ManifestExist(repository string, reference string) (bool, *distribution.Descriptor, error) {
	ret := _m.Called(repository, reference)

	if len(ret) == 0 {
		panic("no return value specified for ManifestExist")
	}

	var r0 bool
	var r1 *distribution.Descriptor
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, *distribution.Descriptor, error)); ok {
		return rf(repository, reference)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(repository, reference)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) *distribution.Descriptor); ok {
		r1 = rf(repository, reference)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*distribution.Descriptor)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(repository, reference)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MountBlob provides a mock function with given fields: srcRepository, digest, dstRepository
func (_m *Client) MountBlob(srcRepository string, digest string, dstRepository string) error {
	ret := _m.Called(srcRepository, digest, dstRepository)

	if len(ret) == 0 {
		panic("no return value specified for MountBlob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(srcRepository, digest, dstRepository)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields:
func (_m *Client) Ping() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PullBlob provides a mock function with given fields: repository, digest
func (_m *Client) PullBlob(repository string, digest string) (int64, io.ReadCloser, error) {
	ret := _m.Called(repository, digest)

	if len(ret) == 0 {
		panic("no return value specified for PullBlob")
	}

	var r0 int64
	var r1 io.ReadCloser
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (int64, io.ReadCloser, error)); ok {
		return rf(repository, digest)
	}
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(repository, digest)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string) io.ReadCloser); ok {
		r1 = rf(repository, digest)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(repository, digest)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PullBlobChunk provides a mock function with given fields: repository, digest, blobSize, start, end
func (_m *Client) PullBlobChunk(repository string, digest string, blobSize int64, start int64, end int64) (int64, io.ReadCloser, error) {
	ret := _m.Called(repository, digest, blobSize, start, end)

	if len(ret) == 0 {
		panic("no return value specified for PullBlobChunk")
	}

	var r0 int64
	var r1 io.ReadCloser
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, int64, int64, int64) (int64, io.ReadCloser, error)); ok {
		return rf(repository, digest, blobSize, start, end)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, int64, int64) int64); ok {
		r0 = rf(repository, digest, blobSize, start, end)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, int64, int64) io.ReadCloser); ok {
		r1 = rf(repository, digest, blobSize, start, end)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string, int64, int64, int64) error); ok {
		r2 = rf(repository, digest, blobSize, start, end)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PullManifest provides a mock function with given fields: repository, reference, acceptedMediaTypes
func (_m *Client) PullManifest(repository string, reference string, acceptedMediaTypes ...string) (distribution.Manifest, string, error) {
	_va := make([]interface{}, len(acceptedMediaTypes))
	for _i := range acceptedMediaTypes {
		_va[_i] = acceptedMediaTypes[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, repository, reference)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PullManifest")
	}

	var r0 distribution.Manifest
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, ...string) (distribution.Manifest, string, error)); ok {
		return rf(repository, reference, acceptedMediaTypes...)
	}
	if rf, ok := ret.Get(0).(func(string, string, ...string) distribution.Manifest); ok {
		r0 = rf(repository, reference, acceptedMediaTypes...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(distribution.Manifest)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, ...string) string); ok {
		r1 = rf(repository, reference, acceptedMediaTypes...)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string, ...string) error); ok {
		r2 = rf(repository, reference, acceptedMediaTypes...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PushBlob provides a mock function with given fields: repository, digest, size, blob
func (_m *Client) PushBlob(repository string, digest string, size int64, blob io.Reader) error {
	ret := _m.Called(repository, digest, size, blob)

	if len(ret) == 0 {
		panic("no return value specified for PushBlob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64, io.Reader) error); ok {
		r0 = rf(repository, digest, size, blob)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PushBlobChunk provides a mock function with given fields: repository, digest, blobSize, chunk, start, end, location
func (_m *Client) PushBlobChunk(repository string, digest string, blobSize int64, chunk io.Reader, start int64, end int64, location string) (string, int64, error) {
	ret := _m.Called(repository, digest, blobSize, chunk, start, end, location)

	if len(ret) == 0 {
		panic("no return value specified for PushBlobChunk")
	}

	var r0 string
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string, int64, io.Reader, int64, int64, string) (string, int64, error)); ok {
		return rf(repository, digest, blobSize, chunk, start, end, location)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64, io.Reader, int64, int64, string) string); ok {
		r0 = rf(repository, digest, blobSize, chunk, start, end, location)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, int64, io.Reader, int64, int64, string) int64); ok {
		r1 = rf(repository, digest, blobSize, chunk, start, end, location)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(string, string, int64, io.Reader, int64, int64, string) error); ok {
		r2 = rf(repository, digest, blobSize, chunk, start, end, location)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PushManifest provides a mock function with given fields: repository, reference, mediaType, payload
func (_m *Client) PushManifest(repository string, reference string, mediaType string, payload []byte) (string, error) {
	ret := _m.Called(repository, reference, mediaType, payload)

	if len(ret) == 0 {
		panic("no return value specified for PushManifest")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, []byte) (string, error)); ok {
		return rf(repository, reference, mediaType, payload)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, []byte) string); ok {
		r0 = rf(repository, reference, mediaType, payload)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string, string, []byte) error); ok {
		r1 = rf(repository, reference, mediaType, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}