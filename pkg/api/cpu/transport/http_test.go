package transport_test

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/go-pg/pg/orm"
// 	"github.com/labstack/echo"
// 	"github.com/raspibuddy/rpi"
// 	"github.com/raspibuddy/rpi/pkg/api/vcore/transport"
// 	"github.com/raspibuddy/rpi/pkg/utl/server"
// 	"github.com/ribice/gorsk/pkg/api/user"
// 	"github.com/ribice/gorsk/pkg/utl/mock"
// 	"gopkg.in/go-playground/assert.v1"
// )

// func TestList(t *testing.T) {
// 	cases := []struct {
// 		name       string
// 		req        string
// 		wantStatus int
// 		wantResp   []rpi.CPU
// 		udb        *mockdb.CPU
// 	}{
// 		{
// 			name:       "Invalid request",
// 			req:        `?limit=2222&page=-1`,
// 			wantStatus: http.StatusBadRequest,
// 		},
// 		{
// 			name: "Fail on query list",
// 			req:  `?limit=100&page=1`,
// 			rbac: &mock.RBAC{
// 				UserFn: func(c echo.Context) gorsk.AuthUser {
// 					return gorsk.AuthUser{
// 						ID:         1,
// 						CompanyID:  2,
// 						LocationID: 3,
// 						Role:       gorsk.UserRole,
// 						Email:      "john@mail.com",
// 					}
// 				}},
// 			wantStatus: http.StatusForbidden,
// 		},
// 		{
// 			name: "Success",
// 			req:  `?limit=100&page=1`,
// 			rbac: &mock.RBAC{
// 				UserFn: func(c echo.Context) gorsk.AuthUser {
// 					return gorsk.AuthUser{
// 						ID:         1,
// 						CompanyID:  2,
// 						LocationID: 3,
// 						Role:       gorsk.SuperAdminRole,
// 						Email:      "john@mail.com",
// 					}
// 				}},
// 			udb: &mockdb.User{
// 				ListFn: func(db orm.DB, q *gorsk.ListQuery, p gorsk.Pagination) ([]gorsk.User, error) {
// 					if p.Limit == 100 && p.Offset == 100 {
// 						return []gorsk.User{
// 							{
// 								Base: gorsk.Base{
// 									ID:        10,
// 									CreatedAt: mock.TestTime(2001),
// 									UpdatedAt: mock.TestTime(2002),
// 								},
// 								FirstName:  "John",
// 								LastName:   "Doe",
// 								Email:      "john@mail.com",
// 								CompanyID:  2,
// 								LocationID: 3,
// 								Role: &gorsk.Role{
// 									ID:          1,
// 									AccessLevel: 1,
// 									Name:        "SUPER_ADMIN",
// 								},
// 							},
// 							{
// 								Base: gorsk.Base{
// 									ID:        11,
// 									CreatedAt: mock.TestTime(2004),
// 									UpdatedAt: mock.TestTime(2005),
// 								},
// 								FirstName:  "Joanna",
// 								LastName:   "Dye",
// 								Email:      "joanna@mail.com",
// 								CompanyID:  1,
// 								LocationID: 2,
// 								Role: &gorsk.Role{
// 									ID:          2,
// 									AccessLevel: 2,
// 									Name:        "ADMIN",
// 								},
// 							},
// 						}, nil
// 					}
// 					return nil, gorsk.ErrGeneric
// 				},
// 			},
// 			wantStatus: http.StatusOK,
// 			wantResp: &listResponse{
// 				Users: []gorsk.User{
// 					{
// 						Base: gorsk.Base{
// 							ID:        10,
// 							CreatedAt: mock.TestTime(2001),
// 							UpdatedAt: mock.TestTime(2002),
// 						},
// 						FirstName:  "John",
// 						LastName:   "Doe",
// 						Email:      "john@mail.com",
// 						CompanyID:  2,
// 						LocationID: 3,
// 						Role: &gorsk.Role{
// 							ID:          1,
// 							AccessLevel: 1,
// 							Name:        "SUPER_ADMIN",
// 						},
// 					},
// 					{
// 						Base: gorsk.Base{
// 							ID:        11,
// 							CreatedAt: mock.TestTime(2004),
// 							UpdatedAt: mock.TestTime(2005),
// 						},
// 						FirstName:  "Joanna",
// 						LastName:   "Dye",
// 						Email:      "joanna@mail.com",
// 						CompanyID:  1,
// 						LocationID: 2,
// 						Role: &gorsk.Role{
// 							ID:          2,
// 							AccessLevel: 2,
// 							Name:        "ADMIN",
// 						},
// 					},
// 				}, Page: 1},
// 		},
// 	}

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			r := server.New()
// 			rg := r.Group("")
// 			transport.NewHTTP(user.New(nil, tt.udb, tt.rbac, tt.sec), rg)
// 			ts := httptest.NewServer(r)
// 			defer ts.Close()
// 			path := ts.URL + "/users" + tt.req
// 			res, err := http.Get(path)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			defer res.Body.Close()
// 			if tt.wantResp != nil {
// 				response := new(listResponse)
// 				if err := json.NewDecoder(res.Body).Decode(response); err != nil {
// 					t.Fatal(err)
// 				}
// 				assert.Equal(t, tt.wantResp, response)
// 			}
// 			assert.Equal(t, tt.wantStatus, res.StatusCode)
// 		})
// 	}
// }
