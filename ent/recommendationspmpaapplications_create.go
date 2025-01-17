// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"recruit/ent/exam_applications_pmpa"
	"recruit/ent/recommendationspmpaapplications"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecommendationsPMPAApplicationsCreate is the builder for creating a RecommendationsPMPAApplications entity.
type RecommendationsPMPAApplicationsCreate struct {
	config
	mutation *RecommendationsPMPAApplicationsMutation
	hooks    []Hook
}

// SetApplicationID sets the "ApplicationID" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetApplicationID(i int64) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetApplicationID(i)
	return rpac
}

// SetNillableApplicationID sets the "ApplicationID" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableApplicationID(i *int64) *RecommendationsPMPAApplicationsCreate {
	if i != nil {
		rpac.SetApplicationID(*i)
	}
	return rpac
}

// SetEmployeeID sets the "EmployeeID" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetEmployeeID(i int64) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetEmployeeID(i)
	return rpac
}

// SetNillableEmployeeID sets the "EmployeeID" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableEmployeeID(i *int64) *RecommendationsPMPAApplicationsCreate {
	if i != nil {
		rpac.SetEmployeeID(*i)
	}
	return rpac
}

// SetExamNameCode sets the "ExamNameCode" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetExamNameCode(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetExamNameCode(s)
	return rpac
}

// SetNillableExamNameCode sets the "ExamNameCode" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableExamNameCode(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetExamNameCode(*s)
	}
	return rpac
}

// SetExamYear sets the "ExamYear" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetExamYear(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetExamYear(s)
	return rpac
}

// SetNillableExamYear sets the "ExamYear" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableExamYear(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetExamYear(*s)
	}
	return rpac
}

// SetVacancyYear sets the "VacancyYear" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetVacancyYear(i int32) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetVacancyYear(i)
	return rpac
}

// SetNillableVacancyYear sets the "VacancyYear" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableVacancyYear(i *int32) *RecommendationsPMPAApplicationsCreate {
	if i != nil {
		rpac.SetVacancyYear(*i)
	}
	return rpac
}

// SetCARecommendations sets the "CA_Recommendations" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetCARecommendations(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetCARecommendations(s)
	return rpac
}

// SetNillableCARecommendations sets the "CA_Recommendations" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableCARecommendations(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetCARecommendations(*s)
	}
	return rpac
}

// SetCAUpdatedAt sets the "CA_UpdatedAt" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetCAUpdatedAt(t time.Time) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetCAUpdatedAt(t)
	return rpac
}

// SetNillableCAUpdatedAt sets the "CA_UpdatedAt" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableCAUpdatedAt(t *time.Time) *RecommendationsPMPAApplicationsCreate {
	if t != nil {
		rpac.SetCAUpdatedAt(*t)
	}
	return rpac
}

// SetCAUserName sets the "CA_UserName" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetCAUserName(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetCAUserName(s)
	return rpac
}

// SetNillableCAUserName sets the "CA_UserName" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableCAUserName(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetCAUserName(*s)
	}
	return rpac
}

// SetCARemarks sets the "CA_Remarks" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetCARemarks(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetCARemarks(s)
	return rpac
}

// SetNillableCARemarks sets the "CA_Remarks" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableCARemarks(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetCARemarks(*s)
	}
	return rpac
}

// SetNORecommendations sets the "NO_Recommendations" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNORecommendations(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetNORecommendations(s)
	return rpac
}

// SetNillableNORecommendations sets the "NO_Recommendations" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableNORecommendations(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetNORecommendations(*s)
	}
	return rpac
}

// SetNOUpdatedAt sets the "NO_UpdatedAt" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNOUpdatedAt(t time.Time) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetNOUpdatedAt(t)
	return rpac
}

// SetNillableNOUpdatedAt sets the "NO_UpdatedAt" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableNOUpdatedAt(t *time.Time) *RecommendationsPMPAApplicationsCreate {
	if t != nil {
		rpac.SetNOUpdatedAt(*t)
	}
	return rpac
}

// SetNOUserName sets the "NO_UserName" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNOUserName(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetNOUserName(s)
	return rpac
}

// SetNillableNOUserName sets the "NO_UserName" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableNOUserName(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetNOUserName(*s)
	}
	return rpac
}

// SetNORemarks sets the "NO_Remarks" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNORemarks(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetNORemarks(s)
	return rpac
}

// SetNillableNORemarks sets the "NO_Remarks" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableNORemarks(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetNORemarks(*s)
	}
	return rpac
}

