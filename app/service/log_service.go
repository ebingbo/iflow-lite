package service

import (
	"context"
	"strconv"

	"iflow-lite/dao"
	"iflow-lite/type/dto"
	"iflow-lite/type/input"
	"iflow-lite/type/model"
	"iflow-lite/type/output"
)

var DefaultLogService = NewLogService()

type LogService struct {
}

func NewLogService() *LogService {
	return &LogService{}
}

func (this *LogService) LogGet(ctx context.Context, in *input.LogGetInput) (interface{}, error) {
	return dao.DefaultLogDao.LogGet(ctx, in.ID)
}

func (this *LogService) LogAdd(ctx context.Context, in *input.LogAddInput) (interface{}, error) {
	m := model.NewLogBuilder().
		ProcessID(in.ProcessID).
		ProcessCode(in.ProcessCode).
		ExecutionID(in.ExecutionID).
		NodeID(in.NodeID).
		NodeCode(in.NodeCode).
		TaskID(in.TaskID).
		AssigneeID(in.AssigneeID).
		Action(in.Action).
		Remark(in.Remark).
		Build()
	if _, err := dao.DefaultLogDao.LogAdd(ctx, m); err != nil {
		return nil, err
	}
	return m, nil
}

func (this *LogService) LogQuery(ctx context.Context, in *input.LogQueryInput) (interface{}, error) {
	result := new(output.LogQueryOutput)
	cond := map[string]interface{}{}
	if in.ProcessID > 0 {
		cond["process_id"] = in.ProcessID
	}
	if in.ProcessCode != "" {
		cond["process_code"] = in.ProcessCode
	}
	if in.ExecutionID > 0 {
		cond["execution_id"] = in.ExecutionID
	}
	if in.NodeID > 0 {
		cond["node_id"] = in.NodeID
	}
	if in.NodeCode != "" {
		cond["node_code"] = in.NodeCode
	}
	if in.TaskID > 0 {
		cond["task_id"] = in.TaskID
	}
	if in.AssigneeID != "" {
		cond["assignee_id"] = in.AssigneeID
	}
	if in.Page <= 0 {
		in.Page = 1
	}
	if in.Size <= 0 {
		in.Size = 10
	}
	items, total, err := dao.DefaultLogDao.LogQuery(ctx, cond, in.Page, in.Size)
	if err != nil {
		return nil, err
	}

	ids := make([]uint64, 0, len(items))
	seen := make(map[uint64]struct{}, len(items))
	for _, item := range items {
		if item.AssigneeID == "" {
			continue
		}
		id, parseErr := strconv.ParseUint(item.AssigneeID, 10, 64)
		if parseErr != nil || id == 0 {
			continue
		}
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		ids = append(ids, id)
	}

	userMap := map[uint64]*model.User{}
	if len(ids) > 0 {
		users, userErr := dao.DefaultUserDao.UserListByIDs(ctx, ids)
		if userErr != nil {
			return nil, userErr
		}
		userMap = make(map[uint64]*model.User, len(users))
		for _, user := range users {
			userMap[user.ID] = user
		}
	}

	logItems := make([]*dto.Log, 0, len(items))
	for _, item := range items {
		logItem := &dto.Log{Log: item}
		if item.AssigneeID != "" {
			if id, parseErr := strconv.ParseUint(item.AssigneeID, 10, 64); parseErr == nil && id > 0 {
				if user, ok := userMap[id]; ok {
					logItem.AssigneeName = user.Name
				}
			}
		}
		logItems = append(logItems, logItem)
	}

	result.Items = logItems
	result.Total = total
	return result, nil
}
