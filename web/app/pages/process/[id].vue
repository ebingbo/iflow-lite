<script setup lang="ts">
import dayjs from 'dayjs'
import { type Node, type NodeType, nodeTypeConfigs } from '~/types/node'
import type { Transition } from '~/types/transition'

const route = useRoute()
const { getProcess } = useProcessApi()
const id = computed(() => Number(route.params.id))

if (!Number.isInteger(id.value) || id.value <= 0) {
  throw createError({
    statusCode: 400,
    statusMessage: '无效的流程ID'
  })
}

const { data: process, status, error } = await useAsyncData(
  `process-get-${id.value}`,
  () => getProcess(id.value),
  {
    lazy: true
  }
)

const loading = computed(() => ['idle', 'pending'].includes(status.value))

useSeoMeta({
  title: computed(() => process.value?.name ? `流程详情 - ${process.value.name}` : '流程详情'),
  description: '流程详情页面'
})

// 状态
const flowMode = ref<'node' | 'stage'>('node')
const showConfig = ref(true)
const scale = ref(1)
const offset = ref({ x: 0, y: 0 })

// 节点数据
const nodes = ref<Node[]>([
  { id: 1, x: 250, y: 50, name: '流程开始', type: 'start', assignType: null, assignTo: [] },
  { id: 2, x: 250, y: 180, name: '需求评审', type: 'user_task', assignType: 'role', assignTo: ['产品经理'] },
  { id: 3, x: 100, y: 320, name: '前端开发', type: 'user_task', assignType: 'user', assignTo: ['张三'] },
  { id: 4, x: 400, y: 320, name: '后端开发', type: 'user_task', assignType: 'user', assignTo: ['李四'] },
  { id: 5, x: 250, y: 460, name: '开发汇总', type: 'join', assignType: null, assignTo: [] },
  { id: 6, x: 250, y: 590, name: '测试验收', type: 'user_task', assignType: 'role', assignTo: ['测试工程师'] },
  { id: 7, x: 250, y: 720, name: '流程结束', type: 'end', assignType: null, assignTo: [] }
])

const edges = ref<Transition[]>([
  { from_node_id: 1, to_node_id: 2 },
  { from_node_id: 2, to_node_id: 3 },
  { from_node_id: 2, to_node_id: 4 },
  { from_node_id: 3, to_node_id: 5 },
  { from_node_id: 4, to_node_id: 5 },
  { from_node_id: 5, to_node_id: 6 },
  { from_node_id: 6, to_node_id: 7 }
])

// 计算属性
const currentNodes = computed(() => nodes.value)
const currentEdges = computed(() => edges.value)

// 交互状态
const selectedNode = ref<Node | null>(null)
const draggingNode = ref<{ node: Node, offsetX: number, offsetY: number } | null>(null)
const connectingFrom = ref<Node | null>(null)
const connectingLine = ref<{ x1: number, y1: number, x2: number, y2: number } | null>(null)
const isPanning = ref(false)
const panStart = ref({ x: 0, y: 0 })
const canvasRef = ref<HTMLDivElement | null>(null)

// ✅ 核心修改:动态计算 viewBox
const viewBox = computed(() => {
  if (!canvasRef.value || currentNodes.value.length === 0) {
    return '0 0 1000 800'
  }

  const rect = canvasRef.value.getBoundingClientRect()

  // 通过 offset 和 scale 计算当前视口在 SVG 坐标系中的位置和大小
  const viewportX = -offset.value.x / scale.value
  const viewportY = -offset.value.y / scale.value
  const viewportWidth = rect.width / scale.value
  const viewportHeight = rect.height / scale.value

  return `${viewportX} ${viewportY} ${viewportWidth} ${viewportHeight}`
})

// 工具函数
const getNodeHeight = (node: Node): number => {
  return node.type === 'user_task' && node.assignTo && node.assignTo.length > 0 ? 85 : 70
}