// SetApplicationStatus sets the "ApplicationStatus" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetApplicationStatus(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetApplicationStatus(s)
	return rpac
}

// SetNillableApplicationStatus sets the "ApplicationStatus" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableApplicationStatus(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetApplicationStatus(*s)
	}
	return rpac
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetUpdatedAt(t time.Time) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetUpdatedAt(t)
	return rpac
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableUpdatedAt(t *time.Time) *RecommendationsPMPAApplicationsCreate {
	if t != nil {
		rpac.SetUpdatedAt(*t)
	}
	return rpac
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetUpdatedBy(s string) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetUpdatedBy(s)
	return rpac
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableUpdatedBy(s *string) *RecommendationsPMPAApplicationsCreate {
	if s != nil {
		rpac.SetUpdatedBy(*s)
	}
	return rpac
}

// SetGenerateHallTicketFlag sets the "GenerateHallTicketFlag" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetGenerateHallTicketFlag(b bool) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetGenerateHallTicketFlag(b)
	return rpac
}

// SetNillableGenerateHallTicketFlag sets the "GenerateHallTicketFlag" field if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableGenerateHallTicketFlag(b *bool) *RecommendationsPMPAApplicationsCreate {
	if b != nil {
		rpac.SetGenerateHallTicketFlag(*b)
	}
	return rpac
}

// SetID sets the "id" field.
func (rpac *RecommendationsPMPAApplicationsCreate) SetID(i int64) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetID(i)
	return rpac
}

// SetApplnRefID sets the "ApplnRef" edge to the Exam_Applications_PMPA entity by ID.
func (rpac *RecommendationsPMPAApplicationsCreate) SetApplnRefID(id int64) *RecommendationsPMPAApplicationsCreate {
	rpac.mutation.SetApplnRefID(id)
	return rpac
}

// SetNillableApplnRefID sets the "ApplnRef" edge to the Exam_Applications_PMPA entity by ID if the given value is not nil.
func (rpac *RecommendationsPMPAApplicationsCreate) SetNillableApplnRefID(id *int64) *RecommendationsPMPAApplicationsCreate {
	if id != nil {
		rpac = rpac.SetApplnRefID(*id)
	}
	return rpac
}

// SetApplnRef sets the "ApplnRef" edge to the Exam_Applications_PMPA entity.
func (rpac *RecommendationsPMPAApplicationsCreate) SetApplnRef(e *Exam_Applications_PMPA) *RecommendationsPMPAApplicationsCreate {
	return rpac.SetApplnRefID(e.ID)
}

// Mutation returns the RecommendationsPMPAApplicationsMutation object of the builder.
func (rpac *RecommendationsPMPAApplicationsCreate) Mutation() *RecommendationsPMPAApplicationsMutation {
	return rpac.mutation
}

