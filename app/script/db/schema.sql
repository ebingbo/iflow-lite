create database if not exists iflow_db character set utf8mb4 collate utf8mb4_unicode_ci;
-- ===============================================
-- 1. 用户体系
-- ===============================================

use iflow_db;
truncate table process;
truncate table node;
truncate table transition;
truncate table execution;
truncate table task;
truncate table assignment;
truncate table log;

-- 角色表
create table if not exists role
(
    id         bigint auto_increment primary key comment '角色ID',
    name       varchar(50) not null default '' comment '角色名称',
    code       varchar(50) not null default '' comment '角色编码',
    created_at datetime    not null default current_timestamp comment '创建时间',
    updated_at datetime    not null default current_timestamp on update current_timestamp comment '更新时间',
    unique key uk_code (code)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '用户角色表';

-- 用户-角色关系
create table if not exists user_role
(
    id         bigint auto_increment primary key comment '用户角色关系ID',
    user_id    bigint   not null default 0 comment '用户ID',
    role_id    bigint   not null default 0 comment '角色ID',
    created_at datetime not null default current_timestamp comment '创建时间',
    updated_at datetime not null default current_timestamp on update current_timestamp comment '更新时间',
    unique key uk_user_role (user_id, role_id)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '用户-角色映射表';

-- 系统用户
create table if not exists user
(
    id         bigint auto_increment primary key comment '用户ID',
    name       varchar(50)  not null default '' comment '用户姓名',
    email      varchar(100) not null default '' comment '邮箱',
    password   varchar(256) not null default '' comment '密码',
    status     tinyint      not null default 1 comment '状态：1=启用,0=禁用',
    created_at datetime     not null default current_timestamp comment '创建时间',
    updated_at datetime     not null default current_timestamp on update current_timestamp comment '更新时间',
    unique key uk_email (email)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '系统用户表';

-- 流程定义
create table if not exists process
(
    id          bigint auto_increment primary key comment '流程ID',
    name        varchar(100) not null default '' comment '流程名称',
    code        varchar(100) not null default '' comment '流程编码',
    description text comment '描述',
    status      tinyint      not null default 0 comment '状态：1=启用,0=禁用',
    created_by  bigint       not null default 0 comment '创建人',
    updated_by  bigint       not null default 0 comment '更新人',
    created_at  datetime     not null default current_timestamp comment '创建时间',
    updated_at  datetime     not null default current_timestamp on update current_timestamp comment '更新时间',
    unique key uk_code (code)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '流程定义表';

-- 节点定义
create table if not exists node
(
    id           bigint auto_increment primary key comment '节点ID',
    process_id   bigint       not null default 0 comment '所属流程ID',
    process_code varchar(100) not null default '' comment '所属流程编码',
    tag          varchar(100) not null default '' comment '节点标签',
    name         varchar(100) not null default '' comment '节点名称',
    code         varchar(100) not null default '' comment '节点编码',
    type         varchar(50)  not null default 'start' comment '节点类型: start, end, join, user_task, service_task, exclusive_gateway, parallel_gateway',
    description  text comment '描述',
    created_by   bigint       not null default 0 comment '创建人',
    updated_by   bigint       not null default 0 comment '更新人',
    created_at   datetime     not null default current_timestamp comment '创建时间',
    updated_at   datetime     not null default current_timestamp on update current_timestamp comment '更新时间'
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '节点定义表';

-- 节点流转
create table if not exists transition
(
    id           bigint auto_increment primary key comment '流转ID',
    process_id   bigint   not null default 0 comment '所属流程ID',
    from_node_id bigint   not null default 0 comment '源节点ID',
    to_node_id   bigint   not null default 0 comment '目标节点ID',
    created_at   datetime not null default current_timestamp comment '创建时间',
    updated_at   datetime not null default current_timestamp on update current_timestamp comment '更新时间',
    index idx_process_from_node_id (process_id, from_node_id),
    index idx_process_to_node_id (process_id, to_node_id)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '流转规则表';

-- 流程实例
create table if not exists execution
(
    id            bigint auto_increment primary key comment '流程实例ID',
    process_id    bigint        not null default 0 comment '流程ID',
    process_code  varchar(100)  not null default '' comment '流程编码',
    process_name  varchar(100)  not null default '' comment '流程名称',
    business_key  varchar(100)  not null default '' comment '外部业务ID',
    business_type varchar(50)   not null default '' comment '业务类型',
    status        varchar(20)   not null default 'running' comment '流程实例状态：running, completed, terminated',
    progress      decimal(5, 2) not null default 0.00 comment '流程进度百分比，0~100',
    created_by    varchar(100)  not null default '' comment '流程发起人ID',
    created_at    datetime      not null default current_timestamp comment '创建时间',
    updated_at    datetime      not null default current_timestamp on update current_timestamp comment '更新时间',
    index idx_process_code (process_code)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '流程执行实例表';

-- 任务实例
create table if not exists task
(
    id           bigint auto_increment primary key comment '任务ID',
    process_id   bigint       not null default 0 comment '流程ID',
    process_code varchar(100) not null default 0 comment '流程编码',
    process_name varchar(100) not null default '' comment '流程名称',
    execution_id bigint       not null default 0 comment '流程实例ID',
    node_id      bigint       not null default 0 comment '节点ID',
    node_code    varchar(100) not null default '' comment '节点编码',
    node_name    varchar(100) not null default '' comment '任务名称即节点名称',
    assignee_id  varchar(100) not null default '' comment '任务执行人ID',
    status       varchar(20)  not null default 'pending' comment '任务状态：pending, running, completed, skipped',
    started_at   datetime     null comment '开始时间',
    ended_at     datetime     null comment '结束时间',
    description  text comment '描述',
    remark       text comment '备注',
    created_at   datetime     not null default current_timestamp comment '创建时间',
    updated_at   datetime     not null default current_timestamp on update current_timestamp comment '更新时间',
    unique key uk_execution_node (execution_id, node_code),
    index idx_execution_status (execution_id, status),
    index idx_assignee_id (assignee_id)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '任务实例表';

-- 节点分配规则
create table if not exists assignment
(
    id           bigint auto_increment primary key comment '分配规则ID',
    process_id   bigint       not null default 0 comment '流程ID',
    process_code varchar(100) not null default '' comment '流程编码',
    node_id      bigint       not null default 0 comment '节点ID',
    node_code    varchar(100) not null comment '节点编码',
    type         varchar(20)  not null default 'user' comment '分配类型：user, role',
    value        varchar(100) not null default '' comment '分配值，用户ID/角色ID',
    priority     int          not null default 0 comment '优先级，值越小优先级越高',
    strategy     varchar(20)  not null default 'sequential' comment '分配策略：sequential-顺序分配-按优先级顺序，第一条规则满足就停止, parallel-并行分配-所有规则都生成任务',
    created_at   datetime     not null default current_timestamp comment '创建时间',
    updated_at   datetime     not null default current_timestamp on update current_timestamp comment '更新时间',
    key idx_node (node_code)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '节点用户/角色分配规则表';

-- 流程日志
create table if not exists log
(
    id           bigint auto_increment primary key comment '日志ID',
    process_id   bigint       not null default 0 comment '流程ID',
    process_code varchar(100) not null default '' comment '流程编码',
    execution_id bigint       not null default 0 comment '实例ID',
    node_id      bigint       not null default 0 comment '节点ID',
    node_code    varchar(100) not null default 0 comment '节点编码',
    task_id      bigint       not null default 0 comment '任务ID',
    action       varchar(50)  not null default 'start' comment '操作类型：start, complete',
    assignee_id  varchar(100) not null default '' comment '操作人ID',
    remark       text comment '备注',
    created_at   datetime     not null default current_timestamp comment '创建时间',
    key idx_execution_id (execution_id)
) engine = InnoDB
  character set utf8mb4
  collate utf8mb4_unicode_ci comment '流程日志表';