const getEdgePath = (edge: Transition): string => {
  const fromNode = currentNodes.value.find(n => n.id === edge.from_node_id)
  const toNode = currentNodes.value.find(n => n.id === edge.to_node_id)
  if (!fromNode || !toNode) return ''

  const x1 = fromNode.x + 80
  const y1 = fromNode.y + getNodeHeight(fromNode)
  const x2 = toNode.x + 80
  const y2 = toNode.y
  const midY = (y1 + y2) / 2

  return `M ${x1} ${y1} L ${x1} ${midY} L ${x2} ${midY} L ${x2} ${y2}`
}
const getEdgeMidpoint = (edge: Transition): { x: number, y: number } => {
  const fromNode = currentNodes.value.find(n => n.id === edge.from_node_id)
  const toNode = currentNodes.value.find(n => n.id === edge.to_node_id)
  if (!fromNode || !toNode) return { x: 0, y: 0 }
  const x1 = fromNode.x + 80
  const y1 = fromNode.y + getNodeHeight(fromNode)
  const x2 = toNode.x + 80
  const y2 = toNode.y
  return { x: (x1 + x2) / 2, y: (y1 + y2) / 2 }
}
// 事件处理
const handleNodeMouseDown = (e: MouseEvent, node: Node) => {
  e.preventDefault()
  if (e.shiftKey) {
    connectingFrom.value = node
  } else {
    // 注意:现在使用的是 SVG 坐标系,不需要考虑 scale
    draggingNode.value = {
      node,
      offsetX: e.clientX,
      offsetY: e.clientY
    }
  }
  selectedNode.value = node
}
const handleNodeMouseUp = (e: MouseEvent, node: Node) => {
  e.preventDefault() // ✅ 防止默认行为

  if (connectingFrom.value && connectingFrom.value.id !== node.id) {
    const edgeList = edges.value
    const exists = edgeList.some(edge => edge.from_node_id === connectingFrom.value!.id && edge.to_node_id === node.id)

    if (!exists) {
      edges.value.push({ from_node_id: connectingFrom.value.id, to_node_id: node.id })
    }
  }

  // ✅ 清除所有状态
  draggingNode.value = null
  connectingFrom.value = null
  connectingLine.value = null
}
const handleMouseMove = (e: MouseEvent) => {
  if (draggingNode.value && canvasRef.value) {
    const rect = canvasRef.value.getBoundingClientRect()

    // 计算鼠标在 SVG 坐标系中的位置
    const svgX = (e.clientX - rect.left - offset.value.x) / scale.value
    const svgY = (e.clientY - rect.top - offset.value.y) / scale.value

    const nodeList = nodes.value
    const index = nodeList.findIndex(n => n.id === draggingNode.value!.node.id)
    if (index !== -1) {
      // 保持节点中心在鼠标位置
      nodeList[index] = {
        ...nodeList[index],
        x: svgX - 80, // 节点宽度一半
        y: svgY - getNodeHeight(nodeList[index]) / 2
      }
    }
  } else if (connectingFrom.value && canvasRef.value) {
    const rect = canvasRef.value.getBoundingClientRect()
    const svgX = (e.clientX - rect.left - offset.value.x) / scale.value
    const svgY = (e.clientY - rect.top - offset.value.y) / scale.value

    connectingLine.value = {
      x1: connectingFrom.value.x + 80,
      y1: connectingFrom.value.y + getNodeHeight(connectingFrom.value),
      x2: svgX,
      y2: svgY
    }
  } else if (isPanning.value) {
    offset.value = {
      x: e.clientX - panStart.value.x,
      y: e.clientY - panStart.value.y
    }
  }
}
const handleMouseUp = () => {
  // ✅ 彻底清除所有拖拽状态
  draggingNode.value = null
  isPanning.value = false
  if (!connectingFrom.value) {
    connectingLine.value = null
  }
}
const handleCanvasMouseDown = (e: MouseEvent) => {
  const target = e.target as HTMLElement
  if (target === canvasRef.value || target.tagName === 'svg' || target.tagName === 'rect') {
    isPanning.value = true
    panStart.value = {
      x: e.clientX - offset.value.x,
      y: e.clientY - offset.value.y
    }
    selectedNode.value = null
  }
}
// 操作函数
const addNewNode = (type: NodeType) => {
  const newNode: Node = {
    id: Date.now(),
    x: (300 - offset.value.x) / scale.value,
    y: (200 - offset.value.y) / scale.value,
    name: nodeTypeConfigs[type].name,
    type,
    assignType: type === 'user_task' ? 'user' : null,
    assignTo: []
  }
  nodes.value.push(newNode)
}
const deleteNode = () => {
  if (!selectedNode.value) return
  nodes.value = nodes.value.filter(n => n.id !== selectedNode.value!.id)
  edges.value = edges.value.filter(
    e => e.from_node_id !== selectedNode.value!.id && e.to_node_id !== selectedNode.value!.id
  )
  selectedNode.value = null
}
const deleteEdge = (from: number, to: number) => {
  edges.value = edges.value.filter(e => !(e.from_node_id === from && e.to_node_id === to))
}
const updateNodeField = (field: keyof Node, value: any) => {
  if (!selectedNode.value) return
  const nodeList = nodes.value
  const index = nodeList.findIndex(n => n.id === selectedNode.value!.id)
  if (index !== -1) {
    (nodeList[index] as any)[field] = value
  }
}
const addAssignee = () => {
  if (!selectedNode.value || !selectedNode.value.assignTo) return
  selectedNode.value.assignTo.push('新成员')
  updateNodeField('assignTo', selectedNode.value.assignTo)
}
const removeAssignee = (index: number) => {
  if (!selectedNode.value || !selectedNode.value.assignTo) return
  selectedNode.value.assignTo.splice(index, 1)
  updateNodeField('assignTo', selectedNode.value.assignTo)
}
const switchMode = (mode: 'node' | 'stage') => {
  flowMode.value = mode
  selectedNode.value = null
}
const zoomIn = () => {
  if (!canvasRef.value) return
  const rect = canvasRef.value.getBoundingClientRect()
  const centerX = rect.width / 2
  const centerY = rect.height / 2

  const oldScale = scale.value
  const newScale = Math.min(scale.value + 0.1, 3)

  offset.value = {
    x: centerX - (centerX - offset.value.x) * (newScale / oldScale),
    y: centerY - (centerY - offset.value.y) * (newScale / oldScale)
  }

  scale.value = newScale
}

