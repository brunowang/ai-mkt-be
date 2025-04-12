package data

import (
	"ai-mkt-be/internal/lib"
	"context"
	"github.com/pkg/errors"
	"time"

	"ai-mkt-be/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type planRepo struct {
	data *Data
	log  *log.Helper
}

// NewPlanRepo .
func NewPlanRepo(data *Data, logger log.Logger) biz.PlanRepo {
	return &planRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *planRepo) Save(ctx context.Context, data *biz.Plan) error {
	nowt := time.Now()
	if !data.CreateTime.Valid {
		data.CreateTime.Scan(nowt)
	}
	if !data.UpdateTime.Valid {
		data.UpdateTime.Scan(nowt)
	}
	if err := r.data.db.Create(data).Error; err != nil {
		r.log.WithContext(ctx).Errorf("planRepo: save data failed, err: %v", errors.WithStack(err))
		return err
	}
	return nil
}

func (r *planRepo) Update(ctx context.Context, planID string, data *biz.Plan) error {
	where := &biz.Plan{
		PlanID: planID,
	}
	if !data.UpdateTime.Valid {
		data.UpdateTime.Scan(time.Now())
	}
	if err := r.data.db.Where(where).Updates(data).Error; err != nil {
		r.log.WithContext(ctx).Errorf("planRepo: update data failed, err: %v", errors.WithStack(err))
		return err
	}
	return nil
}

func (r *planRepo) FindByID(ctx context.Context, planID string) (*biz.Plan, error) {
	var data biz.Plan
	tm, err := lib.DecodeTime(planID)
	if err != nil {
		r.log.WithContext(ctx).Errorf("planRepo: decode time error: %v", errors.WithStack(err))
		return nil, err
	}
	table := &biz.Plan{}
	if err := r.data.db.WithContext(ctx).Table(table.TableNameSplit(tm)).
		Where("plan_id = ?", planID).Take(&data).Error; err != nil {
		r.log.WithContext(ctx).Errorf("planRepo: take data failed, err: %v", errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return &data, nil
}

func (r *planRepo) ListByUser(ctx context.Context, userID string) ([]*biz.Plan, error) {
	var list []*biz.Plan
	nowt := time.Now()
	table := &biz.Plan{}
	if err := r.data.db.WithContext(ctx).Table(table.TableNameSplit(nowt)).
		Where("user_id = ?", userID).Find(&list).Error; err != nil {
		r.log.WithContext(ctx).Errorf("planRepo: find list failed, err: %v", errors.WithStack(err))
		return nil, errors.WithStack(err)
	}
	return list, nil
}

func (r *planRepo) ListAll(context.Context) ([]*biz.Plan, error) {
	return nil, nil
}
