-- 添加用户
INSERT INTO workflow_db.user (id, name, email, password, status) VALUES (1, '吴作顺', '123456789@163.com', '79d746398f26e0d903f4069946fdd79dad053a26ad69578a76e14ff8630fd93e', 1);

-- 添加角色
INSERT INTO workflow_db.role (id, name, code) VALUES (1, '管理员', 'admin');

-- 添加用户角色
INSERT INTO workflow_db.user_role (id, user_id, role_id) VALUES (1, 1, 1);

-- 添加流程
INSERT INTO workflow_db.process (id, name, code, description, status) VALUES (1, '信创适配默认流程', 'adapter-process', '信创适配默认流程', 1);

-- 添加阶段
INSERT INTO workflow_db.stage (id, process_id, name, code, sequence, description) VALUES (1, 1, '咨询规划', 'consultation', 1, '咨询规划是项目管理流程的初始阶段，核心是通过需求分析、可行性研究和方案设计，为项目制定清晰目标、范围及实施路径。该步骤需明确客户需求、评估资源可行性，并输出包含时间、预算、风险应对策略的规划文档，为后续执行提供基准。其关键作用在于确保项目方向与客户预期一致，并具备可操作性。');
INSERT INTO workflow_db.stage (id, process_id, name, code, sequence, description) VALUES (2, 1, '需求分析', 'requirement', 2, '需求分析是项目管理流程中的关键步骤，其核心是通过系统化的方法明确并验证项目目标、范围及干系人期望。该阶段需开展需求调研（如访谈、问卷）、需求整理（分类、优先级排序）和需求确认（形成需求文档并获干系人批准），确保项目交付物与业务目标一致，同时为后续设计、开发提供明确依据。此步骤直接影响项目成功，需注重需求的可追溯性和变更管理。');
INSERT INTO workflow_db.stage (id, process_id, name, code, sequence, description) VALUES (3, 1, '方案设计', 'plan', 3, '方案设计步骤是项目管理中承上启下的关键环节，需将需求分析成果转化为可执行的技术方案和资源计划，包括系统架构设计、任务分解、时间线规划及风险评估，最终形成详细的项目蓝图和验收标准。');
INSERT INTO workflow_db.stage (id, process_id, name, code, sequence, description) VALUES (4, 1, '适配实施', 'implementation', 4, '适配实施是项目管理中将设计方案转化为实际成果的关键阶段，需协调资源、监控进度并确保质量符合标准。该步骤包括设备采购、安装调试、系统集成及验收测试，需严格遵循技术规范和安全要求，同时处理变更请求和风险问题‌。');
INSERT INTO workflow_db.stage (id, process_id, name, code, sequence, description) VALUES (5, 1, '适配测试', 'test', 5, '适配测试是项目管理中确保软件与目标环境兼容的关键步骤，需明确测试范围（如硬件、操作系统、浏览器等）并制定详细计划‌。');
INSERT INTO workflow_db.stage (id, process_id, name, code, sequence, description) VALUES (6, 1, '数据迁移演练', 'drill', 6, '数据迁移演练是项目管理中验证迁移方案可行性的关键阶段，需通过模拟测试环境验证数据完整性、系统兼容性及回退机制，并记录测试结果以优化正式迁移计划。该阶段通常包括数据样本迁移、性能压力测试及错误处理流程的实操验证，确保迁移风险可控‌。');
INSERT INTO workflow_db.stage (id, process_id, name, code, sequence, description) VALUES (7, 1, '上线切换', 'switch', 7, '上线切换是项目管理中确保系统从旧环境平稳过渡到新环境的关键阶段，需协调资源、监控进度并处理突发问题。该步骤包括数据迁移、系统部署、用户培训及最终验收，需严格遵循技术规范和安全要求，同时制定回退预案以应对风险。');

