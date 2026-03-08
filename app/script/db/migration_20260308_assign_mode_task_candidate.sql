use iflow_db;

-- 1) node: 增加分派模式与画布坐标
alter table node
    add column assign_mode varchar(20) not null default 'single' comment '分派模式: single, candidate, multi_instance(预留)' after type,
    add column x int not null default 0 comment '画布X坐标' after assign_mode,
    add column y int not null default 0 comment '画布Y坐标' after x;

-- 2) task: assignee_id 切到 bigint，并增加 claimed_at
update task
set assignee_id = '0'
where assignee_id is null
   or assignee_id = ''
   or assignee_id not regexp '^[0-9]+$';

alter table task
    modify column assignee_id bigint not null default 0 comment '任务执行人ID，candidate模式认领前为0',
    add column claimed_at datetime null comment '认领时间(candidate模式)' after started_at;

-- 3) assignment: type/value 调整为 principal_type/principal_id
update assignment
set `value` = '0'
where `value` is null
   or `value` = ''
   or `value` not regexp '^[0-9]+$';

alter table assignment
    change column `type` principal_type varchar(20) not null default 'user' comment '主体类型: user, role',
    change column `value` principal_id bigint not null default 0 comment '主体ID: user_id/role_id';

-- 去重后再加唯一索引，避免历史重复数据导致建索引失败
delete a1
from assignment a1
         join assignment a2
              on a1.node_id = a2.node_id
                  and a1.principal_type = a2.principal_type
                  and a1.principal_id = a2.principal_id
                  and a1.id > a2.id;

alter table assignment
    add unique key uk_node_principal (node_id, principal_type, principal_id);

-- 4) candidate: 新增任务候选人快照表
create table if not exists task_candidate
(
    id          bigint auto_increment primary key comment '候选记录ID',
    task_id     bigint      not null default 0 comment '任务实例ID',
    user_id     bigint      not null default 0 comment '候选用户ID（仅用户）',
    source_type varchar(20) not null default 'user' comment '来源类型: user, role',
    source_id   bigint      not null default 0 comment '来源ID: user_id/role_id',
    created_at  datetime    not null default current_timestamp comment '创建时间',
    updated_at  datetime    not null default current_timestamp on update current_timestamp comment '更新时间',
    unique key uk_task_user (task_id, user_id),
    key idx_task (task_id),
    key idx_user (user_id)
) engine = InnoDB
  default charset = utf8mb4
  collate = utf8mb4_unicode_ci comment ='任务候选人快照表';