// Save creates the RecommendationsPMPAApplications in the database.
func (rpac *RecommendationsPMPAApplicationsCreate) Save(ctx context.Context) (*RecommendationsPMPAApplications, error) {
	rpac.defaults()
	return withHooks(ctx, rpac.sqlSave, rpac.mutation, rpac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rpac *RecommendationsPMPAApplicationsCreate) SaveX(ctx context.Context) *RecommendationsPMPAApplications {
	v, err := rpac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpac *RecommendationsPMPAApplicationsCreate) Exec(ctx context.Context) error {
	_, err := rpac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpac *RecommendationsPMPAApplicationsCreate) ExecX(ctx context.Context) {
	if err := rpac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpac *RecommendationsPMPAApplicationsCreate) defaults() {
	if _, ok := rpac.mutation.UpdatedAt(); !ok {
		v := recommendationspmpaapplications.DefaultUpdatedAt()
		rpac.mutation.SetUpdatedAt(v)
	}
	if _, ok := rpac.mutation.UpdatedBy(); !ok {
		v := recommendationspmpaapplications.DefaultUpdatedBy
		rpac.mutation.SetUpdatedBy(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpac *RecommendationsPMPAApplicationsCreate) check() error {
	return nil
}

func (rpac *RecommendationsPMPAApplicationsCreate) sqlSave(ctx context.Context) (*RecommendationsPMPAApplications, error) {
	if err := rpac.check(); err != nil {
		return nil, err
	}
	_node, _spec := rpac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rpac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	rpac.mutation.id = &_node.ID
	rpac.mutation.done = true
	return _node, nil
}

func (rpac *RecommendationsPMPAApplicationsCreate) createSpec() (*RecommendationsPMPAApplications, *sqlgraph.CreateSpec) {
	var (
		_node = &RecommendationsPMPAApplications{config: rpac.config}
		_spec = sqlgraph.NewCreateSpec(recommendationspmpaapplications.Table, sqlgraph.NewFieldSpec(recommendationspmpaapplications.FieldID, field.TypeInt64))
	)
	if id, ok := rpac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rpac.mutation.ApplicationID(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldApplicationID, field.TypeInt64, value)
		_node.ApplicationID = value
	}
	if value, ok := rpac.mutation.EmployeeID(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldEmployeeID, field.TypeInt64, value)
		_node.EmployeeID = value
	}
	if value, ok := rpac.mutation.ExamNameCode(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldExamNameCode, field.TypeString, value)
		_node.ExamNameCode = value
	}
	if value, ok := rpac.mutation.ExamYear(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldExamYear, field.TypeString, value)
		_node.ExamYear = value
	}
	if value, ok := rpac.mutation.VacancyYear(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldVacancyYear, field.TypeInt32, value)
		_node.VacancyYear = value
	}
	if value, ok := rpac.mutation.CARecommendations(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldCARecommendations, field.TypeString, value)
		_node.CARecommendations = value
	}
	if value, ok := rpac.mutation.CAUpdatedAt(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldCAUpdatedAt, field.TypeTime, value)
		_node.CAUpdatedAt = value
	}
	if value, ok := rpac.mutation.CAUserName(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldCAUserName, field.TypeString, value)
		_node.CAUserName = value
	}
	if value, ok := rpac.mutation.CARemarks(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldCARemarks, field.TypeString, value)
		_node.CARemarks = value
	}
	if value, ok := rpac.mutation.NORecommendations(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldNORecommendations, field.TypeString, value)
		_node.NORecommendations = value
	}
	if value, ok := rpac.mutation.NOUpdatedAt(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldNOUpdatedAt, field.TypeTime, value)
		_node.NOUpdatedAt = value
	}
	if value, ok := rpac.mutation.NOUserName(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldNOUserName, field.TypeString, value)
		_node.NOUserName = value
	}
	if value, ok := rpac.mutation.NORemarks(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldNORemarks, field.TypeString, value)
		_node.NORemarks = value
	}
	if value, ok := rpac.mutation.ApplicationStatus(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldApplicationStatus, field.TypeString, value)
		_node.ApplicationStatus = value
	}
	if value, ok := rpac.mutation.UpdatedAt(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := rpac.mutation.UpdatedBy(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := rpac.mutation.GenerateHallTicketFlag(); ok {
		_spec.SetField(recommendationspmpaapplications.FieldGenerateHallTicketFlag, field.TypeBool, value)
		_node.GenerateHallTicketFlag = value
	}
	if nodes := rpac.mutation.ApplnRefIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   recommendationspmpaapplications.ApplnRefTable,
			Columns: []string{recommendationspmpaapplications.ApplnRefColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(exam_applications_pmpa.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.exam_applications_pmpa_pmpa_applications_ref = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RecommendationsPMPAApplicationsCreateBulk is the builder for creating many RecommendationsPMPAApplications entities in bulk.
type RecommendationsPMPAApplicationsCreateBulk struct {
	config
	builders []*RecommendationsPMPAApplicationsCreate
}

// Save creates the RecommendationsPMPAApplications entities in the database.
func (rpacb *RecommendationsPMPAApplicationsCreateBulk) Save(ctx context.Context) ([]*RecommendationsPMPAApplications, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rpacb.builders))
	nodes := make([]*RecommendationsPMPAApplications, len(rpacb.builders))
	mutators := make([]Mutator, len(rpacb.builders))
	for i := range rpacb.builders {
		func(i int, root context.Context) {
			builder := rpacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecommendationsPMPAApplicationsMutation)
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
					_, err = mutators[i+1].Mutate(root, rpacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rpacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rpacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rpacb *RecommendationsPMPAApplicationsCreateBulk) SaveX(ctx context.Context) []*RecommendationsPMPAApplications {
	v, err := rpacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpacb *RecommendationsPMPAApplicationsCreateBulk) Exec(ctx context.Context) error {
	_, err := rpacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpacb *RecommendationsPMPAApplicationsCreateBulk) ExecX(ctx context.Context) {
	if err := rpacb.Exec(ctx); err != nil {
		panic(err)
	}
}