-- 添加节点
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (1, 1, 0, '开始', 'start', 'start', null, null, null, null, null, 0);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (2, 1, 1, '前期调研', 'preliminary-research', 'user_task', '前期调研阶段是项目启动前通过系统分析明确需求、评估可行性的关键准备环节，为后续决策提供依据。', null, null, null, null, 1);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (3, 1, 1, '改造策略', 'transformation-strategy', 'user_task', '改造策略阶段是项目管理中针对现状问题制定系统性改进方案的关键环节，需结合前期调研成果明确改造目标、技术路径及资源分配，为后续实施提供清晰指导。', null, null, null, null, 2);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (4, 1, 1, '成本评估', 'cost-assessment', 'user_task', '成本评估阶段是项目管理中通过量化分析确定项目预算和资源分配的关键环节，需综合人力、设备、时间等要素进行精确测算，为决策提供财务依据。', null, null, null, null, 3);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (5, 1, 1, '项目计划', 'project-plan', 'user_task', '项目计划阶段是通过制定范围、时间、成本等子计划，将项目目标转化为可执行路径的核心环节‌。', null, null, null, null, 4);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (6, 1, 0, '第一阶段完成', 'join-1', 'join', null, null, null, null, null, 5);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (7, 1, 2, '应用适配需求', 'application-adaptation', 'user_task', '应用适配需求阶段是项目管理中确保项目成果（如软件、系统）与目标运行环境（如不同设备、操作系统、网络条件）兼容性要求相匹配的关键环节，需通过测试验证功能、性能及用户体验的达标情况‌。', null, null, null, null, 6);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (8, 1, 2, '数据迁移需求', 'data-migration-demand', 'user_task', '数据迁移需求阶段是项目管理中通过评估现有数据、分析业务流程及目标系统要求，明确迁移范围、技术路径和风险控制措施的关键准备环节‌。', null, null, null, null, 7);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (9, 1, 2, '文件迁移需求', 'file-migration', 'user_task', '文件迁移需求阶段是项目管理中通过评估现有文件系统、明确迁移范围及技术标准，为后续数据转移提供依据的关键环节。', null, null, null, null, 8);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (10, 1, 0, '第二阶段完成', 'join-2', 'join', null, null, null, null, null, 9);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (11, 1, 3, '应用适配方案', 'application-adaptation-plan', 'user_task', '应用适配方案阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标环境兼容的关键环节。', null, null, null, null, 10);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (12, 1, 3, '数据迁移方案', 'data-migration-plan', 'user_task', '数据迁移方案阶段是项目管理中通过制定详细的迁移策略、流程和风险应对措施，确保数据从源系统安全、完整地迁移到目标系统的关键环节‌。', null, null, null, null, 11);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (13, 1, 3, '文件迁移方案', 'file-migration-plan', 'user_task', '文件迁移方案阶段是项目管理中通过制定详细的迁移策略、流程和风险应对措施，确保文件从源系统安全、完整地迁移到目标系统的关键环节。', null, null, null, null, 12);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (14, 1, 3, '安全方案', 'security-plan', 'user_task', '安全方案阶段是项目管理中通过制定详细的安全策略、流程和风险应对措施，确保项目成果在目标环境中安全、稳定运行的关键环节。', null, null, null, null, 13);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (15, 1, 3, '密码方案', 'password-plan', 'user_task', '密码方案阶段是项目管理中通过制定加密策略、密钥管理及安全协议，确保项目数据机密性、完整性和可用性的关键环节‌。', null, null, null, null, 14);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (16, 1, 0, '第三阶段完成', 'join-3', 'join', null, null, null, null, null, 15);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (17, 1, 4, '适配环境', 'adapt-environment', 'user_task', '适配环境阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标环境兼容的关键环节‌。', null, null, null, null, 16);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (18, 1, 4, 'OS适配', 'os-adapt', 'user_task', 'OS适配阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标操作系统兼容的关键环节‌。', null, null, null, null, 17);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (19, 1, 4, '中间件适配', 'middleware-adapt', 'user_task', '中间件适配阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标中间件环境兼容的关键环节‌。', null, null, null, null, 18);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (20, 1, 4, '数据库适配', 'database-adapt', 'user_task', '数据库适配阶段是项目管理中通过技术评估和资源调配，确保数据库系统与目标环境（如操作系统、硬件架构、安全标准等）兼容的关键环节。', null, null, null, null, 19);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (21, 1, 4, '终端适配', 'terminal-adapt', 'user_task', '终端适配阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标终端设备兼容的关键环节‌。', null, null, null, null, 20);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (22, 1, 4, '插件适配', 'plugin-adapt', 'user_task', '插件适配阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标插件环境兼容的关键环节。', null, null, null, null, 21);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (23, 1, 4, '应用改造', 'application-transformation', 'user_task', '应用改造阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标应用环境兼容的关键环节。', null, null, null, null, 22);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (24, 1, 4, '性能优化', 'performance-optimization', 'user_task', '性能优化阶段是项目管理中通过技术评估和资源调配，确保项目成果达到预期性能标准的关键环节‌。', null, null, null, null, 23);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (25, 1, 4, '单元测试', 'unit-testing', 'user_task', '单元测试阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标环境兼容的关键环节。', null, null, null, null, 24);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (26, 1, 0, '第四阶段完成', 'join-4', 'join', null, null, null, null, null, 25);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (27, 1, 5, '系统测试', 'system-testing', 'user_task', '系统测试阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标系统环境兼容的关键环节。‌', null, null, null, null, 26);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (28, 1, 5, '性能测试', 'performance-testing', 'user_task', '性能测试阶段是项目管理中通过模拟用户负载、监控系统指标并分析数据，以评估和优化应用程序性能的关键环节‌。', null, null, null, null, 27);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (29, 1, 5, '联调测试', 'joint-debug-testing', 'user_task', '联调测试是项目管理中通过联合调试已通过单元测试的模块或系统，验证接口兼容性和协同工作能力的关键环节‌。', null, null, null, null, 28);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (30, 1, 5, '安全测试', 'security-testing', 'user_task', '安全测试阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标安全环境兼容的关键环节。', null, null, null, null, 29);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (31, 1, 5, '业务测试', 'service-testing', 'user_task', '业务测试阶段是项目管理中通过模拟真实业务场景和用户操作流程，验证系统功能是否符合实际业务需求的关键环节‌。', null, null, null, null, 30);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (32, 1, 0, '第五阶段完成', 'join-5', 'join', null, null, null, null, null, 31);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (33, 1, 6, '环境准备', 'environmental-preparation', 'user_task', '环境准备阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标运行环境兼容的关键环节。', null, null, null, null, 32);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (34, 1, 6, '脚本编制', 'script-compilation', 'user_task', '脚本编制阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标脚本环境兼容的关键环节。', null, null, null, null, 33);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (35, 1, 6, '数据脱敏', 'data-desensitization', 'user_task', '数据脱敏阶段是项目管理中通过技术手段对敏感数据进行变形或加密处理，以保护隐私和符合法规要求的关键环节‌。', null, null, null, null, 34);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (36, 1, 6, '效率优化', 'efficiency-optimization', 'user_task', '效率优化阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标性能标准匹配的关键环节。', null, null, null, null, 35);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (37, 1, 6, '数据库调优', 'database-tuning', 'user_task', '数据库调优阶段是项目管理中通过SQL优化、索引调整、缓存配置及架构升级（如读写分离、分库分表）等手段，系统性提升数据库性能的关键环节‌。', null, null, null, null, 36);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (38, 1, 6, '数据比对', 'data-comparison', 'user_task', '数据比对阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标数据环境兼容的关键环节。', null, null, null, null, 37);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (39, 1, 6, '迁移测试', 'migration-testing', 'user_task', '迁移测试阶段是项目管理中通过模拟真实环境验证数据、系统及功能迁移后完整性与兼容性的关键环节‌。', null, null, null, null, 38);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (40, 1, 6, '时间估算', 'time-estimation', 'user_task', '时间估算阶段是项目管理中通过运用专业工具和技术（如甘特图、PERT三点估算等），预测并确定项目各活动及整体所需持续时间的关键环节‌。', null, null, null, null, 39);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (41, 1, 6, '业务验证', 'business-verification', 'user_task', '业务验证阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标业务环境兼容的关键环节。', null, null, null, null, 40);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (42, 1, 0, '第六阶段完成', 'join-6', 'join', null, null, null, null, null, 41);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (43, 1, 7, '环境准备', 'environmental-of-preparation', 'user_task', '环境准备阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标运行环境兼容的关键环节。', null, null, null, null, 42);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (44, 1, 7, '切换窗口', 'switch-windows', 'user_task', '切换窗口阶段是项目管理中协调资源、监控进度并处理突发问题，确保系统从旧环境平稳过渡到新环境的关键环节。', null, null, null, null, 43);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (45, 1, 7, '程序部署', 'program-deployment', 'user_task', '程序部署阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标运行环境兼容的关键环节。', null, null, null, null, 44);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (46, 1, 7, '数据迁移', 'data-migration', 'user_task', '数据迁移阶段是项目管理中将数据从源系统安全、完整地转移到目标系统，并确保其准确性和可用性的关键环节‌。', null, null, null, null, 45);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (47, 1, 7, '上线切换', 'switch-online', 'user_task', '上线切换阶段是项目管理中通过协调资源、监控进度并处理突发问题，确保系统从旧环境平稳过渡到新环境的关键环节‌。', null, null, null, null, 46);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (48, 1, 7, '试运行重保', 'trial-operation', 'user_task', '试运行重保阶段是项目管理中通过技术评估和资源调配，确保项目成果与目标运行环境兼容的关键环节。', null, null, null, null, 47);
INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (49, 1, 7, '数据补正', 'data-correction', 'user_task', '数据补正阶段是项目管理中通过技术手段对迁移或测试后的数据进行校验、修正和标准化处理，确保数据完整性和准确性的关键环节。', null, null, null, null, 48);

