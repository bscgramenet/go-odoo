package odoo

// ProjectTask represents project.task model.
type ProjectTask struct {
	LastUpdate                   *Time       `xmlrpc:"__last_update,omitempty"`
	AccessToken                  *String     `xmlrpc:"access_token,omitempty"`
	AccessUrl                    *String     `xmlrpc:"access_url,omitempty"`
	AccessWarning                *String     `xmlrpc:"access_warning,omitempty"`
	Active                       *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId      *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline         *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration  *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon        *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                  *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary              *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon             *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId               *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId               *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AllowMilestones              *Bool       `xmlrpc:"allow_milestones,omitempty"`
	AllowRecurringTasks          *Bool       `xmlrpc:"allow_recurring_tasks,omitempty"`
	AllowSubtasks                *Bool       `xmlrpc:"allow_subtasks,omitempty"`
	AllowTaskDependencies        *Bool       `xmlrpc:"allow_task_dependencies,omitempty"`
	AnalyticAccountId            *Many2One   `xmlrpc:"analytic_account_id,omitempty"`
	AncestorId                   *Many2One   `xmlrpc:"ancestor_id,omitempty"`
	AttachmentIds                *Relation   `xmlrpc:"attachment_ids,omitempty"`
	ChildIds                     *Relation   `xmlrpc:"child_ids,omitempty"`
	ChildText                    *String     `xmlrpc:"child_text,omitempty"`
	Color                        *Int        `xmlrpc:"color,omitempty"`
	CommercialPartnerId          *Many2One   `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyId                    *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                   *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                    *Many2One   `xmlrpc:"create_uid,omitempty"`
	DateAssign                   *Time       `xmlrpc:"date_assign,omitempty"`
	DateDeadline                 *Time       `xmlrpc:"date_deadline,omitempty"`
	DateEnd                      *Time       `xmlrpc:"date_end,omitempty"`
	DateLastStageUpdate          *Time       `xmlrpc:"date_last_stage_update,omitempty"`
	DependOnIds                  *Relation   `xmlrpc:"depend_on_ids,omitempty"`
	DependentIds                 *Relation   `xmlrpc:"dependent_ids,omitempty"`
	DependentTasksCount          *Int        `xmlrpc:"dependent_tasks_count,omitempty"`
	Description                  *String     `xmlrpc:"description,omitempty"`
	DisplayName                  *String     `xmlrpc:"display_name,omitempty"`
	DisplayParentTaskButton      *Bool       `xmlrpc:"display_parent_task_button,omitempty"`
	DisplayProjectId             *Many2One   `xmlrpc:"display_project_id,omitempty"`
	DisplaySaleOrderButton       *Bool       `xmlrpc:"display_sale_order_button,omitempty"`
	DisplayedImageId             *Many2One   `xmlrpc:"displayed_image_id,omitempty"`
	EmailCc                      *String     `xmlrpc:"email_cc,omitempty"`
	EmailFrom                    *String     `xmlrpc:"email_from,omitempty"`
	Fri                          *Bool       `xmlrpc:"fri,omitempty"`
	HasLateAndUnreachedMilestone *Bool       `xmlrpc:"has_late_and_unreached_milestone,omitempty"`
	HasMessage                   *Bool       `xmlrpc:"has_message,omitempty"`
	Id                           *Int        `xmlrpc:"id,omitempty"`
	IsAnalyticAccountIdChanged   *Bool       `xmlrpc:"is_analytic_account_id_changed,omitempty"`
	IsBlocked                    *Bool       `xmlrpc:"is_blocked,omitempty"`
	IsClosed                     *Bool       `xmlrpc:"is_closed,omitempty"`
	IsPrivate                    *Bool       `xmlrpc:"is_private,omitempty"`
	KanbanState                  *Selection  `xmlrpc:"kanban_state,omitempty"`
	KanbanStateLabel             *String     `xmlrpc:"kanban_state_label,omitempty"`
	LegendBlocked                *String     `xmlrpc:"legend_blocked,omitempty"`
	LegendDone                   *String     `xmlrpc:"legend_done,omitempty"`
	LegendNormal                 *String     `xmlrpc:"legend_normal,omitempty"`
	ManagerId                    *Many2One   `xmlrpc:"manager_id,omitempty"`
	MessageAttachmentCount       *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds           *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError              *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter       *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError           *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                   *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower            *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId      *Many2One   `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction            *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter     *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds            *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MilestoneId                  *Many2One   `xmlrpc:"milestone_id,omitempty"`
	Mon                          *Bool       `xmlrpc:"mon,omitempty"`
	MyActivityDateDeadline       *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                         *String     `xmlrpc:"name,omitempty"`
	ParentId                     *Many2One   `xmlrpc:"parent_id,omitempty"`
	PartnerCity                  *String     `xmlrpc:"partner_city,omitempty"`
	PartnerEmail                 *String     `xmlrpc:"partner_email,omitempty"`
	PartnerId                    *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerIsCompany             *Bool       `xmlrpc:"partner_is_company,omitempty"`
	PartnerPhone                 *String     `xmlrpc:"partner_phone,omitempty"`
	PersonalStageId              *Many2One   `xmlrpc:"personal_stage_id,omitempty"`
	PersonalStageTypeId          *Many2One   `xmlrpc:"personal_stage_type_id,omitempty"`
	PersonalStageTypeIds         *Relation   `xmlrpc:"personal_stage_type_ids,omitempty"`
	PlannedHours                 *Float      `xmlrpc:"planned_hours,omitempty"`
	PortalUserNames              *String     `xmlrpc:"portal_user_names,omitempty"`
	Priority                     *Selection  `xmlrpc:"priority,omitempty"`
	ProjectAnalyticAccountId     *Many2One   `xmlrpc:"project_analytic_account_id,omitempty"`
	ProjectColor                 *Int        `xmlrpc:"project_color,omitempty"`
	ProjectId                    *Many2One   `xmlrpc:"project_id,omitempty"`
	ProjectPrivacyVisibility     *Selection  `xmlrpc:"project_privacy_visibility,omitempty"`
	ProjectSaleOrderId           *Many2One   `xmlrpc:"project_sale_order_id,omitempty"`
	RatingActive                 *Bool       `xmlrpc:"rating_active,omitempty"`
	RatingAvg                    *Float      `xmlrpc:"rating_avg,omitempty"`
	RatingAvgText                *Selection  `xmlrpc:"rating_avg_text,omitempty"`
	RatingCount                  *Int        `xmlrpc:"rating_count,omitempty"`
	RatingIds                    *Relation   `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback           *String     `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage              *String     `xmlrpc:"rating_last_image,omitempty"`
	RatingLastText               *Selection  `xmlrpc:"rating_last_text,omitempty"`
	RatingLastValue              *Float      `xmlrpc:"rating_last_value,omitempty"`
	RatingPercentageSatisfaction *Float      `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	RecurrenceId                 *Many2One   `xmlrpc:"recurrence_id,omitempty"`
	RecurrenceMessage            *String     `xmlrpc:"recurrence_message,omitempty"`
	RecurrenceUpdate             *Selection  `xmlrpc:"recurrence_update,omitempty"`
	RecurringCount               *Int        `xmlrpc:"recurring_count,omitempty"`
	RecurringTask                *Bool       `xmlrpc:"recurring_task,omitempty"`
	RepeatDay                    *Selection  `xmlrpc:"repeat_day,omitempty"`
	RepeatInterval               *Int        `xmlrpc:"repeat_interval,omitempty"`
	RepeatMonth                  *Selection  `xmlrpc:"repeat_month,omitempty"`
	RepeatNumber                 *Int        `xmlrpc:"repeat_number,omitempty"`
	RepeatOnMonth                *Selection  `xmlrpc:"repeat_on_month,omitempty"`
	RepeatOnYear                 *Selection  `xmlrpc:"repeat_on_year,omitempty"`
	RepeatShowDay                *Bool       `xmlrpc:"repeat_show_day,omitempty"`
	RepeatShowDow                *Bool       `xmlrpc:"repeat_show_dow,omitempty"`
	RepeatShowMonth              *Bool       `xmlrpc:"repeat_show_month,omitempty"`
	RepeatShowWeek               *Bool       `xmlrpc:"repeat_show_week,omitempty"`
	RepeatType                   *Selection  `xmlrpc:"repeat_type,omitempty"`
	RepeatUnit                   *Selection  `xmlrpc:"repeat_unit,omitempty"`
	RepeatUntil                  *Time       `xmlrpc:"repeat_until,omitempty"`
	RepeatWeek                   *Selection  `xmlrpc:"repeat_week,omitempty"`
	RepeatWeekday                *Selection  `xmlrpc:"repeat_weekday,omitempty"`
	SaleLineId                   *Many2One   `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderId                  *Many2One   `xmlrpc:"sale_order_id,omitempty"`
	Sat                          *Bool       `xmlrpc:"sat,omitempty"`
	Sequence                     *Int        `xmlrpc:"sequence,omitempty"`
	StageId                      *Many2One   `xmlrpc:"stage_id,omitempty"`
	SubtaskCount                 *Int        `xmlrpc:"subtask_count,omitempty"`
	SubtaskPlannedHours          *Float      `xmlrpc:"subtask_planned_hours,omitempty"`
	Sun                          *Bool       `xmlrpc:"sun,omitempty"`
	TagIds                       *Relation   `xmlrpc:"tag_ids,omitempty"`
	TaskProperties               interface{} `xmlrpc:"task_properties,omitempty"`
	TaskToInvoice                *Bool       `xmlrpc:"task_to_invoice,omitempty"`
	Thu                          *Bool       `xmlrpc:"thu,omitempty"`
	Tue                          *Bool       `xmlrpc:"tue,omitempty"`
	UserIds                      *Relation   `xmlrpc:"user_ids,omitempty"`
	WebsiteMessageIds            *Relation   `xmlrpc:"website_message_ids,omitempty"`
	Wed                          *Bool       `xmlrpc:"wed,omitempty"`
	WorkingDaysClose             *Float      `xmlrpc:"working_days_close,omitempty"`
	WorkingDaysOpen              *Float      `xmlrpc:"working_days_open,omitempty"`
	WorkingHoursClose            *Float      `xmlrpc:"working_hours_close,omitempty"`
	WorkingHoursOpen             *Float      `xmlrpc:"working_hours_open,omitempty"`
	WriteDate                    *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                     *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// ProjectTasks represents array of project.task model.
type ProjectTasks []ProjectTask

// ProjectTaskModel is the odoo model name.
const ProjectTaskModel = "project.task"

// Many2One convert ProjectTask to *Many2One.
func (pt *ProjectTask) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProjectTask creates a new project.task model and returns its id.
func (c *Client) CreateProjectTask(pt *ProjectTask) (int64, error) {
	ids, err := c.CreateProjectTasks([]*ProjectTask{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTask creates a new project.task model and returns its id.
func (c *Client) CreateProjectTasks(pts []*ProjectTask) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskModel, vv, nil)
}

// UpdateProjectTask updates an existing project.task record.
func (c *Client) UpdateProjectTask(pt *ProjectTask) error {
	return c.UpdateProjectTasks([]int64{pt.Id.Get()}, pt)
}

// UpdateProjectTasks updates existing project.task records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProjectTasks(ids []int64, pt *ProjectTask) error {
	return c.Update(ProjectTaskModel, ids, pt, nil)
}

// DeleteProjectTask deletes an existing project.task record.
func (c *Client) DeleteProjectTask(id int64) error {
	return c.DeleteProjectTasks([]int64{id})
}

// DeleteProjectTasks deletes existing project.task records.
func (c *Client) DeleteProjectTasks(ids []int64) error {
	return c.Delete(ProjectTaskModel, ids)
}

// GetProjectTask gets project.task existing record.
func (c *Client) GetProjectTask(id int64) (*ProjectTask, error) {
	pts, err := c.GetProjectTasks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// GetProjectTasks gets project.task existing records.
func (c *Client) GetProjectTasks(ids []int64) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.Read(ProjectTaskModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTask finds project.task record by querying it with criteria.
func (c *Client) FindProjectTask(criteria *Criteria) (*ProjectTask, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// FindProjectTasks finds project.task records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTasks(criteria *Criteria, options *Options) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTaskIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTaskModel, criteria, options)
}

// FindProjectTaskId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
