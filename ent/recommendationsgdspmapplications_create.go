// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recruit/ent/exam_applications_gdspm"
	"recruit/ent/recommendationsgdspmapplications"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecommendationsGDSPMApplicationsCreate is the builder for creating a RecommendationsGDSPMApplications entity.
type RecommendationsGDSPMApplicationsCreate struct {
	config
	mutation *RecommendationsGDSPMApplicationsMutation
	hooks    []Hook
}

// SetApplicationID sets the "ApplicationID" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetApplicationID(i int64) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetApplicationID(i)
	return rgac
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableApplicationID(i *int64) *RecommendationsGDSPMApplicationsCreate {
	if i != nil {
		rgac.SetApplicationID(*i)
	}
	return rgac
}

// SetEmployeeID sets the "EmployeeID" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetEmployeeID(i int64) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetEmployeeID(i)
	return rgac
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableEmployeeID(i *int64) *RecommendationsGDSPMApplicationsCreate {
	if i != nil {
		rgac.SetEmployeeID(*i)
	}
	return rgac
}

// SetExamNameCode sets the "ExamNameCode" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetExamNameCode(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetExamNameCode(s)
	return rgac
}

// SetNillableExamNameCode sets the "ExamNameCode" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableExamNameCode(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetExamNameCode(*s)
	}
	return rgac
}

// SetExamYear sets the "ExamYear" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetExamYear(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetExamYear(s)
	return rgac
}

// SetNillableExamYear sets the "ExamYear" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableExamYear(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetExamYear(*s)
	}
	return rgac
}

// SetVacancyYear sets the "VacancyYear" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetVacancyYear(i int32) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetVacancyYear(i)
	return rgac
}

// SetNillableVacancyYear sets the "VacancyYear" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableVacancyYear(i *int32) *RecommendationsGDSPMApplicationsCreate {
	if i != nil {
		rgac.SetVacancyYear(*i)
	}
	return rgac
}

// SetPost sets the "Post" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetPost(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetPost(s)
	return rgac
}

// SetNillablePost sets the "Post" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillablePost(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetPost(*s)
	}
	return rgac
}

// SetEligible sets the "Eligible" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetEligible(b bool) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetEligible(b)
	return rgac
}

// SetNillableEligible sets the "Eligible" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableEligible(b *bool) *RecommendationsGDSPMApplicationsCreate {
	if b != nil {
		rgac.SetEligible(*b)
	}
	return rgac
}

// SetCARecommendations sets the "CA_Recommendations" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetCARecommendations(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetCARecommendations(s)
	return rgac
}

// SetNillableCARecommendations sets the "CA_Recommendations" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableCARecommendations(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetCARecommendations(*s)
	}
	return rgac
}

// SetCAUpdatedAt sets the "CA_UpdatedAt" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetCAUpdatedAt(t time.Time) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetCAUpdatedAt(t)
	return rgac
}

// SetNillableCAUpdatedAt sets the "CA_UpdatedAt" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableCAUpdatedAt(t *time.Time) *RecommendationsGDSPMApplicationsCreate {
	if t != nil {
		rgac.SetCAUpdatedAt(*t)
	}
	return rgac
}

// SetCAUserName sets the "CA_UserName" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetCAUserName(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetCAUserName(s)
	return rgac
}

// SetNillableCAUserName sets the "CA_UserName" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableCAUserName(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetCAUserName(*s)
	}
	return rgac
}

// SetCARemarks sets the "CA_Remarks" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetCARemarks(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetCARemarks(s)
	return rgac
}

// SetNillableCARemarks sets the "CA_Remarks" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableCARemarks(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetCARemarks(*s)
	}
	return rgac
}

// SetNORecommendations sets the "NO_Recommendations" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNORecommendations(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetNORecommendations(s)
	return rgac
}

// SetNillableNORecommendations sets the "NO_Recommendations" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableNORecommendations(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetNORecommendations(*s)
	}
	return rgac
}

// SetNOUpdatedAt sets the "NO_UpdatedAt" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNOUpdatedAt(t time.Time) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetNOUpdatedAt(t)
	return rgac
}

// SetNillableNOUpdatedAt sets the "NO_UpdatedAt" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableNOUpdatedAt(t *time.Time) *RecommendationsGDSPMApplicationsCreate {
	if t != nil {
		rgac.SetNOUpdatedAt(*t)
	}
	return rgac
}

// SetNOUserName sets the "NO_UserName" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNOUserName(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetNOUserName(s)
	return rgac
}

