package biz

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Plan is a Plan model.
type Plan struct {
	ID         int64             `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PlanID     string            `gorm:"column:plan_id;index:plan_id" json:"plan_id"`
	UserID     string            `gorm:"column:user_id;index:user_id" json:"user_id"`
	PlanName   string            `gorm:"column:plan_name" json:"plan_name"`
	Images     map[string]string `gorm:"column:images;type:json;serializer:json" json:"images"`
	Prompt     string            `gorm:"column:prompt" json:"prompt"`
	Step       int               `gorm:"column:step;type:int" json:"step"`
	ClipScript string            `gorm:"column:clip_script" json:"clip_script"`
	ClipFrames map[string]string `gorm:"column:clip_frames;type:json;serializer:json" json:"clip_frames"`
	ClipVideos map[string]string `gorm:"column:clip_videos;type:json;serializer:json" json:"clip_videos"`
	Reason     string            `gorm:"column:reason;type:varchar(500)" json:"reason"`
	ExtData    map[string]string `gorm:"column:ext_data;type:json;serializer:json" json:"ext_data"`
	CreateTime sql.NullTime      `gorm:"column:create_time" json:"create_time"`
	UpdateTime sql.NullTime      `gorm:"column:update_time" json:"update_time"`
}

func (m *Plan) TableName() string {
	return m.TableNameSplit(time.Now())
}

func (m *Plan) TableNameSplit(tm time.Time) string {
	return fmt.Sprintf("%s_%s", "plan", tm.Format("200601"))
}

type PlanNext struct {
	Plan `gorm:"embedded" json:",inline"`
}

func (m *PlanNext) TableName() string {
	return m.TableNameSplit(time.Now().AddDate(0, 1, 0))
}

// PlanRepo is a Plan repo.
type PlanRepo interface {
	Save(context.Context, *Plan) error
	Update(context.Context, string, *Plan) error
	FindByID(context.Context, string) (*Plan, error)
	ListByUser(context.Context, string) ([]*Plan, error)
	ListAll(context.Context) ([]*Plan, error)
}

// PlanUsecase is a Plan usecase.
type PlanUsecase struct {
	repo PlanRepo
	log  *log.Helper
}

// NewPlanUsecase new a Plan usecase.
func NewPlanUsecase(repo PlanRepo, logger log.Logger) *PlanUsecase {
	return &PlanUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreatePlan creates a Plan, and returns the new Plan.
func (uc *PlanUsecase) CreatePlan(ctx context.Context, data *Plan) error {
	uc.log.WithContext(ctx).Infof("CreatePlan: %v", data.PlanID)
	return uc.repo.Save(ctx, data)
}

// UpdatePlan updates a Plan, and returns the edited Plan.
func (uc *PlanUsecase) UpdatePlan(ctx context.Context, planID string, data *Plan) error {
	uc.log.WithContext(ctx).Infof("UpdatePlan: %v", planID)
	return uc.repo.Update(ctx, planID, data)
}

// QueryPlan querys a Plan, and returns the specify Plan.
func (uc *PlanUsecase) QueryPlan(ctx context.Context, planID string) (*Plan, error) {
	uc.log.WithContext(ctx).Infof("QueryPlan: %v", planID)
	return uc.repo.FindByID(ctx, planID)
}

// ListPlan list Plans by userID.
func (uc *PlanUsecase) ListPlan(ctx context.Context, userID string) ([]*Plan, error) {
	uc.log.WithContext(ctx).Infof("ListPlan: %v", userID)
	return uc.repo.ListByUser(ctx, userID)
}
