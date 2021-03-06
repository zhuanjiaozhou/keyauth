package mongo

import (
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/infraboard/keyauth/pkg/department"
)

func newQueryDepartmentRequest(req *department.QueryDepartmentRequest) *queryDepartmentRequest {
	return &queryDepartmentRequest{req}
}

type queryDepartmentRequest struct {
	*department.QueryDepartmentRequest
}

func (r *queryDepartmentRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryDepartmentRequest) FindFilter() bson.M {
	filter := bson.M{}

	tk := r.GetToken()
	filter["domain"] = tk.Domain
	if r.ParentID != nil {
		filter["parent_id"] = r.ParentID
	}
	if r.Keywords != "" {
		filter["$or"] = bson.A{
			bson.M{"name": bson.M{"$regex": r.Keywords, "$options": "im"}},
		}
	}

	return filter
}

func newDescribeDepartment(req *department.DescribeDeparmentRequest) (*describeDepartmentRequest, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest(err.Error())
	}

	return &describeDepartmentRequest{req}, nil
}

type describeDepartmentRequest struct {
	*department.DescribeDeparmentRequest
}

func (r *describeDepartmentRequest) FindFilter() bson.M {
	filter := bson.M{}

	if r.ID != "" {
		filter["_id"] = r.ID
	}
	if r.Name != "" {
		filter["name"] = r.Name
	}

	return filter
}

func newQueryApplicationFormRequest(req *department.QueryApplicationFormRequet) *queryApplicationFormRequest {
	return &queryApplicationFormRequest{req}
}

type queryApplicationFormRequest struct {
	*department.QueryApplicationFormRequet
}

func (r *queryApplicationFormRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.PageSize)
	skip := int64(r.PageSize) * int64(r.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryApplicationFormRequest) FindFilter() bson.M {
	tk := r.GetToken()

	filter := bson.M{}
	filter["domain"] = tk.Domain

	if r.Account != "" {
		filter["account"] = r.Account
	}
	if r.DepartmentID != "" {
		filter["department_id"] = r.DepartmentID
	}
	if r.Status != nil {
		filter["status"] = r.Status
	}

	return filter
}

func newDescribeApplicationForm(req *department.DescribeApplicationFormRequet) *describeApplicationForm {
	return &describeApplicationForm{
		DescribeApplicationFormRequet: req,
	}
}

type describeApplicationForm struct {
	*department.DescribeApplicationFormRequet
}

func (r *describeApplicationForm) FindFilter() bson.M {
	tk := r.GetToken()

	filter := bson.M{}
	filter["domain"] = tk.Domain

	if r.ID != "" {
		filter["_id"] = r.ID
	}

	return filter
}