const zoomOut = () => {
  if (!canvasRef.value) return
  const rect = canvasRef.value.getBoundingClientRect()
  const centerX = rect.width / 2
  const centerY = rect.height / 2

  const oldScale = scale.value
  const newScale = Math.max(scale.value - 0.1, 0.2)

  offset.value = {
    x: centerX - (centerX - offset.value.x) * (newScale / oldScale),
    y: centerY - (centerY - offset.value.y) * (newScale / oldScale)
  }

  scale.value = newScale
}

const calculateBounds = () => {
  const padding = 100
  let minX = Infinity, minY = Infinity
  let maxX = -Infinity, maxY = -Infinity

  currentNodes.value.forEach((node) => {
    minX = Math.min(minX, node.x)
    minY = Math.min(minY, node.y)
    maxX = Math.max(maxX, node.x + 160)
    maxY = Math.max(maxY, node.y + getNodeHeight(node))
  })

  return {
    minX: minX - padding,
    minY: minY - padding,
    maxX: maxX + padding,
    maxY: maxY + padding,
    width: maxX - minX + padding * 2,
    height: maxY - minY + padding * 2
  }
}

const resetView = () => {
  if (!canvasRef.value || currentNodes.value.length === 0) {
    scale.value = 1
    offset.value = { x: 0, y: 0 }
    return
  }

  const bounds = calculateBounds()
  const rect = canvasRef.value.getBoundingClientRect()

  const scaleX = (rect.width * 0.9) / bounds.width
  const scaleY = (rect.height * 0.9) / bounds.height
  const newScale = Math.min(scaleX, scaleY, 1)

  const centerX = (bounds.minX + bounds.maxX) / 2
  const centerY = (bounds.minY + bounds.maxY) / 2

  offset.value = {
    x: rect.width / 2 - centerX * newScale,
    y: rect.height / 2 - centerY * newScale
  }

  scale.value = newScale
}

