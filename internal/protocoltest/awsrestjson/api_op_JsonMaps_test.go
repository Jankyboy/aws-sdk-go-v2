// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_JsonMaps_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *JsonMapsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes JSON maps
		"RestJsonJsonMaps": {
			Params: &JsonMapsInput{
				DenseStructMap: map[string]types.GreetingStruct{
					"foo": {
						Hi: ptr.String("there"),
					},
					"baz": {
						Hi: ptr.String("bye"),
					},
				},
				SparseStructMap: map[string]*types.GreetingStruct{
					"foo": {
						Hi: ptr.String("there"),
					},
					"baz": {
						Hi: ptr.String("bye"),
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonMaps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "denseStructMap": {
			        "foo": {
			            "hi": "there"
			        },
			        "baz": {
			            "hi": "bye"
			        }
			    },
			    "sparseStructMap": {
			        "foo": {
			            "hi": "there"
			        },
			        "baz": {
			            "hi": "bye"
			        }
			    }
			}`))
			},
		},
		// Serializes JSON map values in sparse maps
		"RestJsonSerializesNullMapValues": {
			Params: &JsonMapsInput{
				SparseBooleanMap: map[string]*bool{
					"x": nil,
				},
				SparseNumberMap: map[string]*int32{
					"x": nil,
				},
				SparseStringMap: map[string]*string{
					"x": nil,
				},
				SparseStructMap: map[string]*types.GreetingStruct{
					"x": nil,
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonMaps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "sparseBooleanMap": {
			        "x": null
			    },
			    "sparseNumberMap": {
			        "x": null
			    },
			    "sparseStringMap": {
			        "x": null
			    },
			    "sparseStructMap": {
			        "x": null
			    }
			}`))
			},
		},
		// Ensure that 0 and false are sent over the wire in all maps and lists
		"RestJsonSerializesZeroValuesInMaps": {
			Params: &JsonMapsInput{
				DenseNumberMap: map[string]int32{
					"x": 0,
				},
				SparseNumberMap: map[string]*int32{
					"x": ptr.Int32(0),
				},
				DenseBooleanMap: map[string]bool{
					"x": false,
				},
				SparseBooleanMap: map[string]*bool{
					"x": ptr.Bool(false),
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/JsonMaps",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "denseNumberMap": {
			        "x": 0
			    },
			    "sparseNumberMap": {
			        "x": 0
			    },
			    "denseBooleanMap": {
			        "x": false
			    },
			    "sparseBooleanMap": {
			        "x": false
			    }
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if name == "RestJsonSerializesNullMapValues" {
				t.Skip("disabled test aws.protocoltests.restjson#RestJson aws.protocoltests.restjson#JsonMaps")
			}

			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               awshttp.NewBuildableClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.JsonMaps(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_JsonMaps_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *JsonMapsOutput
	}{
		// Deserializes JSON maps
		"RestJsonJsonMaps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "denseStructMap": {
			        "foo": {
			            "hi": "there"
			        },
			        "baz": {
			            "hi": "bye"
			        }
			    },
			    "sparseStructMap": {
			        "foo": {
			            "hi": "there"
			        },
			        "baz": {
			            "hi": "bye"
			        }
			   }
			}`),
			ExpectResult: &JsonMapsOutput{
				DenseStructMap: map[string]types.GreetingStruct{
					"foo": {
						Hi: ptr.String("there"),
					},
					"baz": {
						Hi: ptr.String("bye"),
					},
				},
				SparseStructMap: map[string]*types.GreetingStruct{
					"foo": {
						Hi: ptr.String("there"),
					},
					"baz": {
						Hi: ptr.String("bye"),
					},
				},
			},
		},
		// Deserializes null JSON map values
		"RestJsonDeserializesNullMapValues": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "sparseBooleanMap": {
			        "x": null
			    },
			    "sparseNumberMap": {
			        "x": null
			    },
			    "sparseStringMap": {
			        "x": null
			    },
			    "sparseStructMap": {
			        "x": null
			    }
			}`),
			ExpectResult: &JsonMapsOutput{
				SparseBooleanMap: map[string]*bool{
					"x": nil,
				},
				SparseNumberMap: map[string]*int32{
					"x": nil,
				},
				SparseStringMap: map[string]*string{
					"x": nil,
				},
				SparseStructMap: map[string]*types.GreetingStruct{
					"x": nil,
				},
			},
		},
		// Ensure that 0 and false are sent over the wire in all maps and lists
		"RestJsonDeserializesZeroValuesInMaps": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "denseNumberMap": {
			        "x": 0
			    },
			    "sparseNumberMap": {
			        "x": 0
			    },
			    "denseBooleanMap": {
			        "x": false
			    },
			    "sparseBooleanMap": {
			        "x": false
			    }
			}`),
			ExpectResult: &JsonMapsOutput{
				DenseNumberMap: map[string]int32{
					"x": 0,
				},
				SparseNumberMap: map[string]*int32{
					"x": ptr.Int32(0),
				},
				DenseBooleanMap: map[string]bool{
					"x": false,
				},
				SparseBooleanMap: map[string]*bool{
					"x": ptr.Bool(false),
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params JsonMapsInput
			result, err := client.JsonMaps(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
