// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recruit/ent/exam_application_mtspmmg"
	"recruit/ent/recommendationsmtspmmgapplications"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecommendationsMTSPMMGApplicationsCreate is the builder for creating a RecommendationsMTSPMMGApplications entity.
type RecommendationsMTSPMMGApplicationsCreate struct {
	config
	mutation *RecommendationsMTSPMMGApplicationsMutation
	hooks    []Hook
}

// SetApplicationID sets the "ApplicationID" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetApplicationID(i int64) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetApplicationID(i)
	return rmac
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableApplicationID(i *int64) *RecommendationsMTSPMMGApplicationsCreate {
	if i != nil {
		rmac.SetApplicationID(*i)
	}
	return rmac
}

// SetEmployeeID sets the "EmployeeID" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetEmployeeID(i int64) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetEmployeeID(i)
	return rmac
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableEmployeeID(i *int64) *RecommendationsMTSPMMGApplicationsCreate {
	if i != nil {
		rmac.SetEmployeeID(*i)
	}
	return rmac
}

// SetExamNameCode sets the "ExamNameCode" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetExamNameCode(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetExamNameCode(s)
	return rmac
}

// SetNillableExamNameCode sets the "ExamNameCode" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableExamNameCode(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetExamNameCode(*s)
	}
	return rmac
}

// SetExamYear sets the "ExamYear" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetExamYear(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetExamYear(s)
	return rmac
}

// SetNillableExamYear sets the "ExamYear" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableExamYear(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetExamYear(*s)
	}
	return rmac
}

// SetExamName sets the "ExamName" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetExamName(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetExamName(s)
	return rmac
}

// SetNillableExamName sets the "ExamName" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableExamName(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetExamName(*s)
	}
	return rmac
}

// SetVacancyYear sets the "VacancyYear" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetVacancyYear(i int32) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetVacancyYear(i)
	return rmac
}

// SetNillableVacancyYear sets the "VacancyYear" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableVacancyYear(i *int32) *RecommendationsMTSPMMGApplicationsCreate {
	if i != nil {
		rmac.SetVacancyYear(*i)
	}
	return rmac
}

// SetCARecommendations sets the "CA_Recommendations" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetCARecommendations(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetCARecommendations(s)
	return rmac
}

// SetNillableCARecommendations sets the "CA_Recommendations" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableCARecommendations(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetCARecommendations(*s)
	}
	return rmac
}

// SetCAUpdatedAt sets the "CA_UpdatedAt" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetCAUpdatedAt(t time.Time) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetCAUpdatedAt(t)
	return rmac
}

// SetNillableCAUpdatedAt sets the "CA_UpdatedAt" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableCAUpdatedAt(t *time.Time) *RecommendationsMTSPMMGApplicationsCreate {
	if t != nil {
		rmac.SetCAUpdatedAt(*t)
	}
	return rmac
}

// SetCAUserName sets the "CA_UserName" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetCAUserName(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetCAUserName(s)
	return rmac
}

// SetNillableCAUserName sets the "CA_UserName" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableCAUserName(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetCAUserName(*s)
	}
	return rmac
}

// SetCARemarks sets the "CA_Remarks" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetCARemarks(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetCARemarks(s)
	return rmac
}

// SetNillableCARemarks sets the "CA_Remarks" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableCARemarks(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetCARemarks(*s)
	}
	return rmac
}

// SetNORecommendations sets the "NO_Recommendations" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNORecommendations(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetNORecommendations(s)
	return rmac
}

// SetNillableNORecommendations sets the "NO_Recommendations" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableNORecommendations(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetNORecommendations(*s)
	}
	return rmac
}

// SetNOUpdatedAt sets the "NO_UpdatedAt" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNOUpdatedAt(t time.Time) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetNOUpdatedAt(t)
	return rmac
}

// SetNillableNOUpdatedAt sets the "NO_UpdatedAt" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableNOUpdatedAt(t *time.Time) *RecommendationsMTSPMMGApplicationsCreate {
	if t != nil {
		rmac.SetNOUpdatedAt(*t)
	}
	return rmac
}

// SetNOUserName sets the "NO_UserName" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNOUserName(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetNOUserName(s)
	return rmac
}

// SetNillableNOUserName sets the "NO_UserName" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableNOUserName(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetNOUserName(*s)
	}
	return rmac
}