// SetNillableNOUserName sets the "NO_UserName" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableNOUserName(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetNOUserName(*s)
	}
	return rgac
}

// SetNORemarks sets the "NO_Remarks" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNORemarks(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetNORemarks(s)
	return rgac
}

// SetNillableNORemarks sets the "NO_Remarks" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableNORemarks(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetNORemarks(*s)
	}
	return rgac
}

// SetApplicationStatus sets the "ApplicationStatus" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetApplicationStatus(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetApplicationStatus(s)
	return rgac
}

// SetNillableApplicationStatus sets the "ApplicationStatus" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableApplicationStatus(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetApplicationStatus(*s)
	}
	return rgac
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetUpdatedAt(t time.Time) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetUpdatedAt(t)
	return rgac
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableUpdatedAt(t *time.Time) *RecommendationsGDSPMApplicationsCreate {
	if t != nil {
		rgac.SetUpdatedAt(*t)
	}
	return rgac
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetUpdatedBy(s string) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetUpdatedBy(s)
	return rgac
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableUpdatedBy(s *string) *RecommendationsGDSPMApplicationsCreate {
	if s != nil {
		rgac.SetUpdatedBy(*s)
	}
	return rgac
}

// SetGenerateHallTicketFlag sets the "GenerateHallTicketFlag" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetGenerateHallTicketFlag(b bool) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetGenerateHallTicketFlag(b)
	return rgac
}

// SetNillableGenerateHallTicketFlag sets the "GenerateHallTicketFlag" field if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableGenerateHallTicketFlag(b *bool) *RecommendationsGDSPMApplicationsCreate {
	if b != nil {
		rgac.SetGenerateHallTicketFlag(*b)
	}
	return rgac
}

// SetID sets the "id" field.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetID(i int64) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetID(i)
	return rgac
}

// SetApplnRefID sets the "ApplnRef" edge to the Exam_Applications_GDSPM entity by ID.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetApplnRefID(id int64) *RecommendationsGDSPMApplicationsCreate {
	rgac.mutation.SetApplnRefID(id)
	return rgac
}

// SetNillableApplnRefID sets the "ApplnRef" edge to the Exam_Applications_GDSPM entity by ID if the given value is not nil.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetNillableApplnRefID(id *int64) *RecommendationsGDSPMApplicationsCreate {
	if id != nil {
		rgac = rgac.SetApplnRefID(*id)
	}
	return rgac
}

// SetApplnRef sets the "ApplnRef" edge to the Exam_Applications_GDSPM entity.
func (rgac *RecommendationsGDSPMApplicationsCreate) SetApplnRef(e *Exam_Applications_GDSPM) *RecommendationsGDSPMApplicationsCreate {
	return rgac.SetApplnRefID(e.ID)
}

// Mutation returns the RecommendationsGDSPMApplicationsMutation object of the builder.
func (rgac *RecommendationsGDSPMApplicationsCreate) Mutation() *RecommendationsGDSPMApplicationsMutation {
	return rgac.mutation
}

