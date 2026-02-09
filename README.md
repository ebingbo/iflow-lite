# iFlow Lite

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)  
[![GitHub stars](https://img.shields.io/github/stars/ebingbo/iflow-lite?style=social)](https://github.com/ebingbo/iflow-lite/stargazers)

iFlow Lite 是一个轻量级、高可扩展的**流程管理与工作流引擎**，适合快速构建内部项目流程、任务分配和执行管理。  
它从企业内部工作流项目中提取了公共、通用部分，开源给开发者社区使用。

---

## 目录

- [特性](#特性)
- [快速开始](#快速开始)
    - [克隆仓库](#克隆仓库)
    - [后端（Go）](#后端go)
    - [前端（Nuxt）](#前端nuxt)
- [分支策略](#分支策略)
- [贡献指南](#贡献指南)
- [License](#license)
- [联系与社区](#联系与社区)

---

## 特性

- **流程定义**：支持多阶段流程和节点定义（start, end, user_task, service_task, 网关等）
- **任务执行**：支持任务分配、进度跟踪、依赖资源管理
- **权限与角色**：内置用户、角色、分配规则系统
- **事件与监听**：支持流程事件监听和自定义扩展
- **日志与审计**：完整操作日志记录，方便追踪流程执行历史
- **前后端分离**：前端基于 Nuxt UI，后端使用 Go，统一代码库管理
- **开源友好**：Apache License 2.0

---

## 快速开始

### 克隆仓库

```bash
git clone https://github.com/ebingbo/iflow-lite.git
cd iflow-lite
```