// ✅ 修复后的缩放函数
const handleWheel = (e: WheelEvent) => {
  if (!canvasRef.value) return
  e.preventDefault()

  const rect = canvasRef.value.getBoundingClientRect()
  const mouseX = e.clientX - rect.left
  const mouseY = e.clientY - rect.top

  const oldScale = scale.value
  const delta = e.deltaY > 0 ? -0.05 : 0.05
  const newScale = Math.max(0.2, Math.min(3, scale.value + delta))

  // 以鼠标位置为中心缩放
  offset.value = {
    x: mouseX - (mouseX - offset.value.x) * (newScale / oldScale),
    y: mouseY - (mouseY - offset.value.y) * (newScale / oldScale)
  }

  scale.value = newScale
}
const exportData = () => {
  const data = {
    mode: flowMode.value,
    nodes: currentNodes.value,
    edges: currentEdges.value
  }
  console.log('流程配置:', data)
  alert('配置已导出到控制台,按F12查看')
}
</script>

<template>
  <UContainer>
    <UPageHeader
      title="流程详情"
      :description="process?.name || '查看流程基本信息'"
    >
      <template #left>
        <UButton
          icon="i-lucide-arrow-left"
          variant="ghost"
          color="neutral"
          to="/process"
        >
          返回列表
        </UButton>
      </template>
    </UPageHeader>

    <UAlert
      v-if="error"
      color="error"
      variant="subtle"
      title="加载失败"
      :description="error.message || '流程信息获取失败'"
      class="mb-4"
    />

    <UCard v-else>
      <template #header>
        <div class="font-medium">
          基本信息
        </div>
      </template>

      <div
        v-if="loading"
        class="text-sm text-muted"
      >
        正在加载...
      </div>

      <div
        v-else-if="process"
        class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm"
      >
        <div>
          <span class="text-muted">ID：</span>{{ process.id }}
        </div>
        <div>
          <span class="text-muted">编码：</span>{{ process.code }}
        </div>
        <div>
          <span class="text-muted">名称：</span>{{ process.name }}
        </div>
        <div>
          <span class="text-muted">状态：</span>{{ process.status === 1 ? '正常' : '禁用' }}
        </div>
        <div class="md:col-span-2">
          <span class="text-muted">描述：</span>{{ process.description || '-' }}
        </div>
        <div>
          <span class="text-muted">创建人：</span>{{ process.created_by_name || '-' }}
        </div>
        <div>
          <span class="text-muted">更新人：</span>{{ process.updated_by_name || '-' }}
        </div>
        <div>
          <span class="text-muted">创建时间：</span>{{ dayjs(process.created_at).format('YYYY-MM-DD HH:mm:ss') }}
        </div>
        <div>
          <span class="text-muted">更新时间：</span>{{ dayjs(process.updated_at).format('YYYY-MM-DD HH:mm:ss') }}
        </div>
      </div>

      <div
        v-else
        class="text-sm text-muted"
      >
        未找到流程信息
      </div>
    </UCard>
    <div class="h-screen flex flex-col bg-base-200">
      <!-- 顶部工具栏 -->
      <div class="navbar bg-base-100 shadow-lg px-6">
        <div class="flex-1">
          <h1 class="text-2xl font-bold">
            信创适配流程配置
          </h1>
          <p class="text-sm text-base-content/60 ml-4">
            拖拽节点 · Shift+点击连线
          </p>
        </div>

        <div class="flex flex-row items-center gap-2">
          <!-- 模式切换 -->
          <div class="join">
            <button
              class="join-item btn btn-sm"
              :class="{ 'btn-primary': flowMode === 'node' }"
              @click="switchMode('node')"
            >
              节点性流程
            </button>
          </div>

          <!-- 缩放控制 -->
          <div class="join">
            <button
              class="join-item btn btn-sm"
              @click="zoomOut"
            >
              <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM13 10H7"
                />
              </svg>
            </button>
            <button
              class="join-item btn btn-sm"
              @click="resetView"
            >
              <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5l-5-5m5 5v-4m0 4h-4"
                />
              </svg>
            </button>
            <button
              class="join-item btn btn-sm"
              @click="zoomIn"
            >
              <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7"
                />
              </svg>
            </button>
          </div>

          <button
            class="btn btn-success btn-sm gap-2"
            @click="exportData"
          >
            <svg
              class="w-4 h-4"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M8 7H5a2 2 0 00-2 2v9a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-3m-1 4l-3 3m0 0l-3-3m3 3V4"
              />
            </svg>
            导出
          </button>

          <button
            class="btn btn-ghost btn-sm"
            @click="showConfig = !showConfig"
          >
            <svg
              class="w-4 h-4"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
              />
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
              />
            </svg>
          </button>
        </div>
      </div>

      <!-- 主内容区 -->
      <div class="flex-1 flex overflow-hidden">
        <!-- 画布区 -->
        <div
          ref="canvasRef"
          class="flex-1 relative bg-base-300 overflow-hidden"
          @mousedown="handleCanvasMouseDown"
          @mousemove="handleMouseMove"
          @mouseup="handleMouseUp"
          @wheel.prevent="handleWheel"
        >
          <!-- 左侧工具栏 -->
          <div class="absolute top-4 left-4 card bg-base-100 shadow-xl p-4 w-48 z-10">
            <h3 class="font-semibold text-sm mb-3">
              添加节点
            </h3>
            <div class="space-y-2">
              <button
                v-for="(config, type) in nodeTypeConfigs"
                :key="type"
                class="btn btn-sm btn-block justify-start gap-2"
                :class="`btn-outline ${config.color}`"
                @click="addNewNode(type as NodeType)"
              >
                <div
                  class="w-2 h-2 rounded-full"
                  :class="config.bgColor"
                />
                {{ config.name }}
              </button>
            </div>
          </div>

          <!-- 连线提示 -->
          <div
            v-if="connectingFrom"
            class="alert alert-warning absolute top-4 left-1/2 transform -translate-x-1/2 w-auto shadow-lg z-10"
          >
            <svg
              class="w-6 h-6"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <span>从 "{{ connectingFrom.name }}" 连线中,点击目标节点...</span>
          </div>

          <!-- 操作提示 -->
          <div class="absolute bottom-4 right-4 card bg-info text-info-content shadow-xl p-4 w-64 z-10">
            <h3 class="font-semibold mb-2">
              💡 操作提示
            </h3>
            <ul class="text-sm space-y-1 opacity-90">
              <li>• 拖拽节点调整位置</li>
              <li>• Shift + 点击开始连线</li>
              <li>• 点击连线的 × 删除</li>
              <li>• 拖动空白区域平移</li>
              <li>• 点击节点查看配置</li>
            </ul>
          </div>

          <!-- 缩放比例显示 -->
          <div class="absolute top-4 right-4 badge badge-lg badge-neutral z-10">
            {{ Math.round(scale * 100) }}%
          </div>

          <!-- SVG 画布 -->
          <svg
            class="w-full h-full"
            :viewBox="viewBox"
          >
            <defs>
              <marker
                id="arrowhead"
                markerWidth="10"
                markerHeight="10"
                refX="9"
                refY="3"
                orient="auto"
              >
                <polygon
                  points="0 0, 10 3, 0 6"
                  fill="#94a3b8"
                />
              </marker>
              <pattern
                id="grid"
                width="20"
                height="20"
                patternUnits="userSpaceOnUse"
              >
                <circle
                  cx="1"
                  cy="1"
                  r="1"
                  fill="currentColor"
                  class="text-base-content/20"
                />
              </pattern>
            </defs>

            <!-- 背景网格 -->
            <rect
              :x="viewBox.split(' ')[0]"
              :y="viewBox.split(' ')[1]"
              :width="viewBox.split(' ')[2]"
              :height="viewBox.split(' ')[3]"
              fill="url(#grid)"
            />

            <!-- 连线 -->
            <g
              v-for="edge in currentEdges"
              :key="`${edge.from}-${edge.to}`"
            >
              <path
                :d="getEdgePath(edge)"
                stroke="currentColor"
                class="stroke-base-content/40"
                stroke-width="2"
                fill="none"
                marker-end="url(#arrowhead)"
              />
              <circle
                :cx="getEdgeMidpoint(edge).x"
                :cy="getEdgeMidpoint(edge).y"
                r="8"
                fill="currentColor"
                class="fill-base-100 stroke-base-content/40 cursor-pointer hover:fill-error/20"
                stroke-width="2"
                @click="deleteEdge(edge.from, edge.to)"
              />
              <text
                :x="getEdgeMidpoint(edge).x"
                :y="getEdgeMidpoint(edge).y"
                text-anchor="middle"
                dominant-baseline="central"
                class="fill-error text-xs font-bold cursor-pointer pointer-events-none"
              >
                ×
              </text>
            </g>

            <!-- 临时连线 -->
            <line
              v-if="connectingLine"
              :x1="connectingLine.x1"
              :y1="connectingLine.y1"
              :x2="connectingLine.x2"
              :y2="connectingLine.y2"
              stroke="currentColor"
              class="stroke-primary"
              stroke-width="2"
              stroke-dasharray="5,5"
            />

            <!-- 节点 -->
            <g
              v-for="node in currentNodes"
              :key="node.id"
              :transform="`translate(${node.x}, ${node.y})`"
              class="cursor-move"
              @mousedown.stop="handleNodeMouseDown($event, node)"
              @mouseup.stop="handleNodeMouseUp($event, node)"
            >
              <rect
                x="0"
                y="0"
                width="160"
                :height="getNodeHeight(node)"
                rx="8"
                fill="currentColor"
                class="fill-base-100"
                :stroke="selectedNode?.id === node.id ? 'currentColor' : 'currentColor'"
                :class="selectedNode?.id === node.id ? 'stroke-primary' : 'stroke-base-content/20'"
                :stroke-width="selectedNode?.id === node.id ? '3' : '2'"
                style="filter: drop-shadow(0 2px 8px rgba(0,0,0,0.1))"
              />

              <!-- 节点头部 -->
              <rect
                x="0"
                y="0"
                width="160"
                height="35"
                rx="8"
                fill="currentColor"
                :class="nodeTypeConfigs[node.type].bgColor"
              />
              <rect
                x="0"
                y="8"
                width="160"
                height="27"
                fill="currentColor"
                :class="nodeTypeConfigs[node.type].bgColor"
              />

              <text
                x="80"
                y="23"
                text-anchor="middle"
                class="fill-base-100 text-sm font-semibold"
              >
                {{ node.name }}
              </text>

              <!-- 执行人信息 -->
              <g v-if="node.type === 'user_task' && node.assignTo && node.assignTo.length > 0">
                <rect
                  x="8"
                  y="45"
                  width="144"
                  height="32"
                  rx="4"
                  class="fill-base-200"
                />
                <text
                  x="14"
                  y="63"
                  class="fill-base-content text-xs"
                >
                  {{ node.assignTo.join(', ') }}
                </text>
              </g>

              <!-- 连接点 -->
              <circle
                cx="80"
                cy="0"
                r="6"
                fill="currentColor"
                :class="nodeTypeConfigs[node.type].bgColor"
                stroke="currentColor"
                class="stroke-base-100"
                stroke-width="2"
                style="cursor: crosshair"
              />
              <circle
                cx="80"
                :cy="getNodeHeight(node)"
                r="6"
                fill="currentColor"
                :class="nodeTypeConfigs[node.type].bgColor"
                stroke="currentColor"
                class="stroke-base-100"
                stroke-width="2"
                style="cursor: crosshair"
              />
            </g>
          </svg>
        </div>

        <!-- 右侧配置面板 -->
        <div
          v-if="showConfig"
          class="w-80 bg-base-100 shadow-xl p-6 overflow-y-auto"
        >
          <h2 class="text-lg font-bold mb-4 flex items-center gap-2">
            <svg
              class="w-5 h-5"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
              />
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
              />
            </svg>
            节点配置
          </h2>

          <div
            v-if="!selectedNode"
            class="text-center py-12"
          >
            <svg
              class="w-12 h-12 mx-auto mb-3 text-base-content/30"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"
              />
            </svg>
            <p class="text-base-content/60">
              点击节点进行配置
            </p>
          </div>

          <div
            v-else
            class="space-y-4"
          >
            <div class="alert alert-info">
              <span class="text-xs">ID: {{ selectedNode.id }}</span>
            </div>

            <div class="form-control">
              <label class="label">
                <span class="label-text font-medium">节点名称</span>
              </label>
              <input
                v-model="selectedNode.name"
                type="text"
                class="input input-bordered"
                @input="updateNodeField('name', selectedNode.name)"
              >
            </div>

            <div class="form-control">
              <label class="label">
                <span class="label-text font-medium">节点类型</span>
              </label>
              <select
                v-model="selectedNode.type"
                class="select select-bordered"
                @change="updateNodeField('type', selectedNode.type)"
              >
                <option
                  v-for="(config, type) in nodeTypeConfigs"
                  :key="type"
                  :value="type"
                >
                  {{ config.name }}
                </option>
              </select>
            </div>

            <template v-if="selectedNode.type === 'user_task'">
              <div class="form-control">
                <label class="label">
                  <span class="label-text font-medium">分配方式</span>
                </label>
                <select
                  v-model="selectedNode.assignType"
                  class="select select-bordered"
                  @change="updateNodeField('assignType', selectedNode.assignType)"
                >
                  <option value="user">
                    指定用户
                  </option>
                  <option value="role">
                    指定角色
                  </option>
                </select>
              </div>

              <div class="form-control">
                <label class="label">
                  <span class="label-text font-medium">
                    {{ selectedNode.assignType === 'role' ? '执行角色' : '执行用户' }}
                  </span>
                </label>
                <div class="space-y-2">
                  <div
                    v-for="(_, idx) in selectedNode.assignTo ?? []"
                    :key="idx"
                    class="join w-full"
                  >
                    <input
                      v-model="selectedNode.assignTo[idx]"
                      type="text"
                      class="input input-bordered join-item flex-1"
                      @input="updateNodeField('assignTo', selectedNode.assignTo)"
                    >
                    <button
                      class="btn btn-error join-item"
                      @click="removeAssignee(idx)"
                    >
                      <svg
                        class="w-4 h-4"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path
                          stroke-linecap="round"
                          stroke-linejoin="round"
                          stroke-width="2"
                          d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                        />
                      </svg>
                    </button>
                  </div>
                  <button
                    class="btn btn-outline btn-block btn-sm"
                    @click="addAssignee"
                  >
                    <svg
                      class="w-4 h-4"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                      />
                    </svg>
                    添加成员
                  </button>
                </div>
              </div>
            </template>

            <button
              class="btn btn-error btn-block gap-2 mt-6"
              @click="deleteNode"
            >
              <svg
                class="w-4 h-4"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                />
              </svg>
              删除节点
            </button>
          </div>
        </div>
      </div>
    </div>
  </UContainer>
</template>