// SetNORemarks sets the "NO_Remarks" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNORemarks(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetNORemarks(s)
	return rmac
}

// SetNillableNORemarks sets the "NO_Remarks" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableNORemarks(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetNORemarks(*s)
	}
	return rmac
}

// SetApplicationStatus sets the "ApplicationStatus" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetApplicationStatus(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetApplicationStatus(s)
	return rmac
}

// SetNillableApplicationStatus sets the "ApplicationStatus" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableApplicationStatus(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetApplicationStatus(*s)
	}
	return rmac
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetUpdatedAt(t time.Time) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetUpdatedAt(t)
	return rmac
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableUpdatedAt(t *time.Time) *RecommendationsMTSPMMGApplicationsCreate {
	if t != nil {
		rmac.SetUpdatedAt(*t)
	}
	return rmac
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetUpdatedBy(s string) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetUpdatedBy(s)
	return rmac
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableUpdatedBy(s *string) *RecommendationsMTSPMMGApplicationsCreate {
	if s != nil {
		rmac.SetUpdatedBy(*s)
	}
	return rmac
}

// SetGenerateHallTicketFlag sets the "GenerateHallTicketFlag" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetGenerateHallTicketFlag(b bool) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetGenerateHallTicketFlag(b)
	return rmac
}

// SetNillableGenerateHallTicketFlag sets the "GenerateHallTicketFlag" field if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableGenerateHallTicketFlag(b *bool) *RecommendationsMTSPMMGApplicationsCreate {
	if b != nil {
		rmac.SetGenerateHallTicketFlag(*b)
	}
	return rmac
}

// SetID sets the "id" field.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetID(i int64) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetID(i)
	return rmac
}

// SetApplnRefID sets the "ApplnRef" edge to the Exam_Application_MTSPMMG entity by ID.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetApplnRefID(id int64) *RecommendationsMTSPMMGApplicationsCreate {
	rmac.mutation.SetApplnRefID(id)
	return rmac
}

// SetNillableApplnRefID sets the "ApplnRef" edge to the Exam_Application_MTSPMMG entity by ID if the given value is not nil.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetNillableApplnRefID(id *int64) *RecommendationsMTSPMMGApplicationsCreate {
	if id != nil {
		rmac = rmac.SetApplnRefID(*id)
	}
	return rmac
}

// SetApplnRef sets the "ApplnRef" edge to the Exam_Application_MTSPMMG entity.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SetApplnRef(e *Exam_Application_MTSPMMG) *RecommendationsMTSPMMGApplicationsCreate {
	return rmac.SetApplnRefID(e.ID)
}

// Mutation returns the RecommendationsMTSPMMGApplicationsMutation object of the builder.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) Mutation() *RecommendationsMTSPMMGApplicationsMutation {
	return rmac.mutation
}