INSERT INTO workflow_db.node (id, process_id, stage_id, name, code, type, description, requirements, deliverables, inputs, resources, sequence) VALUES (50, 1, 0, '结束', 'end', 'end', null, null, null, null, null, 49);

-- 流转规则
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 1, 2);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 1, 3);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 1, 4);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 1, 5);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 2, 6);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 3, 6);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 4, 6);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 5, 6);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 6, 7);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 6, 8);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 6, 9);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 7, 10);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 8, 10);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 9, 10);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 10, 11);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 10, 12);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 10, 13);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 10, 14);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 10, 15);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 11, 16);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 12, 16);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 13, 16);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 14, 16);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 15, 16);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 17);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 18);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 19);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 20);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 21);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 22);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 23);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 24);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 16, 25);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 17, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 18, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 19, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 20, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 21, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 22, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 23, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 24, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 25, 26);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 26, 27);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 26, 28);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 26, 29);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 26, 30);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 26, 31);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 27, 32);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 28, 32);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 29, 32);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 30, 32);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 31, 32);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 33);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 34);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 35);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 36);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 37);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 38);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 39);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 40);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 32, 41);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 33, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 34, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 35, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 36, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 37, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 38, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 39, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 40, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 41, 42);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 42, 43);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 42, 44);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 42, 45);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 42, 46);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 42, 47);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 42, 48);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 42, 49);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 43, 50);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 44, 50);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 45, 50);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 46, 50);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 47, 50);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 48, 50);
INSERT INTO workflow_db.transition (process_id, from_node_id, to_node_id) VALUES (1, 49, 50);