// Save creates the RecommendationsGDSPMApplications in the database.
func (rgac *RecommendationsGDSPMApplicationsCreate) Save(ctx context.Context) (*RecommendationsGDSPMApplications, error) {
	rgac.defaults()
	return withHooks(ctx, rgac.sqlSave, rgac.mutation, rgac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rgac *RecommendationsGDSPMApplicationsCreate) SaveX(ctx context.Context) *RecommendationsGDSPMApplications {
	v, err := rgac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgac *RecommendationsGDSPMApplicationsCreate) Exec(ctx context.Context) error {
	_, err := rgac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgac *RecommendationsGDSPMApplicationsCreate) ExecX(ctx context.Context) {
	if err := rgac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rgac *RecommendationsGDSPMApplicationsCreate) defaults() {
	if _, ok := rgac.mutation.UpdatedAt(); !ok {
		v := recommendationsgdspmapplications.DefaultUpdatedAt()
		rgac.mutation.SetUpdatedAt(v)
	}
	if _, ok := rgac.mutation.UpdatedBy(); !ok {
		v := recommendationsgdspmapplications.DefaultUpdatedBy
		rgac.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rgac *RecommendationsGDSPMApplicationsCreate) check() error {
	return nil
}

func (rgac *RecommendationsGDSPMApplicationsCreate) sqlSave(ctx context.Context) (*RecommendationsGDSPMApplications, error) {
	if err := rgac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rgac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rgac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rgac.mutation.id = &_node.ID
	rgac.mutation.done = true
	return _node, nil
}

func (rgac *RecommendationsGDSPMApplicationsCreate) createSpec() (*RecommendationsGDSPMApplications, *sqlgraph.CreateSpec) {
	var (
		_node = &RecommendationsGDSPMApplications{config: rgac.config}
		_spec = sqlgraph.NewCreateSpec(recommendationsgdspmapplications.Table, sqlgraph.NewFieldSpec(recommendationsgdspmapplications.FieldID, field.TypeInt64))
	)
	if id, ok := rgac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rgac.mutation.ApplicationID(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldApplicationID, field.TypeInt64, value)
		_node.ApplicationID = value
	}
	if value, ok := rgac.mutation.EmployeeID(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := rgac.mutation.ExamNameCode(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldExamNameCode, field.TypeString, value)
		_node.ExamNameCode = value
	}
	if value, ok := rgac.mutation.ExamYear(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldExamYear, field.TypeString, value)
		_node.ExamYear = value
	}
	if value, ok := rgac.mutation.VacancyYear(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldVacancyYear, field.TypeInt32, value)
		_node.VacancyYear = value
	}
	if value, ok := rgac.mutation.Post(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldPost, field.TypeString, value)
		_node.Post = value
	}
	if value, ok := rgac.mutation.Eligible(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldEligible, field.TypeBool, value)
		_node.Eligible = value
	}
	if value, ok := rgac.mutation.CARecommendations(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldCARecommendations, field.TypeString, value)
		_node.CARecommendations = value
	}
	if value, ok := rgac.mutation.CAUpdatedAt(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldCAUpdatedAt, field.TypeTime, value)
		_node.CAUpdatedAt = value
	}
	if value, ok := rgac.mutation.CAUserName(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldCAUserName, field.TypeString, value)
		_node.CAUserName = value
	}
	if value, ok := rgac.mutation.CARemarks(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldCARemarks, field.TypeString, value)
		_node.CARemarks = value
	}
	if value, ok := rgac.mutation.NORecommendations(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldNORecommendations, field.TypeString, value)
		_node.NORecommendations = value
	}
	if value, ok := rgac.mutation.NOUpdatedAt(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldNOUpdatedAt, field.TypeTime, value)
		_node.NOUpdatedAt = value
	}
	if value, ok := rgac.mutation.NOUserName(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldNOUserName, field.TypeString, value)
		_node.NOUserName = value
	}
	if value, ok := rgac.mutation.NORemarks(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldNORemarks, field.TypeString, value)
		_node.NORemarks = value
	}
	if value, ok := rgac.mutation.ApplicationStatus(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldApplicationStatus, field.TypeString, value)
		_node.ApplicationStatus = value
	}
	if value, ok := rgac.mutation.UpdatedAt(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rgac.mutation.UpdatedBy(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := rgac.mutation.GenerateHallTicketFlag(); ok {
		_spec.SetField(recommendationsgdspmapplications.FieldGenerateHallTicketFlag, field.TypeBool, value)
		_node.GenerateHallTicketFlag = value
	}
	if nodes := rgac.mutation.ApplnRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recommendationsgdspmapplications.ApplnRefTable,
			Columns: []string{recommendationsgdspmapplications.ApplnRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_gdspm.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_applications_gdspm_gdspm_applications_ref = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RecommendationsGDSPMApplicationsCreateBulk is the builder for creating many RecommendationsGDSPMApplications entities in bulk.
type RecommendationsGDSPMApplicationsCreateBulk struct {
	config
	builders []*RecommendationsGDSPMApplicationsCreate
}

// Save creates the RecommendationsGDSPMApplications entities in the database.
func (rgacb *RecommendationsGDSPMApplicationsCreateBulk) Save(ctx context.Context) ([]*RecommendationsGDSPMApplications, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rgacb.builders))
	nodes := make([]*RecommendationsGDSPMApplications, len(rgacb.builders))
	mutators := make([]Mutator, len(rgacb.builders))
	for i := range rgacb.builders {
		func(i int, root context.Context) {
			builder := rgacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecommendationsGDSPMApplicationsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rgacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rgacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rgacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rgacb *RecommendationsGDSPMApplicationsCreateBulk) SaveX(ctx context.Context) []*RecommendationsGDSPMApplications {
	v, err := rgacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgacb *RecommendationsGDSPMApplicationsCreateBulk) Exec(ctx context.Context) error {
	_, err := rgacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgacb *RecommendationsGDSPMApplicationsCreateBulk) ExecX(ctx context.Context) {
	if err := rgacb.Exec(ctx); err != nil {
		panic(err)
	}
}