// Save creates the RecommendationsMTSPMMGApplications in the database.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) Save(ctx context.Context) (*RecommendationsMTSPMMGApplications, error) {
	rmac.defaults()
	return withHooks(ctx, rmac.sqlSave, rmac.mutation, rmac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) SaveX(ctx context.Context) *RecommendationsMTSPMMGApplications {
	v, err := rmac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) Exec(ctx context.Context) error {
	_, err := rmac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) ExecX(ctx context.Context) {
	if err := rmac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) defaults() {
	if _, ok := rmac.mutation.UpdatedAt(); !ok {
		v := recommendationsmtspmmgapplications.DefaultUpdatedAt()
		rmac.mutation.SetUpdatedAt(v)
	}
	if _, ok := rmac.mutation.UpdatedBy(); !ok {
		v := recommendationsmtspmmgapplications.DefaultUpdatedBy
		rmac.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmac *RecommendationsMTSPMMGApplicationsCreate) check() error {
	return nil
}

func (rmac *RecommendationsMTSPMMGApplicationsCreate) sqlSave(ctx context.Context) (*RecommendationsMTSPMMGApplications, error) {
	if err := rmac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rmac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rmac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rmac.mutation.id = &_node.ID
	rmac.mutation.done = true
	return _node, nil
}

func (rmac *RecommendationsMTSPMMGApplicationsCreate) createSpec() (*RecommendationsMTSPMMGApplications, *sqlgraph.CreateSpec) {
	var (
		_node = &RecommendationsMTSPMMGApplications{config: rmac.config}
		_spec = sqlgraph.NewCreateSpec(recommendationsmtspmmgapplications.Table, sqlgraph.NewFieldSpec(recommendationsmtspmmgapplications.FieldID, field.TypeInt64))
	)
	if id, ok := rmac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rmac.mutation.EmployeeID(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := rmac.mutation.ExamNameCode(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldExamNameCode, field.TypeString, value)
		_node.ExamNameCode = value
	}
	if value, ok := rmac.mutation.ExamYear(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldExamYear, field.TypeString, value)
		_node.ExamYear = value
	}
	if value, ok := rmac.mutation.ExamName(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldExamName, field.TypeString, value)
		_node.ExamName = value
	}
	if value, ok := rmac.mutation.VacancyYear(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldVacancyYear, field.TypeInt32, value)
		_node.VacancyYear = value
	}
	if value, ok := rmac.mutation.CARecommendations(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldCARecommendations, field.TypeString, value)
		_node.CARecommendations = value
	}
	if value, ok := rmac.mutation.CAUpdatedAt(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldCAUpdatedAt, field.TypeTime, value)
		_node.CAUpdatedAt = value
	}
	if value, ok := rmac.mutation.CAUserName(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldCAUserName, field.TypeString, value)
		_node.CAUserName = value
	}
	if value, ok := rmac.mutation.CARemarks(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldCARemarks, field.TypeString, value)
		_node.CARemarks = value
	}
	if value, ok := rmac.mutation.NORecommendations(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldNORecommendations, field.TypeString, value)
		_node.NORecommendations = value
	}
	if value, ok := rmac.mutation.NOUpdatedAt(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldNOUpdatedAt, field.TypeTime, value)
		_node.NOUpdatedAt = value
	}
	if value, ok := rmac.mutation.NOUserName(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldNOUserName, field.TypeString, value)
		_node.NOUserName = value
	}
	if value, ok := rmac.mutation.NORemarks(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldNORemarks, field.TypeString, value)
		_node.NORemarks = value
	}
	if value, ok := rmac.mutation.ApplicationStatus(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldApplicationStatus, field.TypeString, value)
		_node.ApplicationStatus = value
	}
	if value, ok := rmac.mutation.UpdatedAt(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rmac.mutation.UpdatedBy(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := rmac.mutation.GenerateHallTicketFlag(); ok {
		_spec.SetField(recommendationsmtspmmgapplications.FieldGenerateHallTicketFlag, field.TypeBool, value)
		_node.GenerateHallTicketFlag = value
	}
	if nodes := rmac.mutation.ApplnRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recommendationsmtspmmgapplications.ApplnRefTable,
			Columns: []string{recommendationsmtspmmgapplications.ApplnRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_application_mtspmmg.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ApplicationID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RecommendationsMTSPMMGApplicationsCreateBulk is the builder for creating many RecommendationsMTSPMMGApplications entities in bulk.
type RecommendationsMTSPMMGApplicationsCreateBulk struct {
	config
	builders []*RecommendationsMTSPMMGApplicationsCreate
}

// Save creates the RecommendationsMTSPMMGApplications entities in the database.
func (rmacb *RecommendationsMTSPMMGApplicationsCreateBulk) Save(ctx context.Context) ([]*RecommendationsMTSPMMGApplications, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rmacb.builders))
	nodes := make([]*RecommendationsMTSPMMGApplications, len(rmacb.builders))
	mutators := make([]Mutator, len(rmacb.builders))
	for i := range rmacb.builders {
		func(i int, root context.Context) {
			builder := rmacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecommendationsMTSPMMGApplicationsMutation)
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
					_, err = mutators[i+1].Mutate(root, rmacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rmacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rmacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rmacb *RecommendationsMTSPMMGApplicationsCreateBulk) SaveX(ctx context.Context) []*RecommendationsMTSPMMGApplications {
	v, err := rmacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rmacb *RecommendationsMTSPMMGApplicationsCreateBulk) Exec(ctx context.Context) error {
	_, err := rmacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmacb *RecommendationsMTSPMMGApplicationsCreateBulk) ExecX(ctx context.Context) {
	if err := rmacb.Exec(ctx); err != nil {
		panic(err)
	}
}