-- 添加分配规则
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (2, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (3, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (4, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (5, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (7, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (8, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (9, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (11, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (12, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (13, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (14, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (15, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (17, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (18, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (19, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (20, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (21, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (22, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (23, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (24, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (25, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (27, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (28, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (29, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (30, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (31, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (33, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (34, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (35, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (36, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (37, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (38, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (39, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (40, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (41, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (43, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (44, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (45, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (46, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (47, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (48, 'user', '1');
INSERT INTO workflow_db.assignment (node_id, type, value) VALUES (49, 'user', '1');

-- 初始化平台角色
INSERT INTO workflow_db.`role` (id, name, code, created_at, updated_at) VALUES(1, '系统管理员', 'system-admin', '2025-12-01 16:55:31', '2025-12-01 16:55:31');
INSERT INTO workflow_db.`role` (id, name, code, created_at, updated_at) VALUES(2, '系统用户', 'system-user', '2025-12-01 16:55:31', '2025-12-01 16:55:31');

-- 初始化项目角色
INSERT INTO workflow_db.project_role (id, name, code, created_at, updated_at) VALUES(1, '项目经理', 'project-manager', '2025-12-01 16:56:46', '2025-12-01 16:57:49');
INSERT INTO workflow_db.project_role (id, name, code, created_at, updated_at) VALUES(2, '项目成员', 'project-member', '2025-12-01 16:57:49', '2025-12-01 16:57:49');
INSERT INTO workflow_db.project_role (id, name, code, created_at, updated_at) VALUES(3, '项目管理员', 'project-admin', '2025-12-01 16:56:46', '2025-12-01 16:57:49');





